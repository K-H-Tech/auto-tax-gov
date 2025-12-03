# Go Design Patterns Best Practices Guide

This document outlines best practices for applying design patterns in Go applications, focusing on idiomatic Go principles. It draws from Go’s official documentation, community resources like Dave Cheney’s talks (e.g., "Practical Go" at GopherCon 2016), Russ Cox’s writings on Go simplicity, and widely used patterns in the Go ecosystem. Go favors simplicity, explicitness, and composition over inheritance, so patterns are adapted to avoid over-abstraction and align with Go’s philosophy ("Less is exponentially more"). This guide integrates with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, and dependencies like `pgx`, `mongo-go-driver`, `amqp091-go`, `go-redis`, `sony/gobreaker`, `avast/retry-go`, and `ulule/limiter`. It emphasizes patterns for web services, CLIs, and libraries, suitable for Go 1.21+.

## Core Principles

- **Simplicity First**: Avoid premature abstraction; prefer concrete code until complexity demands patterns.
- **Interfaces Over Inheritance**: Use small, focused interfaces for loose coupling.
- **Composition**: Leverage structs and embedding for behavior reuse, not inheritance.
- **Explicit Errors**: Use error returns over exceptions; handle errors explicitly.
- **Concurrency**: Use goroutines and channels for safe, idiomatic concurrency.
- **Testability**: Design for testing with interfaces and dependency injection.
- **Avoid Overuse**: Patterns like Singleton or Factory are often anti-patterns in Go; use only when justified.

## Recommended Patterns

### 1. Repository Pattern (Data Access Layer)
Encapsulates data access logic, decoupling business logic from storage (e.g., `pgx`, `mongo-go-driver`, `go-redis`).

- **Use Case**: Database operations (SQL/NoSQL), caching.
- **Implementation**:
  ```go
  // internal/app/repository/user.go
  type UserRepository interface {
      GetUser(ctx context.Context, id string) (*User, error)
      SaveUser(ctx context.Context, user *User) error
  }

  type PostgresUserRepo struct {
      db *pgxpool.Pool
  }

  func NewPostgresUserRepo(db *pgxpool.Pool) UserRepository {
      return &PostgresUserRepo{db: db}
  }

  func (r *PostgresUserRepo) GetUser(ctx context.Context, id string) (*User, error) {
      var user User
      err := r.db.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", id).
          Scan(&user.ID, &user.Name)
      if errors.Is(err, pgx.ErrNoRows) {
          return nil, nil
      }
      return &user, err
  }
  ```
- **Benefits**: Testable (mock with `uber-go/mock`), swappable (e.g., Postgres to Mongo), hides DB specifics.
- **Integration**: Use with `sony/gobreaker` for resilience, `avast/retry-go` for retries.
  ```go
  func (r *PostgresUserRepo) GetUser(ctx context.Context, id string) (*User, error) {
      return retry.DoWithData(func() (*User, error) {
          return r.getUser(ctx, id) // Internal method
      }, retry.RetryIf(func(err error) bool { return !errors.Is(err, pgx.ErrNoRows) }))
  }
  ```

### 2. Dependency Injection (DI)
Pass dependencies explicitly via constructors or interfaces, avoiding global state.

- **Use Case**: Inject DB, logger, or metrics clients into services/handlers.
- **Implementation**:
  ```go
  // internal/app/service/user.go
  type UserService struct {
      repo UserRepository
      log  *zap.Logger
  }

  func NewUserService(repo UserRepository, log *zap.Logger) *UserService {
      return &UserService{repo: repo, log: log}
  }

  func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
      s.log.Info("Fetching user", zap.String("id", id))
      return s.repo.GetUser(ctx, id)
  }
  ```
- **Benefits**: Testable, avoids globals (anti-pattern in Go), promotes loose coupling.
- **Integration**: Wire in `cmd/api/main.go` with Viper configs.
  ```go
  db := pgxpool.New(ctx, viper.GetString("db.url"))
  repo := repository.NewPostgresUserRepo(db)
  log := zap.NewProduction()
  svc := service.NewUserService(repo, log)
  ```

### 3. Middleware Pattern (Chain of Responsibility)
Chain handlers for HTTP or message processing (e.g., Gin middleware, RabbitMQ consumers).

- **Use Case**: Authentication, rate limiting (`ulule/limiter`), logging, circuit breaking.
- **Implementation**:
  ```go
  // internal/middleware/ratelimit.go
  func RateLimit(limiter *limiter.Limiter) gin.HandlerFunc {
      return func(c *gin.Context) {
          ctx := c.Request.Context()
          key := c.ClientIP()
          if _, err := limiter.Get(ctx, key); err != nil {
              c.JSON(429, gin.H{"error": "Rate limit exceeded"})
              c.Abort()
              return
          }
          c.Next()
      }
  }

  // cmd/api/main.go
  r := gin.Default()
  store := limiterDrivers.NewStore()
  rate, _ := limiter.NewRateFromFormatted(viper.GetString("ratelimit.rate"))
  r.Use(middleware.RateLimit(limiter.New(store, rate)))
  ```
- **Benefits**: Modular, reusable, stackable (e.g., JWT auth → logging → rate limit).
- **Integration**: Combine with `sony/gobreaker` for external calls, `zap` for logging, `prometheus` for metrics.

### 4. Strategy Pattern (Interface-Based Behavior)
Define interchangeable algorithms via interfaces (e.g., different DB backends, retry strategies).

