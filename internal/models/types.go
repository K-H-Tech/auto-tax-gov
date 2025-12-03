package models

import (
	"net/http"
	"time"
)

// RedirectStep represents a single step in the redirect chain.
type RedirectStep struct {
	StepNumber       int
	URL              string
	Method           string
	RequestHeaders   http.Header
	RequestBody      string
	StatusCode       int
	Status           string
	ResponseHeaders  http.Header
	ResponseBody     string
	RedirectLocation string
	Timestamp        time.Time
	CaptchaData      *CaptchaData
}

// CaptchaData holds captcha information from the SSO login page.
// WARNING: AuthHeader contains sensitive credentials - never expose in logs or API responses.
type CaptchaData struct {
	Key        string `json:"key"`
	ImageData  string `json:"imageData"`
	CSRFToken  string `json:"csrfToken"`
	AuthHeader string `json:"-"` // Excluded from JSON to prevent credential exposure
}

// LoginResult represents the result of a login operation.
type LoginResult struct {
	Message string
	Data    map[string]any
}

// OTPRequest represents the request to send an OTP.
type OTPRequest struct {
	Mobile      string `json:"mobile"`
	CaptchaCode string `json:"captchaCode"`
	CaptchaKey  string `json:"captchaKey"`
	CSRFToken   string `json:"csrfToken"`
}

// OTPVerifyRequest represents the request to verify an OTP.
type OTPVerifyRequest struct {
	Mobile    string `json:"mobile"`
	OTPCode   string `json:"otpCode"`
	CSRFToken string `json:"csrfToken"`
}

// APIResponse is a generic API response structure.
type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// CaptchaResponse is the response for captcha-related endpoints.
type CaptchaResponse struct {
	Success bool `json:"success"`
	Captcha *struct {
		Key       string `json:"key"`
		ImageData string `json:"imageData"`
		CSRFToken string `json:"csrfToken"`
	} `json:"captcha,omitempty"`
	Error string `json:"error,omitempty"`
}

// TaxRegistrationRequest represents the tax file registration request.
type TaxRegistrationRequest struct {
	PostalCode   string `json:"postalCode"`
	BusinessName string `json:"businessName"`
	Type         string `json:"type"` // "Single" or "Company"
}

// DropdownOption represents a single option in a dropdown.
type DropdownOption struct {
	Value   string `json:"value" yaml:"value"`
	Label   string `json:"label" yaml:"label"`
	LabelFa string `json:"labelFa" yaml:"label_fa"`
}

// BasicInfoRequest represents Step 2 basic information form submission.
type BasicInfoRequest struct {
	RegistrationReason   string `json:"registrationReason"`   // دلیل ثبت‌نام
	ActivityType         string `json:"activityType"`         // نوع فعالیت
	StartDate            string `json:"startDate"`            // تاریخ شروع فعالیت
	UnitTitle            string `json:"unitTitle"`            // عنوان واحد
	EightCategoryJob     string `json:"eightCategoryJob"`     // مشاغل هشتگانه
	IndividualJob        string `json:"individualJob"`        // مشاغل انفرادی
	ProfessionalGuild    string `json:"professionalGuild"`    // مجامع صنفی
	ProfessionalAssembly string `json:"professionalAssembly"` // مجامع حرفه‌ای/تشکل های صنفی (searchable dropdown)
	GuildUnion           string `json:"guildUnion"`           // اتحادیه صنفی ذیربط (searchable dropdown)
	NewGuildUnion        string `json:"newGuildUnion"`        // اتحادیه صنفی جدید (cascading dropdown)
	BusinessLicense      string `json:"businessLicense"`      // پروانه کسب
	OwnershipType        string `json:"ownershipType"`        // نوع مالکیت
	Website              string `json:"website,omitempty"`    // پایگاه اینترنتی (optional)
	Mobile               string `json:"mobile,omitempty"`     // شماره موبایل (optional)
	Email                string `json:"email,omitempty"`      // ایمیل (optional)
}

