package cli

import (
	"github.com/spf13/cobra"
)

const Version = "0.0.1"

var rootCmd = &cobra.Command{
	Use:     "diffless",
	Short:   "Diffless is an AI-augmented git workflow CLI builder",
	Long:    `Diffless integrates directly with Git Worktrees to provide safely verified physical sandboxes for autonomous AI agents like Google Antigravity.`,
	Version: Version,
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

	// Register commands generated for Phase 2
	rootCmd.AddCommand(lockdownCmd)
	rootCmd.AddCommand(auditCmd)

	// Register commands generated for Phase 4
	rootCmd.AddCommand(syncCmd)

	// Register commands generated for Phase 5
	rootCmd.AddCommand(proposeCmd)

	// Register commands generated for Phase 6
	rootCmd.AddCommand(toolsCmd)

	// Register commands generated for Phase 10
	rootCmd.AddCommand(runCmd)
}
