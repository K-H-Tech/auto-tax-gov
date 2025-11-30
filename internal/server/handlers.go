package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/helpers"
	"github.com/K-H-Tech/auto-tax-gov/internal/models"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mytax"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/taxregister"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
)

// Handler holds dependencies for HTTP handlers.
type Handler struct {
	cfg         *config.Config
	mytax       *mytax.Service
	taxregister *taxregister.Service
	session     *session.Session
	logger      *slog.Logger
	webDir      string
}

// NewHandler creates a new Handler instance.
func NewHandler(cfg *config.Config, mytaxSvc *mytax.Service, taxregisterSvc *taxregister.Service, logger *slog.Logger, webDir string) *Handler {
	return &Handler{
		cfg:         cfg,
		mytax:       mytaxSvc,
		taxregister: taxregisterSvc,
		session:     session.New(),
		logger:      logger,
		webDir:      webDir,
	}
}

// ServeHome serves the main HTML page.
func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	indexPath := filepath.Join(h.webDir, "index.html")
	content, err := os.ReadFile(indexPath)
	if err != nil {
		h.logger.Error("error loading index.html", "error", err, "path", indexPath)
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}

// HandleStartTracker initiates redirect tracking and returns captcha.
func (h *Handler) HandleStartTracker(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	h.logger.Info("starting MyTax login flow")

	// Reset session for new login
	h.session.Reset()

	captcha, steps, err := h.mytax.InitiateLogin(h.session)
	if err != nil {
		h.logger.Error("error initiating login", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	h.logger.Info("login initiated", "steps", len(steps))

	if captcha == nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No captcha found in redirect chain",
		})
		return
	}

	json.NewEncoder(w).Encode(models.CaptchaResponse{
		Success: true,
		Captcha: &struct {
			Key       string `json:"key"`
			ImageData string `json:"imageData"`
			CSRFToken string `json:"csrfToken"`
		}{
			Key:       captcha.Key,
			ImageData: captcha.ImageData,
			CSRFToken: captcha.CSRFToken,
		},
	})
}

// HandleRefreshCaptcha fetches a new captcha.
func (h *Handler) HandleRefreshCaptcha(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No active session. Please reload the page.",
		})
		return
	}

	captcha, err := h.mytax.RefreshCaptcha(h.session)
	if err != nil {
		h.logger.Error("error refreshing captcha", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.CaptchaResponse{
		Success: true,
		Captcha: &struct {
			Key       string `json:"key"`
			ImageData string `json:"imageData"`
			CSRFToken string `json:"csrfToken"`
		}{
			Key:       captcha.Key,
			ImageData: captcha.ImageData,
			CSRFToken: captcha.CSRFToken,
		},
	})
}

// HandleSendOTP sends an OTP to the user's mobile.
func (h *Handler) HandleSendOTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.OTPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request",
		})
		return
	}

	// Log with masked mobile to avoid PII exposure (Issue 7)
	h.logger.Info("send OTP attempt", "mobile", helpers.MaskMobile(req.Mobile))

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No active session. Please reload the page.",
		})
		return
	}

	result, err := h.mytax.SendOTP(h.session, req.Mobile, req.CaptchaCode, req.CaptchaKey, req.CSRFToken)
	if err != nil {
		h.logger.Error("send OTP error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleVerifyOTP verifies the OTP code.
func (h *Handler) HandleVerifyOTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.OTPVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request",
		})
		return
	}

	// Log with masked mobile to avoid PII exposure (Issue 7)
	h.logger.Info("verify OTP attempt", "mobile", helpers.MaskMobile(req.Mobile))

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No active session. Please reload the page.",
		})
		return
	}

	result, err := h.mytax.VerifyOTP(h.session, req.Mobile, req.OTPCode, req.CSRFToken)
	if err != nil {
		h.logger.Error("verify OTP error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleAccessDashboard accesses the tax dashboard.
func (h *Handler) HandleAccessDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	h.logger.Info("attempting to access dashboard")

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No active session. Please reload the page.",
		})
		return
	}

	result, err := h.mytax.AccessDashboard(h.session)
	if err != nil {
		h.logger.Error("dashboard access error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleStartTaxFile initiates tax file registration.
func (h *Handler) HandleStartTaxFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.TaxRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request",
		})
		return
	}

	// Validate required fields
	if req.PostalCode == "" || req.BusinessName == "" || req.Type == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "لطفاً تمام فیلدها را پر کنید",
		})
		return
	}

	h.logger.Info("starting tax file registration", "postalCode", req.PostalCode, "type", req.Type)

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "No active session. Please reload the page.",
		})
		return
	}

	result, err := h.mytax.StartTaxFileRegistration(h.session, req.PostalCode, req.BusinessName, req.Type)
	if err != nil {
		h.logger.Error("tax file registration error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleGetFormOptions returns form dropdown options.
func (h *Handler) HandleGetFormOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	options := h.mytax.GetFormOptions()

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    options,
	})
}

