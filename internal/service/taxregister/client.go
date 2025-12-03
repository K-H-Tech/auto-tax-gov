package taxregister

import (
	"net/http"
	"strings"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// Client wraps an HTTP client for tax registration operations.
// It handles requests to both my.tax.gov.ir and register.tax.gov.ir.
type Client struct {
	httpClient *http.Client
	mytaxCfg   *client.HeaderConfig
	regCfg     *client.HeaderConfig
}

// NewClient creates a new tax registration HTTP client.
func NewClient(cfg *config.Config) *Client {
	return &Client{
		httpClient: client.NewHTTPClient(&cfg.HTTP),
		mytaxCfg: &client.HeaderConfig{
			Origin:  cfg.Services.MyTax.BaseURL,
			BaseURL: cfg.Services.MyTax.BaseURL,
		},
		regCfg: &client.HeaderConfig{
			Origin:  cfg.Services.RegisterTax.BaseURL,
			BaseURL: cfg.Services.RegisterTax.BaseURL,
		},
	}
}

// Do executes an HTTP request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// SetCommonHeaders sets browser-like headers for my.tax.gov.ir.
func (c *Client) SetCommonHeaders(req *http.Request) {
	client.SetCommonHeadersWithConfig(req, c.mytaxCfg)
}

// SetRegisterHeaders sets browser-like headers for register.tax.gov.ir.
func (c *Client) SetRegisterHeaders(req *http.Request) {
	client.SetCommonHeadersWithConfig(req, c.regCfg)
}

// SetAPIHeaders sets API request headers for my.tax.gov.ir.
func (c *Client) SetAPIHeaders(req *http.Request, csrfToken string) {
	client.SetAPIHeadersWithConfig(req, csrfToken, c.mytaxCfg)
}

// SetNavigationHeaders sets headers that browsers send during page navigation.
// These Sec-Fetch-* headers are required for cross-domain auth.
func (c *Client) SetNavigationHeaders(req *http.Request, referer string) {
	// Determine which config to use based on request URL
	if strings.Contains(req.URL.Host, "register.tax.gov.ir") {
		client.SetCommonHeadersWithConfig(req, c.regCfg)
	} else {
		client.SetCommonHeadersWithConfig(req, c.mytaxCfg)
	}

	// Browser fetch metadata headers
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")

	// Determine Sec-Fetch-Site based on referer domain
	if strings.Contains(referer, "my.tax.gov.ir") && strings.Contains(req.URL.Host, "register.tax.gov.ir") {
		req.Header.Set("Sec-Fetch-Site", "same-site") // Same *.tax.gov.ir domain
	} else if strings.Contains(referer, "register.tax.gov.ir") && strings.Contains(req.URL.Host, "register.tax.gov.ir") {
		req.Header.Set("Sec-Fetch-Site", "same-origin")
	} else {
		req.Header.Set("Sec-Fetch-Site", "cross-site")
	}

	if referer != "" {
		req.Header.Set("Referer", referer)
	}
}

// SetAjaxHeaders sets headers for ASP.NET AJAX partial postback requests.
// These headers are required for the X-MicrosoftAjax: Delta=true pattern.
func (c *Client) SetAjaxHeaders(req *http.Request, referer string) {
	client.SetCommonHeadersWithConfig(req, c.regCfg)

	// AJAX-specific headers from step_08.raw
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-MicrosoftAjax", "Delta=true")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	// Sec-Fetch headers for AJAX
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")

	req.Header.Set("Origin", c.regCfg.Origin)
	if referer != "" {
		req.Header.Set("Referer", referer)
	}
}

// SetFormSubmitHeaders sets headers for full form POST submission.
func (c *Client) SetFormSubmitHeaders(req *http.Request, referer string) {
	client.SetCommonHeadersWithConfig(req, c.regCfg)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", c.regCfg.Origin)
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// Sec-Fetch headers for form submission
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")

	if referer != "" {
		req.Header.Set("Referer", referer)
	}
}

// AddCookies adds cookies to the request.
func (c *Client) AddCookies(req *http.Request, cookies []*http.Cookie) {
	client.AddCookies(req, cookies)
}

// MyTaxOrigin returns the my.tax.gov.ir origin.
func (c *Client) MyTaxOrigin() string {
	return c.mytaxCfg.Origin
}

// MyTaxBaseURL returns the my.tax.gov.ir base URL.
func (c *Client) MyTaxBaseURL() string {
	return c.mytaxCfg.BaseURL
}

// RegisterOrigin returns the register.tax.gov.ir origin.
func (c *Client) RegisterOrigin() string {
	return c.regCfg.Origin
}

// RegisterBaseURL returns the register.tax.gov.ir base URL.
func (c *Client) RegisterBaseURL() string {
	return c.regCfg.BaseURL
}

// HTTPClient returns the underlying http.Client for external use.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}
