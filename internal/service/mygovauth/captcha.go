package mygovauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/models"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
)

// FetchCaptcha retrieves captcha data from the SSO server.
func (s *Service) FetchCaptcha(sess *session.Session, htmlBody, refererURL string) (*models.CaptchaData, error) {
	csrfToken := extractCSRFToken(htmlBody)
	if csrfToken == "" {
		return nil, fmt.Errorf("could not extract CSRF token from HTML")
	}

	req, err := http.NewRequest("GET", s.cfg.Services.MyGovAuth.CaptchaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating captcha request: %w", err)
	}

	s.client.SetAPIHeaders(req, csrfToken)
	req.Header.Set("Referer", refererURL)
	req.Header.Set("Authorization", s.client.AuthHeader())
	s.client.AddCookies(req, sess.GetCookies())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making captcha request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		bodyPreview, _ := io.ReadAll(resp.Body)
		preview := string(bodyPreview)
		if len(preview) > 200 {
			preview = preview[:200]
		}
		return nil, fmt.Errorf("captcha request failed with status: %s, body preview: %s", resp.Status, preview)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading captcha response: %w", err)
	}

	var captchaResp struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := json.Unmarshal(body, &captchaResp); err != nil {
		return nil, fmt.Errorf("error parsing captcha JSON: %w", err)
	}

	captcha := &models.CaptchaData{
		Key:       captchaResp.Key,
		ImageData: captchaResp.Value,
		CSRFToken: csrfToken,
		// AuthHeader intentionally not set - credentials should not be exposed to clients
	}

	sess.SetCaptcha(captcha)
	return captcha, nil
}

// RefreshCaptcha fetches a new captcha using existing session data.
func (s *Service) RefreshCaptcha(sess *session.Session) (*models.CaptchaData, error) {
	csrfToken := sess.GetCSRFToken()
	if csrfToken == "" {
		return nil, fmt.Errorf("CSRF token required for captcha refresh")
	}

	req, err := http.NewRequest("GET", s.cfg.Services.MyGovAuth.CaptchaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating captcha request: %w", err)
	}

	s.client.SetAPIHeaders(req, csrfToken)
	req.Header.Set("Referer", s.cfg.Services.MyGovAuth.LoginURL)
	req.Header.Set("Authorization", s.client.AuthHeader())
	s.client.AddCookies(req, sess.GetCookies())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making captcha request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("captcha request failed with status: %s", resp.Status)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading captcha response: %w", err)
	}

	var captchaResp struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := json.Unmarshal(body, &captchaResp); err != nil {
		return nil, fmt.Errorf("error parsing captcha JSON: %w", err)
	}

	captcha := &models.CaptchaData{
		Key:       captchaResp.Key,
		ImageData: captchaResp.Value,
		CSRFToken: csrfToken,
		// AuthHeader intentionally not set - credentials should not be exposed to clients
	}

	sess.SetCaptcha(captcha)
	return captcha, nil
}

// extractCSRFToken extracts CSRF token from HTML content.
func extractCSRFToken(html string) string {
	patterns := []string{
		`<meta name="_csrf" content="([^"]+)"`,
		`<input[^>]*name="_csrf"[^>]*value="([^"]+)"`,
		`"_csrf":"([^"]+)"`,
		`_csrf["\s:]+["']([^"']+)["']`,
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(html)
		if len(matches) > 1 && matches[1] != "" {
			return matches[1]
		}
	}

	return ""
}
