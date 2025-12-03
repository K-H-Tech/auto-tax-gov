package taxregister

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// Regex patterns for extracting ASP.NET hidden fields.
var (
	viewStatePattern = regexp.MustCompile(
		`<input[^>]*name="__VIEWSTATE"[^>]*value="([^"]*)"`)
	viewStateAltPattern = regexp.MustCompile(
		`<input[^>]*value="([^"]*)"[^>]*name="__VIEWSTATE"`)
	eventValidationPattern = regexp.MustCompile(
		`<input[^>]*name="__EVENTVALIDATION"[^>]*value="([^"]*)"`)
	eventValidationAltPattern = regexp.MustCompile(
		`<input[^>]*value="([^"]*)"[^>]*name="__EVENTVALIDATION"`)
	viewStateGenPattern = regexp.MustCompile(
		`<input[^>]*name="__VIEWSTATEGENERATOR"[^>]*value="([^"]*)"`)
	viewStateGenAltPattern = regexp.MustCompile(
		`<input[^>]*value="([^"]*)"[^>]*name="__VIEWSTATEGENERATOR"`)

	// Pattern to extract GUID from hidden field
	guidPattern = regexp.MustCompile(
		`<input[^>]*name="[^"]*HFGUID"[^>]*value="([^"]*)"`)
	guidAltPattern = regexp.MustCompile(
		`<input[^>]*value="([^"]*)"[^>]*name="[^"]*HFGUID"`)

	// Pattern to extract dropdown options
	selectPattern = regexp.MustCompile(`(?is)<select[^>]*name="([^"]*)"[^>]*>(.*?)</select>`)
	optionPattern = regexp.MustCompile(`(?is)<option[^>]*value="([^"]*)"[^>]*>([^<]*)</option>`)

	// Patterns for parsing AJAX responses
	ajaxViewStatePattern       = regexp.MustCompile(`\|__VIEWSTATE\|([^|]+)\|`)
	ajaxEventValidationPattern = regexp.MustCompile(`\|__EVENTVALIDATION\|([^|]+)\|`)
	ajaxUpdatePanelPattern     = regexp.MustCompile(`\d+\|updatePanel\|([^|]+)\|`)
)

// ParseASPNetForm extracts all ASP.NET form data from HTML.
func ParseASPNetForm(html string) (*PublicDataForm, error) {
	form := &PublicDataForm{
		Fields:          make(map[string]string),
		DropdownOptions: make(map[string][]DropdownOption),
	}

	// Extract __VIEWSTATE
	if match := viewStatePattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewState = match[1]
	} else if match := viewStateAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewState = match[1]
	}

	if form.ViewState == "" {
		return nil, fmt.Errorf("__VIEWSTATE not found in HTML")
	}

	// Extract __EVENTVALIDATION
	if match := eventValidationPattern.FindStringSubmatch(html); len(match) > 1 {
		form.EventValidation = match[1]
	} else if match := eventValidationAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.EventValidation = match[1]
	}

	if form.EventValidation == "" {
		return nil, fmt.Errorf("__EVENTVALIDATION not found in HTML")
	}

	// Extract __VIEWSTATEGENERATOR (optional)
	if match := viewStateGenPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewStateGenerator = match[1]
	} else if match := viewStateGenAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewStateGenerator = match[1]
	}

	// Extract HFGUID
	if match := guidPattern.FindStringSubmatch(html); len(match) > 1 {
		form.GUID = match[1]
	} else if match := guidAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.GUID = match[1]
	}

	// Extract all dropdown options
	selectMatches := selectPattern.FindAllStringSubmatch(html, -1)
	for _, match := range selectMatches {
		if len(match) >= 3 {
			selectName := match[1]
			selectContent := match[2]

			var options []DropdownOption
			optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
			for _, optMatch := range optMatches {
				if len(optMatch) >= 3 {
					options = append(options, DropdownOption{
						Value: optMatch[1],
						Label: strings.TrimSpace(optMatch[2]),
					})
				}
			}
			if len(options) > 0 {
				form.DropdownOptions[selectName] = options
			}
		}
	}

	return form, nil
}

// BuildFormPayload builds URL-encoded form data for full form POST submission.
func BuildFormPayload(form *PublicDataForm, fields map[string]string) url.Values {
	payload := url.Values{}

	// ASP.NET hidden fields
	payload.Set("__VIEWSTATE", form.ViewState)
	if form.ViewStateGenerator != "" {
		payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	}
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")

	// Custom form fields
	for key, value := range fields {
		payload.Set(key, value)
	}

	return payload
}

