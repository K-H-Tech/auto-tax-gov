# RabbitMQ Best Practices Guide with amqp091-go

This document outlines best practices for integrating RabbitMQ into Go applications using the `github.com/rabbitmq/amqp091-go` client library. It draws from the official RabbitMQ tutorials, reliability guides, monitoring documentation, and community recommendations (e.g., reconnection strategies and higher-level wrappers) to ensure robust, secure, and performant messaging. These guidelines are tailored for use with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus for comprehensive Go applications. The library targets AMQP 0-9-1 and is optimized for RabbitMQ servers (versions 2.0+, primarily tested on recent releases).

## Installation and Setup

- Install the `amqp091-go` client using Go modules:
  ```bash:disable-run
  go get github.com/rabbitmq/amqp091-go
  ```
- Import the library (use an alias if migrating from `streadway/amqp`):
  ```go
  import (
      amqp "github.com/rabbitmq/amqp091-go"
  )
  ```
- Supports the two most recent Go releases. Use Go 1.21+ for optimal compatibility.
- For higher-level abstractions (e.g., reconnection logic), consider community wrappers like `github.com/wagslane/go-rabbitmq` as a thin layer over `amqp091-go`, but implement custom reconnection for production to avoid undefined behavior.

## Connection Management

- **Basic Connection**: Dial with a URI including credentials, host, port, and vhost (default "/").
  ```go
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
  if err != nil {
      log.Fatalf("Failed to connect: %v", err)
  }
  defer conn.Close()
  ```
- **Channels**: Use channels for operations (one per logical unit of work). Limit to 1-2 per connection for low-latency; more for high-throughput.
  ```go
  ch, err := conn.Channel()
  if err != nil {
      log.Fatalf("Failed to open channel: %v", err)
  }
  defer ch.Close()
  ```
- **Reconnection Strategy**: The library does not auto-reconnect. Implement a goroutine loop with exponential backoff for production reliability.
  ```go
  func connectWithReconnect(ctx context.Context, uri string) (*amqp.Connection, error) {
      var conn *amqp.Connection
      for {
          select {
          case <-ctx.Done():
              return nil, ctx.Err()
          default:
              var err error
              conn, err = amqp.Dial(uri)
              if err == nil {
                  return conn, nil
              }
              log.Printf("Connection failed: %v. Retrying in 5s...", err)
              time.Sleep(5 * time.Second)
          }
      }
  }
  ```
  - Monitor `conn.NotifyClose(make(chan *amqp.Error))` for failures and trigger reconnection.
  - Redeclare queues/exchanges/bindings on reconnect to synchronize topology.
- **Timeouts**: Use `context.WithTimeout` for publish/consume operations.
- **Heartbeats**: Set via URI (`?heartbeat=10`) or `amqp.Config` to detect dead connections faster than TCP defaults.

## Project Structure

Organize RabbitMQ code for modularity and testability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with RabbitMQ setup
  ├── internal/
  │   ├── messaging/           # RabbitMQ producers/consumers
  │   │   ├── producer.go      # Publishing logic
  │   │   └── consumer.go      # Consumption logic
  │   ├── handlers/            # Gin handlers publishing/consuming
  │   └── services/            # Business logic integrating messaging
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Use a `messaging` package with interfaces for producers/consumers to enable mocking.
  ```go
  type Producer interface {
      Publish(ctx context.Context, msg []byte) error
  }
  ```

## Publishing Messages

- **Basic Publish**: Declare queues/exchanges, then publish with context.
  ```go
  // From Hello World tutorial
  q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
      ContentType: "text/plain",
      Body:        []byte("Hello World!"),
  })
  ```
- **Durable Messages**: Set `DeliveryMode: amqp.Persistent` for persistence across restarts.
- **Publisher Confirms**: Enable with `ch.Confirm(false)` and wait for acks via `ch.WaitForConfirms()`. Use `mandatory` flag for unroutable messages.
  ```go
  err = ch.Confirm(false)
  if err != nil {
      // Handle error
  }
  err = ch.PublishWithContext(ctx, "", q.Name, true, false, publishing)
  if err != nil {
      // Handle publish error
  }
  confirmed := ch.WaitForConfirms(time.Second) // Wait for ack/nack
  if !confirmed {
      // Retransmit
  }
  ```
- **Exchanges**: Use appropriate types (direct, topic, fanout) and bind queues.
  - Fanout (Pub/Sub): `ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)`
  - Topic: `ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil)`; bind with wildcards (e.g., "*.info").

## Consuming Messages

- **Basic Consume**: Use `ch.Consume` with auto-ack for simple cases; manual for reliability.
  ```go
  // From Hello World tutorial
  msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
  for d := range msgs {
      log.Printf("Received: %s", d.Body)
  }
  ```
- **Manual Acknowledgments**: Set `autoAck: false` and ack after processing to ensure at-least-once delivery.
  ```go
  // From Work Queues tutorial
  err = ch.Qos(1, 0, false) // Prefetch 1 for fair dispatch
  msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
  for d := range msgs {
      // Process d.Body
      d.Ack(false) // Manual ack
  }
  ```
- **Prefetch**: Set `ch.Qos(prefetchCount, 0, false)` to limit unacked messages (e.g., 1 for sequential processing).
- **Error Handling**: Handle `d.Nack(false, true)` for requeue on failure; use dead-letter exchanges for poison messages.

## Reliability and Error Handling

