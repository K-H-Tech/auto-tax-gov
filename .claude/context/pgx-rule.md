# PostgreSQL Best Practices Guide with pgx

This document outlines best practices for integrating PostgreSQL into Go applications using the `github.com/jackc/pgx/v5` driver and toolkit. It draws from the official `pgx` documentation, PostgreSQL best practices, the `pgx` getting started guide, and community recommendations to ensure robust, secure, and performant database interactions. These guidelines are tailored for use with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus for comprehensive Go applications. This version updates the previous `rule.md` to prioritize `pgx` over `lib/pq` due to its active development and performance advantages.

## Installation and Setup

- Install the `pgx` driver using Go modules:
  ```bash
  go get github.com/jackc/pgx/v5
  ```
- Import the driver in your code:
  ```go
  import (
      "context"
      "github.com/jackc/pgx/v5"
  )
  ```
- For `database/sql` compatibility, import the `stdlib` package:
  ```go
  import "github.com/jackc/pgx/v5/stdlib"
  ```
- Use Go 1.23 or later and PostgreSQL 13 or higher for full compatibility, as `pgx` supports the latest two Go major releases and PostgreSQL versions from the last 5 years.<grok:render type="render_inline_citation">
<argument name="citation_id">19</argument>
</grok:render>

## Connection Management

- **Direct Connection with pgx**:
  - Use `pgx.Connect` for native `pgx` connections to leverage PostgreSQL-specific features like `LISTEN/NOTIFY` and `COPY`.
  ```go
  ctx := context.Background()
  conn, err := pgx.Connect(ctx, "postgres://user:password@localhost:5432/mydb")
  if err != nil {
      log.Fatalf("Unable to connect to database: %v", err)
  }
  defer conn.Close(ctx)
  ```
- **Connection String**: Use environment variables or Viper for configuration, with `sslmode` set appropriately (e.g., `verify-full` for production).
  ```go
  connStr := viper.GetString("db.url")
  conn, err := pgx.Connect(ctx, connStr)
  ```
- **Connection Pooling**:
  - Use `pgxpool` for concurrent applications to manage multiple connections efficiently.
  ```go
  import "github.com/jackc/pgx/v5/pgxpool"

  pool, err := pgxpool.New(ctx, viper.GetString("db.url"))
  if err != nil {
      log.Fatalf("Unable to create connection pool: %v", err)
  }
  defer pool.Close()
  pool.Config().MaxConns = 25
  pool.Config().MaxConnLifetime = 5 * time.Minute
  ```
  - Set `MaxConns` based on expected load (e.g., 10-50 for moderate traffic).
  - Set `MaxConnLifetime` to recycle connections and avoid stale states.<grok:render type="render_inline_citation">
<argument name="citation_id">20</argument>
</grok:render>
- **Ping Database**: Verify connectivity after establishing a connection.
  ```go
  if err := pool.Ping(ctx); err != nil {
      log.Fatalf("Failed to ping database: %v", err)
  }
  ```
- **After-Connect Hook**: Use `AfterConnect` to set up session-specific settings (e.g., search_path).
  ```go
  pool, err := pgxpool.NewWithConfig(ctx, pgxpool.Config{
      ConnConfig: pgx.ConnConfig{Config: pgx.Config{}},
      AfterConnect: func(ctx context.Context, conn *pgx.Conn) error {
          _, err := conn.Exec(ctx, "SET search_path TO my_schema")
          return err
      },
  })
  ```

## Project Structure

Organize database-related code to promote modularity, testability, and maintainability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with DB initialization
  ├── internal/
  │   ├── db/                  # Database setup and repository logic
  │   │   ├── postgres.go      # Connection setup and queries
  │   │   └── models.go        # Data models and structs
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define database interactions in a `db` package using a repository pattern.
- Use interfaces for data access to enable mocking and testing:
  ```go
  type UserRepository interface {
      GetUserByID(ctx context.Context, id string) (*User, error)
  }
  ```

## Query Best Practices

- **Parameterized Queries**: Use `pgx`’s parameterized queries to prevent SQL injection.
  ```go
  var user User
  err := pool.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", userID).Scan(&user.ID, &user.Name)
  if err == pgx.ErrNoRows {
      return nil, nil
  }
  if err != nil {
      return nil, fmt.Errorf("query failed: %v", err)
  }
  ```
