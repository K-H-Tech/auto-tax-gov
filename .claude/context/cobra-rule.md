# Cobra Best Practices Guide

This document outlines best practices for building command-line interface (CLI) applications in Go using the Cobra library. It draws from the official Cobra documentation, community practices, and real-world usage in projects like Kubernetes and Hugo to ensure your CLI is intuitive, maintainable, and robust. These guidelines are particularly useful when integrating Cobra with other libraries like Viper, Gin, and JWT for comprehensive Go applications.

## Installation and Setup

- Install Cobra and its CLI tool using Go modules:
  ```bash
  go get -u github.com/spf13/cobra@latest
  go install github.com/spf13/cobra-cli@latest
  ```
- Import Cobra in your code:
  ```go
  import "github.com/spf13/cobra"
  ```
- Use the `cobra-cli` tool to scaffold your application:
  ```bash
  cobra-cli init myapp
  cobra-cli add [command]
  ```
  This generates a structured project with a root command and optional subcommands.

## Project Structure

Organize your CLI application to promote clarity and maintainability.

- **Recommended Directory Layout**:
  ```
  myapp/
  ├── cmd/
  │   ├── root.go              # Root command definition
  │   ├── subcommand1.go       # Subcommand definitions
  │   └── subcommand2.go
  ├── internal/
  │   ├── config/              # Configuration handling (e.g., Viper integration)
  │   ├── services/            # Business logic
  │   └── utils/               # Shared utilities
  ├── main.go                  # Application entry point
  ├── go.mod                   # Module definition
  └── go.sum                   # Dependency checksums
  ```
- Place each command in its own file under `cmd/` for modularity.
- Initialize the application in `main.go` by calling the root command's `Execute` method.

## Command Design

Cobra applications are built around commands, arguments, and flags, following the pattern `APPNAME VERB NOUN --FLAG`.

- **Define Commands**:
  - Create a root command and subcommands to represent actions.
  - Use `cobra.Command` to define commands with fields like `Use`, `Short`, `Long`, and `Run`.
  ```go
  var rootCmd = &cobra.Command{
      Use:   "myapp",
      Short: "A brief description of myapp",
      Long:  `A longer description of myapp, including usage details.`,
      Run: func(cmd *cobra.Command, args []string) {
          fmt.Println("Root command executed")
      },
  }
  ```
- **Subcommands**:
  - Add subcommands for specific actions (e.g., `myapp start`, `myapp config`).
  ```go
  var startCmd = &cobra.Command{
      Use:   "start",
      Short: "Start the application",
      RunE: func(cmd *cobra.Command, args []string) error {
          return startApp()
      },
  }
  func init() {
      rootCmd.AddCommand(startCmd)
  }
  ```
  - Use `RunE` for commands that return errors to enable proper error handling.
- **Arguments**:
  - Validate arguments in the `Run` or `RunE` function.
  - Use `Args` field for custom argument validation:
  ```go
  var cmd = &cobra.Command{
      Use:   "greet [name]",
      Args:  cobra.ExactArgs(1),
      Run: func(cmd *cobra.Command, args []string) {
          fmt.Printf("Hello, %s!\n", args[0])
      },
  }
  ```
- **Aliases**: Add aliases to support alternative command names without breaking existing usage.
  ```go
  var cmd = &cobra.Command{
      Use:     "fetch",
      Aliases: []string{"get", "retrieve"},
      Short:   "Fetch resources",
      Run: func(cmd *cobra.Command, args []string) {
          fmt.Println("Fetching resources")
      },
  }
  ```

## Flags

Flags modify command behavior and are provided by the `pflag` library for POSIX compliance.

- **Define Flags**:
  - Use local flags for command-specific options and persistent flags for inheritance by subcommands.
  ```go
  func init() {
      rootCmd.PersistentFlags().String("config", "", "Config file path")
      startCmd.Flags().Int("port", 8080, "Port to run the server on")
  }
  ```
- **Access Flags**:
  - Retrieve flag values in the `Run` function:
  ```go
  port, _ := cmd.Flags().GetInt("port")
  fmt.Printf("Running on port %d\n", port)
  ```
- **Shorthand Flags**: Provide short versions for frequently used flags (e.g., `-p` for `--port`).
  ```go
  rootCmd.PersistentFlags().StringP("config", "c", "", "Config file path")
  ```
- **Mark Flags as Required**: Use `MarkFlagRequired` for essential flags.
  ```go
  cmd.MarkFlagRequired("config")
  ```

## Integration with Viper

Cobra integrates seamlessly with Viper for configuration management.

- **Bind Flags to Viper**:
  - Bind flags to Viper keys to allow configuration via flags, environment variables, or config files.
  ```go
  import (
      "github.com/spf13/cobra"
      "github.com/spf13/viper"
  )

  func init() {
      rootCmd.PersistentFlags().String("port", "8080", "Server port")
      viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
  }
  ```