// HandleSubmitBasicInfo submits the basic information form (Step 2).
func (h *Handler) HandleSubmitBasicInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.BasicInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	h.logger.Info("submitting basic info", "unitTitle", req.UnitTitle)

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست فعال نیست. لطفاً صفحه را بارگذاری مجدد کنید.",
		})
		return
	}

	result, err := h.mytax.SubmitBasicInfo(h.session, &req)
	if err != nil {
		h.logger.Error("basic info submission error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleSubmitPartners submits the partners form (Step 3 - شرکا و اعضا).
func (h *Handler) HandleSubmitPartners(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.PartnersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	h.logger.Info("submitting partners", "count", len(req.Partners))

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست فعال نیست. لطفاً صفحه را بارگذاری مجدد کنید.",
		})
		return
	}

	result, err := h.mytax.SubmitPartners(h.session, &req)
	if err != nil {
		h.logger.Error("partners submission error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleSubmitBankAccounts submits the bank accounts form (Step 4 - حساب‌های بانکی).
func (h *Handler) HandleSubmitBankAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.BankAccountsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	h.logger.Info("submitting bank accounts", "count", len(req.Accounts))

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست فعال نیست. لطفاً صفحه را بارگذاری مجدد کنید.",
		})
		return
	}

	result, err := h.mytax.SubmitBankAccounts(h.session, &req)
	if err != nil {
		h.logger.Error("bank accounts submission error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleDeleteRegistration deletes an incomplete registration.
func (h *Handler) HandleDeleteRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.DeleteRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	if req.RegistrationID == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "شناسه ثبت‌نام الزامی است",
		})
		return
	}

	h.logger.Info("deleting registration", "registrationId", req.RegistrationID)

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست فعال نیست. لطفاً صفحه را بارگذاری مجدد کنید.",
		})
		return
	}

	result, err := h.mytax.DeleteRegistration(h.session, req.RegistrationID)
	if err != nil {
		h.logger.Error("delete registration error", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data:    result.Data,
	})
}

// HandleGetIncompleteRegistrations returns list of incomplete registrations.
func (h *Handler) HandleGetIncompleteRegistrations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	h.logger.Info("fetching incomplete registrations")

	if !h.session.IsActive() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست فعال نیست. لطفاً صفحه را بارگذاری مجدد کنید.",
		})
		return
	}

	registrations, err := h.mytax.GetIncompleteRegistrations(h.session)
	if err != nil {
		h.logger.Error("error fetching incomplete registrations", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    registrations,
	})
}

// ==================== TaxRegister 11-Step Flow Handlers ====================

// HandleNewRegistration creates a new tax registration (Step 1).
func (h *Handler) HandleNewRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req taxregister.RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	h.logger.Info("Step 1: Creating new registration",
		"type", req.Type,
		"postalCode", req.PostalCode)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	result, err := h.taxregister.NewRegistration(h.session, &req)
	if err != nil {
		h.logger.Error("Step 1 failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"guid":    result.GUID,
			"message": result.Message,
		},
	})
}

// HandleGetSSOUrl fetches the SSO redirect URL (Step 2).
func (h *Handler) HandleGetSSOUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	registrationGUID := r.URL.Query().Get("guid")
	if registrationGUID == "" {
		registrationGUID = h.session.GetRegistrationID()
	}

	if registrationGUID == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "GUID ثبت‌نام یافت نشد",
		})
		return
	}

	h.logger.Info("Step 2: Getting SSO URL", "guid", registrationGUID)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده",
		})
		return
	}

	result, err := h.taxregister.GetSSOUrl(h.session, registrationGUID)
	if err != nil {
		h.logger.Error("Step 2 failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"isLogin": result.IsLogin,
			"url":     result.URL,
		},
	})
}

