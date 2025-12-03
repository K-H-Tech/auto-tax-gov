# MongoDB Best Practices Guide with mongo-go-driver

This document outlines best practices for integrating MongoDB into Go applications using the official `go.mongodb.org/mongo-driver/v2` library. It draws from the official MongoDB Go driver documentation, MongoDB best practices, and community recommendations to ensure robust, secure, and performant database interactions. These guidelines are tailored for use with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus for comprehensive Go applications. Follow these practices to leverage MongoDB's document-oriented model effectively in Go.

## Installation and Setup

- Install the MongoDB Go driver using Go modules:
  ```bash
  go get go.mongodb.org/mongo-driver/v2/mongo
  ```
- Import the necessary packages:
  ```go
  import (
      "context"
      "go.mongodb.org/mongo-driver/v2/bson"
      "go.mongodb.org/mongo-driver/v2/mongo"
      "go.mongodb.org/mongo-driver/v2/mongo/options"
      "go.mongodb.org/mongo-driver/v2/mongo/readpref"
  )
  ```
- Use Go 1.19 or higher (Go 1.23+ for running the test suite) and MongoDB 4.0 or higher for compatibility.

## Connection Management

- **Client Creation**: Use `mongo.Connect` with a context and connection options. Always defer `client.Disconnect` to clean up resources.
  ```go
  ctx := context.Background()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if err != nil {
      log.Fatalf("Failed to connect to MongoDB: %v", err)
  }
  defer func() {
      if err = client.Disconnect(ctx); err != nil {
          log.Fatalf("Failed to disconnect: %v", err)
      }
  }()
  ```
- **Connection String**: Use URI format for configuration, including authentication and options like `replicaSet` or `authSource`.
  ```go
  uri := "mongodb+srv://user:password@cluster.mongodb.net/mydb?retryWrites=true&w=majority"
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
  ```
- **Ping for Discovery**: Call `Ping` to ensure the server is reachable and perform initial discovery.
  ```go
  ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
  defer cancel()
  if err := client.Ping(ctx, readpref.Primary()); err != nil {
      log.Fatalf("Failed to ping MongoDB: %v", err)
  }
  ```
- **Connection Pooling**: Configure pool size and timeouts via `ClientOptions`.
  ```go
  opts := options.Client().
      ApplyURI(uri).
      SetMaxPoolSize(50).
      SetMaxConnIdleTime(5 * time.Minute).
      SetServerSelectionTimeout(30 * time.Second)
  client, err := mongo.Connect(ctx, opts)
  ```
  - Set `MaxPoolSize` based on expected load (e.g., 10-100).
  - Use `MaxConnIdleTime` to recycle idle connections.
- **Network Compression**: Enable compression (Snappy, Zlib, Zstd) to reduce bandwidth.
  ```go
  opts := options.Client().SetCompressors([]string{"snappy", "zlib", "zstd"})
  ```

## Project Structure

Organize MongoDB-related code for modularity and testability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Entry point with MongoDB initialization
  ├── internal/
  │   ├── db/                  # Database setup and repository logic
  │   │   ├── mongodb.go       # Connection setup and queries
  │   │   └── models.go        # Data models and structs
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Define interactions in a `db` package using a repository pattern.
- Use interfaces for data access to enable mocking:
  ```go
  type UserRepository interface {
      GetUserByID(ctx context.Context, id string) (*User, error)
  }
  ```

## Query Best Practices

- **BSON Usage**: Use `bson.D` for ordered documents and `bson.M` for unordered maps.
  ```go
  // Insert
  doc := bson.D{{"name", "pi"}, {"value", 3.14159}}
  result, err := collection.InsertOne(ctx, doc)
  id := result.InsertedID.(primitive.ObjectID).Hex()
  ```
