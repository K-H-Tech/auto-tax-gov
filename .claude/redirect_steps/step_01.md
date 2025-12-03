# Step 1: 302 

**Timestamp:** 2025-11-26T16:42:31+03:30

**URL:** `https://sso.my.gov.ir/oauth2/authorize?response_type=code&scope=openid%20profile&client_id=my.tax&state=state1&redirect_uri=https://my.tax.gov.ir/myiran/sso`

## Request

**Method:** GET

### Request Headers

```http
Accept-Language: en-US,en;q=0.5
Connection: keep-alive
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8
```

## Response

**Status Code:** 302

**Status:** 302 

**Redirect Location:** `http://sso.my.gov.ir/login`

### Response Headers

```http
Referrer-Policy: no-referrer-when-downgrade
Access-Control-Allow-Headers: Content-Type,Sec-Fetch-Site, Sec-Fetch-Mode, Sec-Fetch-Dest, Origin, Cross-Domain, Cache-Control, Referer, User-Agent, Host, Accept, Accept-Encoding, Accept-Language, Authorization, cachecontrol,X-HTTP-Method-Override
X-Xss-Protection: 0
X-Xss-Protection: 1; mode=block
Cache-Control: no-cache, no-store, max-age=0, must-revalidate
Access-Control-Max-Age: 3600
X-Frame-Options: DENY
X-Frame-Options: DENY
Expires: 0
Server: ITO
Date: Wed, 26 Nov 2025 13:12:49 GMT
Access-Control-Expose-Headers: *
Pragma: no-cache
X-Powered-By: ITO
Access-Control-Allow-Origin: https://my.gov.ir
Strict-Transport-Security: max-age=31536000; includeSubDomains; preload
Content-Length: 0
Connection: keep-alive
Access-Control-Allow-Methods: *
X-Content-Type-Options: nosniff
X-Content-Type-Options: nosniff
Location: http://sso.my.gov.ir/login
Set-Cookie: SESSION=MTg0Y2Y4NzctZTQxYi00YjYwLTkzNjMtZTU3Yjg4Y2Y4YjU2; Path=/; HttpOnly; SameSite=Lax
Set-Cookie: f5avraaaaaaaaaaaaaaaa_session_=HGPHMFNKOBMOMLFAGEJHMPHDBFIEJJKKMMGKAMLLGNCICDJCEIPBEGMJPCODPONANAEDGCDDCMDFAFBFJAKAGKLKDPGKJMKHJAHLFKPELICGAJDMOGHOMINMBPEBDPLE; HttpOnly; secure;
```

### Response Body

*(empty)*

---

[Summary](summary.md) | [Next Step →](step_02.md)
