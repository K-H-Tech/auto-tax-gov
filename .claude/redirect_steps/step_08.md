# Step 8: 302 Found

**Timestamp:** (Tax Application OAuth Callback Handler)

**URL:** `https://my.tax.gov.ir/myiran/sso?code=vnNXRMJs2NRmhW_BHj2p8P7xaXY5h5lRSQMjdu4RK_dwfJYRgLY9MIaTVR2XdkIHD571h6r5a-EkwFVq6iGLGSfH6IFzeJO-4MA14kCvNW2q4cfyB06bb5kM29MdqDuM&state=state1`

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

**Redirect Location:** `/`

### Response Headers

```http
Location: /
Set-Cookie: [Session cookie for my.tax.gov.ir]
```

### Notes

The tax application:
1. Receives the authorization code
2. Validates the state parameter (CSRF check)
3. Exchanges the code for an access token (backend server-to-server call)
4. Creates a user session
5. Sets session cookies
6. Redirects to the application home page (`/`)

**This is the final step before reaching the application.**

### Response Body

*(redirect response)*

---

[← Previous Step](step_07.md) | [Summary](summary_full.md) | [Next Step →](step_09.md)
