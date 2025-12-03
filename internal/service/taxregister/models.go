package taxregister

// RegistrationRequest contains data for creating a new tax registration.
type RegistrationRequest struct {
	Type         string `json:"type"`         // "Multiple" or "Single"
	PostalCode   string `json:"postalCode"`   // 10-digit postal code
	BusinessName string `json:"businessName"` // URL-encoded Persian business name
}

// RegistrationResponse is the response from NewRegistration API.
type RegistrationResponse struct {
	Success bool   `json:"isSuccess"`
	Message string `json:"msg"` // Contains GUID on success
	GUID    string `json:"guid,omitempty"`
}

// SSOResponse is the response from SSODoc API.
type SSOResponse struct {
	IsLogin bool   `json:"isLogin"`
	URL     string `json:"url"` // Full URL to TokenLoginProcessWithSignout
}

// HomePageData contains parsed data from the HomePage.
type HomePageData struct {
	GUID          string            `json:"guid"`
	Status        string            `json:"status"`
	StatusMessage string            `json:"statusMessage"`
	Sections      []string          `json:"sections"` // List of completed sections
	RawHTML       string            `json:"-"`        // Raw HTML for debugging
	FormData      map[string]string `json:"formData"` // Any hidden form fields
}

// PublicDataForm contains the ASP.NET form state and fields from PublicData page.
type PublicDataForm struct {
	// ASP.NET hidden fields
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Registration GUID
	GUID string `json:"guid"`

	// Form field values (current state)
	Fields map[string]string `json:"fields"`

	// Available dropdown options
	DropdownOptions map[string][]DropdownOption `json:"dropdownOptions"`
}

// DropdownOption represents a single option in a dropdown.
type DropdownOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// PublicDataRequest contains form data for submitting the PublicData form.
type PublicDataRequest struct {
	// Required fields from step_09.raw
	RegistrationCause   string `json:"registrationCause"`   // DDLNewRegistrationCause
	IsTejari            string `json:"isTejari"`            // DDLIsTejari
	FinancialStartDate  string `json:"financialStartDate"`  // TextBoxFinantialStartDate (Jalali)
	BusinessName        string `json:"businessName"`        // TextBoxPDName
	GroupOneType        string `json:"groupOneType"`        // DDLGroupOneTypes
	LegalType           string `json:"legalType"`           // DDLPDLegalType
	NewLegalGroup       string `json:"newLegalGroup"`       // DDLPDNewLegalGroup
	NewLegalType        string `json:"newLegalType"`        // DDLPDNewLegalType
	HasJobLicense       string `json:"hasJobLicense"`       // DDLHasJobLicence
	Ownership           string `json:"ownership"`           // DDLPDOwnership
	FinancialDayStart   string `json:"financialDayStart"`   // DDLFinantialDayStart
	FinancialMonthStart string `json:"financialMonthStart"` // DDLFinantialMonthStart

	// Optional fields
	Website string `json:"website,omitempty"` // TextBoxAddressWebsite
	Email   string `json:"email,omitempty"`   // TextBoxAddressEmail
}

// AjaxPostbackRequest contains data for an AJAX partial postback.
type AjaxPostbackRequest struct {
	EventTarget string            // The control that triggered the postback (e.g., "ctl00$CPC$DDLPDNewLegalGroup")
	UpdatePanel string            // The update panel to refresh (e.g., "ctl00$CPC$UPGuildGroup")
	FormState   *PublicDataForm   // Current form state
	FieldValues map[string]string // Current field values to include
}

// AjaxPostbackResponse contains the result of an AJAX partial postback.
type AjaxPostbackResponse struct {
	// Updated ASP.NET state
	ViewState       string `json:"viewState"`
	EventValidation string `json:"eventValidation"`

	// Updated dropdown options (if any changed)
	UpdatedDropdowns map[string][]DropdownOption `json:"updatedDropdowns,omitempty"`

	// Raw response for debugging
	RawResponse string `json:"-"`
}

// RegistrationFlowRequest contains all data needed for the complete 11-step flow.
type RegistrationFlowRequest struct {
	// Step 1: New Registration
	Type         string `json:"type"`         // "Multiple" or "Single"
	PostalCode   string `json:"postalCode"`   // 10-digit postal code
	BusinessName string `json:"businessName"` // Business name

	// Steps 9-10: PublicData form submission
	PublicData *PublicDataRequest `json:"publicData,omitempty"`
}

