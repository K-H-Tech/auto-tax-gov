# Viper Best Practices Guide

This document outlines best practices for managing configuration in Go applications using the Viper library. It draws from the official Viper documentation and community recommendations to ensure configuration is robust, maintainable, and secure. Follow these guidelines to effectively integrate Viper into your Go projects, particularly when used with frameworks like Gin.

## Installation and Setup

- Install Viper using Go modules: `go get github.com/spf13/viper`.
- For remote key/value store support (e.g., etcd, Consul), include the remote package: `import _ "github.com/spf13/viper/remote"`.
- Initialize a new Viper instance for each application context to avoid global state issues: `v := viper.New()`. Avoid using the global Viper instance to improve testability and prevent unexpected behavior.

## Configuration Sources

Viper supports multiple configuration sources with a defined precedence order: explicit `Set`, flags, environment variables, config files, key/value stores, and defaults. Use this hierarchy to manage configuration effectively.

- **Defaults**: Set default values for all required configuration keys to ensure fallback behavior.
  ```go
  v.SetDefault("server.port", 8080)
  v.SetDefault("database.host", "localhost")
  ```
- **Config Files**: Support JSON, YAML, TOML, or HCL formats. Prefer YAML or JSON for readability and widespread tool support.
  ```go
  v.SetConfigName("config")
  v.SetConfigType("yaml")
  v.AddConfigPath("/etc/myapp/")
  v.AddConfigPath("$HOME/.myapp")
  v.AddConfigPath(".")
  if err := v.ReadInConfig(); err != nil {
      if _, ok := err.(viper.ConfigFileNotFoundError); ok {
          log.Println("No config file found, using defaults")
      } else {
          log.Fatalf("Error reading config: %v", err)
      }
  }
  ```
- **Environment Variables**: Use `SetEnvPrefix` and `AutomaticEnv` for automatic environment variable binding. Use `BindEnv` for specific variables.
  ```go
  v.SetEnvPrefix("MYAPP")
  v.AutomaticEnv()
  v.BindEnv("db_port", "MYAPP_DB_PORT")
  ```
- **Flags**: Integrate with `pflag` (used by Cobra) for command-line flags.
  ```go
  import "github.com/spf13/pflag"
  pflag.Int("port", 8080, "Server port")
  pflag.Parse()
  v.BindPFlags(pflag.CommandLine)
  ```
- **Remote Stores**: Use etcd or Consul for dynamic configuration. Ensure encryption for sensitive data.
  ```go
  v.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/myapp.json")
  v.SetConfigType("json")
  if err := v.ReadRemoteConfig(); err != nil {
      log.Fatalf("Error reading remote config: %v", err)
  }
  ```

## Project Integration

- **Dependency Injection**: Pass Viper instances to modules or services to encapsulate configuration.
  ```go
  type Server struct {
      Port int
  }
  func NewServer(v *viper.Viper) *Server {
      return &Server{Port: v.GetInt("server.port")}
  }
  ```
- **Sub-Configurations**: Use `v.Sub()` to extract configuration subsets for specific modules.
  ```go
  dbConfig := v.Sub("database")
  if dbConfig == nil {
      log.Fatal("Database configuration not found")
  }
  dbHost := dbConfig.GetString("host")
  ```
- **Unmarshaling**: Unmarshal configurations into structs for type safety. Use `mapstructure` tags to map keys.
  ```go
  type Config struct {
      Server struct {
          Port int    `mapstructure:"port"`
          Host string `mapstructure:"host"`
      } `mapstructure:"server"`
  }
  var cfg Config
  if err := v.Unmarshal(&cfg); err != nil {
      log.Fatalf("Error unmarshaling config: %v", err)
  }
  ```

## Best Practices

- **Case Insensitivity**: Viper keys are case-insensitive. Use consistent naming (e.g., `server.port` or `SERVER_PORT`) across sources.
- **Validation**: Validate configuration values after loading. Use custom validation logic or libraries like `validator`.
  ```go
  port := v.GetInt("server.port")
  if port <= 0 || port > 65535 {
      log.Fatal("Invalid port number")
  }
  ```
- **Dynamic Updates**: Enable live reloading for config files or remote stores using `WatchConfig` or `WatchRemoteConfig`.
  ```go
  v.WatchConfig()
  v.OnConfigChange(func(e fsnotify.Event) {
      log.Printf("Config file changed: %s", e.Name)
  })
  ```
- **Error Handling**: Always check for errors when reading configs and handle `ConfigFileNotFoundError` gracefully.
- **Security**: Avoid storing sensitive data (e.g., API keys) in config files unless encrypted. Use remote stores with encryption for sensitive data.
  ```go
  v.AddSecureRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/myapp.json", "/etc/secrets/mykeyring.gpg")
  ```
- **Testing**: Mock Viper instances in tests by setting values explicitly or using in-memory configs.
  ```go
  v := viper.New()
  v.SetConfigType("yaml")
  yamlConfig := []byte(`server: {port: 8080}`)
  v.ReadConfig(bytes.NewBuffer(yamlConfig))
  ```
- **Thread Safety**: Synchronize access to Viper instances in concurrent applications using `sync.RWMutex`.
  ```go
  var mu sync.RWMutex
  mu.RLock()
  port := v.GetInt("server.port")
  mu.RUnlock()
  ```

## Integration with Gin

When using Viper with Gin, centralize configuration in a Viper instance and inject it into your application's components.

- **Example Setup**:
  ```go
  package main

  import (
      "log"
      "github.com/gin-gonic/gin"
      "github.com/spf13/viper"
  )

  type Config struct {
      Server struct {
          Port int `mapstructure:"port"`
      } `mapstructure:"server"`
  }

  func main() {
      v := viper.New()
      v.SetConfigName("config")
      v.SetConfigType("yaml")
      v.AddConfigPath(".")
      v.SetDefault("server.port", 8080)

      if err := v.ReadInConfig(); err != nil {
          if _, ok := err.(viper.ConfigFileNotFoundError); ok {
              log.Println("No config file, using defaults")
          } else {
              log.Fatalf("Config error: %v", err)
          }
      }

      var cfg Config
      if err := v.Unmarshal(&cfg); err != nil {
          log.Fatalf("Unmarshal error: %v", err)
      }

      r := gin.Default()
      r.GET("/ping", func(c *gin.Context) {
          c.JSON(200, gin.H{"message": "pong"})
      })

      r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
  }
  ```
- **Config File (config.yaml)**:
  ```yaml
  server:
    port: 8080
  ```

## Additional Resources

- Official Viper GitHub: https://github.com/spf13/viper
- Viper Feedback Form for v2: https://forms.gle/R6faU74qPRPAzchZ9

Adhere to these practices to ensure a robust configuration setup with Viper. Update this document as new features or best practices emerge.