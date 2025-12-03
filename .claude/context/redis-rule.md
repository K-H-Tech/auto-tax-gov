# Redis Best Practices Guide with go-redis

This document outlines best practices for integrating Redis into Go applications using the `github.com/redis/go-redis/v9` client library. It draws from the official `go-redis` documentation, Redis best practices, community recommendations, and the Redis documentation to ensure robust, secure, and performant key-value store interactions. These guidelines are tailored for integration with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus for comprehensive Go applications. The library supports Redis 7.2, 7.4, and 8.0, with Go 1.18+ (preferably 1.23+ for testing).

## Installation and Setup

- Install the `go-redis` client using Go modules:
  ```bash
  go get github.com/redis/go-redis/v9
  ```
- Import the necessary packages:
  ```go
  import (
      "context"
      "github.com/redis/go-redis/v9"
  )
  ```
- Use Go 1.23+ for testing compatibility and Redis 7.2+ for full feature support, including modules like RediSearch.<grok:render type="render_inline_citation">
<argument name="citation_id">21</argument>
</grok:render>

## Connection Management

- **Client Creation**: Initialize a client with `redis.NewClient` and connection options. Always close the client to release resources.
  ```go
  ctx := context.Background()
  rdb := redis.NewClient(&redis.Options{
      Addr:     "localhost:6379",
      Password: "", // No password by default
      DB:       0,  // Default DB
  })
  defer rdb.Close()
  ```
- **Connection String**: Use Viper for configuration or environment variables.
  ```go
  rdb := redis.NewClient(&redis.Options{
      Addr: viper.GetString("redis.addr"),
      Password: viper.GetString("redis.password"),
      DB: viper.GetInt("redis.db"),
  })
  ```
- **Ping for Connectivity**: Verify server reachability after connecting.
  ```go
  ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
  defer cancel()
  if err := rdb.Ping(ctx).Err(); err != nil {
      log.Fatalf("Failed to ping Redis: %v", err)
  }
  ```
- **Connection Pooling**: Configure pool settings for concurrency.
  ```go
  rdb := redis.NewClient(&redis.Options{
      Addr:         "localhost:6379",
      PoolSize:     50,              // Max connections
      MinIdleConns: 10,              // Minimum idle connections
      MaxConnAge:   30 * time.Minute, // Connection lifetime
      PoolTimeout:  5 * time.Second,  // Timeout for acquiring connection
  })
  ```
  - Set `PoolSize` based on load (e.g., 10-100 for moderate traffic).
  - Use `MinIdleConns` to keep connections ready.
  - Set `MaxConnAge` to recycle connections.
- **Authentication**: Prioritize secure credential providers (e.g., `StreamingCredentialsProvider` for dynamic credentials like Entra ID).
  ```go
  import "github.com/redis/go-redis-entraid"

  provider := entraid.NewDefaultAzureIdentityProvider()
  rdb := redis.NewClient(&redis.Options{
      Addr:                        "redis-server.redis.cache.windows.net:6380",
      StreamingCredentialsProvider: provider,
      TLSConfig:                   &tls.Config{MinVersion: tls.VersionTLS12},
  })
  ```
  - Fallback to `CredentialsProviderContext`, `CredentialsProvider`, or static `Username/Password` if needed.

## Project Structure

Organize Redis-related code for modularity and testability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with Redis initialization
  ├── internal/
  │   ├── redis/               # Redis client setup and operations
  │   │   ├── redis.go         # Client setup and repository logic
  │   │   └── models.go        # Data models
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define Redis operations in a `redis` package using a repository pattern.
- Use interfaces for data access to enable mocking:
  ```go
  type CacheRepository interface {
      Get(ctx context.Context, key string) (string, error)
      Set(ctx context.Context, key, value string, ttl time.Duration) error
  }
  ```

## Command Best Practices

- **Basic Operations**: Use `Set`, `Get`, and handle `redis.Nil` for missing keys.
  ```go
  err := rdb.Set(ctx, "key", "value", 1*time.Hour).Err()
  if err != nil {
      return err
  }
  val, err := rdb.Get(ctx, "key").Result()
  if err == redis.Nil {
      return fmt.Errorf("key does not exist")
  }
  ```
