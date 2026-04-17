package cli

import (
	"fmt"
	"path/filepath"

	"github.com/SilverShaman/diffless/internal/git"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean [task-id]",
	Short: "Prunes and cleanly removes a completed worktree block from git",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		worktreePath, _ := filepath.Abs(fmt.Sprintf("../.diffless-workspaces/%s", taskID))

		fmt.Printf("Cleaning up Sandbox data for '%s'...\n", taskID)

		// 1. Remove the worktree logically
		if err := git.RemoveWorktree(worktreePath); err != nil {
			return err
		}
		fmt.Printf("✓ Removed physical directory and detached worktree link for %s\n", worktreePath)

		// 2. Prune native git references
		if err := git.PruneWorktrees(); err != nil {
			return err
		}

		fmt.Println("✓ Pruned stale worktree graph references")
		return nil
	},
}
