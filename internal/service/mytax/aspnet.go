package mytax

import (
	"fmt"
	"net/url"
	"regexp"
)

// ASPNetFormData holds ASP.NET WebForms state extracted from HTML.
type ASPNetFormData struct {
	ViewState          string
	ViewStateGenerator string
	EventValidation    string
}

// Regex patterns for extracting ASP.NET hidden fields.
// These patterns handle both value="..." and value='...' formats.
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
)

// ExtractASPNetFormData extracts hidden ASP.NET WebForms fields from HTML.
func ExtractASPNetFormData(html string) (*ASPNetFormData, error) {
	formData := &ASPNetFormData{}

	// Extract __VIEWSTATE
	if match := viewStatePattern.FindStringSubmatch(html); len(match) > 1 {
		formData.ViewState = match[1]
	} else if match := viewStateAltPattern.FindStringSubmatch(html); len(match) > 1 {
		formData.ViewState = match[1]
	}

	if formData.ViewState == "" {
		return nil, fmt.Errorf("__VIEWSTATE not found in HTML")
	}

	// Extract __EVENTVALIDATION
	if match := eventValidationPattern.FindStringSubmatch(html); len(match) > 1 {
		formData.EventValidation = match[1]
	} else if match := eventValidationAltPattern.FindStringSubmatch(html); len(match) > 1 {
		formData.EventValidation = match[1]
	}

	if formData.EventValidation == "" {
		return nil, fmt.Errorf("__EVENTVALIDATION not found in HTML")
	}

	// Extract __VIEWSTATEGENERATOR (optional)
	if match := viewStateGenPattern.FindStringSubmatch(html); len(match) > 1 {
		formData.ViewStateGenerator = match[1]
	} else if match := viewStateGenAltPattern.FindStringSubmatch(html); len(match) > 1 {
		formData.ViewStateGenerator = match[1]
	}

	return formData, nil
}

// BuildASPNetPayload builds a form payload with ASP.NET hidden fields and custom fields.
func BuildASPNetPayload(formData *ASPNetFormData, fields map[string]string) url.Values {
	payload := url.Values{}

	// Add ASP.NET WebForms hidden fields
	payload.Set("__VIEWSTATE", formData.ViewState)
	if formData.ViewStateGenerator != "" {
		payload.Set("__VIEWSTATEGENERATOR", formData.ViewStateGenerator)
	}
	payload.Set("__EVENTVALIDATION", formData.EventValidation)
	payload.Set("__EVENTTARGET", "")
	payload.Set("__EVENTARGUMENT", "")

	// Add custom form fields
	for key, value := range fields {
		payload.Set(key, value)
	}

	return payload
}

// MapBasicInfoToASPNet maps our BasicInfoRequest fields to ASP.NET control IDs.
func MapBasicInfoToASPNet(req map[string]string) map[string]string {
	// Field mapping from our JSON fields to ASP.NET control IDs
	mapping := map[string]string{
		"registrationReason":   "DDLNewRegistrationCause",
		"activityType":         "DDLIsTejari",
		"startDate":            "TextBoxFinantialStartDate",
		"unitTitle":            "TextBoxPDName",
		"eightCategoryJob":     "DDLGroupOneTypes",
		"individualJob":        "DDLGroupTwoTypes",
		"professionalGuild":    "DDLPDNewLegalGroup",
		"professionalAssembly": "DDLPDNewLegalGroup", // Same as professionalGuild
		"guildUnion":           "DDLPDLegalType",
		"businessLicense":      "DDLHasJobLicence",
		"ownershipType":        "DDLPDOwnership",
		"website":              "TextBoxWebsite",
		"mobile":               "TextBoxMobile",
		"email":                "TextBoxEmail",
	}

	result := make(map[string]string)
	for ourField, aspnetField := range mapping {
		if value, ok := req[ourField]; ok && value != "" {
			result[aspnetField] = value
		}
	}

	return result
}

// DropdownOption represents a single option extracted from an HTML <select> element.
type DropdownOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// selectPattern matches <select name="...">...</select> blocks
var selectPattern = regexp.MustCompile(`(?is)<select[^>]*name="([^"]*)"[^>]*>(.*?)</select>`)

// optionPattern matches <option value="...">...</option> elements
var optionPattern = regexp.MustCompile(`(?is)<option[^>]*value="([^"]*)"[^>]*>([^<]*)</option>`)

// ExtractDropdownOptions extracts all options from a <select> element by its name attribute.
// Returns a slice of DropdownOption with value and label for each option.
func ExtractDropdownOptions(html, selectName string) []DropdownOption {
	var options []DropdownOption

	// Find the <select> block with the matching name
	matches := selectPattern.FindAllStringSubmatch(html, -1)
	for _, match := range matches {
		if len(match) >= 3 && match[1] == selectName {
			selectContent := match[2]

			// Extract all <option> elements within this select
			optMatches := optionPattern.FindAllStringSubmatch(selectContent, -1)
			for _, optMatch := range optMatches {
				if len(optMatch) >= 3 {
					options = append(options, DropdownOption{
						Value: optMatch[1],
						Label: optMatch[2],
					})
				}
			}
			break
		}
	}

	return options
}

// BuildASPNetPostbackPayload builds a payload for ASP.NET postback (cascade dropdown navigation).
// Unlike regular form submission, postbacks set __EVENTTARGET to the control triggering the postback.
func BuildASPNetPostbackPayload(formData *ASPNetFormData, eventTarget string, fields map[string]string) url.Values {
	payload := url.Values{}

	// Add ASP.NET WebForms hidden fields
	payload.Set("__VIEWSTATE", formData.ViewState)
	if formData.ViewStateGenerator != "" {
		payload.Set("__VIEWSTATEGENERATOR", formData.ViewStateGenerator)
	}
	payload.Set("__EVENTVALIDATION", formData.EventValidation)

	// For postback, set the control that triggered the event
	payload.Set("__EVENTTARGET", eventTarget)
	payload.Set("__EVENTARGUMENT", "")
	payload.Set("__LASTFOCUS", "")

	// Add custom form fields (current dropdown selections)
	for key, value := range fields {
		payload.Set(key, value)
	}

	return payload
}
