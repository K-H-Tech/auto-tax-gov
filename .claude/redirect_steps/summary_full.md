# Complete OAuth2 Login Flow - Full Redirect Chain

**Total Steps:** 9

**Flow Type:** OAuth2 Authorization Code Grant with OTP Authentication

---

## Flow Overview

This document describes the complete authentication flow for the Iranian tax system (my.tax.gov.ir) using OAuth2 with the government SSO (sso.my.gov.ir).

---

## Phase 1: Initial OAuth Request → Login Page (Steps 1-3)

### [Step 1](step_01.md) - OAuth Authorization Request
- **URL:** `https://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso`
- **Status:** 302
- **Redirects to:** `http://sso.my.gov.ir/login`
- **Purpose:** Tax application initiates OAuth2 authorization request

### [Step 2](step_02.md) - HTTP to HTTPS Redirect
- **URL:** `http://sso.my.gov.ir/login`
- **Status:** 302 Moved Temporarily
- **Redirects to:** `https://sso.my.gov.ir/login`
- **Purpose:** Security upgrade from HTTP to HTTPS

### [Step 3](step_03.md) - Login Page Display
- **URL:** `https://sso.my.gov.ir/login`
- **Status:** 200
- **Purpose:** User sees login form with captcha
- **User Actions Required:**
  1. Enter mobile number
  2. Solve captcha
  3. Click "Send OTP"
  4. Enter received OTP code
  5. Submit OTP verification

---

## Phase 2: Post-Login OAuth Flow (Steps 4-9)

### [Step 4](step_04.md) - Post-Login Redirect
- **URL:** `http://sso.my.gov.ir/logged-in`
- **Status:** 302 Found
- **Redirects to:** `https://sso.my.gov.ir/logged-in`
- **Purpose:** Initial redirect after successful OTP verification

### [Step 5](step_05.md) - Logged-in Handler
- **URL:** `https://sso.my.gov.ir/logged-in`
- **Status:** 302 Found
- **Redirects to:** `http://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso&continue`
- **Purpose:** Redirect back to OAuth authorization endpoint with authenticated session
- **Note:** `continue` parameter indicates to proceed with authorization

### [Step 6](step_06.md) - OAuth Authorize (HTTP)
- **URL:** `http://sso.my.gov.ir/oauth2/authorize?...&continue`
- **Status:** 302 Found
- **Redirects to:** `https://sso.my.gov.ir/oauth2/authorize?...&continue`
- **Purpose:** HTTP to HTTPS redirect for OAuth endpoint

### [Step 7](step_07.md) - Authorization Code Generation
- **URL:** `https://sso.my.gov.ir/oauth2/authorize?...&continue`
- **Status:** 302 Found
- **Redirects to:** `https://my.tax.gov.ir/myiran/sso?code=vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM&state=state1`
- **Purpose:** SSO generates authorization code and redirects to tax app callback
- **Authorization Code:** `vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM`

### [Step 8](step_08.md) - OAuth Callback Handler
- **URL:** `https://my.tax.gov.ir/myiran/sso?code=...&state=state1`
- **Status:** 302 Found
- **Redirects to:** `/`
- **Purpose:** Tax application:
  - Receives authorization code
  - Validates state parameter (CSRF protection)
  - Exchanges code for access token (backend call)
  - Creates user session
  - Redirects to home page

### [Step 9](step_09.md) - Authenticated Home Page
- **URL:** `https://my.tax.gov.ir/`
- **Status:** 200 OK
- **Purpose:** User successfully logged in, viewing tax application home page
- **Final destination** - no further redirects

---

## Quick Navigation

- [Step 1: OAuth Authorization Request](step_01.md)
- [Step 2: HTTP to HTTPS Redirect](step_02.md)
- [Step 3: Login Page with Captcha & OTP](step_03.md)
- [Step 4: Post-Login Redirect (HTTP)](step_04.md)
- [Step 5: Logged-in Handler](step_05.md)
- [Step 6: OAuth Authorize HTTP→HTTPS](step_06.md)
- [Step 7: Authorization Code Generation](step_07.md)
- [Step 8: OAuth Callback Handler](step_08.md)
- [Step 9: Authenticated Home Page](step_09.md)

---

## Technical Summary

**Authentication Method:** OAuth2 Authorization Code Grant Flow

**Identity Provider:** sso.my.gov.ir (Government SSO)

**Client Application:** my.tax.gov.ir (Tax System)

**Login Method:** Mobile OTP with Captcha

**Total Redirects:** 8

**Final Destination:** https://my.tax.gov.ir/ (authenticated)

---

## OAuth2 Flow Diagram

```
┌─────────────────┐
│  Tax App Client │
│ (my.tax.gov.ir) │
└────────┬────────┘
         │
         │ 1. Initiate OAuth (Step 1)
         ▼
┌─────────────────────────┐
│   SSO Authorization     │
│  sso.my.gov.ir/oauth2   │
└────────┬────────────────┘
         │
         │ 2. Redirect to Login (Steps 2-3)
         ▼
┌─────────────────────────┐
│     Login Page          │
│  - Enter Mobile         │
│  - Solve Captcha        │
│  - Enter OTP            │
└────────┬────────────────┘
         │
         │ 3. OTP Verified (Step 4-5)
         ▼
┌─────────────────────────┐
│   SSO Authorization     │
│  (with session cookie)  │
└────────┬────────────────┘
         │
         │ 4. Generate Auth Code (Step 7)
         ▼
┌─────────────────────────┐
│   Tax App Callback      │
│  /myiran/sso?code=...   │
│  - Exchange code→token  │
│  - Create session       │
└────────┬────────────────┘
         │
         │ 5. Redirect to Home (Step 8)
         ▼
┌─────────────────────────┐
│   Tax App Home Page     │
│  (Authenticated)        │
└─────────────────────────┘
```

---

## Key Security Features

1. **CSRF Protection:** State parameter validated throughout flow
2. **HTTPS Enforcement:** Automatic HTTP→HTTPS redirects
3. **Captcha Verification:** Prevents automated attacks
4. **OTP Authentication:** Mobile-based two-factor authentication
5. **Authorization Code:** One-time use code for token exchange
6. **Session Management:** Secure session cookies set by tax application

---

## Notes

- The flow follows standard OAuth2 Authorization Code Grant specification
- Mobile OTP is used as the authentication factor instead of password
- Captcha is required to prevent brute-force attacks on OTP sending
- All redirects maintain session cookies for continuity
- The authorization code is single-use and expires quickly
- Tax application exchanges code for token in backend (not visible in browser)
