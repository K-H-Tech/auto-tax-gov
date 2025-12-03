# Prometheus client_golang Best Practices Guide

This document outlines best practices for instrumenting Go applications with Prometheus metrics using the `github.com/prometheus/client_golang` library. It draws from the official Prometheus Go client documentation, the instrumenting guide, and community recommendations to ensure metrics are reliable, performant, and aligned with Prometheus conventions. These guidelines are particularly useful when integrating with frameworks like Gin, Viper, Cobra, JWT, and Zap for comprehensive observability in Go projects.

## Installation and Setup

- Install the core Prometheus client libraries using Go modules:
  ```bash
  go get github.com/prometheus/client_golang/prometheus
  go get github.com/prometheus/client_golang/prometheus/promauto
  go get github.com/prometheus/client_golang/prometheus/promhttp
  ```
- Import the necessary packages:
  ```go
  import (
      "github.com/prometheus/client_golang/prometheus"
      "github.com/prometheus/client_golang/prometheus/promauto"
      "github.com/prometheus/client_golang/prometheus/promhttp"
  )
  ```
- Use Go 1.21 or later for full compatibility, as the library supports only the two most recent major Go releases.

## Project Structure

Organize metrics code to promote reusability and separation of concerns.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with metrics setup
  ├── internal/
  │   ├── metrics/             # Metrics definitions and collectors
  │   │   ├── registry.go      # Custom registry if needed
  │   │   └── http.go          # HTTP metrics middleware
  │   ├── handlers/            # Gin handlers with metric instrumentation
  │   └── services/            # Business logic with metrics
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define metrics in a dedicated `metrics` package to centralize registration and avoid global state pollution.
- Initialize metrics in `main.go` or a setup function, registering them with the default registry or a custom one for isolation.

## Metric Types

Prometheus supports four core metric types. Choose based on what you want to measure: counts (Counter), current states (Gauge), distributions (Histogram/Summary).

- **Counter**: Monotonically increasing values (e.g., total requests). Use for irreversible counts.
  ```go
  var (
      requestsTotal = promauto.NewCounter(prometheus.CounterOpts{
          Name: "myapp_http_requests_total",
          Help: "Total number of HTTP requests",
      })
  )

  func handleRequest() {
      requestsTotal.Inc()  // Increment by 1
      // Or: requestsTotal.Add(5.0) for fractional increments
  }
  ```
- **Gauge**: Values that can go up or down (e.g., active connections, CPU usage).
  ```go
  var (
      activeUsers = promauto.NewGauge(prometheus.GaugeOpts{
          Name: "myapp_active_users",
          Help: "Number of active users",
      })
  )

  func updateUsers(count float64) {
      activeUsers.Set(count)  // Set to absolute value
      // Or: activeUsers.Inc() / activeUsers.Dec() for increments/decrements
  }
  ```
- **Histogram**: Sample distributions (e.g., request durations) with configurable buckets. Provides sum, count, and quantiles via PromQL.
  ```go
  var (
      requestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
          Name:    "myapp_http_request_duration_seconds",
          Help:    "Duration of HTTP requests",
          Buckets: prometheus.DefBuckets,  // Default buckets: 0.005, 0.01, ..., 1m
      })
  )

  func measureDuration(start time.Time) {
      requestDuration.Observe(time.Since(start).Seconds())
  }
  ```
- **Summary**: Similar to Histogram but pre-computes quantiles client-side (less flexible for aggregation).
  ```go
  var (
      responseSize = promauto.NewSummary(prometheus.SummaryOpts{
          Name:       "myapp_http_response_size_bytes",
          Help:       "Size of HTTP responses",
          Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
      })
  )

  func recordSize(size float64) {
      responseSize.Observe(size)
  }
  ```
- Use `promauto` for automatic registration with the default registry in simple apps. For advanced use, manually register with `prometheus.MustRegister(metric)` to handle errors explicitly.

## Exposing Metrics

Expose metrics via an HTTP endpoint for Prometheus scraping.

- **Basic Setup with net/http**:
  ```go
  func main() {
      http.Handle("/metrics", promhttp.Handler())  // Exposes all registered metrics
      http.ListenAndServe(":2112", nil)
  }
  ```
- **With Gin Integration**:
  ```go
  import "github.com/gin-gonic/gin"

  func main() {
      r := gin.Default()
      r.GET("/metrics", gin.WrapH(promhttp.Handler()))  // Wrap promhttp.Handler for Gin
      r.Run(":8080")
  }
  ```
- Use `promhttp.HandlerFor(registry, promhttp.HandlerOpts{})` for custom registries or options (e.g., disabling Go collector).