- **Load Config with Viper**:
  - Initialize Viper in `main.go` or a dedicated `config` package.
  ```go
  viper.SetConfigName("config")
  viper.AddConfigPath(".")
  viper.ReadInConfig()
  ```
- **Access Config**:
  - Use Viper to retrieve configuration values in commands.
  ```go
  port := viper.GetInt("server.port")
  ```

## Integration with Gin and JWT

When building a CLI that manages a Gin-based server with JWT authentication, use Cobra to control server lifecycle and configuration.

- **Example CLI with Gin and JWT**:
  ```go
  package main

  import (
      "fmt"
      "github.com/gin-gonic/gin"
      "github.com/golang-jwt/jwt/v5"
      "github.com/spf13/cobra"
      "github.com/spf13/viper"
  )

  var rootCmd = &cobra.Command{
      Use:   "myapp",
      Short: "A CLI for managing myapp",
  }

  var serverCmd = &cobra.Command{
      Use:   "server",
      Short: "Start the Gin server",
      RunE: func(cmd *cobra.Command, args []string) error {
          port, _ := cmd.Flags().GetInt("port")
          secret := viper.GetString("jwt.secret")
          r := gin.Default()
          r.GET("/ping", func(c *gin.Context) {
              c.JSON(200, gin.H{"message": "pong"})
          })
          r.POST("/login", func(c *gin.Context) {
              claims := jwt.RegisteredClaims{Subject: "user123"}
              token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
              tokenString, _ := token.SignedString([]byte(secret))
              c.JSON(200, gin.H{"token": tokenString})
          })
          return r.Run(fmt.Sprintf(":%d", port))
      },
  }

  func init() {
      serverCmd.Flags().IntP("port", "p", 8080, "Server port")
      viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
      viper.SetConfigName("config")
      viper.AddConfigPath(".")
      viper.ReadInConfig()
      rootCmd.AddCommand(serverCmd)
  }

  func main() {
      if err := rootCmd.Execute(); err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
  }
  ```
- **Config File (config.yaml)**:
  ```yaml
  server:
    port: 8080
  jwt:
    secret: your-256-bit-secret
  ```

## Help and Documentation

- **Automatic Help**: Cobra generates help text automatically via `-h` or `--help`.
- **Custom Help**: Customize help output by overriding `HelpFunc` or `UsageTemplate` for specific formatting.
- **Subcommand Grouping**: Group related subcommands for better help organization.
  ```go
  var adminCmd = &cobra.Command{
      Use:   "admin",
      Short: "Administrative commands",
  }
  rootCmd.AddCommand(adminCmd)
  adminCmd.AddCommand(createUserCmd, deleteUserCmd)
  ```
- **Man Pages and Autocomplete**: Generate man pages and shell autocomplete scripts:
  ```bash
  cobra-cli add completion
  ```

## Error Handling

- Use `RunE` to return errors and let Cobra handle them.
- Provide clear error messages for user feedback.
  ```go
  RunE: func(cmd *cobra.Command, args []string) error {
      if len(args) == 0 {
          return fmt.Errorf("missing required argument")
      }
      return nil
  }
  ```

## Testing

- **Unit Tests**: Test commands using `cobra.Command.Execute` with mocked inputs.
  ```go
  func TestRootCommand(t *testing.T) {
      cmd := &cobra.Command{
          Use: "test",
          Run: func(cmd *cobra.Command, args []string) {
              fmt.Fprintln(cmd.OutOrStdout(), "test executed")
          },
      }
      output := &bytes.Buffer{}
      cmd.SetOut(output)
      cmd.Execute()
      assert.Contains(t, output.String(), "test executed")
  }
  ```
- **Flag Testing**: Verify flag parsing and binding to Viper.
- **Integration Tests**: Test CLI interactions with the server using `os/exec` or by invoking commands directly.

## Best Practices

- **Command Naming**: Use clear, verb-noun patterns (e.g., `myapp create user`).
- **Flag Consistency**: Use consistent flag names across commands (e.g., `--port` for all server-related commands).
- **Modularity**: Keep command logic in separate files and reuse utilities in `internal/` packages.
- **Validation**: Validate arguments and flags early in the `Run` function.
- **Versioning**: Add a `version` command to display the application version.
  ```go
  var versionCmd = &cobra.Command{
      Use:   "version",
      Short: "Print the version number",
      Run: func(cmd *cobra.Command, args []string) {
          fmt.Println("myapp v1.0.0")
      },
  }
  ```
- **Environment Variables**: Use Viper’s `AutomaticEnv` with a prefix (e.g., `MYAPP_`) for environment variable support.
- **Concurrency Safety**: Ensure shared resources (e.g., Viper configs) are accessed safely with `sync.RWMutex`.

## Additional Resources

- Official Cobra GitHub: https://github.com/spf13/cobra
- Cobra User Guide: https://github.com/spf13/cobra/blob/main/user_guide.md
- Cobra-CLI README: https://github.com/spf13/cobra-cli/blob/main/README.md

Adhere to these practices to build intuitive and robust CLI applications with Cobra. Update this document as new best practices or library updates emerge.