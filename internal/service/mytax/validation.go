package mytax

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/models"
)

// ValidationError represents a validation error with Persian message.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationResult holds the result of validating basic info request.
type ValidationResult struct {
	Valid  bool              `json:"valid"`
	Errors []ValidationError `json:"errors,omitempty"`
}

// persianDateRegex matches Persian date format: YYYY/MM/DD
var persianDateRegex = regexp.MustCompile(`^1[34]\d{2}/(?:0[1-9]|1[0-2])/(?:0[1-9]|[12]\d|3[01])$`)

// websiteRegex matches basic website URL patterns.
var websiteRegex = regexp.MustCompile(`^(https?://)?[\w\-]+(\.[\w\-]+)+(/[\w\-._~:/?#\[\]@!$&'()*+,;=]*)?$`)

// ValidateBasicInfo validates the BasicInfoRequest with full field validation.
func ValidateBasicInfo(req *models.BasicInfoRequest, opts *config.FormOptionsConfig) *ValidationResult {
	result := &ValidationResult{Valid: true}

	// Nil guard for opts (Issue 23)
	if opts == nil {
		result.addError("_form", "خطای پیکربندی: گزینه‌های فرم بارگذاری نشده‌اند")
		return result
	}

	// Required field validations
	if req.RegistrationReason == "" {
		result.addError("registrationReason", "دلیل ثبت‌نام الزامی است")
	} else if !isValidOption(req.RegistrationReason, opts.RegistrationReasons) {
		result.addError("registrationReason", "دلیل ثبت‌نام انتخاب شده معتبر نیست")
	}

	if req.ActivityType == "" {
		result.addError("activityType", "نوع فعالیت الزامی است")
	} else if !isValidOption(req.ActivityType, opts.ActivityTypes) {
		result.addError("activityType", "نوع فعالیت انتخاب شده معتبر نیست")
	}

	if req.StartDate == "" {
		result.addError("startDate", "تاریخ شروع فعالیت الزامی است")
	} else if !isValidPersianDate(req.StartDate) {
		result.addError("startDate", "فرمت تاریخ صحیح نیست (مثال: 1403/01/01)")
	}

	if req.UnitTitle == "" {
		result.addError("unitTitle", "عنوان واحد الزامی است")
	} else if utf8.RuneCountInString(req.UnitTitle) < 2 {
		// Use rune count for proper UTF-8 character counting (Issue 23)
		result.addError("unitTitle", "عنوان واحد باید حداقل ۲ کاراکتر باشد")
	} else if utf8.RuneCountInString(req.UnitTitle) > 100 {
		result.addError("unitTitle", "عنوان واحد نباید بیش از ۱۰۰ کاراکتر باشد")
	}

	if req.EightCategoryJob == "" {
		result.addError("eightCategoryJob", "مشاغل هشتگانه الزامی است")
	} else if !isValidOption(req.EightCategoryJob, opts.EightCategoryJobs) {
		result.addError("eightCategoryJob", "مشاغل هشتگانه انتخاب شده معتبر نیست")
	}

	if req.IndividualJob == "" {
		result.addError("individualJob", "مشاغل انفرادی الزامی است")
	} else if !isValidOption(req.IndividualJob, opts.IndividualJobs) {
		result.addError("individualJob", "مشاغل انفرادی انتخاب شده معتبر نیست")
	}

	if req.ProfessionalGuild == "" {
		result.addError("professionalGuild", "مجامع صنفی الزامی است")
	} else if !isValidOption(req.ProfessionalGuild, opts.ProfessionalGuilds) {
		result.addError("professionalGuild", "مجامع صنفی انتخاب شده معتبر نیست")
	}

	if req.ProfessionalAssembly == "" {
		result.addError("professionalAssembly", "مجامع صنفی (انتخاب) الزامی است")
	} else if !isValidOption(req.ProfessionalAssembly, opts.ProfessionalAssemblies) {
		result.addError("professionalAssembly", "مجامع صنفی انتخاب شده معتبر نیست")
	}

	if req.GuildUnion == "" {
		result.addError("guildUnion", "اتحادیه صنفی الزامی است")
	} else if !isValidOption(req.GuildUnion, opts.GuildUnions) {
		result.addError("guildUnion", "اتحادیه صنفی انتخاب شده معتبر نیست")
	}

	if req.BusinessLicense == "" {
		result.addError("businessLicense", "پروانه کسب الزامی است")
	} else if !isValidOption(req.BusinessLicense, opts.BusinessLicenses) {
		result.addError("businessLicense", "پروانه کسب انتخاب شده معتبر نیست")
	}

	if req.OwnershipType == "" {
		result.addError("ownershipType", "نوع مالکیت الزامی است")
	} else if !isValidOption(req.OwnershipType, opts.OwnershipTypes) {
		result.addError("ownershipType", "نوع مالکیت انتخاب شده معتبر نیست")
	}

	// Optional field validation
	if req.Website != "" && !isValidWebsite(req.Website) {
		result.addError("website", "آدرس پایگاه اینترنتی معتبر نیست")
	}

	return result
}

