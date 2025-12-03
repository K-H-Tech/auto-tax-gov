package main

import (
	"fmt"
	"path/filepath"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port   int
	webDir string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long: `Start an HTTP server with a web UI for tracking redirects
and automating the tax portal login flow.

The server provides endpoints for:
  - Viewing captcha and entering codes
  - Sending and verifying OTP
  - Accessing the tax dashboard`,
	RunE: runServe,
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "port to listen on")
	serveCmd.Flags().StringVarP(&webDir, "web-dir", "w", "web", "directory containing web assets")

	// Bind flags to viper
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Override port if flag was set
	if cmd.Flags().Changed("port") {
		cfg.Server.Port = port
	}

	// Resolve web directory path
	webPath, err := filepath.Abs(webDir)
	if err != nil {
		return fmt.Errorf("failed to resolve web directory path: %w", err)
	}

	logger.Info("starting server",
		"port", cfg.Server.Port,
		"webDir", webPath,
	)

	srv := server.New(cfg, logger, webPath)
	return srv.Start()
}
