# GoMock Best Practices Guide

This document outlines best practices for using the `go.uber.org/mock` (GoMock) framework to create mocks for Go interfaces in testing. It draws from the official GoMock documentation, community recommendations, and testing best practices to ensure maintainable, readable, and effective unit tests. These guidelines are tailored for integration with frameworks like Gin, Viper, Cobra, JWT, Zap, and Prometheus, as well as dependencies like `database/sql`, `pgx`, `mongo-go-driver`, `amqp091-go`, and `go-redis`. The library supports the two most recent Go releases (1.23+).

## Installation and Setup

- Install the `mockgen` tool using Go modules:
  ```bash
  go install go.uber.org/mock/mockgen@latest
  ```
- Ensure `GOPATH/bin` is in your `PATH`:
  ```bash
  export PATH=$PATH:$(go env GOPATH)/bin
  ```
- Verify installation:
  ```bash
  mockgen -version
  ```
- Import the mocking library:
  ```go
  import (
      "go.uber.org/mock/gomock"
  )
  ```
- Use Go 1.23+ for compatibility with GoMock's support for the latest two Go releases.

## Project Structure

Organize mocks to align with the application’s structure for clarity and maintainability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   └── app/
  │       └── main.go          # Application entry point
  ├── internal/
  │   ├── db/                  # Database logic
  │   │   ├── postgres.go      # Database repository
  │   │   └── mocks/           # Mock implementations
  │   │       └── mock_postgres.go  # Generated mocks
  │   ├── messaging/           # RabbitMQ or Redis logic
  │   ├── handlers/            # Gin handlers
  │   └── services/            # Business logic
  ├── tests/
  │   └── integration/         # Integration tests
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Generate mocks in a `mocks` subpackage within the relevant package (e.g., `internal/db/mocks`).
- Use interfaces to define dependencies, enabling mock generation:
  ```go
  // internal/db/postgres.go
  type UserRepository interface {
      GetUser(ctx context.Context, id string) (*User, error)
  }
  ```

## Mock Generation

- **Source Mode (Recommended)**: Generate mocks from a source file containing interfaces.
  ```bash
  mockgen -source=internal/db/postgres.go -destination=internal/db/mocks/mock_postgres.go -package=mocks
  ```
- **Package Mode**: Generate mocks from a package’s interfaces.
  ```bash
  mockgen -package=mocks . UserRepository
  ```
- **Use with go:generate**: Add to source files for automation.
  ```go
  //go:generate mockgen -source=postgres.go -destination=mocks/mock_postgres.go -package=mocks
  ```
  Run `go generate ./...` to regenerate mocks.
- **Flags for Customization**:
  - `-package`: Set custom package name (e.g., `mocks` instead of `mock_db`).
  - `-mock_names`: Customize mock names (e.g., `UserRepository=MockUserRepo`).
  - `-typed`: Enable type-safe `Return`, `Do`, `DoAndReturn` methods for better IDE support.
  - `-copyright_file`: Add a copyright header.
  - `-exclude_interfaces`: Skip specific interfaces if needed.
  ```bash
  mockgen -source=postgres.go -destination=mocks/mock_postgres.go -package=mocks -mock_names=UserRepository=MockUserRepo -typed
  ```
- **Avoid Import Cycles**: Use `-self_package` if the mock package matches the input package to prevent cycles.

## Mocking Best Practices

- **Controller Lifecycle**: Create a `gomock.Controller` per test and finish it to verify expectations.
  ```go
  func TestGetUser(t *testing.T) {
      ctrl := gomock.NewController(t)
      defer ctrl.Finish()
      // Mock setup and test
  }
  ```
- **Expectations**: Set precise expectations using matchers like `gomock.Eq`, `gomock.Any`, or custom matchers.
  ```go
  mockRepo := mocks.NewMockUserRepository(ctrl)
  mockRepo.EXPECT().
      GetUser(gomock.Eq(context.Background()), gomock.Eq("123")).
      Return(&User{ID: "123", Name: "Alice"}, nil)
  ```
- **Flexible Matching**: Use `gomock.Any()` for less strict matching or custom matchers for complex types.
  ```go
  mockRepo.EXPECT().
      GetUser(gomock.Any(), gomock.Eq("123")).
      Return(nil, nil).
      AnyTimes()
  ```
- **Stubbing Behavior**: Use `DoAndReturn` for dynamic responses.
  ```go
  mockRepo.EXPECT().
      GetUser(gomock.Any(), gomock.Eq("123")).
      DoAndReturn(func(ctx context.Context, id string) (*User, error) {
          return &User{ID: id, Name: "Test"}, nil
      })
  ```
- **Ordered Calls**: Use `gomock.InOrder` for strict call sequences.
  ```go
  gomock.InOrder(
      mockRepo.EXPECT().GetUser(gomock.Any(), "1").Return(nil, nil),
      mockRepo.EXPECT().GetUser(gomock.Any(), "2").Return(&User{}, nil),
  )
  ```
