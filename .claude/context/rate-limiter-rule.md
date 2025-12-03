# Rate Limiting Best Practices Guide

This document outlines best practices for implementing rate limiting in Go applications. It draws from the official Go `golang.org/x/time/rate` package (token bucket algorithm), community recommendations, and middleware libraries like `github.com/ulule/limiter` (generic cell-rate) and `github.com/JGLTechnologies/gin-rate-limit` (Gin-specific). Rate limiting prevents abuse, ensures fair resource usage, and protects against DDoS attacks. For Gin web frameworks, use middleware; for general use, leverage `x/time/rate`. These guidelines integrate with tools like Viper, Zap, and Prometheus.

## Installation and Setup

- **Core Package**: Use Go's standard `golang.org/x/time/rate` for token bucket rate limiting.
  ```bash:disable-run
  go get golang.org/x/time/rate
  ```
- **Gin Middleware**: For HTTP rate limiting, use `github.com/ulule/limiter/v3` (flexible stores) or `github.com/JGLTechnologies/gin-rate-limit` (simple Redis/in-memory).
  ```bash
  go get github.com/ulule/limiter/v3
  go get github.com/redis/go-redis/v9  # For Redis store
  ```
- Import packages:
  ```go
  import (
      "golang.org/x/time/rate"
      "github.com/ulule/limiter/v3"
      limiterDrivers "github.com/ulule/limiter/v3/drivers/store/memory"  // Or redis
      "github.com/gin-gonic/gin"
  )
  ```
- Supports Go 1.21+; configure via Viper for production.

## Project Structure

Organize rate limiters for reusability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with limiter setup
  ├── internal/
  │   ├── middleware/          # Rate limiting middleware
  │   │   └── ratelimit.go     # Limiter factories
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic with client-side limiting
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define limiters in a `middleware` package with factories (e.g., `NewAPILimiter`, `NewClientLimiter`).
- Use interfaces for custom limiters: `type Limiter interface { Allow(ctx context.Context) bool }`.

## Configuration

- **Token Bucket (x/time/rate)**: Set rate (tokens/second) and burst (max burst size).
  ```go
  limiter := rate.NewLimiter(10, 50)  // 10 req/s, burst 50
  ```
  - **Rate**: Average requests per second (e.g., 100 for APIs).
  - **Burst**: Allow initial spikes (e.g., 200 for login pages).
  - Infinite rate: `rate.Inf` ignores burst.
- **ulule/limiter**: Use rates like "100-M" (100/minute); supports stores (memory, Redis).
  ```go
  import (
      "github.com/ulule/limiter/v3"
      drivers "github.com/ulule/limiter/v3/drivers/store/memory"
  )

  rate, _ := limiter.NewRateFromFormatted("10-M")  // 10/min
  store := drivers.NewStore()  // Or Redis
  limiterInstance := limiter.New(store, rate)
  ```
  - **Stores**: In-memory for single-instance; Redis for distributed.
  - **Keys**: Limit by IP (`c.ClientIP()`), user ID, or custom headers.
- **Viper Integration**: Load limits dynamically.
  ```go
  rateStr := viper.GetString("ratelimit.api.rate")  // e.g., "100-H"
  rate, _ := limiter.NewRateFromFormatted(rateStr)
  ```
  - **Config File (config.yaml)**:
    ```yaml
    ratelimit:
      api:
        rate: "100-M"  # 100/min
      client:
        rate: 5  # tokens/s
        burst: 10
    ```

## Usage

- **Server-Side (Gin Middleware with ulule/limiter)**: Apply to routes; returns 429 on exceed.
  ```go
  import "github.com/ulule/limiter/v3/middleware/gin"

  func main() {
      r := gin.Default()
      rate, _ := limiter.NewRateFromFormatted("10-M")
      store := drivers.NewStore()
      instance := limiter.New(store, rate)
      r.Use(ginmiddleware.NewMiddleware(instance))  // Limits by IP by default
      r.GET("/api", func(c *gin.Context) { c.JSON(200, gin.H{}) })
      r.Run(":8080")
  }
  ```
  - **Custom Key**: `WithKeyGetter(func(c *gin.Context) string { return c.GetHeader("API-Key") })`.
  - **Trust Forwarded Headers**: Enable carefully: `WithTrustForwardHeader(true)` for proxies like AWS ALB.
- **Client-Side (x/time/rate)**: Limit outgoing requests.
  ```go
  limiter := rate.NewLimiter(5, 10)  // 5 req/s, burst 10
  if limiter.Allow() {  // Or limiter.Wait(ctx) for blocking
      // Make request
  } else {
      // Delay or skip
  }
  ```
  - **WaitN**: Block until N tokens available: `limiter.WaitN(ctx, 3)`.
  - **Reservation**: Pre-allocate: `res := limiter.ReserveN(1); if res.OK() { /* use */ } else { res.Cancel() }`.
