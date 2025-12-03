package client

import "net/http"

// HeaderConfig holds service-specific header configuration.
type HeaderConfig struct {
	Origin    string
	BaseURL   string
	UserAgent string
}

// DefaultUserAgent is the browser User-Agent string.
const DefaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36"

// getUserAgent returns the user agent from config or default.
func getUserAgent(cfg *HeaderConfig) string {
	if cfg != nil && cfg.UserAgent != "" {
		return cfg.UserAgent
	}
	return DefaultUserAgent
}

// SetCommonHeadersWithConfig sets browser-like headers using config.
func SetCommonHeadersWithConfig(req *http.Request, cfg *HeaderConfig) {
	req.Header.Set("User-Agent", getUserAgent(cfg))
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa;q=0.8,ar;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
}

// SetAPIHeadersWithConfig sets headers for API requests using config.
func SetAPIHeadersWithConfig(req *http.Request, csrfToken string, cfg *HeaderConfig) {
	req.Header.Set("User-Agent", getUserAgent(cfg))
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa;q=0.8,ar;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("sec-ch-ua", `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	if csrfToken != "" {
		req.Header.Set("x-csrf-token", csrfToken)
	}
}

// SetFormHeadersWithConfig sets headers for form submissions using config.
func SetFormHeadersWithConfig(req *http.Request, referer string, cfg *HeaderConfig) {
	req.Header.Set("User-Agent", getUserAgent(cfg))
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,fa;q=0.8,ar;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if cfg != nil && cfg.Origin != "" {
		req.Header.Set("Origin", cfg.Origin)
	}
	req.Header.Set("Referer", referer)
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua", `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
}