## Labels and Dimensions

Labels add dimensionality to metrics for filtering and aggregation in PromQL.

- Use low-cardinality labels (e.g., HTTP method, status code) to avoid cardinality explosion.
  ```go
  var (
      requestsWithLabels = promauto.NewCounterVec(
          prometheus.CounterOpts{
              Name: "myapp_http_requests_total",
              Help: "Total HTTP requests by method and status",
          },
          []string{"method", "status"},  // Label names
      )
  )

  func handleRequest(c *gin.Context) {
      requestsWithLabels.WithLabelValues(c.Request.Method, "200").Inc()
  }
  ```
- Best Practice: Limit labels to 4-5 per metric; avoid dynamic values like user IDs. Use consistent casing (snake_case) and units (e.g., `_seconds`, `_bytes`).

## Best Practices

- **Naming Conventions**: Prefix with app name (e.g., `myapp_`), use snake_case, include units in name (e.g., `_total`, `_seconds`). Always provide a `Help` string for documentation.
- **Avoid High Cardinality**: Monitor label values to prevent memory issues; use exemplars for tracing high-cardinality events.
- **Instrument Business Logic**: Add metrics at service boundaries (e.g., database queries, API calls) rather than deep in code.
  ```go
  // In a service
  func processUser(userID string) {
      timer := prometheus.NewTimer(requestDuration.WithLabelValues("process_user"))
      defer timer.ObserveDuration()
      // Business logic here
  }
  ```
- **Custom Collectors**: Implement `prometheus.Collector` for complex metrics (e.g., aggregating multiple gauges).
- **Non-Global Registry**: Use custom registries for testing or multi-tenant isolation:
  ```go
  reg := prometheus.NewRegistry()
  reg.MustRegister(myMetric)
  http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
  ```
- **Pushgateway**: For short-lived jobs (e.g., Cobra CLI), push metrics instead of exposing HTTP (use `prometheus.NewPushgateway`).
- **Integration with Zap**: Log metric values alongside errors for correlation.
  ```go
  logger.Error("Request failed",
      zap.Error(err),
      zap.Float64("duration_seconds", time.Since(start).Seconds()),
  )
  ```
- **Testing**: Use `prometheus/testutil` to assert metric values in unit tests.
  ```go
  import "github.com/prometheus/client_golang/prometheus/testutil"

  func TestCounter(t *testing.T) {
      assert.Equal(t, 5.0, testutil.ToFloat64(requestsTotal))
  }
  ```
- **Performance**: Metrics are low-overhead, but avoid excessive observations in hot paths. Use `Observe` with `time.Since` for durations.

## Integration with Gin, Viper, Cobra, JWT, and Zap

- **Gin Middleware for HTTP Metrics**:
  ```go
  func MetricsMiddleware() gin.HandlerFunc {
      return func(c *gin.Context) {
          start := time.Now()
          c.Next()
          duration := time.Since(start).Seconds()
          status := fmt.Sprintf("%d", c.Writer.Status())
          requestsTotal.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
          requestDuration.WithLabelValues(c.Request.Method).Observe(duration)
      }
  }

  // In main.go
  r.Use(MetricsMiddleware())
  ```
- **Viper Configuration**: Load metric enablement or bucket sizes from config.
  ```go
  buckets := viper.GetFloat64Slice("metrics.histogram_buckets")
  requestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
      Buckets: buckets,
  })
  ```
- **Cobra CLI Metrics**: For batch jobs, use a custom registry and push to Pushgateway on completion.
- **JWT Metrics**: Track authentication successes/failures.
  ```go
  var jwtValid = promauto.NewCounter(prometheus.CounterOpts{
      Name: "myapp_jwt_valid_total",
      Help: "Total valid JWT tokens",
  })

  // In validation
  if token.Valid {
      jwtValid.Inc()
  }
  ```
- **Zap Correlation**: Use `zap.Any("metric_value", metric.Value)` in logs.

## Prometheus Configuration

Configure Prometheus to scrape your app in `prometheus.yml`:
```yaml
global:
  scrape_interval: 15s
scrape_configs:
  - job_name: myapp
    static_configs:
      - targets: ['localhost:8080']
    metrics_path: /metrics
```

## Additional Resources

- Official client_golang GitHub: https://github.com/prometheus/client_golang
- Go Application Guide: https://prometheus.io/docs/guides/go-application/
- Package Docs: https://pkg.go.dev/github.com/prometheus/client_golang/prometheus

Adhere to these practices to build observable Go applications with Prometheus. Update this document as new features or conventions emerge.