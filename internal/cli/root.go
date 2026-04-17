package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diffless",
	Short: "Diffless is an AI-augmented git workflow CLI builder",
	Long:  `Diffless integrates directly with Git Worktrees to provide safely verified physical sandboxes for autonomous AI agents like Google Antigravity.`,
}

// Execute triggers the root Cobra runner
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Register commands generated for Phase 1
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(cleanCmd)
}
