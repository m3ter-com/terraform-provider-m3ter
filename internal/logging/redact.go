package logging

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Redaction for the provider's TF_LOG=DEBUG HTTP tracing (BR-4209).
//
// The tracing previously wrote every request and response header verbatim,
// including Authorization and x-api-key, plus both bodies in full. TF_LOG=DEBUG
// is a documented Terraform variable that users are routinely asked to enable
// when reporting a problem, and the resulting log is then pasted into an issue,
// a CI job or a support ticket. Anything written here should be assumed to end
// up somewhere public.
//
// The bias throughout is that over-redacting costs a debugging round-trip while
// under-redacting discloses a live credential, so ambiguous cases are redacted.

const redacted = "REDACTED"

// Header names that are always credentials, matched exactly (lower-cased).
var sensitiveHeaders = map[string]bool{
	"authorization":        true,
	"proxy-authorization":  true,
	"cookie":               true,
	"set-cookie":           true,
	"x-api-key":            true,
	"x-amz-security-token": true,
}

// Substrings that make a header name sensitive whatever else it is called. This
// catches vendor- and future-specific names (x-m3ter-token, x-refresh-secret,
// ...) that an exact list would miss, which matters because a header added later
// would otherwise silently start leaking.
var sensitiveHeaderParts = []string{
	"auth", "token", "secret", "credential", "password", "apikey", "api-key",
}

// JSON keys whose values are redacted anywhere in a body, at any depth.
var sensitiveBodyKeys = []string{
	"secret", "password", "token", "apikey", "api_key", "credential",
	"privatekey", "private_key", "clientsecret", "client_secret",
}

// RedactHeader reports the value to log for a header, redacting credentials.
func RedactHeader(name, value string) string {
	lower := strings.ToLower(name)
	if sensitiveHeaders[lower] {
		return redacted
	}
	for _, part := range sensitiveHeaderParts {
		if strings.Contains(lower, part) {
			return redacted
		}
	}
	return value
}

// RedactBody returns a loggable rendering of an HTTP body.
//
// A JSON body is re-marshalled with sensitive keys replaced, which keeps the
// shape useful for debugging while removing the values. Anything that is not
// JSON is not logged at all: the provider speaks JSON, so a non-JSON body is
// unexpected, and guessing at how to sanitise an unknown format is exactly the
// kind of assumption that leaks. The length is kept so its presence is visible.
func RedactBody(body []byte) string {
	if len(body) == 0 {
		return ""
	}

	var parsed any
	if err := json.Unmarshal(body, &parsed); err != nil {
		return fmt.Sprintf("<non-JSON body, %d bytes, not logged>", len(body))
	}

	cleaned, err := json.Marshal(redactValue(parsed))
	if err != nil {
		// Should not happen — it round-tripped from JSON — but never fall back
		// to printing the original.
		return fmt.Sprintf("<body, %d bytes, not logged>", len(body))
	}
	return string(cleaned)
}

func redactValue(v any) any {
	switch typed := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		for key, value := range typed {
			if isSensitiveKey(key) {
				out[key] = redacted
				continue
			}
			out[key] = redactValue(value)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for i, value := range typed {
			out[i] = redactValue(value)
		}
		return out
	default:
		return v
	}
}

func isSensitiveKey(key string) bool {
	// Compared with separators removed so secret, client_secret, clientSecret
	// and client-secret are all caught by one entry.
	normalised := strings.ToLower(key)
	normalised = strings.ReplaceAll(normalised, "_", "")
	normalised = strings.ReplaceAll(normalised, "-", "")
	for _, candidate := range sensitiveBodyKeys {
		if strings.Contains(normalised, strings.ReplaceAll(candidate, "_", "")) {
			return true
		}
	}
	return false
}