- **Error Handling**: Test error paths explicitly.
  ```go
  mockRepo.EXPECT().
      GetUser(gomock.Any(), "invalid").
      Return(nil, fmt.Errorf("invalid ID"))
  ```
- **Custom Formatters**: Modify `Want`/`Got` messages for clearer test failures.
  ```go
  mockRepo.EXPECT().
      GetUser(gomock.Any(), gomock.WantFormatter(
          gomock.StringerFunc(func() string { return "is user ID 123" }),
          gomock.Eq("123"),
      ))
  ```

## Integration with Gin, Viper, Cobra, JWT, Zap, and Prometheus

- **Gin Handler Testing**:
  ```go
  func TestUserHandler(t *testing.T) {
      ctrl := gomock.NewController(t)
      defer ctrl.Finish()
      mockRepo := mocks.NewMockUserRepository(ctrl)
      mockRepo.EXPECT().
          GetUser(gomock.Any(), "123").
          Return(&User{ID: "123", Name: "Alice"}, nil)

      r := gin.Default()
      r.GET("/users/:id", handlers.GetUserHandler(mockRepo))
      req, _ := http.NewRequest("GET", "/users/123", nil)
      w := httptest.NewRecorder()
      r.ServeHTTP(w, req)
      assert.Equal(t, 200, w.Code)
  }
  ```
- **Viper Configuration**: Mock configuration access if needed.
  ```go
  type Config interface {
      GetString(key string) string
  }
  mockConfig := mocks.NewMockConfig(ctrl)
  mockConfig.EXPECT().GetString("db.url").Return("postgres://localhost")
  ```
- **Cobra CLI Testing**: Mock dependencies for commands.
  ```go
  func TestRootCmd(t *testing.T) {
      ctrl := gomock.NewController(t)
      defer ctrl.Finish()
      mockRepo := mocks.NewMockUserRepository(ctrl)
      cmd := &cobra.Command{
          Run: func(cmd *cobra.Command, args []string) {
              user, _ := mockRepo.GetUser(context.Background(), "123")
              // Assert user
          },
      }
      mockRepo.EXPECT().GetUser(gomock.Any(), "123").Return(&User{}, nil)
      cmd.Execute()
  }
  ```
- **JWT**: Mock JWT validation dependencies.
  ```go
  type AuthService interface {
      ValidateToken(token string) (*jwt.Token, error)
  }
  mockAuth := mocks.NewMockAuthService(ctrl)
  mockAuth.EXPECT().ValidateToken("token").Return(&jwt.Token{}, nil)
  ```
- **Zap Logging**: Mock logger to verify log calls or use `zap.NewNop` for no-op logging.
  ```go
  mockLogger := mocks.NewMockLogger(ctrl)
  mockLogger.EXPECT().Info("User fetched", gomock.Any())
  ```
- **Prometheus Metrics**: Mock metric collectors or use test utilities.
  ```go
  mockCounter := mocks.NewMockCounter(ctrl)
  mockCounter.EXPECT().Inc()
  ```

## Testing with External Dependencies

- **Database (pgx, mongo-go-driver)**:
  - Mock `pgxpool.Pool` or `mongo.Collection` interfaces.
  ```go
  type DB interface {
      QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
  }
  mockDB := mocks.NewMockDB(ctrl)
  mockDB.EXPECT().QueryRow(gomock.Any(), gomock.Any(), "123").Return(mockRow)
  ```
- **RabbitMQ (amqp091-go)**:
  - Mock `amqp.Channel` for publishing/consuming.
  ```go
  mockChannel := mocks.NewMockChannel(ctrl)
  mockChannel.EXPECT().PublishWithContext(gomock.Any(), "", "queue", false, false, gomock.Any())
  ```
- **Redis (go-redis)**:
  - Mock `redis.Cmdable` for commands.
  ```go
  mockRedis := mocks.NewMockCmdable(ctrl)
  mockRedis.EXPECT().Get(gomock.Any(), "key").Return(redis.NewStringResult("value", nil))
  ```

## Best Practices

- **One Mock per Interface**: Generate one mock per interface for clarity.
- **Minimal Expectations**: Set only necessary expectations to keep tests focused.
- **Clean Up**: Always call `ctrl.Finish()` to catch unmet expectations.
- **Avoid Over-Mocking**: Mock only external dependencies; test internal logic directly.
- **Use Matchers Sparingly**: Prefer `gomock.Eq` over `gomock.Any` for precision.
- **Regenerate Mocks**: Run `go generate` after interface changes to keep mocks up-to-date.
- **Type-Safe Mocks**: Enable `-typed` flag for better IDE support and type safety.
- **Isolation**: Use separate `gomock.Controller` instances for concurrent tests to avoid race conditions.

## Additional Resources

- Official GoMock GitHub: https://github.com/uber-go/mock
- Go Testing Documentation: https://pkg.go.dev/testing
- Community Guides: https://github.com/golang/mock/wiki

Adhere to these practices for effective mocking with GoMock. Update this document as new best practices or library updates emerge.