- **Context with Timeouts**: Always use contexts with timeouts for queries.
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  ```
- **Find Operations**: Use `Find` for multiple documents and iterate with cursors.
  ```go
  cur, err := collection.Find(ctx, bson.D{})
  if err != nil {
      return err
  }
  defer cur.Close(ctx)
  for cur.Next(ctx) {
      var result bson.M
      if err := cur.Decode(&result); err != nil {
          return err
      }
      // Process result
  }
  if err := cur.Err(); err != nil {
      return err
  }
  ```
- **Single Document**: Use `FindOne` for single results, handling `mongo.ErrNoDocuments`.
  ```go
  var result struct {
      Value float64 `bson:"value"`
  }
  filter := bson.D{{"name", "pi"}}
  err := collection.FindOne(ctx, filter).Decode(&result)
  if errors.Is(err, mongo.ErrNoDocuments) {
      return nil, fmt.Errorf("document not found")
  }
  ```
- **Aggregation**: Use pipelines for complex queries.
  ```go
  pipeline := mongo.Pipeline{
      {{"$match", bson.D{{"status", "active"}}}},
      {{"$group", bson.D{{"_id", "$category"}, {"count", bson.D{{"$sum", 1}}}}}},
  }
  cur, err := collection.Aggregate(ctx, pipeline)
  ```
- **Indexes**: Create indexes for performance on frequently queried fields.
  ```go
  indexModel := mongo.IndexModel{
      Keys: bson.D{{"email", 1}},
      Options: options.Index().SetUnique(true),
  }
  _, err := collection.Indexes().CreateOne(ctx, indexModel)
  ```
- **Transactions**: Use multi-document transactions for atomicity (requires replica set).
  ```go
  session, err := client.StartSession()
  if err != nil {
      return err
  }
  defer session.EndSession(ctx)
  err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
      _, err := collection.InsertOne(sc, doc)
      if err != nil {
          return err
      }
      return session.CommitTransaction(sc)
  })
  ```

## Security

- **Authentication**: Use SCRAM-SHA-256 or X.509 in the connection URI.
  ```go
  uri := "mongodb://user:password@localhost:27017/mydb?authMechanism=SCRAM-SHA-256&authSource=admin"
  ```
- **TLS/SSL**: Enable TLS for secure connections.
  ```go
  opts := options.Client().ApplyURI(uri).SetTLSConfig(&tls.Config{InsecureSkipVerify: false})
  ```
- **Role-Based Access**: Assign minimal roles to the application user (e.g., readWrite on specific databases).
- **Input Validation**: Validate and sanitize inputs before queries to prevent injection (though BSON parameterization helps).
- **Encryption**: Use MongoDB's client-side field level encryption for sensitive data.

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Repository Pattern**:
  ```go
  package db

  import (
      "context"
      "go.mongodb.org/mongo-driver/v2/mongo"
  )

  type MongoRepo struct {
      Collection *mongo.Collection
  }

  func (r *MongoRepo) GetUser(ctx context.Context, id string) (*User, error) {
      var user User
      objID, _ := primitive.ObjectIDFromHex(id)
      err := r.Collection.FindOne(ctx, bson.D{{"_id", objID}}).Decode(&user)
      if err == mongo.ErrNoDocuments {
          return nil, nil
      }
      return &user, err
  }

  // In handlers
  func GetUserHandler(repo *MongoRepo) gin.HandlerFunc {
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
  viper.SetDefault("db.uri", "mongodb://localhost:27017")
  uri := viper.GetString("db.uri")
  ```
  - **Config File (config.yaml)**:
    ```yaml
    db:
      uri: mongodb://user:password@localhost:27017/mydb
    ```
- **Cobra CLI**: Initialize the client in the root command.
  ```go
  var client *mongo.Client

  var rootCmd = &cobra.Command{
      Use: "myapp",
      PersistentPreRun: func(cmd *cobra.Command, args []string) {
          ctx := context.Background()
          var err error
          client, err = mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("db.uri")))
          if err != nil {
              log.Fatalf("Failed to connect to MongoDB: %v", err)
          }
          if err := client.Ping(ctx, readpref.Primary()); err != nil {
              log.Fatalf("Failed to ping MongoDB: %v", err)
          }
      },
      PersistentPostRun: func(cmd *cobra.Command, args []string) {
          client.Disconnect(context.Background())
      },
  }
  ```
- **JWT Authentication**: Store and query user sessions in MongoDB.
  ```go
  func ValidateUser(repo *MongoRepo, token *jwt.Token) error {
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
- **Zap Logging**: Log MongoDB operations and errors.
  ```go
  logger, _ := zap.NewProduction()
  defer logger.Sync()
  if err := client.Ping(ctx, readpref.Primary()); err != nil {
      logger.Error("MongoDB ping failed", zap.Error(err))
  } else {
      logger.Info("MongoDB connected")
  }
  ```
- **Prometheus Metrics**: Track query performance.
  ```go
  var (
      queryDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
          Name:    "myapp_mongo_query_duration_seconds",
          Help:    "Duration of MongoDB queries",
          Buckets: prometheus.DefBuckets,
      }, []string{"operation"})
  )

  func (r *MongoRepo) GetUser(ctx context.Context, id string) (*User, error) {
      start := time.Now()
      defer func() {
          queryDuration.WithLabelValues("find_one").Observe(time.Since(start).Seconds())
      }()
      // Query logic
  }
  ```

## Testing

- **Unit Tests**: Use mocks or `mongo-driver/mongo/integration/mtest` for integration testing.
  ```go
  import "go.mongodb.org/mongo-driver/v2/mongo/integration/mtest"

  func TestGetUser(t *testing.T) {
      mtest.TestWithSingleDB(t, func(mtest *mtest.Test) {
          collection := mtest.Coll
          // Insert test data
          _, _ = collection.InsertOne(mtest.Ctx, bson.D{{"_id", "testid"}, {"name", "Test"}})
          repo := &MongoRepo{Collection: collection}
          user, err := repo.GetUser(mtest.Ctx, "testid")
          assert.NoError(t, err)
          assert.Equal(t, "Test", user.Name)
      })
  }
  ```
- **Integration Tests**: Use a test MongoDB instance (e.g., Dockerized MongoDB).

## Performance Tips

- **Connection Pool Tuning**: Adjust `MaxPoolSize` and timeouts based on load.
- **Indexes**: Create compound indexes for common query patterns.
- **Projection**: Use `Projection` in queries to fetch only needed fields.
  ```go
  opts := options.Find().SetProjection(bson.D{{"name", 1}})
  cur, _ := collection.Find(ctx, filter, opts)
  ```
- **Read Preferences**: Use `readpref.Secondary` for read-heavy workloads.
- **Bulk Writes**: Use `BulkWrite` for multiple operations.
  ```go
  ops := []mongo.WriteModel{&mongo.InsertOneModel{Document: doc1}, &mongo.InsertOneModel{Document: doc2}}
  _, err := collection.BulkWrite(ctx, ops)
  ```

## Additional Resources

- Official MongoDB Go Driver GitHub: https://github.com/mongodb/mongo-go-driver
- MongoDB Documentation: https://www.mongodb.com/docs/drivers/go/current/
- Tutorials: https://www.mongodb.com/docs/drivers/go/current/fundamentals/

Adhere to these practices to ensure secure and efficient MongoDB integration with the Go driver. Update this document as new best practices or library updates emerge.