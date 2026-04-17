package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/SilverShaman/diffless/internal/security"
	"github.com/spf13/cobra"
)

var enableJail bool

var runCmd = &cobra.Command{
	Use:   "run [task-id] [command...]",
	Short: "Executes a command inside the workspace (optionally inside an OS-level jail)",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		worktreePath, _ := filepath.Abs(fmt.Sprintf("../.diffless-workspaces/%s", taskID))
		command := args[1:]

		if enableJail {
			fmt.Printf("Jailing execution for '%s'...\n", taskID)
			if err := security.RunInJail(worktreePath, command); err != nil {
				return fmt.Errorf("jail execution failed: %w", err)
			}
			return nil
		}

		fmt.Printf("Executing %v normally within '%s'...\n", command, taskID)
		
		c := exec.Command(command[0], command[1:]...)
		c.Dir = worktreePath
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			return fmt.Errorf("execution failed: %w", err)
		}

		fmt.Println("✓ Process completed.")
		return nil
	},
}

func init() {
	runCmd.Flags().BoolVarP(&enableJail, "jail", "j", false, "Enable Advanced OS-level sandboxing (Phase 10)")
}
