package taxregister

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/K-H-Tech/auto-tax-gov/internal/client"
	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
)

// uuidPattern matches UUID in URLs and responses.
var uuidPattern = regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)

// isErrorRedirect checks if a redirect location points to an error page.
// The portal returns HTTP 302 to /Pages/Error/ when form validation fails.
func isErrorRedirect(location string) bool {
	return strings.Contains(location, "/Pages/Error/") ||
		strings.Contains(location, "/Error/") ||
		strings.Contains(strings.ToLower(location), "error")
}

// Service handles the complete 11-step tax registration flow.
type Service struct {
	cfg    *config.Config
	client *Client
	logger *slog.Logger
}

// New creates a new TaxRegister service instance.
func New(cfg *config.Config, logger *slog.Logger) *Service {
	return &Service{
		cfg:    cfg,
		client: NewClient(cfg),
		logger: logger,
	}
}

// Step 1: NewRegistration creates a new tax registration.
// POST https://my.tax.gov.ir/Page/NewRegistration/
// Returns the registration GUID on success.
func (s *Service) NewRegistration(sess *session.Session, req *RegistrationRequest) (*RegistrationResponse, error) {
	if !sess.IsAuthenticated() {
		return nil, fmt.Errorf("session not authenticated")
	}

	// Build form payload matching step_01.raw
	payload := url.Values{}
	payload.Set("NewRegistrationType", req.Type)
	payload.Set("NewRegistrationPostalCode", req.PostalCode)
	payload.Set("NewRegistrationBusinessName", req.BusinessName)

	httpReq, err := http.NewRequest("POST", s.cfg.Services.MyTax.RegistrationURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating registration request: %w", err)
	}

	// Set headers matching step_01.raw
	s.client.SetCommonHeaders(httpReq)
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	httpReq.Header.Set("X-Requested-With", "XMLHttpRequest")
	httpReq.Header.Set("Origin", s.client.MyTaxOrigin())
	httpReq.Header.Set("Referer", s.cfg.Services.MyTax.DashboardURL)
	httpReq.Header.Set("Sec-Fetch-Site", "same-origin")
	httpReq.Header.Set("Sec-Fetch-Mode", "cors")
	httpReq.Header.Set("Sec-Fetch-Dest", "empty")
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Step 1: Creating new registration",
		"url", s.cfg.Services.MyTax.RegistrationURL,
		"type", req.Type,
		"postalCode", req.PostalCode)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error creating registration: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading registration response: %w", err)
	}

	s.logger.Debug("Step 1 response", "status", resp.StatusCode, "body", string(body))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("registration returned status %d", resp.StatusCode)
	}

	// Parse JSON response
	var result RegistrationResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing registration response: %w", err)
	}

	if !result.Success {
		return &result, fmt.Errorf("registration failed: %s", result.Message)
	}

	// The GUID is in the "msg" field
	result.GUID = result.Message
	sess.SetRegistrationID(result.GUID)

	s.logger.Info("Step 1 complete: Registration created", "guid", result.GUID)

	return &result, nil
}

// RecoverRegistration recovers a deleted registration using the UndoDelete endpoint.
// This is needed when a registration with the same postal code/national ID already exists.
// GET https://my.tax.gov.ir/Page/UndoDelete/{guid}
func (s *Service) RecoverRegistration(sess *session.Session, guid string) (*RegistrationResponse, error) {
	if !sess.IsAuthenticated() {
		return nil, fmt.Errorf("session not authenticated")
	}

	undoURL := fmt.Sprintf("%s/Page/UndoDelete/%s", s.cfg.Services.MyTax.BaseURL, guid)

	s.logger.Info("RecoverRegistration: Recovering deleted registration",
		"guid", guid,
		"url", undoURL)

	req, err := http.NewRequest("GET", undoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating recovery request: %w", err)
	}

	s.client.SetCommonHeaders(req)
	s.client.AddCookies(req, sess.GetCookies())

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing recovery request: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	// Check for redirect (success indicator)
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		s.logger.Info("RecoverRegistration: Redirected after recovery", "location", location)
	}

	// Read response body to check for errors
	body, err := client.ReadResponseBody(resp)
	if err != nil {
		s.logger.Warn("RecoverRegistration: Could not read response body", "error", err)
	}

	// Check if response indicates an error
	if resp.StatusCode != 200 && resp.StatusCode < 300 {
		return nil, fmt.Errorf("recovery request failed with status %d", resp.StatusCode)
	}

	// Check for error content in body
	if strings.Contains(string(body), "خطا") {
		return nil, fmt.Errorf("recovery failed: response contains error")
	}

	sess.SetRegistrationID(guid)

	s.logger.Info("RecoverRegistration: Registration recovered successfully", "guid", guid)

	return &RegistrationResponse{
		Success: true,
		GUID:    guid,
		Message: "Registration recovered successfully",
	}, nil
}

// Step 2: GetSSOUrl fetches the SSO redirect URL for cross-domain authentication.
// GET https://my.tax.gov.ir/Page/SSODoc/{registrationGUID}
// Returns the URL to TokenLoginProcessWithSignout with a different token UUID.
func (s *Service) GetSSOUrl(sess *session.Session, registrationGUID string) (*SSOResponse, error) {
	if !sess.IsAuthenticated() {
		return nil, fmt.Errorf("session not authenticated")
	}

	ssoDocURL := strings.TrimSuffix(s.cfg.Services.MyTax.SSODocURL, "/") + "/" + registrationGUID

	httpReq, err := http.NewRequest("GET", ssoDocURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating SSODoc request: %w", err)
	}

	// Set headers matching step_02.raw
	s.client.SetCommonHeaders(httpReq)
	httpReq.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	httpReq.Header.Set("Content-Type", "application/json;charset=utf-8")
	httpReq.Header.Set("X-Requested-With", "XMLHttpRequest")
	httpReq.Header.Set("Referer", s.cfg.Services.MyTax.DashboardURL)
	httpReq.Header.Set("Sec-Fetch-Site", "same-origin")
	httpReq.Header.Set("Sec-Fetch-Mode", "cors")
	httpReq.Header.Set("Sec-Fetch-Dest", "empty")
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Step 2: Fetching SSO URL", "url", ssoDocURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching SSO URL: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading SSO response: %w", err)
	}

	s.logger.Debug("Step 2 response", "status", resp.StatusCode, "body", string(body))

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("SSODoc returned status %d", resp.StatusCode)
	}

	// Parse JSON response
	var result SSOResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing SSO response: %w", err)
	}

	if !result.IsLogin {
		return nil, fmt.Errorf("SSO login check failed - user not authenticated")
	}

	s.logger.Info("Step 2 complete: Got SSO URL", "url", result.URL)

	return &result, nil
}

// Steps 3-5: AuthenticateToRegister follows the redirect chain for cross-domain authentication.
// This handles the 302 redirect chain from TokenLoginProcessWithSignout to HomePage.
func (s *Service) AuthenticateToRegister(sess *session.Session, ssoURL string) error {
	if !sess.IsAuthenticated() {
		return fmt.Errorf("session not authenticated")
	}

	s.logger.Info("Steps 3-5: Starting cross-domain auth redirect chain", "url", ssoURL)

	currentURL := ssoURL
	maxRedirects := 10
	referer := s.cfg.Services.MyTax.BaseURL + "/"

	for i := 0; i < maxRedirects; i++ {
		httpReq, err := http.NewRequest("GET", currentURL, nil)
		if err != nil {
			return fmt.Errorf("error creating request at step %d: %w", i+3, err)
		}

		// Set navigation headers matching step_03/04/05.raw
		s.client.SetNavigationHeaders(httpReq, referer)
		s.client.AddCookies(httpReq, sess.GetCookies())

		// Log cookies being sent (first step only)
		if i == 0 {
			s.logger.Debug("Cookies being sent", "header", httpReq.Header.Get("Cookie"))
		}

		resp, err := s.client.Do(httpReq)
		if err != nil {
			return fmt.Errorf("error during redirect chain at step %d: %w", i+3, err)
		}

		// Save cookies from response - critical for ASP.NET_SessionId and PreregisterTAX.*
		if respCookies := resp.Cookies(); len(respCookies) > 0 {
			sess.MergeCookies(respCookies)
			s.logger.Debug("Saved cookies from redirect",
				"step", i+3,
				"count", len(respCookies))
		}

		_, _ = client.ReadResponseBody(resp) // Drain response body
		resp.Body.Close()

		redirectLocation := resp.Header.Get("Location")
		s.logger.Info("Redirect step",
			"step", i+3,
			"url", currentURL,
			"status", resp.StatusCode,
			"redirectTo", redirectLocation)

		if resp.StatusCode == 200 {
			// We've reached a final destination
			if strings.Contains(currentURL, "/Pages/Preaction/") {
				s.logger.Info("Steps 3-5 complete: Reached Preaction page", "finalURL", currentURL)
				return nil
			}

			// Check for login page (auth failure)
			if strings.Contains(currentURL, "/Pages/Login") && !strings.Contains(currentURL, "TokenLogin") {
				return fmt.Errorf("authentication failed - redirected to login page")
			}

			s.logger.Info("Steps 3-5 complete: Reached final destination", "finalURL", currentURL)
			return nil
		}

		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location := resp.Header.Get("Location")
			if location == "" {
				return fmt.Errorf("redirect without Location header at step %d", i+3)
			}

			// Update referer for next request
			referer = currentURL

			// Handle relative URLs
			if strings.HasPrefix(location, "/") {
				parsedURL, err := url.Parse(currentURL)
				if err == nil {
					currentURL = parsedURL.Scheme + "://" + parsedURL.Host + location
				} else {
					currentURL = s.cfg.Services.RegisterTax.BaseURL + location
				}
			} else {
				currentURL = location
			}
			continue
		}

		return fmt.Errorf("unexpected status %d at step %d", resp.StatusCode, i+3)
	}

	return fmt.Errorf("too many redirects (max %d)", maxRedirects)
}