// BuildAjaxPayload builds URL-encoded form data for AJAX partial postback.
// This follows the pattern from step_08.raw with ScriptManager (SM2) and UpdatePanel.
func BuildAjaxPayload(form *PublicDataForm, eventTarget, updatePanel string, fields map[string]string) url.Values {
	payload := url.Values{}

	// ScriptManager identifier for partial postback
	// Format: SM2=UpdatePanelID|ControlID
	payload.Set("ctl00$SM2", fmt.Sprintf("%s|%s", updatePanel, eventTarget))

	// ASP.NET hidden fields
	payload.Set("__VIEWSTATE", form.ViewState)
	if form.ViewStateGenerator != "" {
		payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	}
	payload.Set("__EVENTVALIDATION", form.EventValidation)

	// For AJAX postback, set the control that triggered the event
	payload.Set("__EVENTTARGET", eventTarget)
	payload.Set("__EVENTARGUMENT", "")
	payload.Set("__LASTFOCUS", "")

	// Custom form fields
	for key, value := range fields {
		payload.Set(key, value)
	}

	return payload
}

// ParseAjaxResponse parses an ASP.NET AJAX partial postback response.
// The response format uses pipe-delimited sections with updated form state.
func ParseAjaxResponse(response string) (*AjaxPostbackResponse, error) {
	result := &AjaxPostbackResponse{
		RawResponse:      response,
		UpdatedDropdowns: make(map[string][]DropdownOption),
	}

	// Extract updated __VIEWSTATE
	if match := ajaxViewStatePattern.FindStringSubmatch(response); len(match) > 1 {
		result.ViewState = match[1]
	}

	// Extract updated __EVENTVALIDATION
	if match := ajaxEventValidationPattern.FindStringSubmatch(response); len(match) > 1 {
		result.EventValidation = match[1]
	}

	// Extract UpdatePanel content and parse for dropdown options
	panelMatches := ajaxUpdatePanelPattern.FindAllStringSubmatch(response, -1)
	for _, match := range panelMatches {
		if len(match) >= 2 {
			panelID := match[1]
			// Find the HTML content for this panel
			// Format: length|updatePanel|panelID|HTML content|...
			contentPattern := regexp.MustCompile(fmt.Sprintf(`\d+\|updatePanel\|%s\|([^|]+)\|`, regexp.QuoteMeta(panelID)))
			if contentMatch := contentPattern.FindStringSubmatch(response); len(contentMatch) > 1 {
				// Parse dropdowns from the HTML content
				htmlContent := contentMatch[1]
				selectMatches := selectPattern.FindAllStringSubmatch(htmlContent, -1)
				for _, selMatch := range selectMatches {
					if len(selMatch) >= 3 {
						selectName := selMatch[1]
						selectContent := selMatch[2]

						var options []DropdownOption
						optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
						for _, optMatch := range optMatches {
							if len(optMatch) >= 3 {
								options = append(options, DropdownOption{
									Value: optMatch[1],
									Label: strings.TrimSpace(optMatch[2]),
								})
							}
						}
						if len(options) > 0 {
							result.UpdatedDropdowns[selectName] = options
						}
					}
				}
			}
		}
	}

	// If we couldn't parse VIEWSTATE from AJAX format, try standard HTML format
	// (sometimes the server returns full HTML instead of AJAX delta)
	if result.ViewState == "" {
		if match := viewStatePattern.FindStringSubmatch(response); len(match) > 1 {
			result.ViewState = match[1]
		} else if match := viewStateAltPattern.FindStringSubmatch(response); len(match) > 1 {
			result.ViewState = match[1]
		}
	}

	if result.EventValidation == "" {
		if match := eventValidationPattern.FindStringSubmatch(response); len(match) > 1 {
			result.EventValidation = match[1]
		} else if match := eventValidationAltPattern.FindStringSubmatch(response); len(match) > 1 {
			result.EventValidation = match[1]
		}
	}

	return result, nil
}

// UpdateFormState updates the form state with values from an AJAX response.
func UpdateFormState(form *PublicDataForm, ajaxResp *AjaxPostbackResponse) {
	if ajaxResp.ViewState != "" {
		form.ViewState = ajaxResp.ViewState
	}
	if ajaxResp.EventValidation != "" {
		form.EventValidation = ajaxResp.EventValidation
	}

	// Update dropdown options
	for name, options := range ajaxResp.UpdatedDropdowns {
		form.DropdownOptions[name] = options
	}
}

