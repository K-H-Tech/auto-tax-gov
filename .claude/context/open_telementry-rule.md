# OpenTelemetry Best Practices Guide for Go

This document outlines best practices for implementing OpenTelemetry (OTel) in Go applications to enable observability through traces, metrics, and logs. It draws from the official OpenTelemetry Go documentation, instrumentation guides, and community examples to ensure effective telemetry collection, export, and analysis. OTel provides a vendor-agnostic framework for generating, collecting, and exporting telemetry data, supporting traces (distributed tracing), metrics (performance counters), and logs (structured events). These guidelines are tailored for integration with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, as well as dependencies like `pgx`, `mongo-go-driver`, `amqp091-go`, `go-redis`, `minio-go`, `sony/gobreaker`, and `avast/retry-go`. Focus on Go 1.21+ and OTel v0.50+ (stable for traces/metrics; beta for logs).

## Installation and Setup

- Install core OTel packages using Go modules:
  ```bash:disable-run
  go get go.opentelemetry.io/otel
  go get go.opentelemetry.io/otel/sdk/trace
  go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
  go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc
  go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace  # For console export (dev)
  go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp  # HTTP instrumentation
  ```
- For specific libraries:
  - Gin: `go get go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin`
  - pgx: `go get github.com/exaring/otelpgx` (community instrumentation)
  - go-redis: `go get github.com/redis/go-redis/extra/redisotel/v9`
  - RabbitMQ (amqp091-go): No official; use manual spans or community wrappers (e.g., propagate via headers).
  - MinIO: Manual spans or `go.opentelemetry.io/contrib/instrumentation/github.com/minio/minio-go/v7/otelminio`.
- Import packages:
  ```go
  import (
      "go.opentelemetry.io/otel"
      "go.opentelemetry.io/otel/attribute"
      sdktrace "go.opentelemetry.io/otel/sdk/trace"
      "go.opentelemetry.io/otel/trace"
  )
  ```
- Use Go 1.21+ for full compatibility; initialize SDK early in `main()` or a dedicated init function.

## Project Structure

Organize OTel code to promote reusability and minimal boilerplate.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with OTel SDK init
  ├── internal/
  │   ├── observability/       # OTel setup and exporters
  │   │   ├── otel.go          # SDK initialization
  │   │   └── metrics.go       # Custom metrics
  │   ├── handlers/            # Gin handlers with manual spans
  │   └── services/            # Business logic with instrumentation
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Centralize OTel init in `/internal/observability/otel.go` to avoid duplication.
- Use interfaces for traceable components (e.g., `TracedService`) for mocking with `uber-go/mock`.

## Configuration

- **SDK Initialization**: Set up TracerProvider and MeterProvider once at startup. Use OTLP for production export (e.g., to Jaeger, Zipkin, or SigNoz).
  ```go
  // internal/observability/otel.go
  package observability

  import (
      "context"
      "go.opentelemetry.io/otel"
      "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
      "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
      "go.opentelemetry.io/otel/sdk/resource"
      "go.opentelemetry.io/otel/sdk/trace"
      semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
      sdkmetric "go.opentelemetry.io/otel/sdk/metric"
  )

  func InitTracer(ctx context.Context) (shutdown func(context.Context) error, err error) {
      exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))
      if err != nil {
          return nil, err
      }
      tp := trace.NewTracerProvider(
          trace.WithBatcher(exporter),
          trace.WithResource(resource.NewWithAttributes(
              semconv.SchemaURL,
              semconv.ServiceNameKey.String("my-go-app"),
          )),
      )
      otel.SetTracerProvider(tp)

      // Metrics exporter
      mexp, _ := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint("localhost:4317"))
      mp := sdkmetric.NewMeterProvider(
          sdkmetric.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceNameKey.String("my-go-app"))),
          sdkmetric.WithReader(sdkmetric.NewPeriodicReader(mexp)),
      )
      otel.SetMeterProvider(mp)

      return func(ctx context.Context) error {
          return multierr.Combine(tp.Shutdown(ctx), mp.Shutdown(ctx))
      }, nil
  }
  ```
  - **Exporter**: OTLP/gRPC for production; stdout for dev (`go get go.opentelemetry.io/otel/exporters/stdout/stdouttrace`).
  - **Resource**: Set service name, version; use `semconv` for semantic attributes.
  - **Batching**: Use `trace.WithBatcher` for efficient export.
- **Viper Integration**: Load exporter endpoint and service name.
  ```go
  endpoint := viper.GetString("otel.endpoint")
  serviceName := viper.GetString("otel.service_name")
  ```
  - **Config File (config.yaml)**:
    ```yaml
    otel:
      endpoint: "localhost:4317"
      service_name: "my-go-app"
    ```
- **Propagation**: Set global propagator (e.g., W3C TraceContext).
  ```go
  import "go.opentelemetry.io/otel/propagation"
  otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
  ```

## Instrumentation

- **Manual Instrumentation**: Use `tracer.Start` for custom spans; add attributes and events.
  ```go
  tracer := otel.Tracer("my-service")
  _, span := tracer.Start(ctx, "process-user", trace.WithAttributes(attribute.String("user.id", "123")))
  defer span.End()
  // Business logic
  if err != nil {
      span.RecordError(err)
      span.SetStatus(codes.Error, err.Error())
  }
  ```