// Partner represents a business partner (Step 3 - شرکا و اعضا).
// Fields map to register.tax.gov.ir/Pages/Preaction/MembersEdit ASP.NET form.
// The basic fields (NationalID, FullName, SharePercent, Role) are sent from the frontend.
// Other fields have defaults applied server-side if not provided.
type Partner struct {
	// Basic fields from frontend (required)
	NationalID   string `json:"nationalId"`   // کد ملی - 10 digits (TextBoxMemberNationalID)
	FullName     string `json:"fullName"`     // نام و نام خانوادگی (for display, not used in ASP.NET form)
	SharePercent int    `json:"sharePercent"` // درصد سهم (TextBoxMemberShares) - sent as int from frontend
	Role         string `json:"role"`         // نقش from frontend: مدیر، شریک (maps to Position)

	// Identity fields (optional - defaults applied server-side)
	PersonType         string `json:"personType,omitempty"`         // نوع شخص: 1=حقیقی/Real, 2=حقوقی/Legal (DDLMemberType)
	Nationality        string `json:"nationality,omitempty"`        // ملیت: 33=Iran (DDLMemberNationality)
	BirthDate          string `json:"birthDate,omitempty"`          // تاریخ تولد - Jalali: 1370/01/01 (TextBoxMemberBirthdate)
	BirthCountry       string `json:"birthCountry,omitempty"`       // کشور محل تولد: 33=Iran (DDLMemberCountryOfBorn)
	IdentityNumber     string `json:"identityNumber,omitempty"`     // شماره گذرنامه - for non-Iranians (TextBoxMemberIdentityNumber)
	NationalCardType   string `json:"nationalCardType,omitempty"`   // نوع کارت ملی: 1=old, 2=new smart (DDLMemberNationalCardType)
	NationalCardSerial string `json:"nationalCardSerial,omitempty"` // سریال کارت ملی (TextboxMemberNationalCardSerial)

	// Financial/membership fields (optional - defaults applied server-side)
	MembershipType     string `json:"membershipType,omitempty"`     // نوع عضویت: 0=اختیاری/Optional, 1=قهری/Forced (DDLMembershipType)
	IsResponsible      string `json:"isResponsible,omitempty"`      // مسئول: 1=Yes, 0=No (DDLMemberResponsible)
	SignatureAuthority string `json:"signatureAuthority,omitempty"` // حق امضاء مالی: 0=No, 1=Yes (DDLMemberRightSignFinancial)
	ResponsibilityType string `json:"responsibilityType,omitempty"` // نوع مسئولیت: 0-5 (DDLMemberRespondibilityType)
	Position           string `json:"position,omitempty"`           // سمت: 1=Representative, 7=Accountant, 8=Partner, 9=Customs Rep (DDLMemberPosition)
	StartDate          string `json:"startDate,omitempty"`          // تاریخ شروع - Jalali (TextBoxMemberStartDate)
	EndDate            string `json:"endDate,omitempty"`            // تاریخ پایان: 0 for ongoing (TextBoxMemberEndDate)
	LicenseNumber      string `json:"licenseNumber,omitempty"`      // شماره مجوز وکالت (TextBoxMemberLicenseNumber)

	// Spouse fields (for partnership cases)
	SpouseNationalID string `json:"spouseNationalId,omitempty"` // شماره ملی همسر در صورت شراکت (TextBoxMemberHusbandWifeFidaCode)
	SpouseBirthDate  string `json:"spouseBirthDate,omitempty"`  // تاریخ تولد همسر در صورت شراکت (TextBoxMemberHusbandWifeBirthdate)

	// Contact fields (optional)
	PostalCode string `json:"postalCode,omitempty"` // کد پستی (TextBoxMemberPostalCode)
	Address    string `json:"address,omitempty"`    // آدرس (TextBoxMemberAddress)
	Phone      string `json:"phone,omitempty"`      // تلفن (TextBoxMemberTel)
	AreaCode   string `json:"areaCode,omitempty"`   // کد تلفن (TextBoxMemberTelCode)
	Mobile     string `json:"mobile,omitempty"`     // موبایل (TextBoxMemberMobile)
	Email      string `json:"email,omitempty"`      // ایمیل (TextBoxMemberEmail)
}

// PartnersRequest represents Step 3 partners form submission.
type PartnersRequest struct {
	RegistrationID string    `json:"registrationId"` // شناسه ثبت‌نام
	Partners       []Partner `json:"partners"`       // لیست شرکا
}

// BankAccount represents a business bank account (Step 4 - حساب‌های بانکی).
type BankAccount struct {
	IBAN      string `json:"iban"`      // شماره شبا (بدون IR - 24 رقم)
	StartDate string `json:"startDate"` // تاریخ شروع استفاده از حساب
}

// BankAccountsRequest represents Step 4 bank accounts form submission.
type BankAccountsRequest struct {
	RegistrationID string        `json:"registrationId"` // شناسه ثبت‌نام
	Accounts       []BankAccount `json:"accounts"`       // لیست حساب‌های بانکی
}

// DeleteRegistrationRequest represents a request to delete a registration.
type DeleteRegistrationRequest struct {
	RegistrationID string `json:"registrationId"` // شناسه ثبت‌نام
}

// IncompleteRegistration represents a tax file registration that is not completed (before Step 4).
type IncompleteRegistration struct {
	UUID         string `json:"uuid"`         // شناسه یکتا (از URL)
	TrackingCode string `json:"trackingCode"` // شماره رهگیری
	Type         string `json:"type"`         // نوع پرونده: حقیقی/حقوقی
	Status       string `json:"status"`       // وضعیت: گام1, گام2, گام3
	BusinessName string `json:"businessName"` // نام کسب‌وکار
	PostalCode   string `json:"postalCode"`   // کد پستی
}

// INTACodeLevel represents a dropdown level selection in the INTA code cascade.
// Level 1 = Main category (تولید/بازرگانی/خدمات)
// Level 2, 3, etc. = Sub-categories loaded dynamically via ASP.NET postback.
type INTACodeLevel struct {
	Level int    `json:"level"`           // Level number (1, 2, 3, etc.)
	Value string `json:"value"`           // Selected option value
	Label string `json:"label,omitempty"` // Display label (optional, for reference)
}

// INTACodeActivity represents a single INTA code activity to submit.
// Each activity requires navigating through the dropdown cascade,
// then submitting with a description and percentage.
type INTACodeActivity struct {
	Levels      []INTACodeLevel `json:"levels"`      // Cascade selections from Level 1 to final
	Description string          `json:"description"` // شرح فعالیت (mandatory)
	Percent     int             `json:"percent"`     // درصد فعالیت (1-100)
}

// INTACodeRequest represents the INTA code activities submission request.
// Supports multiple activities where percentages must total 100%.
type INTACodeRequest struct {
	RegistrationID string             `json:"registrationId"` // شناسه ثبت‌نام
	Activities     []INTACodeActivity `json:"activities"`     // لیست فعالیت‌ها
}

// INTACodeOptionsRequest represents a request to get dropdown options for a level.
type INTACodeOptionsRequest struct {
	RegistrationID string          `json:"registrationId"` // شناسه ثبت‌نام
	Levels         []INTACodeLevel `json:"levels"`         // Previous level selections
}

// INTACodeOptionsResponse returns available options for the next dropdown level.
type INTACodeOptionsResponse struct {
	Level   int              `json:"level"`   // The level number these options belong to
	Options []DropdownOption `json:"options"` // Available options for this level
}
