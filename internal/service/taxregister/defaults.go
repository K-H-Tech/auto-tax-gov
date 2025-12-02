package taxregister

import (
	"fmt"
	"time"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
)

// GetCurrentJalaliDate returns current date in Jalali format (YYYY/MM/DD).
func GetCurrentJalaliDate() string {
	// Simple Gregorian to Jalali conversion for current date
	// For a production app, use a proper Jalali calendar library
	now := time.Now()
	year, month, day := gregorianToJalali(now.Year(), int(now.Month()), now.Day())
	return formatJalaliDate(year, month, day)
}

// gregorianToJalali converts Gregorian date to Jalali date.
// This is a simplified conversion - for production, use a proper library.
func gregorianToJalali(gy, gm, gd int) (jy, jm, jd int) {
	var g_days_in_month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var j_days_in_month = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	gy2 := gy - 1600
	gm2 := gm - 1
	gd2 := gd - 1

	g_day_no := 365*gy2 + (gy2+3)/4 - (gy2+99)/100 + (gy2+399)/400
	for i := 0; i < gm2; i++ {
		g_day_no += g_days_in_month[i]
	}
	if gm2 > 1 && ((gy%4 == 0 && gy%100 != 0) || gy%400 == 0) {
		g_day_no++
	}
	g_day_no += gd2

	j_day_no := g_day_no - 79

	j_np := j_day_no / 12053
	j_day_no = j_day_no % 12053

	jy = 979 + 33*j_np + 4*(j_day_no/1461)
	j_day_no = j_day_no % 1461

	if j_day_no >= 366 {
		jy += (j_day_no - 1) / 365
		j_day_no = (j_day_no - 1) % 365
	}

	for i := 0; i < 11 && j_day_no >= j_days_in_month[i]; i++ {
		j_day_no -= j_days_in_month[i]
		jm = i + 2
	}
	if jm == 0 {
		jm = 1
	}
	jd = j_day_no + 1

	return jy, jm, jd
}

// formatJalaliDate formats Jalali date as YYYY/MM/DD.
func formatJalaliDate(year, month, day int) string {
	return fmt.Sprintf("%04d/%02d/%02d", year, month, day)
}

// ApplyBasicInfoDefaults applies default values to empty basic info fields.
func ApplyBasicInfoDefaults(req *BasicInfoFormData, defaults config.BasicInfoDefaults) {
	if req.RegistrationReason == "" {
		req.RegistrationReason = defaults.RegistrationReason
	}
	if req.ActivityType == "" {
		req.ActivityType = defaults.ActivityType
	}
	if req.StartDate == "" || req.StartDate == "1xxx/xx/xx" {
		req.StartDate = GetCurrentJalaliDate()
	}
	if req.EightCategoryJob == "" || req.EightCategoryJob == "نامشخص" {
		req.EightCategoryJob = defaults.EightCategoryJob
	}
	if req.ProfessionalAssembly == "" {
		req.ProfessionalAssembly = defaults.ProfessionalAssembly
	}
	if req.BusinessLicense == "" {
		req.BusinessLicense = defaults.BusinessLicense
	}
	if req.OwnershipType == "" || req.OwnershipType == "نامشخص" {
		req.OwnershipType = defaults.OwnershipType
	}
}

// ApplyMemberDefaults applies default values to empty member fields.
func ApplyMemberDefaults(req *MemberSubmitRequest, defaults config.MemberDefaults) {
	if req.PersonType == "" {
		req.PersonType = defaults.PersonType
	}
	if req.Nationality == "" {
		req.Nationality = defaults.Nationality
	}
	if req.BirthCountry == "" {
		req.BirthCountry = defaults.BirthCountry
	}
	if req.NationalCardType == "" {
		req.NationalCardType = defaults.NationalCardType
	}
	// MembershipType maps to PartnershipType in config
	if req.MembershipType == "" {
		req.MembershipType = defaults.PartnershipType
	}
	// IsResponsible maps to IsEmployed in config
	if req.IsResponsible == "" {
		req.IsResponsible = defaults.IsEmployed
	}
	if req.SignatureAuthority == "" {
		req.SignatureAuthority = defaults.SignatureAuthority
	}
	if req.ResponsibilityType == "" {
		req.ResponsibilityType = defaults.ResponsibilityType
	}
	if req.Position == "" {
		req.Position = defaults.Position
	}
	if req.StartDate == "" || req.StartDate == "1xxx/xx/xx" {
		req.StartDate = GetCurrentJalaliDate()
	}
	if req.EndDate == "" || req.EndDate == "1xxx/xx/xx" {
		req.EndDate = defaults.EndDate
	}
}