// RegistrationFlowResponse contains the result of the complete registration flow.
type RegistrationFlowResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	// Registration details
	GUID         string `json:"guid,omitempty"`
	Status       string `json:"status,omitempty"`
	FinalPageURL string `json:"finalPageUrl,omitempty"`

	// Step-by-step results
	Steps []StepResult `json:"steps,omitempty"`
}

// StepResult contains the result of a single step in the flow.
type StepResult struct {
	Step    int    `json:"step"`
	Name    string `json:"name"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	URL     string `json:"url,omitempty"`
}

// MembersFormData contains the ASP.NET form state from MembersEdit page.
type MembersFormData struct {
	// ASP.NET hidden fields
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Member ID (GUID) - empty for new member
	MemberID string `json:"memberId,omitempty"`

	// Current field values (if editing existing member)
	Fields map[string]string `json:"fields,omitempty"`

	// Available dropdown options
	DropdownOptions map[string][]DropdownOption `json:"dropdownOptions,omitempty"`
}

// MemberSubmitRequest contains form data for submitting a member.
type MemberSubmitRequest struct {
	// ASP.NET state from MembersFormData
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Identity fields (اطلاعات هویتی)
	PersonType         string `json:"personType"`                   // DDLMemberType: 1=حقیقی, 2=حقوقی
	Nationality        string `json:"nationality"`                  // DDLMemberNationality: 33=ایران
	NationalID         string `json:"nationalId"`                   // TextBoxMemberNationalID
	BirthDate          string `json:"birthDate"`                    // TextBoxMemberBirthdate (Jalali)
	BirthCountry       string `json:"birthCountry,omitempty"`       // DDLMemberCountryOfBorn
	NationalCardType   string `json:"nationalCardType,omitempty"`   // DDLMemberNationalCardType: 1=قدیم, 2=هوشمند
	NationalCardSerial string `json:"nationalCardSerial,omitempty"` // TextboxMemberNationalCardSerial

	// Financial fields (اطلاعات مالی)
	MembershipType     string `json:"membershipType"`               // DDLMembershipType: 0=اختیاری, 1=قهری
	IsResponsible      string `json:"isResponsible,omitempty"`      // DDLMemberResponsible: 0=خیر, 1=بله
	SignatureAuthority string `json:"signatureAuthority,omitempty"` // DDLMemberRightSignFinancial: 0=ندارد, 1=دارد
	ResponsibilityType string `json:"responsibilityType"`           // DDLMemberRespondibilityType: 0-5
	SharePercent       string `json:"sharePercent"`                 // TextBoxMemberShares (percentage)
	Position           string `json:"position,omitempty"`           // DDLMemberPosition: 1,7,8,9
	StartDate          string `json:"startDate"`                    // TextBoxMemberStartDate (Jalali)
	EndDate            string `json:"endDate,omitempty"`            // TextBoxMemberEndDate: 0 for ongoing
	LicenseNumber      string `json:"licenseNumber,omitempty"`      // TextBoxMemberLicenseNumber
	SpouseNationalID   string `json:"spouseNationalId,omitempty"`   // TextBoxMemberHusbandWifeFidaCode
	SpouseBirthDate    string `json:"spouseBirthDate,omitempty"`    // TextBoxMemberHusbandWifeBirthdate

	// Contact fields (اطلاعات تماس)
	PostalCode string `json:"postalCode"`         // TextBoxMemberPostalCode
	Address    string `json:"address,omitempty"`  // TextBoxMemberAddress
	Phone      string `json:"phone,omitempty"`    // TextBoxMemberTel
	AreaCode   string `json:"areaCode,omitempty"` // TextBoxMemberTelCode
	Mobile     string `json:"mobile"`             // TextBoxMemberMobile
	Email      string `json:"email,omitempty"`    // TextBoxMemberEmail
}

// MemberSubmitResponse is the response from member submission.
type MemberSubmitResponse struct {
	Success  bool   `json:"success"`
	MemberID string `json:"memberId,omitempty"`
	Message  string `json:"message,omitempty"`
}

// MemberInfo represents a member in the list view (from MembersEdit page).
type MemberInfo struct {
	ID           string `json:"id"`           // Member GUID
	PersonType   string `json:"personType"`   // حقیقی/حقوقی
	Name         string `json:"name"`         // نام و نام خانوادگی
	NationalID   string `json:"nationalId"`   // شماره ملی/شناسه ملی
	Position     string `json:"position"`     // سمت
	SharePercent int    `json:"sharePercent"` // درصد سهام
	Status       string `json:"status"`       // وضعیت
}

// BankAccountInfo represents a bank account in the list view (from AddShebaNumber page).
type BankAccountInfo struct {
	ID        string `json:"id"`        // Account row ID
	IBAN      string `json:"iban"`      // شماره شبا
	StartDate string `json:"startDate"` // تاریخ شروع استفاده
	Status    string `json:"status"`    // وضعیت
}

// INTASearchResult represents a search result from the INTA code search.
type INTASearchResult struct {
	Code       string `json:"code"`       // e.g., "3190130"
	FullPath   string `json:"fullPath"`   // e.g., "خدمات/فعالیت های سینمایی.../جلوه های ویژه"
	EntityType string `json:"entityType"` // e.g., "حقیقی/حقوقی"
}

// INTALevelOption represents an option in a cascade dropdown level.
type INTALevelOption struct {
	Level int    `json:"level"`
	Value string `json:"value"`
	Label string `json:"label"`
}

// INTAActivity represents a complete INTA activity with cascade level selections.
type INTAActivity struct {
	Levels      []INTALevelOption `json:"levels"`      // Selected options at each level
	Code        string            `json:"code"`        // Final INTA code
	Description string            `json:"description"` // شرح فعالیت
	Percent     int               `json:"percent"`     // درصد فعالیت
}

// INTAFormData contains the ASP.NET form state from ActivityINTACode page.
type INTAFormData struct {
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Available dropdown options at each level
	Level1Options []DropdownOption `json:"level1Options,omitempty"`
	Level2Options []DropdownOption `json:"level2Options,omitempty"`
	Level3Options []DropdownOption `json:"level3Options,omitempty"`
	Level4Options []DropdownOption `json:"level4Options,omitempty"`

	// Current activities list
	Activities []INTAActivity `json:"activities,omitempty"`
}

// INTASubmitRequest contains data for submitting INTA activities.
type INTASubmitRequest struct {
	Activities []INTAActivity `json:"activities"`
}

// VATStatusFormData contains the ASP.NET form state from VATStatus page.
type VATStatusFormData struct {
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Current VAT status selection
	EligibilityType string `json:"eligibilityType,omitempty"`
}

// VATStatusRequest contains data for submitting VAT status.
type VATStatusRequest struct {
	EligibilityType string `json:"eligibilityType"` // عدم مشمولیت / مشمول
}

// ShebaFormData contains the ASP.NET form state from AddShebaNumber page.
type ShebaFormData struct {
	ViewState          string `json:"viewState"`
	ViewStateGenerator string `json:"viewStateGenerator"`
	EventValidation    string `json:"eventValidation"`

	// Existing SHEBA accounts
	Accounts []BankAccountInfo `json:"accounts,omitempty"`
}

// ShebaSubmitRequest contains data for submitting a SHEBA number.
type ShebaSubmitRequest struct {
	IBAN      string `json:"iban"`      // شماره شبا (24 رقم بدون IR)
	StartDate string `json:"startDate"` // تاریخ شروع استفاده (Jalali)
}

// PartnerInput represents partner data from the frontend.
type PartnerInput struct {
	NationalID         string `json:"nationalId"`                   // کد ملی (10 رقم)
	SharePercent       int    `json:"sharePercent"`                 // درصد سهم
	Role               string `json:"role,omitempty"`               // مدیر/شریک
	BirthDate          string `json:"birthDate"`                    // تاریخ تولد (1370/01/01)
	NationalCardSerial string `json:"nationalCardSerial,omitempty"` // سریال کارت ملی
	PostalCode         string `json:"postalCode"`                   // کد پستی (10 رقم)
	Mobile             string `json:"mobile"`                       // شماره موبایل (11 رقم)
}

// CompleteRegistrationRequest is the unified request for automated registration.
// User provides only essential PII data; all dropdown/selective values are from config.
type CompleteRegistrationRequest struct {
	// Essential user inputs (PII)
	PostalCode       string `json:"postalCode"`       // کد پستی (10 رقم)
	BusinessName     string `json:"businessName"`     // عنوان واحد/شهرت کسبی
	RegistrationType string `json:"registrationType"` // "individual" or "partnership"
	ShebaNumber      string `json:"shebaNumber"`      // شماره شبا (24 رقم بدون IR)

	// Optional: Partners for partnership type
	Partners []PartnerInput `json:"partners,omitempty"`
}

// CompleteRegistrationResponse is the response from automated registration.
type CompleteRegistrationResponse struct {
	Success      bool         `json:"success"`
	TrackingCode string       `json:"trackingCode,omitempty"` // کد رهگیری
	GUID         string       `json:"guid,omitempty"`         // شناسه ثبت‌نام
	Message      string       `json:"message,omitempty"`
	Steps        []StepResult `json:"steps,omitempty"`
}