// Step 6 & 11: GetHomePage fetches the registration HomePage and parses its content.
// GET https://register.tax.gov.ir/Pages/Preaction/HomePage
func (s *Service) GetHomePage(sess *session.Session) (*HomePageData, error) {
	httpReq, err := http.NewRequest("GET", s.cfg.Services.RegisterTax.HomePageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HomePage request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.BaseURL+"/")
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Step 6/11: Fetching HomePage", "url", s.cfg.Services.RegisterTax.HomePageURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching HomePage: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("HomePage redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HomePage returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading HomePage: %w", err)
	}

	html := string(body)

	result := &HomePageData{
		RawHTML:  html,
		FormData: make(map[string]string),
	}

	// Extract GUID from hidden field
	result.GUID = ExtractGUIDFromHTML(html)

	// Parse status from HTML (basic extraction)
	if strings.Contains(html, "گام1") {
		result.Status = "step1"
		result.StatusMessage = "مرحله اول"
	} else if strings.Contains(html, "گام2") {
		result.Status = "step2"
		result.StatusMessage = "مرحله دوم"
	} else if strings.Contains(html, "گام3") {
		result.Status = "step3"
		result.StatusMessage = "مرحله سوم"
	} else if strings.Contains(html, "تکمیل") || strings.Contains(html, "completed") {
		result.Status = "completed"
		result.StatusMessage = "تکمیل شده"
	}

	s.logger.Info("Step 6/11 complete: HomePage loaded",
		"guid", result.GUID,
		"status", result.Status,
		"bodyLen", len(html))

	return result, nil
}

// Step 7: GetPublicDataForm fetches the PublicData form and extracts ASP.NET state.
// GET https://register.tax.gov.ir/Pages/Preaction/PublicData
func (s *Service) GetPublicDataForm(sess *session.Session) (*PublicDataForm, error) {
	httpReq, err := http.NewRequest("GET", s.cfg.Services.RegisterTax.PublicDataURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating PublicData request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.HomePageURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Step 7: Fetching PublicData form", "url", s.cfg.Services.RegisterTax.PublicDataURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching PublicData form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("PublicData redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("PublicData returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading PublicData form: %w", err)
	}

	html := string(body)

	// Parse ASP.NET form
	form, err := ParseASPNetForm(html)
	if err != nil {
		return nil, fmt.Errorf("error parsing ASP.NET form: %w", err)
	}

	s.logger.Info("Step 7 complete: PublicData form loaded",
		"guid", form.GUID,
		"viewStateLen", len(form.ViewState),
		"dropdownCount", len(form.DropdownOptions))

	return form, nil
}

// Step 8: UpdateDropdown performs an AJAX partial postback to cascade dropdown values.
// POST https://register.tax.gov.ir/Pages/Preaction/PublicData with X-MicrosoftAjax: Delta=true
func (s *Service) UpdateDropdown(sess *session.Session, form *PublicDataForm, eventTarget, updatePanel string, fieldValues map[string]string) (*PublicDataForm, error) {
	payload := BuildAjaxPayload(form, eventTarget, updatePanel, fieldValues)

	httpReq, err := http.NewRequest("POST", s.cfg.Services.RegisterTax.PublicDataURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating AJAX request: %w", err)
	}

	s.client.SetAjaxHeaders(httpReq, s.cfg.Services.RegisterTax.PublicDataURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Step 8: AJAX dropdown update",
		"eventTarget", eventTarget,
		"updatePanel", updatePanel)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error performing AJAX request: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading AJAX response: %w", err)
	}

	s.logger.Debug("Step 8 AJAX response", "status", resp.StatusCode, "bodyLen", len(body))

	// Parse AJAX response and update form state
	ajaxResp, err := ParseAjaxResponse(string(body))
	if err != nil {
		return nil, fmt.Errorf("error parsing AJAX response: %w", err)
	}

	// Update form with new state
	UpdateFormState(form, ajaxResp)

	s.logger.Info("Step 8 complete: Dropdown updated",
		"updatedDropdowns", len(ajaxResp.UpdatedDropdowns))

	return form, nil
}

// Steps 9-10: SubmitPublicData submits the complete PublicData form.
// POST https://register.tax.gov.ir/Pages/Preaction/PublicData
func (s *Service) SubmitPublicData(sess *session.Session, form *PublicDataForm, data *PublicDataRequest) error {
	// Build form fields matching step_09.raw
	fields := map[string]string{
		// Required fields
		"ctl00$CPC$DDLNewRegistrationCause":   data.RegistrationCause,
		"ctl00$CPC$DDLIsTejari":               data.IsTejari,
		"ctl00$CPC$TextBoxFinantialStartDate": data.FinancialStartDate,
		"ctl00$CPC$TextBoxPDName":             data.BusinessName,
		"ctl00$CPC$DDLGroupOneTypes":          data.GroupOneType,
		"ctl00$CPC$DDLEnferadiTypes":          HappyPathDefaults.EnferadiTypes, // "1000" = سایر (غیر انفرادی)
		"ctl00$CPC$DDLPDLegalType":            data.LegalType,
		"ctl00$CPC$DDLPDNewLegalGroup":        data.NewLegalGroup,
		"ctl00$CPC$DDLPDNewLegalType":         data.NewLegalType,
		"ctl00$CPC$DDLHasJobLicence":          data.HasJobLicense,
		"ctl00$CPC$DDLPDOwnership":            data.Ownership,
		"ctl00$CPC$DDLFinantialDayStart":      data.FinancialDayStart,
		"ctl00$CPC$DDLFinantialMonthStart":    data.FinancialMonthStart,

		// Financial audit fields (mandatory for non-commercial activities)
		"ctl00$CPC$DDLFinantialSoratMali":    "2", // خیر
		"ctl00$CPC$DDLFinantialGozareshMali": "2", // خیر

		// Hidden fields
		"ctl00$CPC$HFGUID": form.GUID,

		// Submit button
		"ctl00$CPC$ButtonPRSave": "ثبت",
	}

	// Optional fields
	if data.Website != "" {
		fields["ctl00$CPC$TextBoxAddressWebsite"] = data.Website
	}
	if data.Email != "" {
		fields["ctl00$CPC$TextBoxAddressEmail"] = data.Email
	}

	payload := BuildFormPayload(form, fields)

	httpReq, err := http.NewRequest("POST", s.cfg.Services.RegisterTax.PublicDataURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return fmt.Errorf("error creating form submit request: %w", err)
	}

	s.client.SetFormSubmitHeaders(httpReq, s.cfg.Services.RegisterTax.PublicDataURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("Steps 9-10: Submitting PublicData form",
		"url", s.cfg.Services.RegisterTax.PublicDataURL,
		"guid", form.GUID)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error submitting form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return fmt.Errorf("error reading form response: %w", err)
	}

	s.logger.Debug("Steps 9-10 response", "status", resp.StatusCode, "bodyLen", len(body))

	responseHTML := string(body)

	// Check for success/error indicators
	if resp.StatusCode == 200 {
		if strings.Contains(responseHTML, "خطا") && !strings.Contains(responseHTML, "بدون خطا") {
			s.logger.Warn("Form submission may have errors", "preview", truncateString(responseHTML, 500))
			return fmt.Errorf("form submission returned errors")
		}

		s.logger.Info("Steps 9-10 complete: Form submitted successfully")
		return nil
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		s.logger.Info("Steps 9-10: Form submission redirected", "location", location)
		// Check if redirect is to error page
		if isErrorRedirect(location) {
			return fmt.Errorf("form submission failed - redirected to error page: %s", location)
		}
		return nil
	}

	return fmt.Errorf("form submission returned status %d", resp.StatusCode)
}