// ExtractDropdownOptions extracts options from a specific dropdown by name.
func ExtractDropdownOptions(html, selectName string) []DropdownOption {
	var options []DropdownOption

	matches := selectPattern.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		if len(match) >= 3 && match[1] == selectName {
			selectContent := match[2]

			optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
			for _, optMatch := range optMatches {
				if len(optMatch) >= 3 {
					options = append(options, DropdownOption{
						Value: optMatch[1],
						Label: strings.TrimSpace(optMatch[2]),
					})
				}
			}
			break
		}
	}

	return options
}

// ExtractGUIDFromHTML extracts the registration GUID from HTML content.
func ExtractGUIDFromHTML(html string) string {
	if match := guidPattern.FindStringSubmatch(html); len(match) > 1 {
		return match[1]
	}
	if match := guidAltPattern.FindStringSubmatch(html); len(match) > 1 {
		return match[1]
	}
	return ""
}

// ExtractHiddenField extracts a hidden field value from HTML by field name.
func ExtractHiddenField(html, fieldName string) string {
	// Pattern: <input type="hidden" name="fieldName" value="..." />
	pattern := regexp.MustCompile(fmt.Sprintf(`<input[^>]*name="%s"[^>]*value="([^"]*)"`, regexp.QuoteMeta(fieldName)))
	if match := pattern.FindStringSubmatch(html); len(match) > 1 {
		return match[1]
	}

	// Alternative pattern: value before name
	altPattern := regexp.MustCompile(fmt.Sprintf(`<input[^>]*value="([^"]*)"[^>]*name="%s"`, regexp.QuoteMeta(fieldName)))
	if match := altPattern.FindStringSubmatch(html); len(match) > 1 {
		return match[1]
	}

	return ""
}

// ParseDropdownOptions extracts options from a specific dropdown by partial ID match.
func ParseDropdownOptions(html, partialID string) []DropdownOption {
	var options []DropdownOption

	// Find select elements containing the partial ID
	selectRegex := regexp.MustCompile(fmt.Sprintf(`(?is)<select[^>]*id="[^"]*%s[^"]*"[^>]*>(.*?)</select>`, regexp.QuoteMeta(partialID)))
	matches := selectRegex.FindAllStringSubmatch(html, -1)

	for _, match := range matches {
		if len(match) >= 2 {
			selectContent := match[1]

			optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
			for _, optMatch := range optMatches {
				if len(optMatch) >= 3 {
					value := strings.TrimSpace(optMatch[1])
					label := strings.TrimSpace(optMatch[2])

					// Skip empty/default options
					if value == "" || value == "-1" {
						continue
					}

					options = append(options, DropdownOption{
						Value: value,
						Label: label,
					})
				}
			}
			break // Use first matching select
		}
	}

	return options
}

// ParseMembersForm extracts all ASP.NET form data from MembersEdit page HTML.
func ParseMembersForm(html string) (*MembersFormData, error) {
	form := &MembersFormData{
		Fields:          make(map[string]string),
		DropdownOptions: make(map[string][]DropdownOption),
	}

	// Extract __VIEWSTATE
	if match := viewStatePattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewState = match[1]
	} else if match := viewStateAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewState = match[1]
	}

	if form.ViewState == "" {
		return nil, fmt.Errorf("__VIEWSTATE not found in HTML")
	}

	// Extract __EVENTVALIDATION
	if match := eventValidationPattern.FindStringSubmatch(html); len(match) > 1 {
		form.EventValidation = match[1]
	} else if match := eventValidationAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.EventValidation = match[1]
	}

	if form.EventValidation == "" {
		return nil, fmt.Errorf("__EVENTVALIDATION not found in HTML")
	}

	// Extract __VIEWSTATEGENERATOR (optional)
	if match := viewStateGenPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewStateGenerator = match[1]
	} else if match := viewStateGenAltPattern.FindStringSubmatch(html); len(match) > 1 {
		form.ViewStateGenerator = match[1]
	}

	// Extract all dropdown options
	selectMatches := selectPattern.FindAllStringSubmatch(html, -1)
	for _, match := range selectMatches {
		if len(match) >= 3 {
			selectName := match[1]
			selectContent := match[2]

			var options []DropdownOption
			optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
			for _, optMatch := range optMatches {
				if len(optMatch) >= 3 {
					options = append(options, DropdownOption{
						Value: optMatch[1],
						Label: strings.TrimSpace(optMatch[2]),
					})
				}
			}
			if len(options) > 0 {
				form.DropdownOptions[selectName] = options
			}
		}
	}

	// Extract existing field values from input elements
	inputPattern := regexp.MustCompile(`<input[^>]*name="([^"]*)"[^>]*value="([^"]*)"[^>]*>`)
	inputAltPattern := regexp.MustCompile(`<input[^>]*value="([^"]*)"[^>]*name="([^"]*)"[^>]*>`)

	inputMatches := inputPattern.FindAllStringSubmatch(html, -1)
	for _, match := range inputMatches {
		if len(match) >= 3 {
			form.Fields[match[1]] = match[2]
		}
	}
	inputAltMatches := inputAltPattern.FindAllStringSubmatch(html, -1)
	for _, match := range inputAltMatches {
		if len(match) >= 3 {
			form.Fields[match[2]] = match[1]
		}
	}

	return form, nil
}

