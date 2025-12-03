package mytax

import (
	"net/http"
	"strings"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// Client wraps an HTTP client with MyTax-specific configuration.
type Client struct {
	httpClient *http.Client
	config     *client.HeaderConfig
}

// NewClient creates a new MyTax HTTP client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: client.NewHTTPClient(&cfg.HTTP),
		config: &client.HeaderConfig{
			Origin:  cfg.Services.MyTax.BaseURL,
			BaseURL: cfg.Services.MyTax.BaseURL,
		},
	}
}

// Do executes an HTTP request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// SetCommonHeaders sets browser-like headers.
func (c *Client) SetCommonHeaders(req *http.Request) {
	client.SetCommonHeadersWithConfig(req, c.config)
}

// SetAPIHeaders sets API request headers.
func (c *Client) SetAPIHeaders(req *http.Request, csrfToken string) {
	client.SetAPIHeadersWithConfig(req, csrfToken, c.config)
}

// SetFormHeaders sets form submission headers.
func (c *Client) SetFormHeaders(req *http.Request, referer string) {
	client.SetFormHeadersWithConfig(req, referer, c.config)
}

// SetNavigationHeaders sets headers that browsers send during page navigation.
// These Sec-Fetch-* headers are required for cross-domain auth to register.tax.gov.ir.
// The server checks these headers to validate legitimate browser navigations.
func (c *Client) SetNavigationHeaders(req *http.Request, referer string) {
	c.SetCommonHeaders(req)

	// Browser fetch metadata headers
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")

	// Determine Sec-Fetch-Site based on referer domain
	if strings.Contains(referer, "my.tax.gov.ir") {
		req.Header.Set("Sec-Fetch-Site", "same-site") // Same *.tax.gov.ir domain
	} else if strings.Contains(referer, "register.tax.gov.ir") {
		req.Header.Set("Sec-Fetch-Site", "same-origin")
	} else {
		req.Header.Set("Sec-Fetch-Site", "cross-site")
	}

	if referer != "" {
		req.Header.Set("Referer", referer)
	}
}

// AddCookies adds cookies to the request.
func (c *Client) AddCookies(req *http.Request, cookies []*http.Cookie) {
	client.AddCookies(req, cookies)
}

// Origin returns the configured origin.
func (c *Client) Origin() string {
	return c.config.Origin
}

// BaseURL returns the base URL.
func (c *Client) BaseURL() string {
	return c.config.BaseURL
}

// HTTPClient returns the underlying http.Client for external use.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
