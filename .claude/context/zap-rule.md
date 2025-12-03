# Zap Best Practices Guide

This document outlines best practices for implementing structured, leveled logging in Go applications using the `go.uber.org/zap` library. It draws from the official Zap documentation, performance benchmarks, and community recommendations to ensure logging is fast, maintainable, and effective, particularly when integrated with frameworks like Gin, Viper, Cobra, and JWT. Follow these guidelines to achieve high-performance logging in your Go projects.

## Installation and Setup

- Install Zap using Go modules:
  ```bash
  go get -u go.uber.org/zap
  ```
- Import Zap in your code:
  ```go
  import "go.uber.org/zap"
  ```
- Support only the two most recent minor versions of Go to ensure compatibility and security, as Zap aligns with Go’s version release policy.

## Logger Types

Zap provides two main logger types: `Logger` for performance-critical applications and `SugaredLogger` for ease of use with a slight performance trade-off.

- **Logger**: Use for high-performance, type-safe, structured logging with minimal allocations.
  ```go
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  logger.Info("Failed to fetch URL",
      zap.String("url", "http://example.com"),
      zap.Int("attempt", 3),
      zap.Duration("backoff", time.Second),
  )
  ```
- **SugaredLogger**: Use for less performance-critical contexts where a `printf`-style API is preferred.
  ```go
  sugar := logger.Sugar()
  sugar.Infow("Failed to fetch URL",
      "url", "http://example.com",
      "attempt", 3,
      "backoff", time.Second,
  )
  sugar.Infof("Failed to fetch URL: %s", "http://example.com")
  ```

## Configuration

- **Production Config**: Use `zap.NewProduction()` for most applications, which provides JSON-structured logging with `Info` level and above, suitable for production environments.
  ```go
  logger, err := zap.NewProduction()
  if err != nil {
      log.Fatalf("Failed to initialize logger: %v", err)
  }
  defer logger.Sync()
  ```
- **Development Config**: Use `zap.NewDevelopment()` for human-readable console output during development.
  ```go
  logger, _ := zap.NewDevelopment()
  defer logger.Sync()
  ```
- **Custom Config**: Use `zap.Config` for fine-tuned control over output format, level, and fields.
  ```go
  config := zap.NewProductionConfig()
  config.OutputPaths = []string{"stdout", "app.log"}
  config.ErrorOutputPaths = []string{"stderr"}
  config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
  logger, _ := config.Build()
  defer logger.Sync()
  ```

## Logging Levels

Use appropriate logging levels to categorize log messages:

- **Debug**: Detailed information for debugging (e.g., variable states).
- **Info**: General operational information (e.g., service started).
- **Warn**: Non-critical issues that may need attention (e.g., deprecated API usage).
- **Error**: Errors that affect functionality but allow the application to continue (e.g., failed request).
- **Fatal**: Critical errors that terminate the application (use `logger.Fatal`).
  ```go
  logger.Debug("Debugging request", zap.String("url", url))
  logger.Info("Server started", zap.Int("port", 8080))
  logger.Warn("Deprecated endpoint used", zap.String("path", "/old"))
  logger.Error("Failed to process request", zap.Error(err))
  logger.Fatal("Critical failure", zap.Error(err))
  ```

## Structured Logging

- Always use structured logging with key-value pairs to enable machine-readable logs.
  ```go
  logger.Info("User login",
      zap.String("user_id", "123"),
      zap.Time("timestamp", time.Now()),
  )
  ```
- Use specific field types (e.g., `zap.String`, `zap.Int`, `zap.Duration`) for type safety and performance.
- Avoid reflection-based serialization (e.g., `encoding/json`) for fields to minimize CPU usage and allocations.

## Performance Optimization

- **Minimize Allocations**: Use `Logger` over `SugaredLogger` in hot paths to reduce allocations (e.g., 0 allocs/op for static strings vs. 1 alloc/op for `SugaredLogger`).
- **Batch Logging**: Defer `logger.Sync()` to flush buffered logs at program exit or critical points to avoid I/O overhead.
- **Sampling**: Use `zapcore.NewSampler` to reduce logging in high-frequency scenarios.
  ```go
  config := zap.NewProductionConfig()
  config.Sampling = &zap.SamplingConfig{
      Initial:    100,
      Thereafter: 100,
  }
  logger, _ := config.Build()
  ```
- **Avoid Over-Logging**: Limit debug logs in production to reduce overhead.

## Integration with Gin

Integrate Zap with Gin for structured request logging.

