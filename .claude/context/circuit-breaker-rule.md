# Circuit Breaker Best Practices Guide with gobreaker

This document outlines best practices for implementing circuit breakers in Go applications using the `github.com/sony/gobreaker/v2` library. It draws from the official gobreaker documentation, circuit breaker patterns (e.g., Michael Nygard's principles), and community recommendations to ensure resilience, fault tolerance, and observability. These guidelines are tailored for integration with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, as well as dependencies like `pgx`, `mongo-go-driver`, `amqp091-go`, and `go-redis`. The library implements a state machine (Closed, Open, Half-Open) to prevent cascading failures.

## Installation and Setup

- Install gobreaker using Go modules:
  ```bash:disable-run
  go get github.com/sony/gobreaker/v2
  ```
- Import the library:
  ```go
  import (
      "github.com/sony/gobreaker/v2"
  )
  ```
- Use Go 1.21+ for compatibility, as the library supports recent Go versions.

## Project Structure

Organize circuit breakers to promote reusability and separation of concerns.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with CB initialization
  ├── internal/
  │   ├── resilience/          # Circuit breaker wrappers
  │   │   ├── breaker.go       # CB factories and configs
  │   │   └── http_client.go   # Wrapped HTTP client with CB
  │   ├── handlers/            # Gin handlers using CBs
  │   └── services/            # Business logic with external calls
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define circuit breakers in a `resilience` package with factories for different services (e.g., `NewDBBreaker`, `NewHTTPBreaker`).
- Use generics for type-safe wrappers: `CircuitBreaker[T]` where `T` is the response type.

## Configuration

- **Settings Struct**: Customize via `gobreaker.Settings` for fine-tuned behavior.
  ```go
  settings := gobreaker.Settings{
      Name:          "http-client-breaker",
      MaxRequests:   3,              // Max requests in half-open state
      Interval:      10 * time.Second, // Closed-state reset interval
      Timeout:       30 * time.Second, // Open-state timeout to half-open
      ReadyToTrip:   func(counts gobreaker.Counts) bool {
          failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
          return counts.Requests >= 5 && failureRatio > 0.6 // Trip if >60% failures after 5 requests
      },
      OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
          logger.Info("Circuit breaker state changed",
              zap.String("name", name),
              zap.String("from", from.String()),
              to.String("to", to.String()),
          )
      },
      IsSuccessful: func(err error) bool {
          // Customize success criteria (e.g., HTTP 2xx)
          return err == nil || status.IsSuccess(status.Code(err))
      },
  }
  cb := gobreaker.NewCircuitBreaker[[]byte](settings) // T = []byte for HTTP response
  ```
  - **Name**: Unique identifier for logging/monitoring.
  - **MaxRequests**: Limit in half-open (default 1 if 0).
  - **Interval**: Rolling window for counts (use with `BucketPeriod` for sliding window).
  - **Timeout**: Open duration before probing (default 60s if 0).
  - **ReadyToTrip**: Custom failure threshold (default: 5 consecutive failures).
  - **OnStateChange**: Hook for observability (e.g., log state transitions).
  - **IsSuccessful**: Define what counts as success (default: no error).
- **Counts**: Internal struct tracks requests, successes, failures (total/consecutive). Clears on state change or interval.
- **Viper Integration**: Load settings dynamically.
  ```go
  settings := gobreaker.Settings{
      MaxRequests: uint32(viper.GetInt("cb.max_requests")),
      Timeout:     viper.GetDuration("cb.timeout"),
      // ...
  }
  ```

## Usage

- **Execute Method**: Wrap external calls with `cb.Execute` for automatic state management.
  ```go
  func GetUserData(url string) ([]byte, error) {
      return cb.Execute(func() ([]byte, error) {
          resp, err := http.Get(url)
          if err != nil {
              return nil, err
          }
          defer resp.Body.Close()
          body, err := io.ReadAll(resp.Body)
          if err != nil {
              return nil, err
          }
          if resp.StatusCode >= 400 {
              return nil, fmt.Errorf("HTTP %d", resp.StatusCode)
          }
          return body, nil
      })
  }
  ```
  - Returns `gobreaker.ErrOpenState` if open (reject immediately).
  - Handles panics as errors and re-raises them.
- **State Management**: Query state with `cb.State()`.
  ```go
  if cb.State() == gobreaker.StateOpen {
      logger.Warn("Circuit open, using fallback")
      return fallbackData()
  }
  ```

## Best Practices

- **Per-Service Breakers**: Use separate breakers for different dependencies (e.g., one for DB, one for external API) to isolate failures.
- **Granular Configuration**: Tune thresholds based on service SLAs (e.g., lower `MaxRequests` for latency-sensitive calls).
- **Fallbacks**: Always provide fallbacks or cached responses when open.
  ```go
  data, err := cb.Execute(getExternalData)
  if err == gobreaker.ErrOpenState {
      return getCachedData() // Or default values
  }
  ```
- **State Transitions**:
  - **Closed**: Normal operation; track failures.
  - **Open**: Reject requests; use `Timeout` to transition to half-open.
  - **Half-Open**: Allow `MaxRequests` probes; success → Closed, failure → Open.
- **Rolling vs Fixed Window**: Set `BucketPeriod` > 0 for sliding window to smooth metrics.
- **Error Classification**: Customize `IsSuccessful` to ignore transient errors (e.g., timeouts as failures, but 429 as success for rate limits).
- **Avoid Over-Tripping**: Start conservative (e.g., 50% failure rate) and monitor.
- **Thread Safety**: `Execute` is safe for concurrent use.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin HTTP Client Wrapper**:
  ```go
  func ResilientGet(cb *gobreaker.CircuitBreaker[[]byte]) gin.HandlerFunc {
      return func(c *gin.Context) {
          data, err := cb.Execute(func() ([]byte, error) {
              resp, err := http.Get("https://api.example.com/data")
              // ... handle response
              return body, nil
          })
          if err != nil {
              c.JSON(503, gin.H{"error": "Service unavailable"})
              return
          }
          c.Data(200, "application/json", data)
      }
  }
  ```
- **Viper Configuration**: As shown above; default to conservative values.
- **Cobra CLI**: Wrap batch operations.
  ```go
  var cb *gobreaker.CircuitBreaker[string]
  // In PersistentPreRun
  settings := gobreaker.Settings{ /* Viper-loaded */ }
  cb = gobreaker.NewCircuitBreaker[string](settings)
  // In command Run
  result, err := cb.Execute(processBatch)
  ```
- **JWT Validation**: Wrap token validation calls if external.
  ```go
  func ValidateWithCB(cb *gobreaker.CircuitBreaker[*jwt.Token], tokenStr string) (*jwt.Token, error) {
      return cb.Execute(func() (*jwt.Token, error) {
          return jwt.Parse(tokenStr, keyFunc)
      })
  }
  ```
- **Zap Logging**: Use `OnStateChange` hook.
  ```go
  settings.OnStateChange = func(name string, from, to gobreaker.State) {
      logger.Info("CB state change",
          zap.String("breaker", name),
          zap.String("from", from.String()),
          zap.String("to", to.String()),
      )
  }
  ```
- **Prometheus Metrics**: Track states and failures.
  ```go
  var (
      cbStateGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
          Name: "myapp_cb_state",
          Help: "Circuit breaker state (0=closed, 1=open, 2=half-open)",
      }, []string{"name"})
      cbFailures = promauto.NewCounterVec(prometheus.CounterOpts{
          Name: "myapp_cb_failures_total",
          Help: "Circuit breaker failures",
      }, []string{"name"})
  )
  // In Execute wrapper
  if err != nil {
      cbFailures.WithLabelValues("http-client").Inc()
  }
  // Periodic scrape or hook to update gauge
  cbStateGauge.WithLabelValues("http-client").Set(float64(cb.State()))
  ```

## Testing

- **Unit Tests**: Mock wrapped functions or use `gobreaker` directly.
  ```go
  func TestCircuitBreaker(t *testing.T) {
      settings := gobreaker.Settings{MaxRequests: 1, Timeout: time.Second}
      cb := gobreaker.NewCircuitBreaker[string](settings)
      // Simulate failures
      _, err := cb.Execute(func() (string, error) { return "", errors.New("fail") })
      assert.Error(t, err)
      // Should trip to open
      assert.Equal(t, gobreaker.StateOpen, cb.State())
  }
  ```
- **Integration Tests**: Use with mocked external services (e.g., `httptest.Server`).
- **Mock External Calls**: Combine with GoMock for dependencies.

## Performance Tips

- **Low Overhead**: `Execute` adds minimal latency in closed state.
- **Concurrent Safety**: Safe for multiple goroutines.
- **Memory**: Counts use fixed-size buckets; tune `BucketPeriod` for balance.

## Additional Resources

- Official gobreaker GitHub: https://github.com/sony/gobreaker
- Example: https://github.com/sony/gobreaker/tree/master/v2/example
- Circuit Breaker Pattern: https://martinfowler.com/bliki/CircuitBreaker.html

Adhere to these practices for resilient Go applications with gobreaker. Update this document as new best practices emerge.
```