// ExecuteFullFlow executes the complete 11-step registration flow.
func (s *Service) ExecuteFullFlow(sess *session.Session, req *RegistrationFlowRequest) (*RegistrationFlowResponse, error) {
	response := &RegistrationFlowResponse{
		Steps: make([]StepResult, 0, 11),
	}

	// Step 1: Create new registration
	regReq := &RegistrationRequest{
		Type:         req.Type,
		PostalCode:   req.PostalCode,
		BusinessName: req.BusinessName,
	}

	regResp, err := s.NewRegistration(sess, regReq)
	if err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 1, Name: "NewRegistration", Success: false, Message: err.Error()})
		return response, err
	}
	response.Steps = append(response.Steps, StepResult{Step: 1, Name: "NewRegistration", Success: true, URL: s.cfg.Services.MyTax.RegistrationURL})
	response.GUID = regResp.GUID

	// Step 2: Get SSO URL
	ssoResp, err := s.GetSSOUrl(sess, regResp.GUID)
	if err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 2, Name: "GetSSOUrl", Success: false, Message: err.Error()})
		return response, err
	}
	response.Steps = append(response.Steps, StepResult{Step: 2, Name: "GetSSOUrl", Success: true, URL: ssoResp.URL})

	// Steps 3-5: Authenticate to register.tax.gov.ir
	if err := s.AuthenticateToRegister(sess, ssoResp.URL); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 3, Name: "AuthenticateToRegister", Success: false, Message: err.Error()})
		return response, err
	}
	response.Steps = append(response.Steps, StepResult{Step: 3, Name: "AuthenticateToRegister", Success: true, Message: "Redirect chain completed"})

	// Step 6: Get HomePage
	homeData, err := s.GetHomePage(sess)
	if err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 6, Name: "GetHomePage", Success: false, Message: err.Error()})
		return response, err
	}
	response.Steps = append(response.Steps, StepResult{Step: 6, Name: "GetHomePage", Success: true, URL: s.cfg.Services.RegisterTax.HomePageURL})

	// Update GUID from HomePage if available
	if homeData.GUID != "" {
		response.GUID = homeData.GUID
	}

	// Step 7: Get PublicData form
	form, err := s.GetPublicDataForm(sess)
	if err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 7, Name: "GetPublicDataForm", Success: false, Message: err.Error()})
		return response, err
	}
	response.Steps = append(response.Steps, StepResult{Step: 7, Name: "GetPublicDataForm", Success: true, URL: s.cfg.Services.RegisterTax.PublicDataURL})

	// Steps 8-10: Submit form if data provided
	if req.PublicData != nil {
		// Step 8: AJAX dropdown updates (if needed)
		// This would be called multiple times for cascade dropdowns
		// For now, we skip AJAX if all values are pre-selected
		response.Steps = append(response.Steps, StepResult{Step: 8, Name: "UpdateDropdown", Success: true, Message: "Skipped (values pre-selected)"})

		// Steps 9-10: Submit form
		if err := s.SubmitPublicData(sess, form, req.PublicData); err != nil {
			response.Steps = append(response.Steps, StepResult{Step: 9, Name: "SubmitPublicData", Success: false, Message: err.Error()})
			return response, err
		}
		response.Steps = append(response.Steps, StepResult{Step: 9, Name: "SubmitPublicData", Success: true, URL: s.cfg.Services.RegisterTax.PublicDataURL})

		// Step 11: Verify on HomePage
		finalHomeData, err := s.GetHomePage(sess)
		if err != nil {
			response.Steps = append(response.Steps, StepResult{Step: 11, Name: "FinalHomePage", Success: false, Message: err.Error()})
			return response, err
		}
		response.Steps = append(response.Steps, StepResult{Step: 11, Name: "FinalHomePage", Success: true, URL: s.cfg.Services.RegisterTax.HomePageURL})
		response.Status = finalHomeData.Status
		response.FinalPageURL = s.cfg.Services.RegisterTax.HomePageURL
	}

	response.Success = true
	response.Message = "Registration flow completed successfully"

	return response, nil
}

// GetConfig returns the service configuration.
func (s *Service) GetConfig() *config.Config {
	return s.cfg
}

