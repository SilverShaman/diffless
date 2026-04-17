package cli

import (
	"fmt"
	"path/filepath"

	"github.com/SilverShaman/diffless/internal/antigravity"
	"github.com/SilverShaman/diffless/internal/git"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [task-id]",
	Short: "Create an isolated worktree sandbox for an AI agent",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		worktreePath, _ := filepath.Abs(fmt.Sprintf("../.diffless-workspaces/%s", taskID))

		fmt.Printf("Starting Diffless sandbox for '%s'...\n", taskID)

		// 1. Create Git Worktree
		if err := git.AddWorktree(taskID, worktreePath); err != nil {
			return err
		}
		fmt.Printf("✓ Created git worktree physically isolated at %s\n", worktreePath)

		// 2. Engage Antigravity Sandbox Mode
		if err := antigravity.BindWorkspace(worktreePath); err != nil {
			return err
		}

		fmt.Printf("\nSuccess! Sandbox is ready. Run 'diffless switch %s' to begin work.\n", taskID)
		return nil
	},
}