// HandleAuthenticateToRegister performs cross-domain authentication (Steps 3-5).
func (h *Handler) HandleAuthenticateToRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req struct {
		SSOURL string `json:"ssoUrl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	if req.SSOURL == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "URL SSO الزامی است",
		})
		return
	}

	h.logger.Info("Steps 3-5: Authenticating to register.tax.gov.ir")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده",
		})
		return
	}

	err := h.taxregister.AuthenticateToRegister(h.session, req.SSOURL)
	if err != nil {
		h.logger.Error("Steps 3-5 failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "احراز هویت به register.tax.gov.ir موفق بود",
	})
}

// HandleGetRegisterHomePage fetches the registration HomePage (Steps 6, 11).
func (h *Handler) HandleGetRegisterHomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	h.logger.Info("Step 6/11: Fetching HomePage")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده",
		})
		return
	}

	result, err := h.taxregister.GetHomePage(h.session)
	if err != nil {
		h.logger.Error("Step 6/11 failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"guid":          result.GUID,
			"status":        result.Status,
			"statusMessage": result.StatusMessage,
		},
	})
}

// HandleGetPublicDataForm fetches the PublicData form (Step 7).
func (h *Handler) HandleGetPublicDataForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	h.logger.Info("Step 7: Fetching PublicData form")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده",
		})
		return
	}

	result, err := h.taxregister.GetPublicDataForm(h.session)
	if err != nil {
		h.logger.Error("Step 7 failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"guid":            result.GUID,
			"dropdownOptions": result.DropdownOptions,
			"hasViewState":    result.ViewState != "",
		},
	})
}

// HandleExecuteFullFlow executes the complete 11-step registration flow.
func (h *Handler) HandleExecuteFullFlow(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req taxregister.RegistrationFlowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	h.logger.Info("Executing full 11-step registration flow",
		"type", req.Type,
		"postalCode", req.PostalCode)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	result, err := h.taxregister.ExecuteFullFlow(h.session, &req)
	if err != nil {
		h.logger.Error("Full flow failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
			Data:    result, // Include partial results
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: result.Message,
		Data:    result,
	})
}

// ==================== Members Form Handlers ====================

// HandleGetMembersForm fetches the MembersEdit form for creating or editing a member.
// GET /api/register/members/form?memberId={optional}
func (h *Handler) HandleGetMembersForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	memberID := r.URL.Query().Get("memberId")

	h.logger.Info("Fetching MembersEdit form", "memberId", memberID)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	result, err := h.taxregister.GetMembersForm(h.session, memberID)
	if err != nil {
		h.logger.Error("GetMembersForm failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"memberId":        result.MemberID,
			"dropdownOptions": result.DropdownOptions,
			"fields":          result.Fields,
			"hasViewState":    result.ViewState != "",
		},
	})
}

// HandleSubmitMember submits a member form (create or update).
// POST /api/register/members/submit
func (h *Handler) HandleSubmitMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req taxregister.MemberSubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	// Validate required fields
	if req.NationalID == "" || req.Mobile == "" || req.PostalCode == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "کد ملی، موبایل و کد پستی الزامی هستند",
		})
		return
	}

	h.logger.Info("Submitting member form", "nationalId", req.NationalID)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	// First get the form to obtain ASP.NET state
	memberID := r.URL.Query().Get("memberId")
	form, err := h.taxregister.GetMembersForm(h.session, memberID)
	if err != nil {
		h.logger.Error("Failed to get members form", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "خطا در دریافت فرم: " + err.Error(),
		})
		return
	}

	// Submit the member data
	result, err := h.taxregister.SubmitMember(h.session, form, &req)
	if err != nil {
		h.logger.Error("SubmitMember failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
		Data: map[string]interface{}{
			"memberId": result.MemberID,
		},
	})
}

// ==================== INTA Code Handlers ====================

// HandleGetINTACodeOptions returns dropdown options for the INTA code cascade.
// GET /api/register/inta-code/options
func (h *Handler) HandleGetINTACodeOptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	var req models.INTACodeOptionsRequest

	// Support both GET with query params and POST with body
	if r.Method == http.MethodPost {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			json.NewEncoder(w).Encode(models.APIResponse{
				Success: false,
				Error:   "درخواست نامعتبر است",
			})
			return
		}
	} else {
		// GET request - parse from query string
		req.RegistrationID = r.URL.Query().Get("registrationId")
		// Levels would need to be parsed from JSON query param if needed
	}

	h.logger.Info("Fetching INTA code options", "registrationId", req.RegistrationID, "levels", len(req.Levels))

	result, err := h.mytax.GetINTACodeOptions(h.session, &req)
	if err != nil {
		h.logger.Error("GetINTACodeOptions failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    result,
	})
}

// HandleSubmitINTACode submits INTA code activities.
// POST /api/register/inta-code/submit
func (h *Handler) HandleSubmitINTACode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.INTACodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	// Validate activities
	if len(req.Activities) == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "حداقل یک فعالیت باید وارد شود",
		})
		return
	}

	// Validate total percentage = 100%
	totalPercent := 0
	for _, activity := range req.Activities {
		totalPercent += activity.Percent
	}
	if totalPercent != 100 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "مجموع درصد فعالیت‌ها باید ۱۰۰٪ باشد",
		})
		return
	}

	h.logger.Info("Submitting INTA code activities",
		"registrationId", req.RegistrationID,
		"activityCount", len(req.Activities),
	)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	result, err := h.mytax.SubmitINTACode(h.session, &req)
	if err != nil {
		h.logger.Error("SubmitINTACode failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
	})
}

// ==================== Bank Account (SHEBA) Handlers ====================

// HandleSubmitShebaNumber submits bank account (SHEBA) information.
// POST /api/register/sheba/submit
func (h *Handler) HandleSubmitShebaNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req models.BankAccountsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "درخواست نامعتبر است",
		})
		return
	}

	// Validate accounts
	if len(req.Accounts) == 0 {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "حداقل یک حساب بانکی باید وارد شود",
		})
		return
	}

	// Validate each IBAN
	for i, account := range req.Accounts {
		iban := account.IBAN
		// Remove IR prefix if present
		if len(iban) >= 2 && (iban[:2] == "IR" || iban[:2] == "ir") {
			iban = iban[2:]
		}
		if len(iban) != 24 {
			json.NewEncoder(w).Encode(models.APIResponse{
				Success: false,
				Error:   fmt.Sprintf("شماره شبا حساب %d باید ۲۴ رقم باشد", i+1),
			})
			return
		}
	}

	h.logger.Info("Submitting bank accounts",
		"registrationId", req.RegistrationID,
		"accountCount", len(req.Accounts),
	)

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	result, err := h.mytax.SubmitBankAccounts(h.session, &req)
	if err != nil {
		h.logger.Error("SubmitBankAccounts failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: result.Success,
		Message: result.Message,
	})
}

// HandleGetShebaList returns the list of bank accounts for a registration.
// GET /api/register/sheba/list
func (h *Handler) HandleGetShebaList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	registrationID := r.URL.Query().Get("registrationId")
	h.logger.Info("Fetching bank accounts list", "registrationId", registrationID)

	accounts, err := h.taxregister.GetShebaList(h.session, registrationID)
	if err != nil {
		h.logger.Error("GetShebaList failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    accounts,
	})
}

// HandleDeleteSheba deletes a bank account from a registration.
// DELETE /api/register/sheba/delete
func (h *Handler) HandleDeleteSheba(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	registrationID := r.URL.Query().Get("registrationId")
	shebaID := r.URL.Query().Get("shebaId")

	if shebaID == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "شناسه حساب بانکی الزامی است",
		})
		return
	}

	h.logger.Info("Deleting bank account", "registrationId", registrationID, "shebaId", shebaID)

	err := h.taxregister.DeleteSheba(h.session, registrationID, shebaID)
	if err != nil {
		h.logger.Error("DeleteSheba failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "حساب بانکی با موفقیت حذف شد",
	})
}

// ==================== Member List/Delete Handlers ====================

// HandleListMembers returns the list of members for a registration.
// GET /api/register/members/list
func (h *Handler) HandleListMembers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	registrationID := r.URL.Query().Get("registrationId")
	h.logger.Info("Fetching members list", "registrationId", registrationID)

	members, err := h.taxregister.ListMembers(h.session, registrationID)
	if err != nil {
		h.logger.Error("ListMembers failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    members,
	})
}

// HandleDeleteMember deletes a member from a registration.
// DELETE /api/register/members/delete
func (h *Handler) HandleDeleteMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if !h.session.IsAuthenticated() {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "نشست احراز هویت نشده. لطفاً ابتدا وارد شوید.",
		})
		return
	}

	registrationID := r.URL.Query().Get("registrationId")
	memberID := r.URL.Query().Get("memberId")

	if memberID == "" {
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "شناسه عضو الزامی است",
		})
		return
	}

	h.logger.Info("Deleting member", "registrationId", registrationID, "memberId", memberID)

	err := h.taxregister.DeleteMember(h.session, registrationID, memberID)
	if err != nil {
		h.logger.Error("DeleteMember failed", "error", err)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "عضو با موفقیت حذف شد",
	})
}
