package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/m3ter-com/m3ter-sdk-go/option"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func Middleware(ctx context.Context) option.Middleware {
	return func(req *http.Request, next option.MiddlewareNext) (*http.Response, error) {
		if req != nil {
			if err := LogRequest(ctx, req); err != nil {
				return nil, err
			}
		}

		resp, err := next(req)

		if resp != nil {
			if err := LogResponse(ctx, resp); err != nil {
				return nil, err
			}
		}

		return resp, err
	}
}

func LogRequest(ctx context.Context, req *http.Request) error {
	lines, err := requestLines(req)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

// requestLines builds the log lines for a request. Split out from LogRequest so
// the redaction can be asserted end to end: testing RedactHeader alone would
// still pass if someone dropped the call from here, which is the mistake most
// likely to reintroduce the leak.
func requestLines(req *http.Request) ([]string, error) {
	lines := []string{fmt.Sprintf("\n%s %s %s", req.Method, req.URL.Path, req.Proto)}

	// Log headers, redacting credentials — see redact.go. TF_LOG output is
	// routinely pasted into issues and CI logs.
	for name, values := range req.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("> %s: %s", strings.ToLower(name), RedactHeader(name, value)))
		}
	}

	if req.Body != nil {
		// Read the body without mutating the original response
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}

		// Restore the original body to the response so it can be read again
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Log the body with sensitive values removed; a non-JSON body is not
		// logged at all (see RedactBody).
		lines = append(lines, ">\n", RedactBody(bodyBytes), "\n")
	}

	return lines, nil
}

func LogResponse(ctx context.Context, resp *http.Response) error {
	lines, err := responseLines(resp)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, strings.Join(lines, "\n"))

	return nil
}

// responseLines builds the log lines for a response; split out for the same
// reason as requestLines.
func responseLines(resp *http.Response) ([]string, error) {
	// Log the status code
	lines := []string{fmt.Sprintf("\n< %s %s", resp.Proto, resp.Status)}

	// Log headers, redacting credentials — a response can carry Set-Cookie or a
	// refreshed token just as a request carries Authorization.
	for name, values := range resp.Header {
		for _, value := range values {
			lines = append(lines, fmt.Sprintf("< %s: %s", strings.ToLower(name), RedactHeader(name, value)))
		}
	}

	// Read the body without mutating the original response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Restore the original body to the response so it can be read again
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	lines = append(lines, "<\n", RedactBody(bodyBytes), "\n")

	return lines, nil
}