- **At-Least-Once Delivery**: Combine durable queues (`durable: true`), persistent messages, manual acks, and publisher confirms.
- **Idempotency**: Design consumers to handle duplicates (check `d.Redelivered` flag).
- **Dead-Letter Queues**: Configure via queue arguments (`amqp.Table{"x-dead-letter-exchange": "dlx"}`).
- **Failure Recovery**: On channel/connection close, log errors and reconnect. Use `conn.NotifyClose` and `ch.NotifyClose` channels.
- **Timeouts and Contexts**: Always use timeouts to prevent hanging operations.
- **Testing**: Use `github.com/rabbitmq/amqp091-go` test utilities or mock with `testcontainers-go` for integration tests.

## Security

- **TLS**: Enable with URI (`amqps://`) or custom dialer.
  ```go
  config := amqp.Config{
      TLSClientConfig: &tls.Config{
          InsecureSkipVerify: false,
          RootCAs:            caCertPool,
      },
  }
  conn, err := amqp.DialTLSConfig("amqps://localhost:5671/", config)
  ```
- **Credentials**: Use Viper for secrets; support PLAIN/EXTERNAL auth. Avoid defaults in production.
- **VHost Permissions**: Restrict users to specific vhosts with read/write/configure permissions.
- **Network**: Bind RabbitMQ to localhost or use VPN; enable firewall rules.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Publisher Middleware**: Publish events from handlers.
  ```go
  // In a handler
  func OrderCreatedHandler(c *gin.Context) {
      // Process order
      err := producer.Publish(c.Request.Context(), orderBytes)
      if err != nil {
          c.JSON(500, gin.H{"error": "Publish failed"})
          return
      }
      c.JSON(200, gin.H{"status": "published"})
  }
  ```
- **Viper Configuration**:
  ```go
  viper.SetDefault("rabbitmq.uri", "amqp://guest:guest@localhost:5672/")
  uri := viper.GetString("rabbitmq.uri")
  ```
  - **Config File (config.yaml)**:
    ```yaml
    rabbitmq:
      uri: amqp://user:pass@host:5672/vhost
      prefetch: 1
    ```
- **Cobra CLI**: Initialize in root command for batch jobs.
  ```go
  var conn *amqp.Connection
  var rootCmd = &cobra.Command{
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          var err error
          conn, err = amqp.Dial(viper.GetString("rabbitmq.uri"))
          // Handle error
      },
      PersistentPostRun: func(cmd *cobra.Command, args []string) {
          conn.Close()
      },
  }
  ```
- **JWT**: Publish authenticated events with user claims in headers.
  ```go
  publishing := amqp.Publishing{
      Headers: amqp.Table{"user_id": claims.Subject},
      Body:    payload,
  }
  ```
- **Zap Logging**: Log deliveries and errors.
  ```go
  logger.Info("Message published", zap.String("queue", q.Name), zap.Int("size", len(body)))
  logger.Error("Consume failed", zap.Error(err))
  ```
- **Prometheus Metrics**: Enable `rabbitmq_prometheus` plugin on server (port 15692). Scrape every 30s; monitor queue lengths, rates.
  ```go
  // Custom metrics
  var publishDuration = promauto.NewHistogram(prometheus.HistogramOpts{
      Name: "myapp_rabbitmq_publish_duration_seconds",
      Help: "Duration of RabbitMQ publishes",
  })
  // In publish func
  start := time.Now()
  defer func() { publishDuration.Observe(time.Since(start).Seconds()) }()
  ```

## Examples

### Hello World (Basic Queue)
See sending/receiving code from tutorials .

### Work Queues (with Acks and Prefetch)
See producer/worker from .

### Pub/Sub (Fanout Exchange)
See emit_log/receive_logs from .

### Routing (Direct Exchange)
See emit_log_direct/receive_logs_direct from .

### Topics (Wildcard Routing)
See emit_log_topic/receive_logs_topic from .

### RPC (Request-Reply)
See rpc_client/rpc_server from .

## Testing

- **Unit Tests**: Mock `amqp.Connection`/`Channel` interfaces.
- **Integration Tests**: Use Dockerized RabbitMQ via `testcontainers-go`.
  ```go
  import "github.com/testcontainers/testcontainers-go/modules/rabbitmq"
  // In test
  ctx := context.Background()
  rabbitmqContainer, err := rabbitmq.Run(ctx, "rabbitmq:3-management")
  uri, _ := rabbitmqContainer.AmqpURL(ctx)
  conn, _ := amqp.Dial(uri.String())
  ```

## Performance Tips

- **Prefetch**: Set low (1-10) for latency-sensitive; higher for throughput.
- **Batch Publishes**: Use transactions (`ch.Tx()`) for small batches; avoid for high-volume.
- **Quorum Queues**: Prefer over classic for high availability (RabbitMQ 3.8+).
- **Monitoring**: Scrape Prometheus endpoint; alert on high queue lengths (>1000).

## Additional Resources

- Official amqp091-go GitHub: https://github.com/rabbitmq/amqp091-go
- RabbitMQ Tutorials: https://www.rabbitmq.com/tutorials/tutorial-one-go.html
- Reliability Guide: https://www.rabbitmq.com/docs/reliability
- Monitoring: https://www.rabbitmq.com/docs/monitoring
- Community Wrapper: https://github.com/wagslane/go-rabbitmq

Adhere to these practices for reliable RabbitMQ integration. Update as new features emerge.
```