- **Context with Timeouts**: Always use `context.WithTimeout` for operations.
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  ```
- **Pipelines**: Use `Pipeline` or `TxPipeline` for batch operations to reduce round-trips.
  ```go
  pipe := rdb.Pipeline()
  pipe.Set(ctx, "key1", "value1", 0)
  pipe.Set(ctx, "key2", "value2", 0)
  _, err := pipe.Exec(ctx)
  ```
- **Transactions**: Use `TxPipeline` with `WATCH` for optimistic locking.
  ```go
  err := rdb.Watch(ctx, func(tx *redis.Tx) error {
      val, _ := tx.Get(ctx, "key").Int()
      _, err := tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
          pipe.Set(ctx, "key", val+1, 0)
          return nil
      })
      return err
  }, "key")
  ```
- **RediSearch (Optional)**: Use `FTSearchWithArgs` for search; prefer RESP2 for stability until RESP3 matures.
  ```go
  res, err := rdb.FTSearchWithArgs(ctx, "idx:users", "@name:alice", &redis.FTSearchOptions{}).Result()
  ```
  - Set `DialectVersion` (e.g., 2 or 3) if needed for specific queries.

## Security

- **TLS**: Enable TLS for production with `TLSConfig`.
  ```go
  rdb := redis.NewClient(&redis.Options{
      Addr:      "redis-server:6380",
      TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12},
  })
  ```
- **Credentials**: Use Viper for secure credential management; avoid hardcoding.
  ```go
  rdb := redis.NewClient(&redis.Options{
      Addr:     viper.GetString("redis.addr"),
      Password: viper.GetString("redis.password"),
  })
  ```
- **Access Control**: Use Redis ACLs to restrict commands and keys per user.
- **Data Sensitivity**: Avoid storing sensitive data (e.g., JWT tokens) unless encrypted.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Cache Middleware**:
  ```go
  package redis

  type CacheRepo struct {
      Client *redis.Client
  }

  func (r *CacheRepo) Get(ctx context.Context, key string) (string, error) {
      return r.Client.Get(ctx, key).Result()
  }

  // In handlers
  func CacheHandler(repo *CacheRepo) gin.HandlerFunc {
      return func(c *gin.Context) {
          val, err := repo.Get(c.Request.Context(), "key")
          if err == redis.Nil {
              c.JSON(404, gin.H{"error": "key not found"})
              return
          }
          c.JSON(200, gin.H{"value": val})
      }
  }
  ```
- **Viper Configuration**:
  ```go
  viper.SetDefault("redis.addr", "localhost:6379")
  viper.SetDefault("redis.db", 0)
  ```
  - **Config File (config.yaml)**:
    ```yaml
    redis:
      addr: localhost:6379
      password: ""
      db: 0
    ```
- **Cobra CLI**: Initialize Redis in the root command.
  ```go
  var rdb *redis.Client

  var rootCmd = &cobra.Command{
      Use: "myapp",
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          rdb = redis.NewClient(&redis.Options{
              Addr:     viper.GetString("redis.addr"),
              Password: viper.GetString("redis.password"),
              DB:       viper.GetInt("redis.db"),
          })
          if err := rdb.Ping(context.Background()).Err(); err != nil {
              log.Fatalf("Redis ping failed: %v", err)
          }
      },
      PersistentPostRun: func(cmd *cobra.Command, args []string) {
          rdb.Close()
      },
  }
  ```
- **JWT Caching**: Cache user sessions in Redis.
  ```go
  func CacheSession(ctx context.Context, rdb *redis.Client, userID string, token string) error {
      return rdb.Set(ctx, "session:"+userID, token, 24*time.Hour).Err()
  }
  ```
- **Zap Logging**: Log Redis operations.
  ```go
  logger.Info("Redis set", zap.String("key", "key"), zap.String("value", "value"))
  if err != nil {
      logger.Error("Redis operation failed", zap.Error(err))
  }
  ```
- **Prometheus Metrics**: Track Redis command latency and errors.
  ```go
  var (
      redisDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
          Name:    "myapp_redis_command_duration_seconds",
          Help:    "Duration of Redis commands",
          Buckets: prometheus.DefBuckets,
      }, []string{"command"})
  )

  func (r *CacheRepo) Get(ctx context.Context, key string) (string, error) {
      start := time.Now()
      defer func() { redisDuration.WithLabelValues("get").Observe(time.Since(start).Seconds()) }()
      return r.Client.Get(ctx, key).Result()
  }
  ```

## Testing

- **Unit Tests**: Mock Redis with `miniredis` or similar libraries.
  ```go
  import "github.com/alicebob/miniredis/v2"

  func TestCache(t *testing.T) {
      mr, _ := miniredis.Run()
      defer mr.Close()
      rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
      repo := &CacheRepo{Client: rdb}
      repo.Set(context.Background(), "key", "value", 0)
      val, err := repo.Get(context.Background(), "key")
      assert.NoError(t, err)
      assert.Equal(t, "value", val)
  }
  ```
- **Integration Tests**: Use Dockerized Redis via `testcontainers-go`.
  ```go
  import "github.com/testcontainers/testcontainers-go/modules/redis"

  func TestIntegration(t *testing.T) {
      ctx := context.Background()
      redisContainer, err := redis.Run(ctx, "redis:7.2")
      if err != nil {
          t.Fatalf("Failed to start Redis container: %v", err)
      }
      defer redisContainer.Terminate(ctx)
      uri, _ := redisContainer.ConnectionString(ctx)
      rdb := redis.NewClient(&redis.Options{Addr: uri})
      // Run tests
  }
  ```

## Performance Tips

- **Pipelining**: Use `Pipeline` for batch operations to minimize latency.
- **Connection Pool Tuning**: Adjust `PoolSize` and `MinIdleConns` based on load.
- **Command Optimization**: Use `MGET`/`MSET` for multiple keys.
- **TTL**: Set expiration (`SetEX`) for transient data to manage memory.
- **Redis Modules**: Leverage RediSearch for search but stick to RESP2 until RESP3 stabilizes.

## Additional Resources

- Official go-redis GitHub: https://github.com/redis/go-redis
- Redis Documentation: https://redis.io/docs/
- Redis University: https://university.redis.com/

Adhere to these practices for secure and efficient Redis integration with `go-redis`. Update this document as new features or best practices emerge.