// ApplyINTACodeDefaults applies default values to empty INTA code fields.
func ApplyINTACodeDefaults(req *INTACodeSubmitRequest, defaults config.INTACodeDefaults) {
	if req.Code == "" {
		req.Code = defaults.Code
	}
	if req.Description == "" {
		req.Description = defaults.Description
	}
	if req.Percent == 0 {
		req.Percent = defaults.Percent
	}
}

// BasicInfoFormData represents the basic info form data structure.
type BasicInfoFormData struct {
	RegistrationReason   string `json:"registrationReason"`
	ActivityType         string `json:"activityType"`
	StartDate            string `json:"startDate"`
	UnitTitle            string `json:"unitTitle"`
	EightCategoryJob     string `json:"eightCategoryJob"`
	IndividualJob        string `json:"individualJob"`
	ProfessionalGuild    string `json:"professionalGuild"`
	ProfessionalAssembly string `json:"professionalAssembly"`
	GuildUnion           string `json:"guildUnion"`
	NewGuildUnion        string `json:"newGuildUnion"`
	BusinessLicense      string `json:"businessLicense"`
	LicenseAuthority     string `json:"licenseAuthority"`
	OwnershipType        string `json:"ownershipType"`
	Phone                string `json:"phone"`
	AreaCode             string `json:"areaCode"`
	Fax                  string `json:"fax"`
	FaxAreaCode          string `json:"faxAreaCode"`
	Email                string `json:"email"`
	Website              string `json:"website"`
	HasBusinessCard      bool   `json:"hasBusinessCard"`
	HasCoinTax           bool   `json:"hasCoinTax"`
}

// INTACodeSubmitRequest represents an INTA code submission request.
type INTACodeSubmitRequest struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Percent     int    `json:"percent"`
}

// GetDefaultBasicInfo returns a BasicInfoFormData with all defaults applied.
func GetDefaultBasicInfo(defaults config.BasicInfoDefaults, businessName string) BasicInfoFormData {
	return BasicInfoFormData{
		RegistrationReason:   defaults.RegistrationReason,
		ActivityType:         defaults.ActivityType,
		StartDate:            GetCurrentJalaliDate(),
		UnitTitle:            businessName,
		EightCategoryJob:     defaults.EightCategoryJob,
		ProfessionalAssembly: defaults.ProfessionalAssembly,
		BusinessLicense:      defaults.BusinessLicense,
		OwnershipType:        defaults.OwnershipType,
	}
}

// GetDefaultMember returns a MemberSubmitRequest with all defaults applied.
func GetDefaultMember(defaults config.MemberDefaults) MemberSubmitRequest {
	return MemberSubmitRequest{
		PersonType:         defaults.PersonType,
		Nationality:        defaults.Nationality,
		BirthCountry:       defaults.BirthCountry,
		NationalCardType:   defaults.NationalCardType,
		MembershipType:     defaults.PartnershipType,
		IsResponsible:      defaults.IsEmployed,
		SignatureAuthority: defaults.SignatureAuthority,
		ResponsibilityType: defaults.ResponsibilityType,
		Position:           defaults.Position,
		StartDate:          GetCurrentJalaliDate(),
		EndDate:            defaults.EndDate,
	}
}

// GetDefaultINTACode returns an INTACodeSubmitRequest with all defaults applied.
func GetDefaultINTACode(defaults config.INTACodeDefaults) INTACodeSubmitRequest {
	return INTACodeSubmitRequest{
		Code:        defaults.Code,
		Description: defaults.Description,
		Percent:     defaults.Percent,
	}
}
