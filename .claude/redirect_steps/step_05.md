# Step 5: 302 Found

**Timestamp:** (After HTTPS Upgrade)

**URL:** `https://sso.my.gov.ir/logged-in`

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

**Redirect Location:** `http://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso&continue`

### Response Headers

```http
Location: http://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso&continue
```

### Notes

This redirect returns to the OAuth2 authorize endpoint, but now with an authenticated session cookie. The `continue` parameter indicates the authorization flow should proceed.

### Response Body

*(redirect response)*

---

[← Previous Step](step_04.md) | [Summary](summary_full.md) | [Next Step →](step_06.md)
