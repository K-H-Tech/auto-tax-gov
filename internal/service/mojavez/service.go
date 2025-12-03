package mojavez

import (
	"log/slog"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/models"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mygovauth"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
)

// Service handles Mojavez portal operations.
// Mojavez is the Iranian business licensing service.
type Service struct {
	cfg    *config.Config
	auth   *mygovauth.Service
	logger *slog.Logger
}

// New creates a new Mojavez service instance.
func New(cfg *config.Config, auth *mygovauth.Service, logger *slog.Logger) *Service {
	return &Service{
		cfg:    cfg,
		auth:   auth,
		logger: logger,
	}
}

// Result represents the result of a Mojavez operation.
type Result struct {
	Success bool
	Message string
	Data    map[string]any
}

// GetAuthURL returns the OAuth2 authorization URL for this service.
func (s *Service) GetAuthURL() string {
	return s.cfg.GetMojavezAuthURL()
}

// InitiateLogin starts the login flow using mygovauth service.
func (s *Service) InitiateLogin(sess *session.Session) (*models.CaptchaData, []models.RedirectStep, error) {
	authURL := s.GetAuthURL()
	s.logger.Info("initiating Mojavez login", "auth_url", authURL)
	return s.auth.InitiateAuth(sess, authURL)
}

// SendOTP sends OTP through mygovauth service.
func (s *Service) SendOTP(sess *session.Session, mobile, captchaCode, captchaKey, csrfToken string) (*mygovauth.AuthResult, error) {
	return s.auth.SendOTP(sess, mobile, captchaCode, captchaKey, csrfToken)
}

// VerifyOTP verifies OTP through mygovauth service.
func (s *Service) VerifyOTP(sess *session.Session, mobile, otpCode, csrfToken string) (*mygovauth.AuthResult, error) {
	return s.auth.VerifyOTP(sess, mobile, otpCode, csrfToken)
}

// RefreshCaptcha refreshes captcha through mygovauth service.
func (s *Service) RefreshCaptcha(sess *session.Session) (*models.CaptchaData, error) {
	return s.auth.RefreshCaptcha(sess)
}

// GetConfig returns the service configuration.
func (s *Service) GetConfig() *config.MojavezConfig {
	return &s.cfg.Services.Mojavez
}

// TODO: Add Mojavez-specific operations here when requirements are known.
// Examples:
// - ApplyForLicense()
// - CheckLicenseStatus()
// - RenewLicense()
