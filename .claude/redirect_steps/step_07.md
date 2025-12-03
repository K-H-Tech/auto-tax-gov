# Step 7: 302 Found

**Timestamp:** (OAuth2 Authorization Code Grant)

**URL:** `https://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso&continue`

## Request

**Method:** GET

### Request Headers

```http
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.9,fa;q=0.8,ar;q=0.7
```

## Response

**Status Code:** 302

**Status:** 302 Found

**Redirect Location:** `https://my.tax.gov.ir/myiran/sso?code=vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM&state=state1`

### Response Headers

```http
Location: https://my.tax.gov.ir/myiran/sso?code=vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM&state=state1
```

### Notes

**CRITICAL STEP**: This is where the OAuth2 authorization server (sso.my.gov.ir) generates an authorization code and redirects back to the tax application's callback URL.

**Authorization Code:** `vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM`

**State:** `state1` (matches the original request - CSRF protection)

This redirect transfers control from the SSO server to the tax.gov.ir application, which will exchange the code for an access token.

### Response Body

*(redirect response)*

---

[← Previous Step](step_06.md) | [Summary](summary_full.md) | [Next Step →](step_08.md)