// GetMembersForm fetches the MembersEdit form and extracts ASP.NET state.
// GET https://register.tax.gov.ir/Pages/Preaction/MembersEdit or MembersEdit/{memberID}
func (s *Service) GetMembersForm(sess *session.Session, memberID string) (*MembersFormData, error) {
	membersURL := s.cfg.Services.RegisterTax.MembersEditURL
	if memberID != "" {
		membersURL = strings.TrimSuffix(membersURL, "/") + "/" + memberID
	}

	httpReq, err := http.NewRequest("GET", membersURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating MembersEdit request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.HomePageURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("GetMembersForm: Fetching members edit form",
		"url", membersURL,
		"memberID", memberID)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching MembersEdit form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("MembersEdit redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("MembersEdit returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading MembersEdit form: %w", err)
	}

	html := string(body)

	// Parse ASP.NET form
	form, err := ParseMembersForm(html)
	if err != nil {
		return nil, fmt.Errorf("error parsing ASP.NET form: %w", err)
	}

	// Set member ID if provided
	form.MemberID = memberID

	s.logger.Info("GetMembersForm complete: MembersEdit form loaded",
		"memberID", memberID,
		"viewStateLen", len(form.ViewState),
		"dropdownCount", len(form.DropdownOptions),
		"fieldCount", len(form.Fields))

	return form, nil
}

// SubmitMember submits member data to the MembersEdit form.
// POST https://register.tax.gov.ir/Pages/Preaction/MembersEdit or MembersEdit/{memberID}
func (s *Service) SubmitMember(sess *session.Session, form *MembersFormData, req *MemberSubmitRequest) (*MemberSubmitResponse, error) {
	membersURL := s.cfg.Services.RegisterTax.MembersEditURL
	if form.MemberID != "" {
		membersURL = strings.TrimSuffix(membersURL, "/") + "/" + form.MemberID
	}

	// Build form payload
	payload := BuildMemberFormPayload(form, req)

	httpReq, err := http.NewRequest("POST", membersURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating member submit request: %w", err)
	}

	s.client.SetFormSubmitHeaders(httpReq, membersURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("SubmitMember: Submitting member form",
		"url", membersURL,
		"memberID", form.MemberID,
		"nationalID", req.NationalID)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error submitting member form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading member response: %w", err)
	}

	s.logger.Debug("SubmitMember response", "status", resp.StatusCode, "bodyLen", len(body))

	responseHTML := string(body)
	result := &MemberSubmitResponse{}

	// Check for success/error indicators
	if resp.StatusCode == 200 {
		// Check for error messages in response
		if strings.Contains(responseHTML, "خطا") && !strings.Contains(responseHTML, "بدون خطا") {
			s.logger.Warn("Member submission may have errors", "preview", truncateString(responseHTML, 500))
			result.Success = false
			result.Message = "Form submission returned errors"
			return result, fmt.Errorf("member form submission returned errors")
		}

		// Try to extract member ID from response if new member
		if form.MemberID == "" {
			// Look for GUID in response
			if match := uuidPattern.FindString(responseHTML); match != "" {
				result.MemberID = match
			}
		} else {
			result.MemberID = form.MemberID
		}

		result.Success = true
		result.Message = "Member submitted successfully"
		s.logger.Info("SubmitMember complete: Member submitted", "memberID", result.MemberID)
		return result, nil
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		s.logger.Info("SubmitMember: Redirected after submission", "location", location)
		// Check if redirect is to error page
		if isErrorRedirect(location) {
			result.Success = false
			result.Message = fmt.Sprintf("Member submission failed - redirected to error page: %s", location)
			return result, fmt.Errorf("member submission failed - redirected to error page: %s", location)
		}
		result.Success = true
		result.Message = "Member submission redirected"
		return result, nil
	}

	result.Success = false
	result.Message = fmt.Sprintf("member form submission returned status %d", resp.StatusCode)
	return result, fmt.Errorf("%s", result.Message)
}

// truncateString truncates a string to maxLen characters.
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// ==================== Member List/Delete Methods ====================

// ListMembers fetches all members for a registration from the MembersEdit page.
// It parses the HTML table to extract member information.
func (s *Service) ListMembers(sess *session.Session, registrationID string) ([]MemberInfo, error) {
	s.logger.Info("Fetching members list", "registrationId", registrationID)

	// Navigate to the HomePage first to get the members table
	homePageURL := s.cfg.Services.RegisterTax.HomePageURL
	if homePageURL == "" {
		homePageURL = "https://register.tax.gov.ir/Pages/Preaction/HomePage"
	}

	req, err := http.NewRequest(http.MethodGet, homePageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	s.client.SetCommonHeaders(req)

	// Add session cookies
	for _, cookie := range sess.GetCookies() {
		req.AddCookie(cookie)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch home page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse members from the HTML response
	members := parseMembersTable(string(body))

	s.logger.Info("Members list fetched", "count", len(members))
	return members, nil
}

// parseMembersTable extracts member information from the HomePage HTML table.
func parseMembersTable(html string) []MemberInfo {
	var members []MemberInfo

	// Look for member rows in the table
	// Pattern: <tr> with member data
	memberRowRegex := regexp.MustCompile(`(?s)<tr[^>]*>.*?<td[^>]*>([^<]*)</td>.*?<td[^>]*>([^<]*)</td>.*?<td[^>]*>([^<]*)</td>.*?<td[^>]*>([^<]*)</td>.*?<td[^>]*>([^<]*)</td>.*?</tr>`)
	matches := memberRowRegex.FindAllStringSubmatch(html, -1)

	for i, match := range matches {
		if len(match) >= 5 {
			// Extract share percentage as int
			shareStr := strings.TrimSpace(match[4])
			shareStr = strings.ReplaceAll(shareStr, "%", "")
			shareStr = strings.ReplaceAll(shareStr, "٪", "")
			share := 0
			fmt.Sscanf(shareStr, "%d", &share)

			member := MemberInfo{
				ID:           fmt.Sprintf("member-%d", i),
				PersonType:   strings.TrimSpace(match[1]),
				Name:         strings.TrimSpace(match[2]),
				NationalID:   strings.TrimSpace(match[3]),
				SharePercent: share,
				Position:     strings.TrimSpace(match[5]),
				Status:       "فعال",
			}
			members = append(members, member)
		}
	}

	return members
}

// DeleteMember removes a member from a registration.
// It navigates to the MembersEdit page and triggers the delete postback.
func (s *Service) DeleteMember(sess *session.Session, registrationID, memberID string) error {
	s.logger.Info("Deleting member", "registrationId", registrationID, "memberId", memberID)

	// Note: This is a placeholder implementation.
	// The actual implementation would need to:
	// 1. Navigate to the member's edit page
	// 2. Extract ASP.NET state
	// 3. Submit a delete postback

	// For now, return a not implemented error
	return fmt.Errorf("delete member not fully implemented yet - memberID: %s", memberID)
}

// ==================== Bank Account (SHEBA) List/Delete Methods ====================

// GetShebaList fetches all bank accounts for a registration from the AddShebaNumber page.
func (s *Service) GetShebaList(sess *session.Session, registrationID string) ([]BankAccountInfo, error) {
	s.logger.Info("Fetching bank accounts list", "registrationId", registrationID)

	// Navigate to the HomePage to get the bank accounts table
	homePageURL := s.cfg.Services.RegisterTax.HomePageURL
	if homePageURL == "" {
		homePageURL = "https://register.tax.gov.ir/Pages/Preaction/HomePage"
	}

	req, err := http.NewRequest(http.MethodGet, homePageURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	s.client.SetCommonHeaders(req)

	// Add session cookies
	for _, cookie := range sess.GetCookies() {
		req.AddCookie(cookie)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch home page: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse bank accounts from the HTML response
	accounts := parseBankAccountsTable(string(body))

	s.logger.Info("Bank accounts list fetched", "count", len(accounts))
	return accounts, nil
}

// parseBankAccountsTable extracts bank account information from the HomePage HTML table.
func parseBankAccountsTable(html string) []BankAccountInfo {
	var accounts []BankAccountInfo

	// Look for SHEBA rows in the bank accounts table
	// Pattern: IR + 24 digits for IBAN
	shebaRegex := regexp.MustCompile(`(?s)<tr[^>]*>.*?<td[^>]*>(IR\d{24}|\d{24})</td>.*?<td[^>]*>([^<]*)</td>.*?</tr>`)
	matches := shebaRegex.FindAllStringSubmatch(html, -1)

	for i, match := range matches {
		if len(match) >= 3 {
			iban := strings.TrimSpace(match[1])
			// Ensure IR prefix
			if !strings.HasPrefix(iban, "IR") {
				iban = "IR" + iban
			}

			account := BankAccountInfo{
				ID:        fmt.Sprintf("sheba-%d", i),
				IBAN:      iban,
				StartDate: strings.TrimSpace(match[2]),
				Status:    "فعال",
			}
			accounts = append(accounts, account)
		}
	}

	return accounts
}

// DeleteSheba removes a bank account from a registration.
// It navigates to the AddShebaNumber page and triggers the delete postback.
func (s *Service) DeleteSheba(sess *session.Session, registrationID, shebaID string) error {
	s.logger.Info("Deleting bank account", "registrationId", registrationID, "shebaId", shebaID)

	// Note: This is a placeholder implementation.
	// The actual implementation would need to:
	// 1. Navigate to the AddShebaNumber page
	// 2. Find the delete link for the specific SHEBA
	// 3. Extract ASP.NET state
	// 4. Submit a delete postback

	// For now, return a not implemented error
	return fmt.Errorf("delete SHEBA not fully implemented yet - shebaID: %s", shebaID)
}

// ==================== INTA Code Methods ====================

// GetINTACodeForm fetches the ActivityINTACode form and extracts ASP.NET state.
// GET https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/
func (s *Service) GetINTACodeForm(sess *session.Session) (*INTAFormData, error) {
	intaURL := s.cfg.Services.RegisterTax.ActivityINTACodeURL
	if intaURL == "" {
		intaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/"
	}

	httpReq, err := http.NewRequest("GET", intaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating INTACode request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.HomePageURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("GetINTACodeForm: Fetching INTA code form", "url", intaURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching INTACode form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("INTACode redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("INTACode returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading INTACode form: %w", err)
	}

	html := string(body)

	// Parse ASP.NET form state
	form := &INTAFormData{}
	form.ViewState = ExtractHiddenField(html, "__VIEWSTATE")
	form.ViewStateGenerator = ExtractHiddenField(html, "__VIEWSTATEGENERATOR")
	form.EventValidation = ExtractHiddenField(html, "__EVENTVALIDATION")

	// Parse Level 1 dropdown options (always available)
	form.Level1Options = ParseDropdownOptions(html, "DDLActivityINTACodeCategory")
	if len(form.Level1Options) == 0 {
		// Fallback: hardcoded Level 1 options based on portal investigation
		form.Level1Options = []DropdownOption{
			{Value: "1", Label: "[1] تولید"},
			{Value: "2", Label: "[2] بازرگانی"},
			{Value: "3", Label: "[3] خدمات"},
		}
	}

	// Parse existing activities from the page
	form.Activities = parseExistingActivities(html)

	s.logger.Info("GetINTACodeForm complete",
		"viewStateLen", len(form.ViewState),
		"level1Options", len(form.Level1Options),
		"existingActivities", len(form.Activities))

	return form, nil
}

// GetINTACodeOptions fetches dropdown options for a specific cascade level.
// level: 1-4 (depth in cascade)
// parentValue: value of parent dropdown (empty for level 1)
// parentLevels: all parent level values for deep cascade
func (s *Service) GetINTACodeOptions(sess *session.Session, level int, parentLevels []string) ([]DropdownOption, error) {
	s.logger.Info("GetINTACodeOptions: Fetching options",
		"level", level,
		"parentLevels", parentLevels)

	// Level 1 is static
	if level == 1 {
		return []DropdownOption{
			{Value: "1", Label: "[1] تولید"},
			{Value: "2", Label: "[2] بازرگانی"},
			{Value: "3", Label: "[3] خدمات"},
		}, nil
	}

	// For levels 2+, we need to do AJAX postback to get options
	// First, get the current form state
	form, err := s.GetINTACodeForm(sess)
	if err != nil {
		return nil, fmt.Errorf("failed to get INTA form: %w", err)
	}

	// Build the event target based on level
	eventTargets := map[int]string{
		2: "ctl00$CPC$DDLActivityINTACodeCategory",
		3: "ctl00$CPC$DDLActivityINTACodeSubCategory1",
		4: "ctl00$CPC$DDLActivityINTACodeSubCategory2",
	}

	eventTarget, ok := eventTargets[level]
	if !ok {
		return nil, fmt.Errorf("invalid level: %d", level)
	}

	// Perform cascade selection up to requested level
	for i := 1; i < level && i <= len(parentLevels); i++ {
		fieldValues := make(map[string]string)

		// Set all parent values
		for j := 0; j < i && j < len(parentLevels); j++ {
			switch j {
			case 0:
				fieldValues["ctl00$CPC$DDLActivityINTACodeCategory"] = parentLevels[j]
			case 1:
				fieldValues["ctl00$CPC$DDLActivityINTACodeSubCategory1"] = parentLevels[j]
			case 2:
				fieldValues["ctl00$CPC$DDLActivityINTACodeSubCategory2"] = parentLevels[j]
			}
		}

		// Make AJAX postback to trigger cascade
		options, err := s.doINTACascadePostback(sess, form, eventTarget, fieldValues, level)
		if err != nil {
			return nil, fmt.Errorf("cascade postback failed at level %d: %w", i+1, err)
		}

		if i == level-1 {
			return options, nil
		}
	}

	return nil, fmt.Errorf("failed to get options for level %d", level)
}

// doINTACascadePostback performs an AJAX postback to get cascade dropdown options.
func (s *Service) doINTACascadePostback(sess *session.Session, form *INTAFormData, eventTarget string, fieldValues map[string]string, targetLevel int) ([]DropdownOption, error) {
	intaURL := s.cfg.Services.RegisterTax.ActivityINTACodeURL
	if intaURL == "" {
		intaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/"
	}

	// Build AJAX payload
	payload := url.Values{}
	payload.Set("ctl00$SMaster", "ctl00$CPC$UPActivityINTACode|"+eventTarget)
	payload.Set("__EVENTTARGET", eventTarget)
	payload.Set("__EVENTARGUMENT", "")
	payload.Set("__VIEWSTATE", form.ViewState)
	payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__ASYNCPOST", "true")

	// Add field values
	for k, v := range fieldValues {
		payload.Set(k, v)
	}

	httpReq, err := http.NewRequest("POST", intaURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating AJAX request: %w", err)
	}

	s.client.SetAjaxHeaders(httpReq, intaURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error performing AJAX request: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading AJAX response: %w", err)
	}

	// Parse the AJAX response to extract dropdown options
	options := parseAjaxDropdownOptions(string(body), targetLevel)

	s.logger.Debug("INTA cascade postback complete",
		"targetLevel", targetLevel,
		"optionsCount", len(options))

	return options, nil
}

// parseAjaxDropdownOptions extracts dropdown options from AJAX response.
func parseAjaxDropdownOptions(ajaxResponse string, targetLevel int) []DropdownOption {
	var options []DropdownOption

	// AJAX responses contain UpdatePanel content with dropdown HTML
	// Look for <option> tags in the response
	optionRegex := regexp.MustCompile(`<option[^>]*value="([^"]*)"[^>]*>([^<]*)</option>`)
	matches := optionRegex.FindAllStringSubmatch(ajaxResponse, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			value := strings.TrimSpace(match[1])
			label := strings.TrimSpace(match[2])

			// Skip empty/default options
			if value == "" || value == "-1" || label == "انتخاب شود..." || label == "انتخاب کنید" {
				continue
			}

			options = append(options, DropdownOption{
				Value: value,
				Label: label,
			})
		}
	}

	return options
}

// parseExistingActivities extracts existing INTA activities from the form HTML.
func parseExistingActivities(html string) []INTAActivity {
	var activities []INTAActivity

	// Look for activity rows in the table
	// The pattern depends on the actual HTML structure
	activityRowRegex := regexp.MustCompile(`(?s)<tr[^>]*data-code="([^"]*)"[^>]*>.*?<td[^>]*>([^<]*)</td>.*?<td[^>]*>(\d+)%?</td>.*?</tr>`)
	matches := activityRowRegex.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		if len(match) >= 4 {
			percent := 0
			fmt.Sscanf(match[3], "%d", &percent)

			activity := INTAActivity{
				Code:        match[1],
				Description: strings.TrimSpace(match[2]),
				Percent:     percent,
			}
			activities = append(activities, activity)
		}
	}

	return activities
}

// SearchINTACodes searches INTA codes by keyword.
func (s *Service) SearchINTACodes(sess *session.Session, keyword string) ([]INTASearchResult, error) {
	s.logger.Info("SearchINTACodes: Searching INTA codes", "keyword", keyword)

	if len(keyword) < 3 {
		return nil, fmt.Errorf("keyword must be at least 3 characters")
	}

	// Get the form state first
	form, err := s.GetINTACodeForm(sess)
	if err != nil {
		return nil, fmt.Errorf("failed to get INTA form: %w", err)
	}

	intaURL := s.cfg.Services.RegisterTax.ActivityINTACodeURL
	if intaURL == "" {
		intaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/"
	}

	// Build search payload
	// ASP.NET Web Forms control naming: ctl00$CPC$TextboxActivityINTACode$TextboxINTASearch (nested user control)
	// ScriptManager format: UpdatePanel ID|Button ID
	payload := url.Values{}
	payload.Set("ctl00$SMaster", "ctl00$CPC$TextboxActivityINTACode$UPINTA|ctl00$CPC$TextboxActivityINTACode$ButtonINTASearch")
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")
	payload.Set("__VIEWSTATE", form.ViewState)
	payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__ASYNCPOST", "true")
	payload.Set("ctl00$CPC$TextboxActivityINTACode$TextboxINTASearch", keyword)
	payload.Set("ctl00$CPC$TextboxActivityINTACode$ButtonINTASearch", "جستجو")

	httpReq, err := http.NewRequest("POST", intaURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error creating search request: %w", err)
	}

	s.client.SetAjaxHeaders(httpReq, intaURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error performing search request: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading search response: %w", err)
	}

	bodyStr := string(body)

	// Debug: log response details
	s.logger.Debug("SearchINTACodes: AJAX response received",
		"responseLength", len(bodyStr),
		"hasUpdatePanel", strings.Contains(bodyStr, "|updatePanel|"),
		"hasLB", strings.Contains(bodyStr, "LB"),
		"hasDoPostBack", strings.Contains(bodyStr, "__doPostBack"),
		"hasTextboxActivityINTACode", strings.Contains(bodyStr, "TextboxActivityINTACode"),
		"first500chars", truncateString(bodyStr, 500))

	// Parse search results from AJAX response
	results := parseINTASearchResults(bodyStr)

	// If no results, log more details to help debug
	if len(results) == 0 {
		s.logger.Warn("SearchINTACodes: No results found, logging response for debugging",
			"keyword", keyword,
			"responseLength", len(bodyStr),
			"hasUpdatePanel", strings.Contains(bodyStr, "|updatePanel|"),
			"hasLB", strings.Contains(bodyStr, "LB"),
			"hasDoPostBack", strings.Contains(bodyStr, "__doPostBack"),
			"first1000chars", truncateString(bodyStr, 1000))
	}

	s.logger.Info("SearchINTACodes complete",
		"keyword", keyword,
		"resultsCount", len(results))

	return results, nil
}

// parseINTASearchResults extracts search results from the AJAX response.
// ASP.NET AJAX partial postback responses use pipe-delimited format:
// length|type|id|content|length|type|id|content|...
// Portal HTML format: <a href="javascript:__doPostBack('ctl00$CPC$TextboxActivityINTACode$LB3210050',”)">• [3210050] [type] description</a>
func parseINTASearchResults(ajaxResponse string) []INTASearchResult {
	var results []INTASearchResult

	// First, extract updatePanel content from AJAX response
	// The response format is: length|updatePanel|panelID|<html content>|...
	htmlContent := extractUpdatePanelContent(ajaxResponse)

	// If no updatePanel found, use the raw response (might be regular HTML)
	if htmlContent == "" {
		htmlContent = ajaxResponse
	}

	// Pattern 1: Match __doPostBack with LB code - most specific pattern
	// Matches: javascript:__doPostBack('ctl00$CPC$TextboxActivityINTACode$LB3210050','')">• [3210050]...
	resultRegex := regexp.MustCompile(`__doPostBack\([^)]*\$LB(\d{7})[^)]*\)[^>]*>([^<]+)</a>`)
	matches := resultRegex.FindAllStringSubmatch(htmlContent, -1)

	for _, match := range matches {
		if len(match) >= 3 {
			// Clean up the description - remove bullet point
			description := strings.TrimSpace(match[2])
			description = strings.TrimPrefix(description, "•")
			description = strings.TrimSpace(description)

			result := INTASearchResult{
				Code:     match[1],
				FullPath: description,
			}
			results = append(results, result)
		}
	}

	// Pattern 2: Alternative - match LB code with any surrounding context
	if len(results) == 0 {
		altRegex := regexp.MustCompile(`LB(\d{7})['"][^>]*>([^<]+)</a>`)
		altMatches := altRegex.FindAllStringSubmatch(htmlContent, -1)

		for _, match := range altMatches {
			if len(match) >= 3 {
				description := strings.TrimSpace(match[2])
				description = strings.TrimPrefix(description, "•")
				description = strings.TrimSpace(description)

				result := INTASearchResult{
					Code:     match[1],
					FullPath: description,
				}
				results = append(results, result)
			}
		}
	}

	// Pattern 3: Even simpler - just look for LB followed by 7 digits and text
	if len(results) == 0 {
		simpleRegex := regexp.MustCompile(`LB(\d{7})[^>]+>([^<]+)</a>`)
		simpleMatches := simpleRegex.FindAllStringSubmatch(htmlContent, -1)

		for _, match := range simpleMatches {
			if len(match) >= 3 {
				description := strings.TrimSpace(match[2])
				description = strings.TrimPrefix(description, "•")
				description = strings.TrimSpace(description)

				result := INTASearchResult{
					Code:     match[1],
					FullPath: description,
				}
				results = append(results, result)
			}
		}
	}

	// Pattern 4: Last resort - look for [code] pattern in text
	if len(results) == 0 {
		lastRegex := regexp.MustCompile(`\[(\d{7})\]\s*(\[[^\]]*\]\s*[^\n<]+)`)
		lastMatches := lastRegex.FindAllStringSubmatch(htmlContent, -1)

		for _, match := range lastMatches {
			if len(match) >= 3 {
				result := INTASearchResult{
					Code:     match[1],
					FullPath: strings.TrimSpace(match[2]),
				}
				results = append(results, result)
			}
		}
	}

	return results
}

// extractUpdatePanelContent extracts HTML content from ASP.NET AJAX partial postback response.
// Format: length|type|id|content|length|type|id|content|...
// We look for updatePanel sections that contain our search results.
func extractUpdatePanelContent(response string) string {
	// Check if this looks like an AJAX response (contains |updatePanel|)
	if len(response) == 0 {
		return ""
	}

	// If not an AJAX response, return empty (caller will use raw response as fallback)
	if !strings.Contains(response, "|updatePanel|") {
		return ""
	}

	var allContent strings.Builder

	// More robust parsing: find updatePanel markers and extract content between them
	// The format is: length|updatePanel|panelID|content|...
	// Content may contain pipe characters, so we use marker-based extraction

	// Split by |updatePanel| to find all update panel sections
	sections := strings.Split(response, "|updatePanel|")

	for i, section := range sections {
		if i == 0 {
			// First section is before any updatePanel, skip it
			continue
		}

		// After |updatePanel|, the format is: panelID|content|nextLength|...
		// Find the panel ID (ends at first |)
		pipeIdx := strings.Index(section, "|")
		if pipeIdx == -1 {
			continue
		}

		// Content starts after panel ID
		remaining := section[pipeIdx+1:]

		// Find where content ends - it's before the next section marker (digit|type|)
		// Look for pattern like: |digit| which starts next section
		endPattern := regexp.MustCompile(`\|\d+\|`)
		endLoc := endPattern.FindStringIndex(remaining)

		var content string
		if endLoc != nil {
			content = remaining[:endLoc[0]]
		} else {
			// Last section - take everything
			content = remaining
		}

		// Only include content that looks like it has search results
		if strings.Contains(content, "__doPostBack") ||
			strings.Contains(content, "LB") ||
			strings.Contains(content, "TextboxActivityINTACode") {
			allContent.WriteString(content)
			allContent.WriteString("\n")
		}
	}

	return allContent.String()
}

// SubmitINTACodes submits INTA activities to the portal.
func (s *Service) SubmitINTACodes(sess *session.Session, activities []INTAActivity) error {
	s.logger.Info("SubmitINTACodes: Submitting activities", "count", len(activities))

	if len(activities) == 0 {
		return fmt.Errorf("at least one activity is required")
	}

	// Validate total percentage
	totalPercent := 0
	for _, a := range activities {
		totalPercent += a.Percent
	}
	if totalPercent != 100 {
		return fmt.Errorf("total percentage must be 100, got %d", totalPercent)
	}

	// Get the form state
	form, err := s.GetINTACodeForm(sess)
	if err != nil {
		return fmt.Errorf("failed to get INTA form: %w", err)
	}

	intaURL := s.cfg.Services.RegisterTax.ActivityINTACodeURL
	if intaURL == "" {
		intaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/"
	}

	// Submit each activity one by one
	for i, activity := range activities {
		s.logger.Debug("Submitting activity",
			"index", i+1,
			"code", activity.Code,
			"percent", activity.Percent)

		// Build payload for adding an activity
		payload := url.Values{}
		payload.Set("__VIEWSTATE", form.ViewState)
		payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
		payload.Set("__EVENTVALIDATION", form.EventValidation)

		// Set the INTA code via postback (simulating click on search result)
		payload.Set("__EVENTTARGET", "ctl00$CPC$TextboxActivityINTACode$LB"+activity.Code)
		payload.Set("__EVENTARGUMENT", "")

		// Set description and percentage
		payload.Set("ctl00$CPC$TextBoxActivityDescription", activity.Description)
		payload.Set("ctl00$CPC$TextBoxActivityPercent", fmt.Sprintf("%d", activity.Percent))

		// Add button
		payload.Set("ctl00$CPC$ButtonActivityINTACodeAdd", "افزودن")

		httpReq, err := http.NewRequest("POST", intaURL, strings.NewReader(payload.Encode()))
		if err != nil {
			return fmt.Errorf("error creating submit request for activity %d: %w", i+1, err)
		}

		s.client.SetFormSubmitHeaders(httpReq, intaURL)
		s.client.AddCookies(httpReq, sess.GetCookies())

		resp, err := s.client.Do(httpReq)
		if err != nil {
			return fmt.Errorf("error submitting activity %d: %w", i+1, err)
		}

		// Save cookies
		if cookies := resp.Cookies(); len(cookies) > 0 {
			sess.MergeCookies(cookies)
		}

		// Check for redirect to error page
		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			location := resp.Header.Get("Location")
			resp.Body.Close()
			if isErrorRedirect(location) {
				return fmt.Errorf("activity %d submission failed - redirected to error page: %s", i+1, location)
			}
		}

		body, err := client.ReadResponseBody(resp)
		resp.Body.Close()

		if err != nil {
			return fmt.Errorf("error reading response for activity %d: %w", i+1, err)
		}

		// Check for errors in response
		responseHTML := string(body)
		if strings.Contains(responseHTML, "خطا") && !strings.Contains(responseHTML, "بدون خطا") {
			return fmt.Errorf("error submitting activity %d: form returned errors", i+1)
		}

		// Update form state for next activity
		form.ViewState = ExtractHiddenField(responseHTML, "__VIEWSTATE")
		form.EventValidation = ExtractHiddenField(responseHTML, "__EVENTVALIDATION")

		s.logger.Debug("Activity submitted successfully", "index", i+1)
	}

	s.logger.Info("SubmitINTACodes complete: All activities submitted")
	return nil
}

// ==================== VAT Status Methods ====================

// GetVATStatusForm fetches the VatStatus form and extracts ASP.NET state.
// GET https://register.tax.gov.ir/Pages/Preaction/Edit/VatStatus/
func (s *Service) GetVATStatusForm(sess *session.Session) (*VATStatusFormData, error) {
	vatURL := s.cfg.Services.RegisterTax.VATStatusURL
	if vatURL == "" {
		vatURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/VatStatus/"
	}

	httpReq, err := http.NewRequest("GET", vatURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating VATStatus request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.HomePageURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("GetVATStatusForm: Fetching VAT status form", "url", vatURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching VATStatus form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("VATStatus redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("VATStatus returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading VATStatus form: %w", err)
	}

	html := string(body)

	// Parse ASP.NET form state
	form := &VATStatusFormData{}
	form.ViewState = ExtractHiddenField(html, "__VIEWSTATE")
	form.ViewStateGenerator = ExtractHiddenField(html, "__VIEWSTATEGENERATOR")
	form.EventValidation = ExtractHiddenField(html, "__EVENTVALIDATION")

	s.logger.Info("GetVATStatusForm complete",
		"viewStateLen", len(form.ViewState))

	return form, nil
}

// SubmitVATStatus submits VAT eligibility status.
// POST https://register.tax.gov.ir/Pages/Preaction/Edit/VatStatus/
func (s *Service) SubmitVATStatus(sess *session.Session, req *VATStatusRequest) error {
	s.logger.Info("SubmitVATStatus: Submitting VAT status", "eligibility", req.EligibilityType)

	// Get the form state first
	form, err := s.GetVATStatusForm(sess)
	if err != nil {
		return fmt.Errorf("failed to get VAT status form: %w", err)
	}

	vatURL := s.cfg.Services.RegisterTax.VATStatusURL
	if vatURL == "" {
		vatURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/VatStatus/"
	}

	// Build form payload
	payload := url.Values{}
	payload.Set("__VIEWSTATE", form.ViewState)
	payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")

	// VAT status dropdown - ctl00$CPC$DDLVatStatus
	payload.Set("ctl00$CPC$DDLVatStatus", req.EligibilityType)

	// Submit button
	payload.Set("ctl00$CPC$ButtonVatStatusSave", "ثبت")

	httpReq, err := http.NewRequest("POST", vatURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return fmt.Errorf("error creating VAT status submit request: %w", err)
	}

	s.client.SetFormSubmitHeaders(httpReq, vatURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error submitting VAT status: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return fmt.Errorf("error reading VAT status response: %w", err)
	}

	responseHTML := string(body)

	// Check for errors
	if resp.StatusCode == 200 {
		if strings.Contains(responseHTML, "خطا") && !strings.Contains(responseHTML, "بدون خطا") {
			s.logger.Warn("VAT status submission may have errors", "preview", truncateString(responseHTML, 500))
			return fmt.Errorf("VAT status submission returned errors")
		}
		s.logger.Info("SubmitVATStatus complete: VAT status submitted successfully")
		return nil
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		s.logger.Info("SubmitVATStatus: Redirected after submission", "location", location)
		// Check if redirect is to error page
		if isErrorRedirect(location) {
			return fmt.Errorf("VAT status submission failed - redirected to error page: %s", location)
		}
		return nil
	}

	return fmt.Errorf("VAT status submission returned status %d", resp.StatusCode)
}

// ==================== SHEBA (Bank Account) Methods ====================

// GetShebaForm fetches the AddShebaNumber form and extracts ASP.NET state.
// GET https://register.tax.gov.ir/Pages/Preaction/Edit/AddShebaNumber/
func (s *Service) GetShebaForm(sess *session.Session) (*ShebaFormData, error) {
	shebaURL := s.cfg.Services.RegisterTax.AddShebaNumberURL
	if shebaURL == "" {
		shebaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/AddShebaNumber/"
	}

	httpReq, err := http.NewRequest("GET", shebaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating SHEBA form request: %w", err)
	}

	s.client.SetNavigationHeaders(httpReq, s.cfg.Services.RegisterTax.HomePageURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	s.logger.Info("GetShebaForm: Fetching SHEBA form", "url", shebaURL)

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error fetching SHEBA form: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	if resp.StatusCode == 302 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, "/Login") {
			return nil, fmt.Errorf("not authenticated - redirected to login page")
		}
		return nil, fmt.Errorf("SHEBA form redirected to %s", location)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("SHEBA form returned status %d", resp.StatusCode)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("error reading SHEBA form: %w", err)
	}

	html := string(body)

	// Parse ASP.NET form state
	form := &ShebaFormData{}
	form.ViewState = ExtractHiddenField(html, "__VIEWSTATE")
	form.ViewStateGenerator = ExtractHiddenField(html, "__VIEWSTATEGENERATOR")
	form.EventValidation = ExtractHiddenField(html, "__EVENTVALIDATION")

	// Parse existing SHEBA accounts
	form.Accounts = parseBankAccountsTable(html)

	s.logger.Info("GetShebaForm complete",
		"viewStateLen", len(form.ViewState),
		"existingAccounts", len(form.Accounts))

	return form, nil
}

// SubmitSheba submits a new SHEBA (bank account) number.
// POST https://register.tax.gov.ir/Pages/Preaction/Edit/AddShebaNumber/
func (s *Service) SubmitSheba(sess *session.Session, req *ShebaSubmitRequest) error {
	s.logger.Info("SubmitSheba: Submitting SHEBA number", "iban", req.IBAN)

	// Clean IBAN - remove IR prefix if present
	iban := strings.TrimPrefix(strings.TrimPrefix(req.IBAN, "IR"), "ir")
	if len(iban) != 24 {
		return fmt.Errorf("invalid SHEBA number: must be 24 digits (got %d)", len(iban))
	}

	// Get the form state first
	form, err := s.GetShebaForm(sess)
	if err != nil {
		return fmt.Errorf("failed to get SHEBA form: %w", err)
	}

	shebaURL := s.cfg.Services.RegisterTax.AddShebaNumberURL
	if shebaURL == "" {
		shebaURL = "https://register.tax.gov.ir/Pages/Preaction/Edit/AddShebaNumber/"
	}

	// Build form payload
	payload := url.Values{}
	payload.Set("__VIEWSTATE", form.ViewState)
	payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")

	// SHEBA fields - verified via Playwright manual testing
	// HTML ID: CPC_TextBoxShebaNumber, ASP.NET name: ctl00$CPC$TextBoxShebaNumber
	payload.Set("ctl00$CPC$TextBoxShebaNumber", iban)
	payload.Set("ctl00$CPC$TextBoxShebaStartDate", req.StartDate)

	// Submit button
	payload.Set("ctl00$CPC$ButtonShebaSave", "ثبت")

	httpReq, err := http.NewRequest("POST", shebaURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return fmt.Errorf("error creating SHEBA submit request: %w", err)
	}

	s.client.SetFormSubmitHeaders(httpReq, shebaURL)
	s.client.AddCookies(httpReq, sess.GetCookies())

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("error submitting SHEBA: %w", err)
	}
	defer resp.Body.Close()

	// Save cookies
	if cookies := resp.Cookies(); len(cookies) > 0 {
		sess.MergeCookies(cookies)
	}

	body, err := client.ReadResponseBody(resp)
	if err != nil {
		return fmt.Errorf("error reading SHEBA response: %w", err)
	}

	responseHTML := string(body)

	// Check for errors
	if resp.StatusCode == 200 {
		if strings.Contains(responseHTML, "خطا") && !strings.Contains(responseHTML, "بدون خطا") {
			// Extract specific error message if present
			if strings.Contains(responseHTML, "الگو شماره شبا نادرست است") {
				return fmt.Errorf("invalid SHEBA format: الگو شماره شبا نادرست است")
			}
			s.logger.Warn("SHEBA submission may have errors", "preview", truncateString(responseHTML, 500))
			return fmt.Errorf("SHEBA submission returned errors")
		}
		s.logger.Info("SubmitSheba complete: SHEBA submitted successfully")
		return nil
	}

	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		s.logger.Info("SubmitSheba: Redirected after submission", "location", location)
		// Check if redirect is to error page
		if isErrorRedirect(location) {
			return fmt.Errorf("SHEBA submission failed - redirected to error page: %s", location)
		}
		return nil
	}

	return fmt.Errorf("SHEBA submission returned status %d", resp.StatusCode)
}

// ==================== Complete Registration Flow ====================

// ExecuteCompleteRegistration executes the entire registration process automatically.
// User provides only essential PII data; all dropdown/selective values use config defaults.
//
// Flow:
//  1. Create registration on my.tax.gov.ir
//  2. Cross-domain auth to register.tax.gov.ir
//  3. Submit PublicData (defaults from config)
//  4. Submit INTA code (defaults from config)
//  5. Submit VAT status (defaults from config)
//  6. Submit SHEBA number (user input)
//  7. [If partnership] Submit members
func (s *Service) ExecuteCompleteRegistration(sess *session.Session, req *CompleteRegistrationRequest) (*CompleteRegistrationResponse, error) {
	response := &CompleteRegistrationResponse{
		Steps: make([]StepResult, 0, 8),
	}

	s.logger.Info("ExecuteCompleteRegistration: Starting automated registration",
		"postalCode", req.PostalCode,
		"businessName", req.BusinessName,
		"type", req.RegistrationType)

	// Determine registration type
	regType := "Single"
	if req.RegistrationType == "partnership" {
		regType = "Multiple"
	}

	// Step 1: Create new registration
	regReq := &RegistrationRequest{
		Type:         regType,
		PostalCode:   req.PostalCode,
		BusinessName: req.BusinessName,
	}

	regResp, err := s.NewRegistration(sess, regReq)
	if err != nil {
		// Check if it's a duplicate registration error with recovery option
		errMsg := err.Error()
		if strings.Contains(errMsg, "قبلا انجام شده است") {
			// Extract recovery GUID from error message (looks for UndoDelete/{guid} or just a UUID)
			var recoveryGUID string
			if matches := uuidPattern.FindStringSubmatch(errMsg); len(matches) > 0 {
				recoveryGUID = matches[0]
			}

			if recoveryGUID != "" {
				s.logger.Info("Duplicate registration detected, attempting auto-recovery",
					"recoveryGUID", recoveryGUID)

				// Attempt recovery
				_, recoverErr := s.RecoverRegistration(sess, recoveryGUID)
				if recoverErr != nil {
					s.logger.Error("Auto-recovery failed", "error", recoverErr)
					response.Steps = append(response.Steps, StepResult{
						Step:    1,
						Name:    "NewRegistration",
						Success: false,
						Message: fmt.Sprintf("Duplicate found but recovery failed: %s", recoverErr.Error()),
					})
					return response, fmt.Errorf("step 1 (NewRegistration) failed: duplicate registration exists, recovery failed: %w", recoverErr)
				}

				// Recovery successful, use the recovered GUID
				s.logger.Info("Auto-recovery successful, continuing with existing registration", "guid", recoveryGUID)
				response.Steps = append(response.Steps, StepResult{
					Step:    1,
					Name:    "NewRegistration",
					Success: true,
					Message: fmt.Sprintf("Recovered existing registration: %s", recoveryGUID),
				})
				response.GUID = recoveryGUID

				// Continue to Step 2 with the recovered GUID
				goto step2
			}
		}

		response.Steps = append(response.Steps, StepResult{Step: 1, Name: "NewRegistration", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 1 (NewRegistration) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 1, Name: "NewRegistration", Success: true, URL: s.cfg.Services.MyTax.RegistrationURL})
	response.GUID = regResp.GUID

step2:

	// Step 2: Get SSO URL
	ssoResp, err := s.GetSSOUrl(sess, response.GUID)
	if err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 2, Name: "GetSSOUrl", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 2 (GetSSOUrl) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 2, Name: "GetSSOUrl", Success: true, URL: ssoResp.URL})

	// Step 3: Authenticate to register.tax.gov.ir
	if err := s.AuthenticateToRegister(sess, ssoResp.URL); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 3, Name: "AuthenticateToRegister", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 3 (AuthenticateToRegister) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 3, Name: "AuthenticateToRegister", Success: true, Message: "Cross-domain auth completed"})

	// Step 4: Submit PublicData with defaults
	if err := s.submitPublicDataWithDefaults(sess, req.BusinessName); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 4, Name: "SubmitPublicData", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 4 (SubmitPublicData) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 4, Name: "SubmitPublicData", Success: true, Message: "Basic info submitted with defaults"})

	// Step 5: Submit INTA code with defaults
	if err := s.submitINTACodeWithDefaults(sess); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 5, Name: "SubmitINTACode", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 5 (SubmitINTACode) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 5, Name: "SubmitINTACode", Success: true, Message: "INTA code submitted with defaults"})

	// Step 6: Submit VAT status with defaults
	if err := s.submitVATStatusWithDefaults(sess); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 6, Name: "SubmitVATStatus", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 6 (SubmitVATStatus) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 6, Name: "SubmitVATStatus", Success: true, Message: "VAT status submitted with defaults"})

	// Step 7: Submit SHEBA number
	shebaReq := &ShebaSubmitRequest{
		IBAN:      req.ShebaNumber,
		StartDate: GetCurrentJalaliDate(),
	}
	if err := s.SubmitSheba(sess, shebaReq); err != nil {
		response.Steps = append(response.Steps, StepResult{Step: 7, Name: "SubmitSheba", Success: false, Message: err.Error()})
		return response, fmt.Errorf("step 7 (SubmitSheba) failed: %w", err)
	}
	response.Steps = append(response.Steps, StepResult{Step: 7, Name: "SubmitSheba", Success: true, Message: "SHEBA submitted successfully"})

	// Step 8: Submit members (only for partnership)
	if req.RegistrationType == "partnership" && len(req.Partners) > 0 {
		if err := s.submitPartnersWithDefaults(sess, req.Partners); err != nil {
			response.Steps = append(response.Steps, StepResult{Step: 8, Name: "SubmitPartners", Success: false, Message: err.Error()})
			return response, fmt.Errorf("step 8 (SubmitPartners) failed: %w", err)
		}
		response.Steps = append(response.Steps, StepResult{Step: 8, Name: "SubmitPartners", Success: true, Message: fmt.Sprintf("%d partners submitted", len(req.Partners))})
	} else {
		response.Steps = append(response.Steps, StepResult{Step: 8, Name: "SubmitPartners", Success: true, Message: "Skipped (individual registration)"})
	}

	// Get final status from HomePage
	homeData, err := s.GetHomePage(sess)
	if err != nil {
		s.logger.Warn("Failed to get final status from HomePage", "error", err)
	} else {
		response.TrackingCode = homeData.GUID
	}

	response.Success = true
	response.Message = "Registration completed successfully"

	s.logger.Info("ExecuteCompleteRegistration: Registration completed",
		"guid", response.GUID,
		"trackingCode", response.TrackingCode)

	return response, nil
}

