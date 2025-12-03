package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile         string
	logLevel        string
	logLevelFlagSet bool // Track if --log-level flag was explicitly provided (Issue 2)
	logger          *slog.Logger
)

var rootCmd = &cobra.Command{
	Use:   "auto-tax-gov",
	Short: "Tax.gov.ir redirect tracker and automation tool",
	Long: `A CLI tool to track HTTP redirects on the Iranian tax portal (my.tax.gov.ir)
and automate the login flow with OTP verification.

Supports two modes:
  - serve: Start a web server with UI for interactive use
  - track: Run redirect tracking from the command line`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Check if --log-level flag was explicitly set (Issue 2)
		logLevelFlagSet = cmd.Flags().Changed("log-level")
		return initConfig()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level (debug, info, warn, error)")

	// Version command
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("auto-tax-gov %s (commit: %s, built: %s)\n", version, commit, date)
		},
	})
}

func initConfig() error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.auto-tax-gov")
	}

	// Read environment variables (Issue 1: use underscores for valid env var names)
	viper.SetEnvPrefix("AUTO_TAX_GOV")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found is OK, we'll use defaults
	}

	// Initialize logger
	logger = initLogger()

	return nil
}

func initLogger() *slog.Logger {
	var level slog.Level
	switch strings.ToLower(viper.GetString("log.level")) {
	case "debug":
		level = slog.LevelDebug
	case "warn", "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// Override with flag only if explicitly provided (Issue 2)
	if logLevelFlagSet {
		switch strings.ToLower(logLevel) {
		case "debug":
			level = slog.LevelDebug
		case "warn", "warning":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		}
	}

	opts := &slog.HandlerOptions{
		Level: level,
	}

	format := viper.GetString("log.format")
	var handler slog.Handler
	if format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