// addError adds an error to the validation result.
func (r *ValidationResult) addError(field, message string) {
	r.Valid = false
	r.Errors = append(r.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

// ErrorMessage returns a combined error message for all validation errors.
func (r *ValidationResult) ErrorMessage() string {
	if r.Valid {
		return ""
	}
	var messages []string
	for _, err := range r.Errors {
		messages = append(messages, err.Message)
	}
	return strings.Join(messages, "، ")
}

// FirstError returns the first validation error message.
func (r *ValidationResult) FirstError() string {
	if r.Valid || len(r.Errors) == 0 {
		return ""
	}
	return r.Errors[0].Message
}

// isValidOption checks if a value exists in the dropdown options.
func isValidOption(value string, options []config.DropdownOption) bool {
	for _, opt := range options {
		if opt.Value == value {
			return true
		}
	}
	return false
}

// isValidPersianDate validates Persian date format (YYYY/MM/DD).
func isValidPersianDate(date string) bool {
	return persianDateRegex.MatchString(date)
}

// isValidWebsite validates website URL format.
func isValidWebsite(url string) bool {
	return websiteRegex.MatchString(url)
}

// ValidateBasicInfoSimple performs simple validation without config options.
func ValidateBasicInfoSimple(req *models.BasicInfoRequest) error {
	if req.RegistrationReason == "" {
		return fmt.Errorf("دلیل ثبت‌نام الزامی است")
	}
	if req.ActivityType == "" {
		return fmt.Errorf("نوع فعالیت الزامی است")
	}
	if req.StartDate == "" {
		return fmt.Errorf("تاریخ شروع فعالیت الزامی است")
	}
	if !isValidPersianDate(req.StartDate) {
		return fmt.Errorf("فرمت تاریخ صحیح نیست (مثال: 1403/01/01)")
	}
	if req.UnitTitle == "" {
		return fmt.Errorf("عنوان واحد الزامی است")
	}
	if req.EightCategoryJob == "" {
		return fmt.Errorf("مشاغل هشتگانه الزامی است")
	}
	if req.IndividualJob == "" {
		return fmt.Errorf("مشاغل انفرادی الزامی است")
	}
	if req.ProfessionalGuild == "" {
		return fmt.Errorf("مجامع صنفی الزامی است")
	}
	if req.ProfessionalAssembly == "" {
		return fmt.Errorf("مجامع صنفی (انتخاب) الزامی است")
	}
	if req.GuildUnion == "" {
		return fmt.Errorf("اتحادیه صنفی الزامی است")
	}
	if req.BusinessLicense == "" {
		return fmt.Errorf("پروانه کسب الزامی است")
	}
	if req.OwnershipType == "" {
		return fmt.Errorf("نوع مالکیت الزامی است")
	}
	if req.Website != "" && !isValidWebsite(req.Website) {
		return fmt.Errorf("آدرس پایگاه اینترنتی معتبر نیست")
	}
	return nil
}
