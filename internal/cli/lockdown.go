package cli

import (
	"fmt"
	"path/filepath"

	"github.com/SilverShaman/diffless/internal/security"
	"github.com/spf13/cobra"
)

var lockdownCmd = &cobra.Command{
	Use:   "lockdown [task-id]",
	Short: "Hardens the sandbox permissions and isolates credentials",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		worktreePath, _ := filepath.Abs(fmt.Sprintf("../.diffless-workspaces/%s", taskID))

		fmt.Printf("Initiating security lockdown for '%s'...\n", taskID)

		// Enforce strict 0700 directory permissions
		if err := security.HardenDirectory(worktreePath); err != nil {
			return fmt.Errorf("failed to harden directory: %w", err)
		}
		fmt.Printf("✓ Hardened directory permissions (0700) for %s\n", worktreePath)

		// Strip out production credentials and insert ephemeral keys
		if err := security.GenerateEphemeralEnv(worktreePath); err != nil {
			return fmt.Errorf("failed to generate ephemeral env: %w", err)
		}
		fmt.Println("✓ Generated isolated ephemeral .env")

		fmt.Println("✓ Sandbox lockdown complete. Zero-Trust boundaries active.")
		return nil
	},
}
