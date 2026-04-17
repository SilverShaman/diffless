package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [task-id]",
	Short: "Provides instructions or IDE automation to switch to the isolated worktree",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		
		// Because a child binary cannot natively change the CWD of the parent shell,
		// we output explicit directives. In an IDE like Antigravity, this acts as an IPC hook.
		fmt.Printf("[Antigravity] Internal IDE hook -> switch workspace context to: ../.diffless-workspaces/%s\n", taskID)
		
		fmt.Printf("\nManual Terminal Hook:\n  cd ../.diffless-workspaces/%s\n", taskID)
		return nil
	},
}