// submitPublicDataWithDefaults submits the PublicData form with happy path defaults.
// Uses hardcoded values extracted from curl.md that are known to work.
func (s *Service) submitPublicDataWithDefaults(sess *session.Session, businessName string) error {
	// Get the form
	form, err := s.GetPublicDataForm(sess)
	if err != nil {
		return fmt.Errorf("failed to get PublicData form: %w", err)
	}

	// Use happy path defaults (hardcoded values that work)
	publicData := GetHappyPathPublicData(businessName)

	s.logger.Info("Submitting PublicData with happy path defaults",
		"registrationCause", publicData.RegistrationCause,
		"isTejari", publicData.IsTejari,
		"groupOneType", publicData.GroupOneType,
		"legalType", publicData.LegalType,
		"newLegalGroup", publicData.NewLegalGroup,
		"hasJobLicense", publicData.HasJobLicense,
		"ownership", publicData.Ownership)

	return s.SubmitPublicData(sess, form, publicData)
}

// submitINTACodeWithDefaults submits INTA code with happy path defaults.
// Uses hardcoded values extracted from curl.md that are known to work.
func (s *Service) submitINTACodeWithDefaults(sess *session.Session) error {
	// Use happy path defaults (hardcoded values that work)
	activity := GetHappyPathINTACode()

	s.logger.Info("Submitting INTA code with happy path defaults",
		"code", activity.Code,
		"description", activity.Description,
		"percent", activity.Percent)

	// Step 1: Search for the INTA code first to populate the search results
	// This is required because the submit uses LB{code} which only exists after search
	_, err := s.SearchINTACodes(sess, activity.Description)
	if err != nil {
		s.logger.Warn("INTA code search failed, continuing anyway", "error", err)
		// Continue anyway - the code might already be in the form from a previous search
	}

	// Step 2: Submit the activity
	return s.SubmitINTACodes(sess, []INTAActivity{activity})
}

