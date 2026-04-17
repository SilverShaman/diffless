package cli

import (
	"fmt"
	"path/filepath"

	"github.com/SilverShaman/diffless/internal/artifact"
	"github.com/spf13/cobra"
)

var proposeCmd = &cobra.Command{
	Use:   "propose [task-id]",
	Short: "Orchestrates an artifact-driven PR rather than a raw code diff push",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		taskID := args[0]
		worktreePath, _ := filepath.Abs(fmt.Sprintf("../.diffless-workspaces/%s", taskID))

		cmd.Printf("Orchestrating Artifact-Driven PR generation for '%s'...\n", taskID)

		if err := artifact.Generate(worktreePath, taskID); err != nil {
			return fmt.Errorf("failed to mathematically generate artifact package bounds: %w", err)
		}

		cmd.Println("✓ Generated comprehensive Markdown execution summary.")
		cmd.Println("✓ Generated Mermaid system architecture trace.")
		cmd.Println("✓ Triggered browser subagent and packaged .mp4 regression video.")
		cmd.Println("✓ Bundled Artifact Package successfully. PR is safely staged for human review without code diff fatigue.")

		return nil
	},
}
