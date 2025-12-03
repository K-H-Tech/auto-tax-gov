# Retry Mechanism Best Practices Guide with retry-go

This document outlines best practices for implementing retry mechanisms in Go applications using the `github.com/avast/retry-go/v4` library. It draws from the official `retry-go` documentation, community recommendations, and general retry patterns to ensure resilient and maintainable code. These guidelines are tailored for integration with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, as well as dependencies like `pgx`, `mongo-go-driver`, `amqp091-go`, `go-redis`, and `sony/gobreaker`. The library supports Go 1.21+ and provides a simple, flexible retry mechanism with exponential backoff, fixed delays, and context support.

## Installation and Setup

- Install `retry-go` using Go modules:
  ```bash
  go get github.com/avast/retry-go/v4
  ```
- Import the library:
  ```go
  import (
      "github.com/avast/retry-go/v4/retry"
  )
  ```
- Use Go 1.21+ for compatibility with recent Go versions.

## Project Structure

Organize retry logic to promote reusability and separation of concerns.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with retry initialization
  ├── internal/
  │   ├── resilience/          # Retry and circuit breaker logic
  │   │   ├── retry.go         # Retry factories and wrappers
  │   │   └── breaker.go       # Circuit breaker logic
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic with retries
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define retry logic in a `resilience` package with factories (e.g., `NewHTTPRetry`, `NewDBRetry`) for different use cases.
- Use generics for type-safe retries: `retry.DoWithData[T]` where `T` is the response type.

## Configuration

- **Basic Retry with Defaults**: Use `retry.Do` for simple retries with default settings (100ms backoff, 3 attempts).
  ```go
  err := retry.Do(func() error {
      resp, err := http.Get("http://example.com")
      if err != nil {
          return err
      }
      defer resp.Body.Close()
      return nil
  })
  ```
- **With Data**: Use `retry.DoWithData` for operations returning data.
  ```go
  body, err := retry.DoWithData(func() ([]byte, error) {
      resp, err := http.Get("http://example.com")
      if err != nil {
          return nil, err
      }
      defer resp.Body.Close()
      return io.ReadAll(resp.Body)
  })
  ```
- **Custom Options**:
  ```go
  err := retry.Do(
      func() error { return http.Get("http://example.com") },
      retry.Attempts(5),                    // Max attempts
      retry.Delay(500 * time.Millisecond),  // Initial delay
      retry.MaxDelay(2 * time.Second),      // Cap delay
      retry.DelayType(retry.BackOffDelay),  // Exponential backoff
      retry.MaxJitter(100 * time.Millisecond), // Random jitter
      retry.OnRetry(func(n uint, err error) {
          logger.Info("Retry attempt", zap.Uint("attempt", n), zap.Error(err))
      }),
      retry.RetryIf(func(err error) bool {
          // Retry on transient errors (e.g., timeouts, 503)
          return errors.Is(err, syscall.ECONNREFUSED) || strings.Contains(err.Error(), "timeout")
      }),
      retry.Context(ctx), // Cancel via context
  )
  ```
  - **Attempts**: Default 3; use `retry.UntilSucceeded()` or `retry.Attempts(0)` for infinite retries (with caution).
  - **Delay**: Initial delay (default 100ms).
  - **MaxDelay**: Cap exponential backoff.
  - **DelayType**: `BackOffDelay` (exponential), `FixedDelay`, `RandomDelay`, or `CombineDelay`.
  - **MaxJitter**: Randomize delays to avoid thundering herd.
  - **OnRetry**: Log or monitor retries.
  - **RetryIf**: Skip retries for non-transient errors (e.g., `retry.Unrecoverable`).
  - **Context**: Use for timeouts/cancellation.
- **Viper Integration**: Load settings dynamically.
  ```go
  opts := []retry.Option{
      retry.Attempts(uint(viper.GetInt("retry.attempts"))),
      retry.Delay(viper.GetDuration("retry.delay")),
      retry.MaxDelay(viper.GetDuration("retry.max_delay")),
  }
  ```
  - **Config File (config.yaml)**:
    ```yaml
    retry:
      attempts: 5
      delay: 500ms
      max_delay: 2s
    ```

## Best Practices

- **Selective Retries**: Use `RetryIf` to retry only transient errors (e.g., network issues, 429, 503).
  ```go
  retry.RetryIf(func(err error) bool {
      if strings.Contains(err.Error(), "429") || errors.Is(err, context.DeadlineExceeded) {
          return true
      }
      return false
  })
  ```
- **Unrecoverable Errors**: Wrap fatal errors with `retry.Unrecoverable`.
  ```go
  if resp.StatusCode == 400 {
      return retry.Unrecoverable(fmt.Errorf("bad request"))
  }
  ```
