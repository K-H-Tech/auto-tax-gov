# JWT Best Practices Guide

This document outlines best practices for implementing JSON Web Tokens (JWT) in Go applications using the `github.com/golang-jwt/jwt/v5` library. It is based on the official documentation, RFC 7519, and community recommendations to ensure secure, maintainable, and efficient JWT usage, particularly when integrated with frameworks like Gin. Follow these guidelines to handle authentication and authorization securely in your Go projects.

## Installation and Setup

- Install the JWT package using Go modules:
  ```bash
  go get -u github.com/golang-jwt/jwt/v5
  ```
- Import it in your code:
  ```go
  import "github.com/golang-jwt/jwt/v5"
  ```
- Use Go version 1.15 or higher to avoid known security issues in `crypto/elliptic`.

## JWT Basics

JWTs consist of three parts: **Header**, **Payload (Claims)**, and **Signature**, base64url-encoded and separated by dots (`.`). They are commonly used as Bearer tokens for authentication in OAuth 2.0 flows.

- **Header**: Specifies the signing algorithm (e.g., HS256, RS256) and token type (`JWT`).
- **Payload**: Contains claims (e.g., `sub`, `exp`, `iat`) as per RFC 7519.
- **Signature**: Ensures token integrity using a secret or key pair.

## Signing Methods

The `jwt/v5` library supports HMAC SHA (HS256, HS384, HS512), RSA, RSA-PSS, and ECDSA. Choose the appropriate method based on your security and performance needs.

- **HMAC SHA (HS256, HS384, HS512)**: Use for symmetric key signing. Ensure the secret is strong (at least 32 bytes for HS256) and securely stored.
  ```go
  secret := []byte("your-256-bit-secret") // Must be secure and not hardcoded
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "sub": "user123",
      "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenString, err := token.SignedString(secret)
  ```
- **RSA/ECDSA**: Use for asymmetric key signing. Store private keys securely and use public keys for verification.
  ```go
  privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
  token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
      "sub": "user123",
      "exp": time.Now().Add(time.Hour * 24).Unix(),
  })
  tokenString, err := token.SignedString(privateKey)
  ```
- **Avoid `alg=none`**: Only allow it explicitly with `jwt.UnsafeAllowNoneSignatureType` for specific use cases, as it offers no security.

## Token Creation

- Use standard claims from RFC 7519 (`iss`, `sub`, `aud`, `exp`, `iat`, `nbf`, `jti`) to ensure interoperability.
  ```go
  claims := jwt.MapClaims{
      "sub":  "user123",
      "iss":  "myapp",
      "aud":  "myapp-api",
      "exp":  time.Now().Add(time.Hour * 24).Unix(),
      "iat":  time.Now().Unix(),
      "jti":  uuid.New().String(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ```
- Use `jwt.RegisteredClaims` for type-safe standard claims:
  ```go
  claims := jwt.RegisteredClaims{
      Subject:   "user123",
      Issuer:    "myapp",
      Audience:  jwt.ClaimStrings{"myapp-api"},
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
      IssuedAt:  jwt.NewNumericDate(time.Now()),
      ID:        uuid.New().String(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ```
- Include an expiration (`exp`) to limit token lifetime (e.g., 24 hours for access tokens, shorter for sensitive operations).
- Use a unique `jti` (JWT ID) to prevent token reuse and enable revocation tracking.

## Token Validation

- Always validate the signing algorithm to prevent `alg` downgrade attacks:
  ```go
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
      }
      return []byte("your-256-bit-secret"), nil
  })
  ```
- Verify standard claims (`exp`, `nbf`, `iss`, `aud`) using `jwt/v5`’s built-in validation:
  ```go
  claims := jwt.RegisteredClaims{}
  token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
      return []byte("your-256-bit-secret"), nil
  })
  if err != nil {
      if errors.Is(err, jwt.ErrTokenExpired) {
          log.Println("Token has expired")
          return
      }
      log.Fatalf("Token parsing error: %v", err)
  }
  if !token.Valid {
      log.Println("Invalid token")
  }
  ```
- Use `jwt.Keyfunc` to dynamically fetch keys (e.g., from a JWKS endpoint for public key verification).

## Security Best Practices

- **Validate `alg`**: Always check the algorithm matches the expected type to prevent attacks.
- **Secure Key Management**:
  - Store symmetric secrets in environment variables or a secure vault (e.g., AWS Secrets Manager).
  - Use asymmetric keys (RSA/ECDSA) for public-facing APIs, with private keys in secure storage.
