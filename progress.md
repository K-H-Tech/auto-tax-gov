# Implementation Progress: 11-Step Post-Login Registration Flow

## Overview

Implementing the complete post-login registration flow for cross-domain authentication between `my.tax.gov.ir` and `register.tax.gov.ir`.

---

## Step Status

| Step | Description | Status | Notes |
|------|-------------|--------|-------|
| 1 | POST /Page/NewRegistration/ | ✅ Complete | Creates registration, returns GUID |
| 2 | GET /Page/SSODoc/{guid} | ✅ Complete | Gets SSO redirect URL |
| 3 | GET TokenLoginProcessWithSignout | ✅ Complete | 302 redirect, sets ASP.NET_SessionId |
| 4 | GET TokenPreLoginProcess | ✅ Complete | 302 redirect continues |
| 5 | GET /prelogin/edit/{path} | ✅ Complete | 302 redirect, sets PreregisterTAX.* cookies |
| 6 | GET /Pages/Preaction/HomePage | ✅ Complete | 200 OK, HomePage HTML |
| 7 | GET /Pages/Preaction/PublicData | ✅ Complete | 200 OK, Large ASP.NET form |
| 8 | POST PublicData (AJAX) | ✅ Complete | AJAX partial postback for dropdown |
| 9 | POST PublicData (Submit) | ✅ Complete | Full form submit |
| 10 | POST PublicData (Submit) | ✅ Complete | Another form submit |
| 11 | GET HomePage (Final) | ✅ Complete | Final status verification |

**Legend:** ✅ Complete | 🔄 In Progress | ⏳ Pending | ❌ Blocked

---

## Milestones

### Milestone 1: Foundation (Steps 1-2) ✅
- [x] Create `internal/service/taxregister/service.go`
- [x] Implement `NewRegistration()` method
- [x] Implement `GetSSOUrl()` method
- [x] Add configuration for taxregister (already existed)

### Milestone 2: Authentication Flow (Steps 3-5) ✅
- [x] Implement `AuthenticateToRegister()` method
- [x] Handle 302 redirect chain
- [x] Cookie management across domains
- [x] Carry TaxpayerToken from my.tax.gov.ir

### Milestone 3: Form Retrieval (Steps 6-7) ✅
- [x] Implement `GetHomePage()` method
- [x] Implement `GetPublicDataForm()` method
- [x] Create ASP.NET form parser utilities
- [x] Extract __VIEWSTATE, __EVENTVALIDATION

### Milestone 4: Form Interaction (Steps 8-10) ✅
- [x] Implement `UpdateDropdown()` for AJAX postback
- [x] Implement `SubmitPublicData()` for full submit
- [x] Handle X-MicrosoftAjax: Delta=true header
- [x] Parse AJAX response for updated form state

### Milestone 5: API Layer ✅
- [x] Add new handlers to `internal/server/handlers.go`
- [x] Register routes in `internal/server/server.go`
- [ ] Update `config.example.yaml` (optional - defaults work)

### Milestone 6: Testing 🔄
- [ ] Test complete 11-step flow with real credentials
- [ ] Verify cookie management
- [ ] Test error handling

---

## Files Created/Modified

| File | Status | Action |
|------|--------|--------|
| `internal/service/taxregister/service.go` | ✅ | CREATED |
| `internal/service/taxregister/aspnet.go` | ✅ | CREATED |
| `internal/service/taxregister/models.go` | ✅ | CREATED |
| `internal/service/taxregister/client.go` | ✅ | CREATED |
| `internal/server/handlers.go` | ✅ | MODIFIED |
| `internal/server/server.go` | ✅ | MODIFIED |

---

## API Endpoints Added

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/register/new` | POST | Step 1: Create new registration |
| `/api/register/sso-url` | GET | Step 2: Get SSO URL |
| `/api/register/authenticate` | POST | Steps 3-5: Cross-domain auth |
| `/api/register/home-page` | GET | Steps 6/11: Get HomePage |
| `/api/register/public-data-form` | GET | Step 7: Get PublicData form |
| `/api/register/full-flow` | POST | Execute complete 11-step flow |

---

## Service Methods

| Method | Steps | Description |
|--------|-------|-------------|
| `NewRegistration()` | 1 | POST to /Page/NewRegistration/, returns GUID |
| `GetSSOUrl()` | 2 | GET /Page/SSODoc/{guid}, returns SSO URL |
| `AuthenticateToRegister()` | 3-5 | Follow 302 redirect chain, collect cookies |
| `GetHomePage()` | 6, 11 | GET /Pages/Preaction/HomePage |
| `GetPublicDataForm()` | 7 | GET /Pages/Preaction/PublicData, parse ASP.NET form |
| `UpdateDropdown()` | 8 | AJAX partial postback for dropdown cascade |
| `SubmitPublicData()` | 9-10 | Full form POST submission |
| `ExecuteFullFlow()` | 1-11 | Complete registration workflow |

---

## Technical Notes

### Cookie Flow
```
my.tax.gov.ir:
  → TaxpayerToken (JWT)
  → .MyTax.SessionID
  → MyTaxToken

register.tax.gov.ir:
  → ASP.NET_SessionId (from step 3)
  → cookiesession1 (from step 3)
  → PreregisterTAX.username (from step 5)
  → PreregisterTAX.status (from step 5)
```

### ASP.NET Hidden Fields
- `__VIEWSTATE` - Large base64 encoded state
- `__VIEWSTATEGENERATOR` - Generator ID
- `__EVENTVALIDATION` - Form validation
- `__EVENTTARGET` - Event source control
- `__EVENTARGUMENT` - Event arguments

### AJAX Postback Headers
```
X-MicrosoftAjax: Delta=true
X-Requested-With: XMLHttpRequest
Content-Type: application/x-www-form-urlencoded; charset=UTF-8
```

---

## Implementation Summary

The 11-step post-login registration flow has been implemented in a new `taxregister` service. The service:

1. **Creates new registrations** via my.tax.gov.ir API (Step 1)
2. **Obtains SSO tokens** for cross-domain authentication (Step 2)
3. **Follows 302 redirect chains** collecting session cookies (Steps 3-5)
4. **Parses ASP.NET WebForms** extracting VIEWSTATE and EVENTVALIDATION (Steps 6-7)
5. **Handles AJAX partial postbacks** for dropdown cascades (Step 8)
6. **Submits complete forms** with all required fields (Steps 9-10)
7. **Verifies registration status** on final HomePage (Step 11)

---

## Current Session

**Date:** 2025-11-29
**Status:** Implementation complete, ready for testing
**Next:** Test with real credentials to verify end-to-end flow