// BuildMemberFormPayload builds URL-encoded form data for member submission.
func BuildMemberFormPayload(form *MembersFormData, req *MemberSubmitRequest) url.Values {
	payload := url.Values{}

	// ASP.NET hidden fields
	payload.Set("__VIEWSTATE", form.ViewState)
	if form.ViewStateGenerator != "" {
		payload.Set("__VIEWSTATEGENERATOR", form.ViewStateGenerator)
	}
	payload.Set("__EVENTVALIDATION", form.EventValidation)
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")

	// Identity fields (اطلاعات هویتی)
	payload.Set("ctl00$CPC$DDLMemberType", req.PersonType)
	payload.Set("ctl00$CPC$DDLMemberNationality", req.Nationality)
	payload.Set("ctl00$CPC$TextBoxMemberNationalID", req.NationalID)
	payload.Set("ctl00$CPC$TextBoxMemberBirthdate", req.BirthDate)
	if req.BirthCountry != "" {
		payload.Set("ctl00$CPC$DDLMemberCountryOfBorn", req.BirthCountry)
	}
	if req.NationalCardType != "" {
		payload.Set("ctl00$CPC$DDLMemberNationalCardType", req.NationalCardType)
	}
	if req.NationalCardSerial != "" {
		payload.Set("ctl00$CPC$TextboxMemberNationalCardSerial", req.NationalCardSerial)
	}

	// Financial fields (اطلاعات مالی)
	payload.Set("ctl00$CPC$DDLMembershipType", req.MembershipType)
	if req.IsResponsible != "" {
		payload.Set("ctl00$CPC$DDLMemberResponsible", req.IsResponsible)
	}
	if req.SignatureAuthority != "" {
		payload.Set("ctl00$CPC$DDLMemberRightSignFinancial", req.SignatureAuthority)
	}
	payload.Set("ctl00$CPC$DDLMemberRespondibilityType", req.ResponsibilityType)
	payload.Set("ctl00$CPC$TextBoxMemberShares", req.SharePercent)
	if req.Position != "" {
		payload.Set("ctl00$CPC$DDLMemberPosition", req.Position)
	}
	payload.Set("ctl00$CPC$TextBoxMemberStartDate", req.StartDate)
	if req.EndDate != "" {
		payload.Set("ctl00$CPC$TextBoxMemberEndDate", req.EndDate)
	}
	if req.LicenseNumber != "" {
		payload.Set("ctl00$CPC$TextBoxMemberLicenseNumber", req.LicenseNumber)
	}
	if req.SpouseNationalID != "" {
		payload.Set("ctl00$CPC$TextBoxMemberHusbandWifeFidaCode", req.SpouseNationalID)
	}
	if req.SpouseBirthDate != "" {
		payload.Set("ctl00$CPC$TextBoxMemberHusbandWifeBirthdate", req.SpouseBirthDate)
	}

	// Contact fields (اطلاعات تماس)
	payload.Set("ctl00$CPC$TextBoxMemberPostalCode", req.PostalCode)
	if req.Address != "" {
		payload.Set("ctl00$CPC$TextBoxMemberAddress", req.Address)
	}
	if req.Phone != "" {
		payload.Set("ctl00$CPC$TextBoxMemberTel", req.Phone)
	}
	if req.AreaCode != "" {
		payload.Set("ctl00$CPC$TextBoxMemberTelCode", req.AreaCode)
	}
	payload.Set("ctl00$CPC$TextBoxMemberMobile", req.Mobile)
	if req.Email != "" {
		payload.Set("ctl00$CPC$TextBoxMemberEmail", req.Email)
	}

	// Submit button
	payload.Set("ctl00$CPC$ButtonSave", "ثبت")

	return payload
}
