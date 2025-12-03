package mygovauth

import (
	"net/http"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// Client wraps an HTTP client with MyGovAuth-specific configuration.
type Client struct {
	httpClient *http.Client
	config     *client.HeaderConfig
	authHeader string
}

// NewClient creates a new MyGovAuth HTTP client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: client.NewHTTPClient(&cfg.HTTP),
		config: &client.HeaderConfig{
			Origin:  cfg.Services.MyGovAuth.BaseURL,
			BaseURL: cfg.Services.MyGovAuth.BaseURL,
		},
		authHeader: cfg.Services.MyGovAuth.AuthHeader,
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

// AuthHeader returns the authorization header for captcha requests.
func (c *Client) AuthHeader() string {
	return c.authHeader
}

// HTTPClient returns the underlying http.Client for external use.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
