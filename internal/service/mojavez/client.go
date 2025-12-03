package mojavez

import (
	"net/http"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// Client wraps an HTTP client with Mojavez-specific configuration.
type Client struct {
	httpClient *http.Client
	config     *client.HeaderConfig
}

// NewClient creates a new Mojavez HTTP client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: client.NewHTTPClient(&cfg.HTTP),
		config: &client.HeaderConfig{
			Origin:  cfg.Services.Mojavez.BaseURL,
			BaseURL: cfg.Services.Mojavez.BaseURL,
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