// submitVATStatusWithDefaults submits VAT status with happy path defaults.
// Uses hardcoded values extracted from curl.md that are known to work.
func (s *Service) submitVATStatusWithDefaults(sess *session.Session) error {
	// Use happy path defaults (hardcoded values that work)
	req := GetHappyPathVATStatus()

	s.logger.Info("Submitting VAT status with happy path defaults",
		"eligibilityType", req.EligibilityType)

	return s.SubmitVATStatus(sess, req)
}

// submitPartnersWithDefaults submits partners with config defaults applied.
func (s *Service) submitPartnersWithDefaults(sess *session.Session, partners []PartnerInput) error {
	defaults := s.cfg.Defaults.Member

	for i, partner := range partners {
		s.logger.Info("Submitting partner",
			"index", i+1,
			"nationalId", partner.NationalID,
			"share", partner.SharePercent)

		// Get fresh form for each member
		form, err := s.GetMembersForm(sess, "")
		if err != nil {
			return fmt.Errorf("failed to get members form for partner %d: %w", i+1, err)
		}

		// Map role to position value
		position := defaults.Position
		if partner.Role == "مدیر" {
			position = "1" // Representative/Manager
		}

		memberReq := &MemberSubmitRequest{
			ViewState:          form.ViewState,
			ViewStateGenerator: form.ViewStateGenerator,
			EventValidation:    form.EventValidation,

			// Identity
			PersonType:         defaults.PersonType,
			Nationality:        defaults.Nationality,
			NationalID:         partner.NationalID,
			BirthDate:          partner.BirthDate, // Required: from frontend
			BirthCountry:       defaults.BirthCountry,
			NationalCardType:   defaults.NationalCardType,
			NationalCardSerial: partner.NationalCardSerial, // Required: from frontend

			// Financial
			MembershipType:     defaults.PartnershipType,
			IsResponsible:      defaults.IsEmployed,
			SignatureAuthority: defaults.SignatureAuthority,
			ResponsibilityType: defaults.ResponsibilityType,
			SharePercent:       fmt.Sprintf("%d", partner.SharePercent),
			Position:           position,
			StartDate:          GetCurrentJalaliDate(),
			EndDate:            defaults.EndDate,

			// Contact
			PostalCode: partner.PostalCode, // Required: from frontend
			Mobile:     partner.Mobile,     // Required: from frontend
		}

		_, err = s.SubmitMember(sess, form, memberReq)
		if err != nil {
			return fmt.Errorf("failed to submit partner %d: %w", i+1, err)
		}
	}

	return nil
}
