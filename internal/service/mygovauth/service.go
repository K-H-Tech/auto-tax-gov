package mygovauth

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/helpers"
	"github.com/K-H-Tech/auto-tax-gov/internal/models"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
)

// Service handles government SSO authentication operations.
type Service struct {
	cfg    *config.Config
	client *Client
	logger *slog.Logger
}

// New creates a new MyGovAuth service instance.
func New(cfg *config.Config, logger *slog.Logger) *Service {
	return &Service{
		cfg:    cfg,
		client: NewClient(cfg),
		logger: logger,
	}
}

// AuthResult represents the result of an authentication operation.
type AuthResult struct {
	Success     bool
	Message     string
	RedirectURL string
	Data        map[string]any
}

// InitiateAuth starts the OAuth2 authentication flow for a specific service.
// It follows redirects until it reaches the login page and returns captcha data.
func (s *Service) InitiateAuth(sess *session.Session, authURL string) (*models.CaptchaData, []models.RedirectStep, error) {
	s.logger.Info("initiating auth flow", "url", authURL)

	steps, err := s.followRedirects(sess, authURL)
	if err != nil {
		return nil, steps, fmt.Errorf("error following redirects: %w", err)
	}

	// Look for captcha in the steps
	var captcha *models.CaptchaData
	for i := len(steps) - 1; i >= 0; i-- {
		step := steps[i]
		// Check if this is the login page
		if strings.Contains(step.URL, "sso.my.gov.ir/login") && step.StatusCode == 200 {
			captcha, err = s.FetchCaptcha(sess, step.ResponseBody, step.URL)
			if err != nil {
				s.logger.Warn("could not fetch captcha", "error", err)
			} else {
				steps[i].CaptchaData = captcha
				sess.SetCaptcha(captcha)
			}
			break
		}
	}

	if captcha == nil {
		return nil, steps, fmt.Errorf("no captcha found - may not have reached login page")
	}

	return captcha, steps, nil
}

// followRedirects follows the redirect chain until a non-redirect response.
func (s *Service) followRedirects(sess *session.Session, startURL string) ([]models.RedirectStep, error) {
	var steps []models.RedirectStep
	currentURL := startURL
	stepNumber := 1
	maxSteps := s.cfg.HTTP.MaxRedirects

	for stepNumber <= maxSteps {
		s.logger.Debug("tracking step", "step", stepNumber, "url", currentURL)

		step := models.RedirectStep{
			StepNumber: stepNumber,
			URL:        currentURL,
			Method:     "GET",
			Timestamp:  time.Now(),
		}

		req, err := http.NewRequest("GET", currentURL, nil)
		if err != nil {
			return steps, fmt.Errorf("error creating request for step %d: %w", stepNumber, err)
		}

		s.client.SetCommonHeaders(req)
		s.client.AddCookies(req, sess.GetCookies())

		// Store sanitized headers to avoid exposing sensitive data (Issue 19)
		step.RequestHeaders = helpers.SanitizeHeaders(req.Header)

		// Dump request without body to avoid exposing sensitive form data
		if requestDump, err := httputil.DumpRequestOut(req, false); err == nil {
			step.RequestBody = string(requestDump)
		}

		resp, err := s.client.Do(req)
		if err != nil {
			return steps, fmt.Errorf("error making request for step %d: %w", stepNumber, err)
		}

		step.StatusCode = resp.StatusCode
		step.Status = resp.Status
		step.ResponseHeaders = resp.Header.Clone()

		if cookies := resp.Cookies(); len(cookies) > 0 {
			sess.MergeCookies(cookies)
			s.logger.Debug("saved cookies", "count", len(cookies))
		}

		bodyBytes, err := client.ReadResponseBody(resp)
		resp.Body.Close()
		if err != nil {
			return steps, fmt.Errorf("error reading response body for step %d: %w", stepNumber, err)
		}
		step.ResponseBody = string(bodyBytes)

		// Check for redirect
		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location := resp.Header.Get("Location")
			if location == "" {
				s.logger.Info("redirect status but no Location header, stopping", "status", resp.StatusCode)
				steps = append(steps, step)
				break
			}

			step.RedirectLocation = location
			steps = append(steps, step)

			// Handle relative URLs
			currentURL = s.resolveURL(req, location)
			s.logger.Debug("redirect to", "url", currentURL)
			stepNumber++
		} else {
			steps = append(steps, step)
			s.logger.Info("final response", "status", resp.StatusCode)
			break
		}
	}

	if stepNumber > maxSteps {
		return steps, fmt.Errorf("exceeded maximum redirect steps (%d)", maxSteps)
	}

	return steps, nil
}

// FollowRedirectChain follows redirects from a starting URL until a non-redirect response.
// Returns the final URL and updated cookies.
func (s *Service) FollowRedirectChain(sess *session.Session, startURL string) (string, error) {
	currentURL := startURL
	// Use configured maxRedirects with sensible fallback (Issue 16)
	maxRedirects := s.cfg.HTTP.MaxRedirects
	if maxRedirects <= 0 {
		maxRedirects = 10
	}

	for i := 0; i < maxRedirects; i++ {
		req, err := http.NewRequest("GET", currentURL, nil)
		if err != nil {
			return currentURL, err
		}

		s.client.SetCommonHeaders(req)
		s.client.AddCookies(req, sess.GetCookies())

		resp, err := s.client.Do(req)
		if err != nil {
			return currentURL, err
		}
		resp.Body.Close()

		if cookies := resp.Cookies(); len(cookies) > 0 {
			// Log cookie names at INFO level for debugging TaxpayerToken issue
			cookieNames := make([]string, len(cookies))
			for j, c := range cookies {
				cookieNames[j] = c.Name
			}
			s.logger.Info("cookies received in redirect chain",
				"step", i+1,
				"url", currentURL,
				"cookieNames", cookieNames)
			sess.MergeCookies(cookies)
		}

		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location := resp.Header.Get("Location")
			if location == "" {
				return currentURL, nil
			}
			// Resolve relative URLs against current URL
			resolvedLocation := s.resolveURL(req, location)
			s.logger.Info("redirect chain", "step", i+1, "from", currentURL, "to", resolvedLocation, "status", resp.StatusCode)
			currentURL = resolvedLocation
		} else {
			s.logger.Debug("final URL reached", "url", currentURL, "status", resp.StatusCode)
			return currentURL, nil
		}
	}

	return currentURL, nil
}

// resolveURL resolves a potentially relative URL against the request URL.
func (s *Service) resolveURL(req *http.Request, location string) string {
	if strings.HasPrefix(location, "/") {
		if strings.Contains(location, "?") {
			parts := strings.SplitN(location, "?", 2)
			return req.URL.Scheme + "://" + req.URL.Host + parts[0] + "?" + parts[1]
		}
		return req.URL.Scheme + "://" + req.URL.Host + location
	} else if strings.HasPrefix(location, "http") {
		return location
	}
	return req.URL.Scheme + "://" + req.URL.Host + "/" + location
}

// GetConfig returns the service configuration.
func (s *Service) GetConfig() *config.MyGovAuthConfig {
	return &s.cfg.Services.MyGovAuth
}

// GetHTTPClient returns the HTTP client for external use.
func (s *Service) GetHTTPClient() *http.Client {
	return s.client.HTTPClient()
}

// GetClient returns the service client.
func (s *Service) GetClient() *Client {
	return s.client
}
