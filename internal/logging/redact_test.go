package logging

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

// The credential shapes this provider actually sends. If any of these survives
// into a log line, TF_LOG=DEBUG output pasted into an issue discloses a live
// credential — which is the bug this guards (BR-4209).
func TestRedactHeaderRemovesCredentials(t *testing.T) {
	secret := "sk_live_abcdef0123456789abcdef"
	for _, name := range []string{
		"Authorization", "authorization", "X-Api-Key", "x-api-key",
		"Cookie", "Set-Cookie", "Proxy-Authorization", "X-Amz-Security-Token",
		// Not on any list — caught by the substring rule, which is the point:
		// a header added later must not silently start leaking.
		"X-M3ter-Refresh-Token", "X-Custom-Secret", "X-Some-Credential",
	} {
		if got := RedactHeader(name, secret); got != redacted {
			t.Errorf("%s leaked: got %q", name, got)
		}
	}
}

func TestRedactHeaderKeepsDiagnosticHeaders(t *testing.T) {
	// Over-redacting costs a debugging round-trip, so the ordinary headers that
	// make the log worth having must survive.
	for name, value := range map[string]string{
		"Content-Type":   "application/json",
		"User-Agent":     "terraform-provider-m3ter/1.0",
		"X-Request-Id":   "abc-123",
		"Accept":         "application/json",
		"Content-Length": "42",
	} {
		if got := RedactHeader(name, value); got != value {
			t.Errorf("%s should not be redacted: got %q", name, got)
		}
	}
}

func TestRedactBodyRemovesSecretsAtAnyDepth(t *testing.T) {
	body := []byte(`{
	  "name": "svc",
	  "apiKey": "live-key-1",
	  "nested": {"clientSecret": "shhh", "keep": "visible"},
	  "list": [{"password": "hunter2"}, {"id": "ok"}]
	}`)

	got := RedactBody(body)

	for _, leak := range []string{"live-key-1", "shhh", "hunter2"} {
		if strings.Contains(got, leak) {
			t.Errorf("secret %q survived redaction: %s", leak, got)
		}
	}
	// The shape has to stay useful or the log stops being worth writing.
	for _, keep := range []string{"svc", "visible", "ok"} {
		if !strings.Contains(got, keep) {
			t.Errorf("non-secret %q was lost: %s", keep, got)
		}
	}
	var parsed any
	if err := json.Unmarshal([]byte(got), &parsed); err != nil {
		t.Errorf("redacted body is not valid JSON: %v", err)
	}
}

func TestRedactBodyDoesNotLogNonJSON(t *testing.T) {
	// The provider speaks JSON. Guessing how to sanitise an unknown format is
	// how secrets escape, so an unparseable body is described, not printed.
	body := []byte("token=sk_live_abcdef0123456789&user=admin")
	got := RedactBody(body)
	if strings.Contains(got, "sk_live_abcdef0123456789") {
		t.Errorf("non-JSON body was logged verbatim: %s", got)
	}
	if !strings.Contains(got, "not logged") {
		t.Errorf("expected a placeholder describing the body, got: %s", got)
	}
}

func TestRedactBodyHandlesEmpty(t *testing.T) {
	if got := RedactBody(nil); got != "" {
		t.Errorf("empty body should render empty, got %q", got)
	}
}

// End-to-end over the actual log lines. These are the tests that fail if the
// redaction call is dropped from logging.go — the RedactHeader/RedactBody unit
// tests above would keep passing, which is precisely how this leak could come
// back unnoticed.
func TestRequestLinesRedactAuthAndBody(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodPost, "https://api.m3ter.com/v1/things",
		strings.NewReader(`{"name":"svc","apiKey":"live-key-1"}`),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer sk_live_abcdef0123456789")
	req.Header.Set("X-Api-Key", "live-key-1")
	req.Header.Set("Content-Type", "application/json")

	lines, err := requestLines(req)
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(lines, "\n")

	for _, leak := range []string{"sk_live_abcdef0123456789", "live-key-1"} {
		if strings.Contains(joined, leak) {
			t.Errorf("credential %q reached the log lines:\n%s", leak, joined)
		}
	}
	if !strings.Contains(joined, "application/json") {
		t.Errorf("diagnostic headers were lost:\n%s", joined)
	}
}

func TestResponseLinesRedactSetCookieAndBody(t *testing.T) {
	resp := &http.Response{
		Proto:  "HTTP/1.1",
		Status: "200 OK",
		Header: http.Header{
			"Set-Cookie":   []string{"session=sk_live_response_secret"},
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(strings.NewReader(`{"token":"resp-token-1","id":"visible"}`)),
	}

	lines, err := responseLines(resp)
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(lines, "\n")

	for _, leak := range []string{"sk_live_response_secret", "resp-token-1"} {
		if strings.Contains(joined, leak) {
			t.Errorf("credential %q reached the log lines:\n%s", leak, joined)
		}
	}
	if !strings.Contains(joined, "visible") {
		t.Errorf("non-secret response content was lost:\n%s", joined)
	}
}

// The body must still be readable by the caller — redaction happens on a copy.
func TestRequestBodyIsRestoredAfterLogging(t *testing.T) {
	const payload = `{"name":"svc"}`
	req, err := http.NewRequest(http.MethodPost, "https://api.m3ter.com/v1/things",
		strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := requestLines(req); err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != payload {
		t.Errorf("request body was consumed by logging: got %q want %q", got, payload)
	}
}
