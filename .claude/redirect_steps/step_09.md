# Step 9: 200 OK

**Timestamp:** (Tax Application Home Page)

**URL:** `https://my.tax.gov.ir/`

## Request

**Method:** GET

### Request Headers

```http
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
Accept-Language: en-US,en;q=0.9,fa;q=0.8,ar;q=0.7
Cookie: [Session cookie from previous step]
```

## Response

**Status Code:** 200

**Status:** 200 OK

### Response Headers

```http
Content-Type: text/html; charset=utf-8
```

### Notes

**SUCCESS!** The user is now logged in and viewing the tax application home page.

This is the **final destination** in the OAuth2 authorization code flow:

1. **Steps 1-3:** Initial OAuth request → Login page
2. **Step 3:** User enters mobile + captcha, receives OTP
3. **Step 3:** User enters OTP code
4. **Steps 4-7:** Post-login redirects back through OAuth authorize endpoint
5. **Step 8:** Tax app receives authorization code and exchanges it for access token
6. **Step 9:** User lands on authenticated home page

### Response Body

```html
<!-- Tax application home page HTML -->
<!-- User is now authenticated and can access tax services -->
```

---

[← Previous Step](step_08.md) | [Summary](summary_full.md)