- **Custom Middleware**:
  ```go
  package middleware

  import (
      "time"
      "github.com/gin-gonic/gin"
      "go.uber.org/zap"
  )

  func Logger(logger *zap.Logger) gin.HandlerFunc {
      return func(c *gin.Context) {
          start := time.Now()
          path := c.Request.URL.Path
          c.Next()
          latency := time.Since(start)
          logger.Info("HTTP request",
              zap.String("method", c.Request.Method),
              zap.String("path", path),
              zap.Int("status", c.Writer.Status()),
              zap.Duration("latency", latency),
              zap.String("client_ip", c.ClientIP()),
          )
      }
  }
  ```
- **Usage in Gin**:
  ```go
  package main

  import (
      "github.com/gin-gonic/gin"
      "go.uber.org/zap"
  )

  func main() {
      logger, _ := zap.NewProduction()
      defer logger.Sync()
      r := gin.Default()
      r.Use(middleware.Logger(logger))
      r.GET("/ping", func(c *gin.Context) {
          c.JSON(200, gin.H{"message": "pong"})
      })
      r.Run(":8080")
  }
  ```

## Integration with Viper, Cobra, and JWT

- **With Viper**: Load logging configuration (e.g., log level, output paths) from Viper.
  ```go
  import (
      "go.uber.org/zap"
      "github.com/spf13/viper"
  )

  func setupLogger() *zap.Logger {
      config := zap.NewProductionConfig()
      config.Level = zap.NewAtomicLevelAt(zap.Level(viper.GetInt("log.level")))
      config.OutputPaths = viper.GetStringSlice("log.outputs")
      logger, err := config.Build()
      if err != nil {
          log.Fatalf("Failed to initialize logger: %v", err)
      }
      return logger
  }
  ```
  - **Config File (config.yaml)**:
    ```yaml
    log:
      level: 0  # zap.InfoLevel
      outputs: ["stdout", "app.log"]
    ```
- **With Cobra**: Initialize the logger in the root command and pass it to subcommands.
  ```go
  var logger *zap.Logger

  var rootCmd = &cobra.Command{
      Use: "myapp",
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          logger, _ = zap.NewProduction()
      },
      Run: func(cmd *cobra.Command, args []string) {
          logger.Info("Root command executed")
      },
  }
  ```
- **With JWT**: Log authentication events (e.g., token validation failures).
  ```go
  token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
  if err != nil {
      logger.Error("Token validation failed",
          zap.Error(err),
          zap.String("token", tokenString),
      )
      return
  }
  logger.Info("Token validated", zap.String("user_id", claims.Subject))
  ```

## Error Handling

- Use `zap.Error` for logging errors to include stack traces or context.
- Reserve `logger.Fatal` for unrecoverable errors that require program termination.
- Check logger initialization errors and fallback to standard `log` if needed.
  ```go
  logger, err := zap.NewProduction()
  if err != nil {
      log.Fatalf("Failed to initialize logger: %v", err)
  }
  ```

## Testing

- **Mock Logger**: Use `zap.NewNop()` for a no-op logger in tests to avoid logging overhead.
  ```go
  logger := zap.NewNop()
  ```
- **Capture Output**: Test log output by redirecting to a buffer.
  ```go
  func TestLogging(t *testing.T) {
      var buf bytes.Buffer
      config := zap.NewProductionConfig()
      config.OutputPaths = []string{"stdout"}
      logger, _ := config.Build(zap.WrapCore(func(c zapcore.Core) zapcore.Core {
          return zapcore.NewCore(
              zapcore.NewJSONEncoder(config.EncoderConfig),
              zapcore.AddSync(&buf),
              config.Level,
          )
      }))
      logger.Info("Test log", zap.String("key", "value"))
      assert.Contains(t, buf.String(), `"key":"value"`)
  }
  ```
- **Test Levels**: Verify that logs are emitted at the correct level using `zapcore.LevelEnabler`.

## Security

- Avoid logging sensitive data (e.g., JWT tokens, passwords) directly. Use redaction or custom fields.
  ```go
  logger.Info("User action", zap.String("user_id", userID)) // Safe
  logger.Info("Token", zap.Skip()) // Avoid logging raw token
  ```
- Use `zap.Any` sparingly to prevent accidental exposure of sensitive data via reflection.

## Additional Resources

- Official Zap GitHub: https://github.com/uber-go/zap
- Go Package Docs: https://pkg.go.dev/go.uber.org/zap
- FAQ: https://github.com/uber-go/zap/blob/master/FAQ.md

Adhere to these practices to ensure high-performance, structured logging with Zap. Update this document as new best practices or library updates emerge.