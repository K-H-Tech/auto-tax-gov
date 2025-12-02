package config

import (
	"log/slog"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds all application configuration.
type Config struct {
	Server      ServerConfig      `mapstructure:"server"`
	HTTP        HTTPConfig        `mapstructure:"http"`
	Services    ServicesConfig    `mapstructure:"services"`
	Log         LogConfig         `mapstructure:"log"`
	FormOptions FormOptionsConfig `mapstructure:"form_options"`
	Defaults    DefaultsConfig    `mapstructure:"defaults"`
}

// DefaultsConfig holds default values for tax registration forms.
type DefaultsConfig struct {
	BasicInfo BasicInfoDefaults `mapstructure:"basic_info"`
	Member    MemberDefaults    `mapstructure:"member"`
	INTACode  INTACodeDefaults  `mapstructure:"inta_code"`
	VATStatus VATStatusDefaults `mapstructure:"vat_status"`
}

// BasicInfoDefaults holds default values for the Basic Info form (PublicData).
type BasicInfoDefaults struct {
	RegistrationReason   string `mapstructure:"registration_reason"`
	ActivityType         string `mapstructure:"activity_type"`
	EightCategoryJob     string `mapstructure:"eight_category_job"`
	IndividualJob        string `mapstructure:"individual_job"`
	ProfessionalGuild    string `mapstructure:"professional_guild"`
	ProfessionalAssembly string `mapstructure:"professional_assembly"`
	GuildUnion           string `mapstructure:"guild_union"`
	NewGuildUnion        string `mapstructure:"new_guild_union"`
	BusinessLicense      string `mapstructure:"business_license"`
	OwnershipType        string `mapstructure:"ownership_type"`
	FinancialDayStart    string `mapstructure:"financial_day_start"`
	FinancialMonthStart  string `mapstructure:"financial_month_start"`
}

// MemberDefaults holds default values for the Member/Partner form.
type MemberDefaults struct {
	PersonType         string `mapstructure:"person_type"`
	Nationality        string `mapstructure:"nationality"`
	BirthCountry       string `mapstructure:"birth_country"`
	NationalCardType   string `mapstructure:"national_card_type"`
	PartnershipType    string `mapstructure:"partnership_type"`
	IsEmployed         string `mapstructure:"is_employed"`
	SignatureAuthority string `mapstructure:"signature_authority"`
	ResponsibilityType string `mapstructure:"responsibility_type"`
	Position           string `mapstructure:"position"`
	EndDate            string `mapstructure:"end_date"`
}

// INTACodeDefaults holds default values for the INTA Code form.
type INTACodeDefaults struct {
	Level1      string `mapstructure:"level1"`      // Category: تولید/بازرگانی/خدمات
	Level2      string `mapstructure:"level2"`      // Subcategory
	Level3      string `mapstructure:"level3"`      // Sub-subcategory
	Level4      string `mapstructure:"level4"`      // Final INTA code selection
	Code        string `mapstructure:"code"`        // 7-digit INTA code (extracted from Level4)
	Description string `mapstructure:"description"` // Activity description in Persian
	Percent     int    `mapstructure:"percent"`     // Activity percentage (100 for single activity)
}

// VATStatusDefaults holds default values for VAT eligibility status.
type VATStatusDefaults struct {
	EligibilityType string `mapstructure:"eligibility_type"` // عدم مشمولیت, مشمول مرحله اول, etc.
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

// HTTPConfig holds HTTP client configuration.
type HTTPConfig struct {
	Timeout         time.Duration `mapstructure:"timeout"`
	MaxRedirects    int           `mapstructure:"max_redirects"`
	VerifyTLS       bool          `mapstructure:"verify_tls"`
	SaveCookies     bool          `mapstructure:"save_cookies"`
	FollowRedirects bool          `mapstructure:"follow_redirects"`
	OutputDir       string        `mapstructure:"output_dir"`
}

// ServicesConfig holds configuration for all services.
type ServicesConfig struct {
	MyGovAuth   MyGovAuthConfig   `mapstructure:"mygovauth"`
	MyTax       MyTaxConfig       `mapstructure:"mytax"`
	RegisterTax RegisterTaxConfig `mapstructure:"registertax"`
	Enamad      EnamadConfig      `mapstructure:"enamad"`
	Mojavez     MojavezConfig     `mapstructure:"mojavez"`
}

// MyGovAuthConfig holds SSO/government authentication service configuration.
type MyGovAuthConfig struct {
	BaseURL      string `mapstructure:"base_url"`
	AuthURL      string `mapstructure:"auth_url"`
	LoginURL     string `mapstructure:"login_url"`
	LoggedInURL  string `mapstructure:"logged_in_url"`
	CaptchaURL   string `mapstructure:"captcha_url"`
	SendOTPURL   string `mapstructure:"send_otp_url"`
	VerifyOTPURL string `mapstructure:"verify_otp_url"`
	AuthHeader   string `mapstructure:"auth_header"`
}

// MyTaxConfig holds tax service configuration.
type MyTaxConfig struct {
	BaseURL         string `mapstructure:"base_url"`
	CallbackURL     string `mapstructure:"callback_url"`
	DashboardURL    string `mapstructure:"dashboard_url"`
	RegistrationURL string `mapstructure:"registration_url"`
	BasicInfoURL    string `mapstructure:"basic_info_url"`
	PartnersURL     string `mapstructure:"partners_url"`
	AccountsURL     string `mapstructure:"accounts_url"`
	AddAccountURL   string `mapstructure:"add_account_url"`
	DeleteRegURL    string `mapstructure:"delete_reg_url"`
	SSODocURL       string `mapstructure:"sso_doc_url"` // SSO token exchange endpoint for cross-domain auth
	ClientID        string `mapstructure:"client_id"`
	RedirectURI     string `mapstructure:"redirect_uri"`
}

// EnamadConfig holds Enamad service configuration.
type EnamadConfig struct {
	BaseURL     string `mapstructure:"base_url"`
	CallbackURL string `mapstructure:"callback_url"`
	ClientID    string `mapstructure:"client_id"`
	RedirectURI string `mapstructure:"redirect_uri"`
}

// MojavezConfig holds Mojavez service configuration.
type MojavezConfig struct {
	BaseURL     string `mapstructure:"base_url"`
	CallbackURL string `mapstructure:"callback_url"`
	ClientID    string `mapstructure:"client_id"`
	RedirectURI string `mapstructure:"redirect_uri"`
}

// RegisterTaxConfig holds tax registration service configuration (register.tax.gov.ir).
type RegisterTaxConfig struct {
	BaseURL             string `mapstructure:"base_url"`
	TokenLoginURL       string `mapstructure:"token_login_url"`        // Cross-domain auth entry point
	HomePageURL         string `mapstructure:"home_page_url"`          // Final destination after auth
	PublicDataURL       string `mapstructure:"public_data_url"`        // BasicInfo form (Step 2)
	MembersEditURL      string `mapstructure:"members_edit_url"`       // Partners/members form (Step 3)
	AddShebaNumberURL   string `mapstructure:"add_sheba_number_url"`   // Bank accounts form (Step 4)
	ActivityINTACodeURL string `mapstructure:"activity_inta_code_url"` // INTA code activities form
	VATStatusURL        string `mapstructure:"vat_status_url"`         // VAT eligibility status form
}

// LogConfig holds logging configuration.
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// DropdownOption represents a single option in a dropdown.
type DropdownOption struct {
	Value   string `mapstructure:"value" json:"value" yaml:"value"`
	Label   string `mapstructure:"label" json:"label" yaml:"label"`
	LabelFa string `mapstructure:"label_fa" json:"labelFa" yaml:"label_fa"`
}

// FormOptionsConfig holds dropdown options for forms.
type FormOptionsConfig struct {
	RegistrationReasons    []DropdownOption `mapstructure:"registration_reasons"`
	ActivityTypes          []DropdownOption `mapstructure:"activity_types"`
	EightCategoryJobs      []DropdownOption `mapstructure:"eight_category_jobs"`
	IndividualJobs         []DropdownOption `mapstructure:"individual_jobs"`
	ProfessionalGuilds     []DropdownOption `mapstructure:"professional_guilds"`
	ProfessionalAssemblies []DropdownOption `mapstructure:"professional_assemblies"`
	GuildUnions            []DropdownOption `mapstructure:"guild_unions"`
	BusinessLicenses       []DropdownOption `mapstructure:"business_licenses"`
	OwnershipTypes         []DropdownOption `mapstructure:"ownership_types"`
}

// Load reads configuration from file and environment.
func Load(configPath string) (*Config, error) {
	setDefaults()

	if configPath != "" {
		viper.SetConfigFile(configPath)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.auto-tax-gov")
		viper.AddConfigPath("/etc/auto-tax-gov")
	}

	// Read config file (optional - don't fail if not found)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
		slog.Debug("no config file found, using defaults")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Warn if SSO auth_header is not configured
	if cfg.Services.MyGovAuth.AuthHeader == "" {
		slog.Warn("SSO auth_header not configured - set via config file or AUTO_TAX_GOV_SERVICES_MYGOVAUTH_AUTH_HEADER env var")
	}

	return &cfg, nil
}

// setDefaults sets default configuration values.
func setDefaults() {
	// Server defaults
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", "5s")
	viper.SetDefault("server.write_timeout", "10s")
	viper.SetDefault("server.idle_timeout", "120s")

	// HTTP client defaults
	viper.SetDefault("http.timeout", "60s")
	viper.SetDefault("http.max_redirects", 20)
	viper.SetDefault("http.verify_tls", true)
	viper.SetDefault("http.save_cookies", true)
	viper.SetDefault("http.follow_redirects", true)
	viper.SetDefault("http.output_dir", "redirect_steps")

	// MyGovAuth service defaults (SSO)
	viper.SetDefault("services.mygovauth.base_url", "https://sso.my.gov.ir")
	viper.SetDefault("services.mygovauth.auth_url", "https://sso.my.gov.ir/oauth2/authorize")
	viper.SetDefault("services.mygovauth.login_url", "https://sso.my.gov.ir/login")
	viper.SetDefault("services.mygovauth.logged_in_url", "https://sso.my.gov.ir/logged-in")
	viper.SetDefault("services.mygovauth.captcha_url", "https://sso.my.gov.ir/client/v1/captcha")
	viper.SetDefault("services.mygovauth.send_otp_url", "https://sso.my.gov.ir/sendOtp")
	viper.SetDefault("services.mygovauth.verify_otp_url", "https://sso.my.gov.ir/login_otp")
	viper.SetDefault("services.mygovauth.auth_header", "") // Must be set via config file or AUTO_TAX_GOV_SERVICES_MYGOVAUTH_AUTH_HEADER env var

	// MyTax service defaults
	viper.SetDefault("services.mytax.base_url", "https://my.tax.gov.ir")
	viper.SetDefault("services.mytax.callback_url", "https://my.tax.gov.ir/myiran/sso")
	viper.SetDefault("services.mytax.dashboard_url", "https://my.tax.gov.ir/Page/Dashboard")
	viper.SetDefault("services.mytax.registration_url", "https://my.tax.gov.ir/Page/NewRegistration/")
	viper.SetDefault("services.mytax.basic_info_url", "https://my.tax.gov.ir/Page/BasicInfo/")
	viper.SetDefault("services.mytax.partners_url", "https://my.tax.gov.ir/Page/Partners/")
	viper.SetDefault("services.mytax.accounts_url", "https://my.tax.gov.ir/Page/Accounts/")
	viper.SetDefault("services.mytax.add_account_url", "https://my.tax.gov.ir/Page/AddAccount/")
	viper.SetDefault("services.mytax.delete_reg_url", "https://my.tax.gov.ir/Page/DeleteReg/")
	viper.SetDefault("services.mytax.sso_doc_url", "https://my.tax.gov.ir/Page/SSODoc/")
	viper.SetDefault("services.mytax.client_id", "my.tax")
	viper.SetDefault("services.mytax.redirect_uri", "https://my.tax.gov.ir/myiran/sso")

	// Enamad service defaults
	viper.SetDefault("services.enamad.base_url", "https://reg2.enamad.ir")
	viper.SetDefault("services.enamad.callback_url", "https://reg2.enamad.ir/callback")
	viper.SetDefault("services.enamad.client_id", "enamad")
	viper.SetDefault("services.enamad.redirect_uri", "https://reg2.enamad.ir/callback")

	// Mojavez service defaults
	viper.SetDefault("services.mojavez.base_url", "https://mojavez.ir")
	viper.SetDefault("services.mojavez.callback_url", "https://mojavez.ir/callback")
	viper.SetDefault("services.mojavez.client_id", "mojavez")
	viper.SetDefault("services.mojavez.redirect_uri", "https://mojavez.ir/callback")

	// RegisterTax service defaults (tax registration portal)
	viper.SetDefault("services.registertax.base_url", "https://register.tax.gov.ir")
	viper.SetDefault("services.registertax.token_login_url", "https://register.tax.gov.ir/Pages/Login/TokenLoginProcessWithSignout/")
	viper.SetDefault("services.registertax.home_page_url", "https://register.tax.gov.ir/Pages/Preaction/HomePage")
	viper.SetDefault("services.registertax.public_data_url", "https://register.tax.gov.ir/Pages/Preaction/PublicData")
	viper.SetDefault("services.registertax.members_edit_url", "https://register.tax.gov.ir/Pages/Preaction/MembersEdit")
	viper.SetDefault("services.registertax.add_sheba_number_url", "https://register.tax.gov.ir/Pages/Preaction/Edit/AddShebaNumber/")
	viper.SetDefault("services.registertax.activity_inta_code_url", "https://register.tax.gov.ir/Pages/Preaction/Edit/ActivityINTACode/")
	viper.SetDefault("services.registertax.vat_status_url", "https://register.tax.gov.ir/Pages/Preaction/Edit/VatStatus/")

	// Log defaults
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.format", "text")

	// Defaults for tax registration forms (happy path)
	// Basic Info defaults - values match exact dropdown options from register.tax.gov.ir
	viper.SetDefault("defaults.basic_info.registration_reason", "ایجاد یک کسب و کار جدید")
	viper.SetDefault("defaults.basic_info.activity_type", "تجاری (همچنین فعالیتهای معاف از مالیات)")
	viper.SetDefault("defaults.basic_info.eight_category_job", "سایر (گروه دوم و سوم)")
	viper.SetDefault("defaults.basic_info.individual_job", "سایر (غیر انفرادی)")
	viper.SetDefault("defaults.basic_info.professional_guild", "اتاق بازرگانی،صنایع و معادن و کشاورزی ایران")
	viper.SetDefault("defaults.basic_info.professional_assembly", "اتاق بازرگانی،صنایع و معادن و کشاورزی ایران")
	viper.SetDefault("defaults.basic_info.guild_union", "کسب و کاراینترنتی")
	viper.SetDefault("defaults.basic_info.new_guild_union", "اتاق بازرگانی،صنایع،معادن و کشاورزی (کرمانشاه-کرمانشاه) - كرمانشاه")
	viper.SetDefault("defaults.basic_info.business_license", "فعلا فاقد جواز کسب می باشم و در درست اقدام قرار دارد")
	viper.SetDefault("defaults.basic_info.ownership_type", "اجاری") // Must match actual dropdown: ملکی, سرقفلی, اجاری, وقفی
	viper.SetDefault("defaults.basic_info.financial_day_start", "1")
	viper.SetDefault("defaults.basic_info.financial_month_start", "1")

	// Member defaults
	viper.SetDefault("defaults.member.person_type", "حقیقی")
	viper.SetDefault("defaults.member.nationality", "33") // Iran
	viper.SetDefault("defaults.member.birth_country", "33")
	viper.SetDefault("defaults.member.national_card_type", "کارت ملی هوشمند جدید")
	viper.SetDefault("defaults.member.partnership_type", "اختیاری (مستقیم)")
	viper.SetDefault("defaults.member.is_employed", "بله")
	viper.SetDefault("defaults.member.signature_authority", "ندارد")
	viper.SetDefault("defaults.member.responsibility_type", "هیچ کدام")
	viper.SetDefault("defaults.member.position", "شریک")
	viper.SetDefault("defaults.member.end_date", "0") // Ongoing

	// INTA Code defaults - cascade levels and final code
	viper.SetDefault("defaults.inta_code.level1", "[3] خدمات")
	viper.SetDefault("defaults.inta_code.level2", "[11] خدمات امور اداری، حقوقی، حسابداری، کارشناسی، تحقیقاتی و رایانه ای")
	viper.SetDefault("defaults.inta_code.level3", "[0] خدمات")
	viper.SetDefault("defaults.inta_code.level4", "[3110310] [حقیقی/حقوقی] ارائه خدمات رایانه ای نظیر مشاوره، تهیه نقشه، طرح و یا اجرا و پشتیبانی سیستمهای رایانه ای")
	viper.SetDefault("defaults.inta_code.code", "3110310")
	viper.SetDefault("defaults.inta_code.description", "ارائه خدمات رایانه ای و مشاوره فناوری اطلاعات")
	viper.SetDefault("defaults.inta_code.percent", 100)

	// VAT Status defaults
	viper.SetDefault("defaults.vat_status.eligibility_type", "عدم مشمولیت")
}

// GetLogLevel returns the slog level from config.
func (c *Config) GetLogLevel() slog.Level {
	switch strings.ToLower(c.Log.Level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// GetMyTaxAuthURL returns the OAuth2 authorization URL for MyTax service.
func (c *Config) GetMyTaxAuthURL() string {
	u, _ := url.Parse(c.Services.MyGovAuth.AuthURL)
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("scope", "openid profile")
	q.Set("client_id", c.Services.MyTax.ClientID)
	q.Set("state", "state1")
	q.Set("redirect_uri", c.Services.MyTax.RedirectURI)
	u.RawQuery = q.Encode()
	return u.String()
}

// GetEnamadAuthURL returns the OAuth2 authorization URL for Enamad service.
func (c *Config) GetEnamadAuthURL() string {
	u, _ := url.Parse(c.Services.MyGovAuth.AuthURL)
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("scope", "openid profile")
	q.Set("client_id", c.Services.Enamad.ClientID)
	q.Set("state", "state1")
	q.Set("redirect_uri", c.Services.Enamad.RedirectURI)
	u.RawQuery = q.Encode()
	return u.String()
}

// GetMojavezAuthURL returns the OAuth2 authorization URL for Mojavez service.
func (c *Config) GetMojavezAuthURL() string {
	u, _ := url.Parse(c.Services.MyGovAuth.AuthURL)
	q := url.Values{}
	q.Set("response_type", "code")
	q.Set("scope", "openid profile")
	q.Set("client_id", c.Services.Mojavez.ClientID)
	q.Set("state", "state1")
	q.Set("redirect_uri", c.Services.Mojavez.RedirectURI)
	u.RawQuery = q.Encode()
	return u.String()
}
