package main

import (
	"fmt"

	"github.com/K-H-Tech/auto-tax-gov/internal/config"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mygovauth"
	"github.com/K-H-Tech/auto-tax-gov/internal/service/mytax"
	"github.com/K-H-Tech/auto-tax-gov/internal/session"
	"github.com/spf13/cobra"
)

var (
	outputDir  string
	outputJSON bool
)

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track redirects from the command line",
	Long: `Track HTTP redirects on the tax portal and save results to files.

This command performs the redirect chain tracking without starting
a web server. Results are saved to markdown and JSON files.`,
	RunE: runTrack,
}

func init() {
	trackCmd.Flags().StringVarP(&outputDir, "output", "o", ".", "directory to save results")
	trackCmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "output results as JSON")

	rootCmd.AddCommand(trackCmd)
}

func runTrack(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Override output directory if flag was explicitly set (Issue 3)
	if cmd.Flags().Changed("output") {
		cfg.HTTP.OutputDir = outputDir
	}

	// Create services
	authSvc := mygovauth.New(cfg, logger)
	mytaxSvc := mytax.New(cfg, authSvc, logger)

	// Create session
	sess := session.New()

	logger.Info("starting redirect tracking", "startURL", mytaxSvc.GetAuthURL())

	captcha, steps, err := mytaxSvc.InitiateLogin(sess)
	if err != nil {
		return fmt.Errorf("tracking failed: %w", err)
	}

	logger.Info("tracking complete", "steps", len(steps))

	// Print summary
	fmt.Printf("\nRedirect Chain Summary:\n")
	fmt.Printf("=======================\n")
	for i, step := range steps {
		fmt.Printf("%d. [%d] %s\n", i+1, step.StatusCode, step.URL)
		if step.RedirectLocation != "" {
			fmt.Printf("   → %s\n", step.RedirectLocation)
		}
	}
	fmt.Printf("\nTotal steps: %d\n", len(steps))

	// Check for captcha
	if captcha != nil {
		fmt.Printf("\nCaptcha found!\n")
		fmt.Printf("CSRF Token: %s...\n", captcha.CSRFToken[:min(20, len(captcha.CSRFToken))])
	}

	return nil
}