- **Short-Lived Tokens**: Set short expiration times for access tokens (e.g., 15 minutes to 24 hours) and use refresh tokens for longer sessions.
- **Token Revocation**: Maintain a revocation list using `jti` or implement short-lived tokens to minimize revocation needs.
- **Avoid Sensitive Data in Payload**: Do not include sensitive information (e.g., passwords) in JWT claims, as the payload is only base64-encoded, not encrypted.
- **Use HTTPS**: Always transmit JWTs over secure channels to prevent interception.
- **Validate All Claims**: Check `iss`, `aud`, `exp`, and `nbf` to ensure tokens are used as intended.

## Integration with Gin

Use JWTs for authentication in Gin applications by implementing middleware to validate tokens in the `Authorization` header.

- **Example Middleware**:
  ```go
  package middleware

  import (
      "fmt"
      "github.com/gin-gonic/gin"
      "github.com/golang-jwt/jwt/v5"
      "net/http"
      "strings"
  )

  func AuthMiddleware(secret []byte) gin.HandlerFunc {
      return func(c *gin.Context) {
          authHeader := c.GetHeader("Authorization")
          if authHeader == "" {
              c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
              c.Abort()
              return
          }
          parts := strings.Split(authHeader, " ")
          if len(parts) != 2 || parts[0] != "Bearer" {
              c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
              c.Abort()
              return
          }
          tokenString := parts[1]
          claims := &jwt.RegisteredClaims{}
          token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
              if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
              }
              return secret, nil
          })
          if err != nil {
              c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
              c.Abort()
              return
          }
          if !token.Valid {
              c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
              c.Abort()
              return
          }
          c.Set("claims", claims)
          c.Next()
      }
  }
  ```
- **Usage in Gin**:
  ```go
  package main

  import (
      "github.com/gin-gonic/gin"
      "github.com/golang-jwt/jwt/v5"
      "time"
  )

  func main() {
      r := gin.Default()
      secret := []byte("your-256-bit-secret")

      // Public route
      r.POST("/login", func(c *gin.Context) {
          // Simulate user authentication
          claims := jwt.RegisteredClaims{
              Subject:   "user123",
              ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
          }
          token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
          tokenString, err := token.SignedString(secret)
          if err != nil {
              c.JSON(500, gin.H{"error": "Failed to generate token"})
              return
          }
          c.JSON(200, gin.H{"token": tokenString})
      })

      // Protected routes
      protected := r.Group("/api")
      protected.Use(middleware.AuthMiddleware(secret))
      protected.GET("/secure", func(c *gin.Context) {
          claims, _ := c.Get("claims")
          c.JSON(200, gin.H{"message": "Secure endpoint", "user": claims.(*jwt.RegisteredClaims).Subject})
      })

      r.Run(":8080")
  }
  ```

## Testing

- **Unit Tests**: Test token creation and validation using `httptest` for Gin handlers.
  ```go
  func TestAuthMiddleware(t *testing.T) {
      secret := []byte("your-256-bit-secret")
      token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
          Subject:   "user123",
          ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
      })
      tokenString, _ := token.SignedString(secret)

      r := gin.Default()
      r.GET("/test", middleware.AuthMiddleware(secret), func(c *gin.Context) {
          c.JSON(200, gin.H{"message": "success"})
      })

      req, _ := http.NewRequest("GET", "/test", nil)
      req.Header.Set("Authorization", "Bearer "+tokenString)
      w := httptest.NewRecorder()
      r.ServeHTTP(w, req)

      assert.Equal(t, http.StatusOK, w.Code)
  }
  ```
- **Mock Keys**: Use in-memory keys for testing to avoid external dependencies.
- **Edge Cases**: Test expired tokens, invalid signatures, and missing headers.

## Performance Tips

- Cache public keys or JWKS responses to reduce lookup overhead in high-traffic systems.
- Use short-lived tokens to minimize validation overhead and improve security.
- Avoid large payloads to reduce token size and processing time.

## Additional Resources

- Official JWT GitHub: https://github.com/golang-jwt/jwt
- RFC 7519: https://tools.ietf.org/html/rfc7519
- Go Package Docs: https://pkg.go.dev/github.com/golang-jwt/jwt/v5
- Migration Guide: https://github.com/golang-jwt/jwt/blob/main/MIGRATION_GUIDE.md

Adhere to these practices to ensure secure and efficient JWT usage in your Go applications. Update this document as new best practices or library updates emerge.