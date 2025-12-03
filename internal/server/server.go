package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mygovauth"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mytax"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/taxregister"
)

// Server represents the HTTP server.
type Server struct {
	cfg     *config.Config
	handler *Handler
	logger  *slog.Logger
	server  *http.Server
}

// New creates a new Server instance.
func New(cfg *config.Config, logger *slog.Logger, webDir string) *Server {
	// Create services with dependency injection
	authSvc := mygovauth.New(cfg, logger.With("service", "mygovauth"))
	mytaxSvc := mytax.New(cfg, authSvc, logger.With("service", "mytax"))
	taxregisterSvc := taxregister.New(cfg, logger.With("service", "taxregister"))

	h := NewHandler(cfg, mytaxSvc, taxregisterSvc, logger, webDir)

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.ServeHome)
	mux.HandleFunc("/api/start-tracker", h.HandleStartTracker)
	mux.HandleFunc("/api/refresh-captcha", h.HandleRefreshCaptcha)
	mux.HandleFunc("/api/send-otp", h.HandleSendOTP)
	mux.HandleFunc("/api/verify-otp", h.HandleVerifyOTP)
	mux.HandleFunc("/api/access-dashboard", h.HandleAccessDashboard)
	mux.HandleFunc("/api/start-tax-file", h.HandleStartTaxFile)
	mux.HandleFunc("/api/form-options", h.HandleGetFormOptions)
	mux.HandleFunc("/api/defaults", h.HandleGetDefaults)
	mux.HandleFunc("/api/submit-basic-info", h.HandleSubmitBasicInfo)
	mux.HandleFunc("/api/submit-partners", h.HandleSubmitPartners)
	mux.HandleFunc("/api/submit-bank-accounts", h.HandleSubmitBankAccounts)
	mux.HandleFunc("/api/delete-registration", h.HandleDeleteRegistration)
	mux.HandleFunc("/api/incomplete-registrations", h.HandleGetIncompleteRegistrations)

	// TaxRegister 11-Step Flow API routes
	mux.HandleFunc("/api/register/new", h.HandleNewRegistration)
	mux.HandleFunc("/api/register/sso-url", h.HandleGetSSOUrl)
	mux.HandleFunc("/api/register/authenticate", h.HandleAuthenticateToRegister)
	mux.HandleFunc("/api/register/home-page", h.HandleGetRegisterHomePage)
	mux.HandleFunc("/api/register/public-data-form", h.HandleGetPublicDataForm)
	mux.HandleFunc("/api/register/full-flow", h.HandleExecuteFullFlow)

	// Members Form API routes
	mux.HandleFunc("/api/register/members/form", h.HandleGetMembersForm)
	mux.HandleFunc("/api/register/members/submit", h.HandleSubmitMember)
	mux.HandleFunc("/api/register/members/list", h.HandleListMembers)
	mux.HandleFunc("/api/register/members/delete", h.HandleDeleteMember)

	// INTA Code API routes
	mux.HandleFunc("/api/register/inta-code/options", h.HandleGetINTACodeOptions)
	mux.HandleFunc("/api/register/inta-code/submit", h.HandleSubmitINTACode)
	mux.HandleFunc("/api/register/inta-code/search", h.HandleSearchINTACodes)
	mux.HandleFunc("/api/register/inta-code/form", h.HandleGetINTACodeForm)
	mux.HandleFunc("/api/register/inta-code/cascade", h.HandleGetINTACascadeOptions)
	mux.HandleFunc("/api/register/inta-code/activities", h.HandleSubmitINTACodes)

	// Bank Account (SHEBA) API routes
	mux.HandleFunc("/api/register/sheba/submit", h.HandleSubmitShebaNumber)
	mux.HandleFunc("/api/register/sheba/list", h.HandleGetShebaList)
	mux.HandleFunc("/api/register/sheba/delete", h.HandleDeleteSheba)

	// Complete automated registration (single endpoint)
	mux.HandleFunc("/api/register/complete", h.HandleCompleteRegistration)
	mux.HandleFunc("/api/register/recover", h.HandleRecoverRegistration)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return &Server{
		cfg:     cfg,
		handler: h,
		logger:  logger,
		server:  server,
	}
}

// Start starts the HTTP server.
func (s *Server) Start() error {
	// Create channel for shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in goroutine
	go func() {
		s.logger.Info("starting web server",
			"port", s.cfg.Server.Port,
			"url", fmt.Sprintf("http://localhost:%d", s.cfg.Server.Port),
		)

		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for shutdown signal
	<-stop
	s.logger.Info("shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("error during shutdown", "error", err)
		return err
	}

	s.logger.Info("server stopped gracefully")
	return nil
}

// GetHandler returns the handler for testing purposes.
func (s *Server) GetHandler() *Handler {
	return s.handler
}