- **Use Case**: Swap `pgx` for `mongo-go-driver`, or `retry.BackOffDelay` for `retry.FixedDelay`.
- **Implementation**:
  ```go
  // internal/pkg/db/db.go
  type Database interface {
      QueryUser(ctx context.Context, id string) (*User, error)
  }

  type PostgresDB struct { /* pgxpool.Pool */ }
  type MongoDB struct { /* mongo.Collection */ }

  func NewDB(driver string) Database {
      switch driver {
      case "postgres":
          return &PostgresDB{}
      case "mongo":
          return &MongoDB{}
      default:
          panic("unknown driver")
      }
  }
  ```
- **Benefits**: Extensible, testable, avoids tight coupling.
- **Integration**: Configure via Viper (`driver: postgres`); mock with `uber-go/mock`.

### 5. Observer Pattern (Event Notification)
Use channels or callbacks for pub/sub (e.g., RabbitMQ with `amqp091-go`, Prometheus metrics).

- **Use Case**: Notify components of state changes (e.g., circuit breaker state, cache updates).
- **Implementation**:
  ```go
  // internal/pkg/resilience/breaker.go
  type BreakerObserver interface {
      OnStateChange(name string, from, to gobreaker.State)
  }

  type MetricsObserver struct {
      gauge *prometheus.GaugeVec
  }

  func (o *MetricsObserver) OnStateChange(name string, _, to gobreaker.State) {
      o.gauge.WithLabelValues(name).Set(float64(to))
  }

  func NewBreaker(observers ...BreakerObserver) *gobreaker.CircuitBreaker[[]byte] {
      return gobreaker.NewCircuitBreaker[[]byte](gobreaker.Settings{
          OnStateChange: func(name string, from, to gobreaker.State) {
              for _, o := range observers {
                  o.OnStateChange(name, from, to)
              }
          },
      })
  }
  ```
- **Benefits**: Decouples event producers and consumers; scalable for monitoring.
- **Integration**: Use `amqp091-go` for distributed events; `zap` for logging.

### 6. Decorator Pattern (Behavior Extension)
Wrap components to add functionality (e.g., retries, metrics, logging).

- **Use Case**: Add retries (`avast/retry-go`) or metrics to HTTP/DB calls.
- **Implementation**:
  ```go
  // internal/pkg/resilience/retry.go
  type RetryableRepo struct {
      repo  UserRepository
      retry *retry.Config
  }

  func NewRetryableRepo(repo UserRepository) UserRepository {
      return &RetryableRepo{
          repo:  repo,
          retry: &retry.Config{Attempts: 3, Delay: 500 * time.Millisecond},
      }
  }

  func (r *RetryableRepo) GetUser(ctx context.Context, id string) (*User, error) {
      return retry.DoWithData(func() (*User, error) {
          return r.repo.GetUser(ctx, id)
      }, r.retry.Context(ctx), r.retry.Attempts(3))
  }
  ```
- **Benefits**: Non-invasive extension; reusable across components.
- **Integration**: Combine with `sony/gobreaker` for layered resilience.

### 7. Worker Pool Pattern (Concurrency)
Distribute tasks across goroutines with channels for scalability.

- **Use Case**: Process RabbitMQ messages (`amqp091-go`) or batch DB operations.
- **Implementation**:
  ```go
  // internal/pkg/worker/pool.go
  type WorkerPool struct {
      tasks chan Task
  }

  type Task struct { /* payload */ }

  func NewWorkerPool(workers int) *WorkerPool {
      pool := &WorkerPool{tasks: make(chan Task)}
      for i := 0; i < workers; i++ {
          go func() {
              for task := range pool.tasks {
                  // Process task
              }
          }()
      }
      return pool
  }

  func (p *WorkerPool) Submit(task Task) { p.tasks <- task }
  ```
- **Benefits**: Controls concurrency; prevents resource exhaustion.
- **Integration**: Use with `amqp091-go` consumers; monitor with Prometheus.

## Best Practices

- **Keep Interfaces Small**: 1-3 methods; avoid "fat" interfaces.
- **Avoid Singletons**: Use DI or package-level vars with lazy init; avoid global state.
- **Idempotency**: Ensure retries (`avast/retry-go`) and handlers are idempotent.
- **Error Handling**: Wrap errors with context (`fmt.Errorf("op: %w", err)`); use `errors.Is`/`errors.As`.
- **Testing**: Mock dependencies with `uber-go/mock`; test failure paths.
- **Metrics/Logging**: Use Prometheus for metrics, Zap for structured logging in all patterns.
- **Avoid Overuse**: Patterns like Factory or Abstract Factory are rarely needed; prefer functions.
- **Context Propagation**: Pass `context.Context` for cancellation and tracing.

## Integration Example (Gin Web Service)

```go
// cmd/api/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "myproject/internal/app/handlers"
    "myproject/internal/app/repository"
    "myproject/internal/app/service"
    "myproject/internal/pkg/db"
    "github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("configs/config.yaml")
    viper.ReadInConfig()

    log, _ := zap.NewProduction()
    dbConn := db.NewPostgres(viper.GetString("db.url"))
    repo := repository.NewPostgresUserRepo(dbConn)
    svc := service.NewUserService(repo, log)
    r := gin.Default()
    handlers.RegisterUserRoutes(r, svc)
    r.Run(viper.GetString("server.addr"))
}
```

## Additional Resources

- Go Proverbs: https://go-proverbs.github.io/< type="render_inline_citation"name="citation_id">2
- Effective Go: https://go.dev/doc/effective_go< type="render_inline_citation"name="citation_id">1
- GopherCon Talks: Dave Cheney’s "Practical Go" (2016), Peter Bourgon’s "Industrial Go" (2018)< type="render_inline_citation"name="citation_id">7
- Go Patterns: https://github.com/tmrts/go-patterns< type="render_inline_citation"name="citation_id">10

Adhere to these practices for idiomatic Go design patterns. Update as new patterns or tools emerge.