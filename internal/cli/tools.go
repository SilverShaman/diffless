package cli

import (
	"encoding/json"

	"github.com/spf13/cobra"
)

type ToolSchema struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Parameters  []string `json:"parameters"`
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Dumps CLI functions into LLM-readable JSON schemas",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Encapsulate all our CLI actions as pure logic hooks
		schemas := []ToolSchema{
			{Name: "start", Description: "Creates a physically isolated AI sandbox via git worktree.", Parameters: []string{"task-id"}},
			{Name: "switch", Description: "Transitions IDE context to the isolated path.", Parameters: []string{"task-id"}},
			{Name: "lockdown", Description: "Hardens permissions and replaces production keys with ephemeral test configs.", Parameters: []string{"task-id"}},
			{Name: "audit", Description: "Scans for binary anomalies and outbound request traces.", Parameters: []string{"task-id"}},
			{Name: "clean", Description: "Prunes the external worktree safely, erasing isolation footprints.", Parameters: []string{"task-id"}},
			{Name: "sync", Description: "Synthetically and semantically rebases agent code via Antigravity ML.", Parameters: []string{"task-id"}},
			{Name: "propose", Description: "Orchestrates markdown generation, architecture maps, and UI testing video pipelines.", Parameters: []string{"task-id"}},
		}

		b, err := json.MarshalIndent(schemas, "", "  ")
		if err != nil {
			return err
		}

		// Print directly to stdOut without decorative string headers ensuring raw ML parser compliance
		cmd.Println(string(b))
		return nil
	},
}

var toolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "Generates native Natural Language execution bindings for LLM integration",
}

func init() {
	toolsCmd.AddCommand(exportCmd)
}