- **Burst Limiting**: For high initial loads, use larger burst but monitor.

## Best Practices

- **Algorithms**:
  - **Token Bucket** (`x/time/rate`): Smooth rate with bursts; ideal for APIs/clients.
  - **Leaky Bucket** (`ulule/limiter`): Fixed-rate outflow; good for queues.
  - **Fixed Window**: Simple but bursty at boundaries; avoid for user-facing.
  - **Sliding Window**: Most accurate; use Redis for distributed.
- **Keys**: Use IP for anonymous, user ID/JWT for authenticated; combine (e.g., IP+route).
  ```go
  key := c.ClientIP() + "_" + c.Param("user_id")
  ```
- **Placement**: Server-side at ingress (middleware); client-side for external APIs. Prefer API gateways (e.g., Kong) for distributed.
- **Headers**: Set `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` for clients.
- **Error Handling**: Return 429 with `Retry-After`; log throttles with Zap.
- **Distributed**: Use Redis store for multi-instance apps; prefix keys to avoid collisions.
- **Bypass**: Whitelist trusted IPs/users; monitor false positives.
- **Monitoring**: Track throttles with Prometheus (e.g., `ratelimit_throttled_total{key_type="ip"}`).
- **Testing**: Use `httptest` for middleware; simulate bursts.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Middleware (Custom with x/time/rate)**:
  ```go
  func RateLimit(limiter *rate.Limiter) gin.HandlerFunc {
      return func(c *gin.Context) {
          if !limiter.Allow() {
              c.Header("X-RateLimit-Remaining", "0")
              c.JSON(429, gin.H{"error": "Rate limited"})
              c.Abort()
              return
          }
          c.Next()
      }
  }

  // Usage
  r.Use(RateLimit(rate.NewLimiter(100, 200)))  // 100/s, burst 200
  ```
- **Viper**: As shown; validate rates at startup.
- **Cobra CLI**: Limit batch operations.
  ```go
  limiter := rate.NewLimiter(10, 50)
  for _, item := range items {
      limiter.Wait(context.Background())
      // Process item
  }
  ```
- **JWT**: Extract user ID from token for per-user limits.
  ```go
  keyFunc := func(c *gin.Context) string {
      claims := c.Get("claims").(*jwt.RegisteredClaims)
      return claims.Subject + "_" + c.ClientIP()
  }
  ```
- **Zap Logging**: Log throttles.
  ```go
  if !limiter.Allow() {
      logger.Warn("Rate limited", zap.String("ip", c.ClientIP()))
  }
  ```
- **Prometheus Metrics**:
  ```go
  var (
      throttledCounter = promauto.NewCounterVec(prometheus.CounterOpts{
          Name: "app_ratelimit_throttled_total",
          Help: "Total throttled requests",
      }, []string{"key", "reason"})
  )

  // In middleware
  throttledCounter.WithLabelValues(c.ClientIP(), "burst").Inc()
  ```

## Testing

- **Unit Tests**: Mock limiters or use in-memory stores.
  ```go
  func TestRateLimitMiddleware(t *testing.T) {
      limiter := rate.NewLimiter(1, 1)  // 1/s, burst 1
      r := gin.Default()
      r.Use(RateLimit(limiter))
      r.GET("/", func(c *gin.Context) { c.Status(200) })

      // First request succeeds
      w := httptest.NewRecorder()
      req, _ := http.NewRequest("GET", "/", nil)
      r.ServeHTTP(w, req)
      assert.Equal(t, 200, w.Code)

      // Second immediate request throttled
      w = httptest.NewRecorder()
      r.ServeHTTP(w, req)
      assert.Equal(t, 429, w.Code)
  }
  ```
- **Integration**: Use `ab` or `wrk` for load testing; monitor with Prometheus.

## Performance Tips

- **Low Overhead**: `x/time/rate` is lock-free for reads; use for high-throughput.
- **Stores**: In-memory for dev; Redis for prod (TTL cleanup).
- **Burst Management**: Set burst to 2-5x rate for UX without overload.

## Additional Resources

- Go `x/time/rate`: https://pkg.go.dev/golang.org/x/time/rate
- ulule/limiter: https://github.com/ulule/limiter
- Gin-rate-limit: https://github.com/JGLTechnologies/gin-rate-limit
- Go Wiki: https://go.dev/wiki/RateLimiting

Adhere to these practices for effective rate limiting. Update as new libraries emerge.
```