- **Context Support**: Always use `context.Context` for query cancellation and timeouts.
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  rows, err := pool.Query(ctx, "SELECT id, name FROM users WHERE active=$1", true)
  ```
- **Batch Queries**: Use `pgx.Batch` for multiple queries in a single round-trip.
  ```go
  batch := &pgx.Batch{}
  batch.Queue("INSERT INTO users(id, name) VALUES($1, $2)", "1", "Alice")
  batch.Queue("INSERT INTO users(id, name) VALUES($1, $2)", "2", "Bob")
  br := pool.SendBatch(ctx, batch)
  _, err := br.Exec()
  if err != nil {
      log.Fatalf("Batch failed: %v", err)
  }
  br.Close()
  ```
- **Transactions**: Use `pgxpool.Begin` or `pgx.Conn.Begin` for atomic operations.
  ```go
  tx, err := pool.Begin(ctx)
  if err != nil {
      return fmt.Errorf("begin transaction: %v", err)
  }
  defer tx.Rollback(ctx) // Rollback unless committed
  _, err = tx.Exec(ctx, "UPDATE users SET balance = balance - $1 WHERE id = $2", amount, userID)
  if err != nil {
      return fmt.Errorf("update failed: %v", err)
  }
  err = tx.Commit(ctx)
  ```
- **COPY Protocol**: Use `pgx.Conn.CopyFrom` for bulk data insertion.
  ```go
  rows := [][]interface{}{
      {"1", "Alice"},
      {"2", "Bob"},
  }
  _, err := pool.CopyFrom(ctx, pgx.Identifier{"users"}, []string{"id", "name"}, pgx.CopyFromRows(rows))
  if err != nil {
      log.Fatalf("Copy failed: %v", err)
  }
  ```
- **LISTEN/NOTIFY**: Use `pgx.Conn.WaitForNotification` for real-time updates.
  ```go
  _, err := conn.Exec(ctx, "LISTEN user_updates")
  notification, err := conn.WaitForNotification(ctx)
  if err == nil {
      log.Printf("Received notification: %v", notification.Payload)
  }
  ```

## Security

- **SSL/TLS**: Enable SSL with `sslmode=verify-full` for production to secure connections.
  ```go
  connStr := "postgres://user:password@localhost:5432/mydb?sslmode=verify-full&sslcert=/path/to/cert"
  ```
- **Credentials Management**: Store credentials in environment variables or a secure vault, integrated with Viper.
  ```go
  import "github.com/spf13/viper"

  connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
      viper.GetString("db.user"),
      viper.GetString("db.password"),
      viper.GetString("db.host"),
      viper.GetInt("db.port"),
      viper.GetString("db.name"),
      viper.GetString("db.sslmode"),
  )
  ```
- **Least Privilege**: Grant minimal database permissions to the application user (e.g., `SELECT`, `INSERT`, `UPDATE` on specific tables).
- **Sanitize Inputs**: Always use parameterized queries; never concatenate user input into SQL.
- **Sensitive Data**: Use PostgreSQL’s `bytea` or `pgcrypto` for encrypted storage of sensitive data (e.g., passwords with `bcrypt`).

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Repository Pattern**:
  ```go
  package db

  import (
      "context"
      "github.com/jackc/pgx/v5/pgxpool"
  )

  type PostgresRepo struct {
      Pool *pgxpool.Pool
  }

  func (r *PostgresRepo) GetUser(ctx context.Context, id string) (*User, error) {
      var user User
      err := r.Pool.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name)
      if err == pgx.ErrNoRows {
          return nil, nil
      }
      if err != nil {
          return nil, fmt.Errorf("query failed: %v", err)
      }
      return &user, err
  }

  // In handlers
  func GetUserHandler(repo *PostgresRepo) gin.HandlerFunc {
      return func(c *gin.Context) {
          user, err := repo.GetUser(c.Request.Context(), c.Param("id"))
          if err != nil {
              c.JSON(500, gin.H{"error": err.Error()})
              return
          }
          if user == nil {
              c.JSON(404, gin.H{"error": "user not found"})
              return
          }
          c.JSON(200, user)
      }
  }
  ```
- **Viper Configuration**:
  ```go
  viper.SetDefault("db.url", "postgres://user:password@localhost:5432/mydb?sslmode=disable")
  ```
  - **Config File (config.yaml)**:
    ```yaml
    db:
      url: postgres://user:password@localhost:5432/mydb?sslmode=disable
    ```
- **Cobra CLI**: Initialize the connection pool in the root command.
  ```go
  var pool *pgxpool.Pool

  var rootCmd = &cobra.Command{
      Use: "myapp",
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          var err error
          pool, err = pgxpool.New(context.Background(), viper.GetString("db.url"))
          if err != nil {
              log.Fatalf("Failed to create connection pool: %v", err)
          }
      },
      PersistentPostRun: func(cmd *cobra.Command, args []string) {
          pool.Close()
      },
  }
  ```
- **JWT Authentication**: Query user data for JWT validation.
  ```go
  func ValidateUser(repo *PostgresRepo, token *jwt.Token) error {
      claims, ok := token.Claims.(*jwt.RegisteredClaims)
      if !ok {
          return fmt.Errorf("invalid claims")
      }
      user, err := repo.GetUser(context.Background(), claims.Subject)
      if err != nil || user == nil {
          return fmt.Errorf("user not found")
      }
      return nil
  }
  ```
- **Zap Logging**: Log database operations and errors with structured fields.
  ```go
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  if err := pool.Ping(ctx); err != nil {
      logger.Error("Database ping failed", zap.Error(err), zap.String("url", viper.GetString("db.url")))
  } else {
      logger.Info("Database connected", zap.String("url", viper.GetString("db.url")))
  }
  ```
- **Prometheus Metrics**: Track database query performance and errors.
  ```go
  var (
      queryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
          Name:    "myapp_db_query_duration_seconds",
          Help:    "Duration of database queries",
          Buckets: prometheus.DefBuckets,
      }, []string{"query"})
      queryErrors = promauto.NewCounterVec(prometheus.CounterOpts{
          Name: "myapp_db_query_errors_total",
          Help: "Total database query errors",
      }, []string{"query"})
  )

  func (r *PostgresRepo) GetUser(ctx context.Context, id string) (*User, error) {
      start := time.Now()
      var user User
      err := r.Pool.QueryRow(ctx, "SELECT id, name FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name)
      queryDuration.WithLabelValues("get_user").Observe(time.Since(start).Seconds())
      if err != nil {
          queryErrors.WithLabelValues("get_user").Inc()
      }
      return &user, err
  }
  ```

## Testing

- **Unit Tests**: Mock database interactions using `pgxmock` or `pgx.Stdlib` with `sqlmock`.
  ```go
  import "github.com/jackc/pgx/v5/pgxpoolmock"

  func TestGetUser(t *testing.T) {
      mock, _ := pgxpoolmock.NewPoolMock()
      repo := &PostgresRepo{Pool: mock}
      pgxpoolmock.AddMockQuery(mock, "SELECT id, name FROM users WHERE id=$1", []interface{}{"123"}, pgxpoolmock.NewRows([]string{"id", "name"}).AddRow("123", "Alice"))
      user, err := repo.GetUser(context.Background(), "123")
      assert.NoError(t, err)
      assert.Equal(t, "Alice", user.Name)
  }
  ```
- **Integration Tests**: Use a test database or Dockerized PostgreSQL instance (e.g., with `testcontainers-go`).
  ```go
  import "github.com/testcontainers/testcontainers-go/modules/postgres"

  func TestIntegration(t *testing.T) {
      ctx := context.Background()
      container, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:15"))
      if err != nil {
          t.Fatalf("Failed to start container: %v", err)
      }
      defer container.Terminate(ctx)
      connStr, _ := container.ConnectionString(ctx, "sslmode=disable")
      pool, _ := pgxpool.New(ctx, connStr)
      // Run tests
  }
  ```
- **Metrics Testing**: Verify metric values with `prometheus/testutil`.

## Performance Tips

- **Connection Pool Tuning**: Adjust `MaxConns` and `MaxConnLifetime` based on load testing and application needs.
- **Batch Queries**: Use `pgx.Batch` to reduce round-trips for multiple operations.
- **COPY Protocol**: Prefer `CopyFrom` for bulk data loads to minimize overhead.
- **Index Optimization**: Ensure indexes on frequently queried columns (e.g., `id`, `email`) and analyze with `EXPLAIN`.
- **Statement Caching**: Leverage `pgx`’s automatic statement preparation and caching for repeated queries.
- **Binary Format**: Use binary format for custom types to improve encoding/decoding performance.

## Additional Resources

- Official pgx GitHub: https://github.com/jackc/pgx<grok:render type="render_inline_citation">
<argument name="citation_id">19</argument>
</grok:render>
- pgx Getting Started: https://github.com/jackc/pgx/wiki/Getting-Started
- PostgreSQL Documentation: https://www.postgresql.org/docs/

Adhere to these practices to ensure secure, performant, and maintainable PostgreSQL integration with `pgx`. Prefer `pgx` over `lib/pq` for active development, performance, and PostgreSQL-specific features. Update this document as new best practices or library updates emerge.