- **Automatic Instrumentation**: Use contrib libraries for zero-code tracing.
  - **HTTP (otelhttp)**: Wrap handlers.
    ```go
    import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    http.HandleFunc("GET /users", otelhttp.NewHandlerFunc(userHandler, "get-users"))
    ```
  - **Gin (otelgin)**: Add middleware.
    ```go
    import "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
    r.Use(otelgin.Middleware("my-gin-app"))
    ```
  - **pgx (otelpgx)**: Set tracer on connection config.
    ```go
    import "github.com/exaring/otelpgx"
    cfg, _ := pgxpool.ParseConfig(viper.GetString("db.url"))
    cfg.ConnConfig.Tracer = otelpgx.NewTracer()
    pool, _ := pgxpool.NewWithConfig(ctx, cfg)
    if err := otelpgx.RecordStats(pool); err != nil { /* handle */ }
    ```
  - **go-redis (redisotel)**: Add hook.
    ```go
    import "github.com/redis/go-redis/extra/redisotel/v9"
    rdb.AddHook(redisotel.NewTracingHook())
    ```
  - **RabbitMQ (amqp091-go)**: Manual; propagate via headers.
    ```go
    publishing := amqp.Publishing{
        Headers: amqp.Table{"traceparent": otel.GetTextMapCarrier(propagator.TextMapCarrier{})},
        Body:    payload,
    }
    ch.PublishWithContext(ctx, "", queue, false, false, publishing)
    ```
  - **MinIO**: Use contrib instrumentation or manual spans around `PutObject`/`GetObject`.
- **Metrics**: Use `Meter` for counters, gauges, histograms.
  ```go
  meter := otel.Meter("my-service")
  counter, _ := meter.Int64Counter("requests_total", metric.WithDescription("Total requests"))
  counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "GET")))
  ```
- **Logs**: Use `otelzap` bridge with Zap (beta).
  ```go
  import "go.opentelemetry.io/contrib/bridges/otelzap"
  logger := otelzap.New(logger)  // Wrap Zap logger
  logger.Info("Request processed", zap.String("trace_id", span.SpanContext().TraceID().String()))
  ```

## Best Practices

- **Semantic Conventions**: Use `semconv` attributes (e.g., `http.method`, `db.statement`) for interoperability.
- **Span Naming**: Descriptive names (e.g., "GET /users"); use `span.SetAttributes` for context.
- **Error Recording**: Always `span.RecordError(err)` and `span.SetStatus(codes.Error, msg)`.
- **Context Propagation**: Pass `ctx` everywhere; extract/inject with propagators for distributed traces.
- **Sampling**: Configure head-based sampling (e.g., 1% for production) to reduce volume.
  ```go
  trace.WithSampler(trace.TraceIDRatioBased(0.01))
  ```
- **Batching**: Enable for exporters to minimize overhead.
- **Sensitive Data**: Sanitize attributes (e.g., redact passwords) before export.
- **Zero-Code First**: Prefer auto-instrumentation (contrib libs); add manual for custom logic.
- **Performance**: OTel adds ~1-5% overhead; profile with `pprof`.
- **Logs Correlation**: Embed trace/span IDs in logs via `otelzap`.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Middleware**: As shown; integrate with JWT for user attributes.
  ```go
  r.Use(otelgin.Middleware("my-app", otelgin.WithTracerProvider(tp)))
  // In handler
  claims := c.MustGet("claims").(*jwt.RegisteredClaims)
  span.SetAttributes(attribute.String("user.id", claims.Subject))
  ```
- **Viper**: Load OTel config (endpoint, sampling rate).
- **Cobra CLI**: Init SDK in `PersistentPreRun`; trace commands.
  ```go
  rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
      shutdown, _ := observability.InitTracer(context.Background())
      defer shutdown(context.Background())
  }
  ```
- **JWT**: Add trace context to tokens if needed; trace validation.
- **Zap**: Use `otelzap` for correlated logs.
  ```go
  import "go.opentelemetry.io/contrib/bridges/otelzap"
  logger := otelzap.New(zaplogger)
  ```
- **Prometheus**: Export OTel metrics to Prometheus via `otelcol` or direct bridge.
  ```go
  go get go.opentelemetry.io/otel/exporters/prometheus
  // In InitMetrics
  exporter, _ := prometheus.New()
  reader := sdkmetric.NewPeriodicReader(exporter)
  mp := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
  ```

## Testing

- **Unit Tests**: Mock tracers with `gomock`; verify spans added/ended.
  ```go
  mockTracer := mocks.NewMockTracer(ctrl)
  mockSpan := mocks.NewMockSpan(ctrl)
  mockTracer.EXPECT().Start(gomock.Any(), "test").Return(mockSpan, trace.SpanContext{})
  ```
- **Integration Tests**: Use Jaeger all-in-one or SigNoz for trace verification.
- **Load Tests**: Monitor span volume with Prometheus.

## Performance Tips

- **Sampling**: Always sample in prod (e.g., 10-50%) to control data volume.
- **Async Export**: Use batching and periodic readers.
- **Avoid Blocking**: Export in background goroutines.

## Additional Resources

- Official OTel Go Docs: https://opentelemetry.io/docs/languages/go/
- GitHub Repo: https://github.com/open-telemetry/opentelemetry-go
- Gin Instrumentation: https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
- pgx Instrumentation: https://github.com/exaring/otelpgx
- Redis Instrumentation: https://pkg.go.dev/github.com/redis/go-redis/extra/redisotel/v9

Adhere to these practices for robust OTel instrumentation in Go. Update as OTel evolves (e.g., logs stabilization).
```