- **Context Cancellation**: Always use `retry.Context` for timeouts.
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  err := retry.Do(func() error { /* operation */ }, retry.Context(ctx))
  ```
- **Backoff Strategy**: Prefer `BackOffDelay` with `MaxJitter` for production to avoid overwhelming services.
- **Error Aggregation**: Use `retry.LastErrorOnly(true)` to simplify error handling; default aggregates all errors.
- **Circuit Breaker Integration**: Combine with `sony/gobreaker` for layered resilience.
  ```go
  cb := gobreaker.NewCircuitBreaker[[]byte](gobreaker.Settings{})
  data, err := cb.Execute(func() ([]byte, error) {
      return retry.DoWithData(func() ([]byte, error) {
          resp, err := http.Get("http://example.com")
          // Handle response
          return body, nil
      }, retry.Attempts(3))
  })
  ```
- **Avoid Infinite Retries**: Use `UntilSucceeded` only with `Context` and monitoring to prevent hangs.
- **Thread Safety**: `retry.Do` is safe for concurrent use but ensure the retried function is idempotent.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Handler with Retry**:
  ```go
  func GetDataHandler(client *http.Client) gin.HandlerFunc {
      return func(c *gin.Context) {
          body, err := retry.DoWithData(func() ([]byte, error) {
              resp, err := client.Get("http://example.com")
              if err != nil {
                  return nil, err
              }
              defer resp.Body.Close()
              return io.ReadAll(resp.Body)
          }, retry.Attempts(3), retry.Delay(500*time.Millisecond))
          if err != nil {
              c.JSON(503, gin.H{"error": "Service unavailable"})
              return
          }
          c.Data(200, "application/json", body)
      }
  }
  ```
- **Viper**: As shown above; validate settings at startup.
- **Cobra CLI**: Retry batch operations.
  ```go
  var cmd = &cobra.Command{
      Run: func(cmd *cobra.Command, args []string) {
          err := retry.Do(func() error {
              return processBatch()
          }, retry.Attempts(uint(viper.GetInt("retry.attempts"))))
          if err != nil {
              log.Fatalf("Batch failed: %v", err)
          }
      },
  }
  ```
- **JWT**: Retry token validation if external (e.g., OAuth server).
  ```go
  token, err := retry.DoWithData(func() (*jwt.Token, error) {
      return jwt.Parse(tokenStr, keyFunc)
  }, retry.RetryIf(func(err error) bool { return strings.Contains(err.Error(), "timeout") }))
  ```
- **Zap Logging**: Log retries via `OnRetry`.
  ```go
  opts := []retry.Option{
      retry.OnRetry(func(n uint, err error) {
          logger.Warn("Retry", zap.Uint("attempt", n), zap.Error(err))
      }),
  }
  ```
- **Prometheus Metrics**: Track retries and failures.
  ```go
  var (
      retryAttempts = promauto.NewCounterVec(prometheus.CounterOpts{
          Name: "app_retry_attempts_total",
          Help: "Total retry attempts",
      }, []string{"operation"})
      retryFailures = promauto.NewCounterVec(prometheus.CounterOpts{
          Name: "app_retry_failures_total",
          Help: "Total retry failures",
      }, []string{"operation"})
  )

  opts := []retry.Option{
      retry.OnRetry(func(n uint, err error) {
          retryAttempts.WithLabelValues("http_get").Inc()
      }),
  }
  err := retry.Do(func() error { /* operation */ }, opts...)
  if err != nil {
      retryFailures.WithLabelValues("http_get").Inc()
  }
  ```

## Testing

- **Unit Tests**: Use `retry.WithTimer` for mock timers to avoid delays.
  ```go
  type MockTimer struct {
      Ch chan time.Time
  }
  func (t *MockTimer) After(d time.Duration) <-chan time.Time {
      return t.Ch
  }

  func TestRetry(t *testing.T) {
      timer := &MockTimer{Ch: make(chan time.Time)}
      err := retry.Do(
          func() error { return errors.New("fail") },
          retry.Attempts(2),
          retry.WithTimer(timer),
      )
      assert.Error(t, err)
      timer.Ch <- time.Now() // Simulate timer
  }
  ```
- **Integration Tests**: Combine with `httptest.Server` or mock dependencies (e.g., `uber-go/mock`).
- **Load Testing**: Use `wrk` to test retry under failure conditions.

## Performance Tips

- **Idempotency**: Ensure retried operations are safe (e.g., GET requests, idempotent POSTs).
- **Backoff Tuning**: Start with 100ms-500ms delay, 3-5 attempts; adjust based on service SLAs.
- **Jitter**: Always use `MaxJitter` (e.g., 50-100ms) to prevent synchronized retries.
- **Monitoring**: Track retry metrics to detect unstable dependencies.

## Additional Resources

- Official retry-go GitHub: https://github.com/avast/retry-go
- Retry Patterns: https://martinfowler.com/articles/patterns-of-distributed-systems/retry.html
- Go Resilience: https://github.com/sethgrid/pester (HTTP-specific alternative)

Adhere to these practices for robust retry mechanisms with `retry-go`. Update this document as new best practices or library updates emerge.