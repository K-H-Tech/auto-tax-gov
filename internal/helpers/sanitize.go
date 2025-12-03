package helpers

import (
	"net/http"
	"net/url"
	"strings"
)

// MaskMobile masks a mobile number for safe logging (e.g., "09*****34").
// Preserves first 2 and last 2 characters for context while hiding the rest.
func MaskMobile(mobile string) string {
	if len(mobile) < 4 {
		return "****"
	}
	return mobile[:2] + strings.Repeat("*", len(mobile)-4) + mobile[len(mobile)-2:]
}

// SanitizeURL removes query parameters from URL for safe logging.
// Returns only scheme, host, and path without sensitive query params.
func SanitizeURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "[invalid-url]"
	}
	return u.Scheme + "://" + u.Host + u.Path
}

// SanitizeHeaders removes sensitive headers for safe persistence/logging.
// Replaces values of sensitive headers with "[REDACTED]".
func SanitizeHeaders(headers http.Header) http.Header {
	sensitiveHeaders := []string{
		"Authorization",
		"Cookie",
		"Set-Cookie",
		"X-Auth-Token",
		"X-Api-Key",
	}
	sanitized := headers.Clone()
	for _, h := range sensitiveHeaders {
		if sanitized.Get(h) != "" {
			sanitized.Set(h, "[REDACTED]")
		}
	}
	return sanitized
}

// BuildQueryURL safely constructs a URL with encoded query parameters.
// Parameters are properly URL-encoded to prevent injection and malformed URLs.
func BuildQueryURL(baseURL string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
