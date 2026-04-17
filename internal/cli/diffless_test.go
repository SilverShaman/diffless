package cli

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// setupMockRepo creates a temporary directory, runs git init, and returns the path.
// This strictly enforces isolated tests using t.TempDir() natively preventing main repo corruption.
func setupMockRepo(t *testing.T) string {
	t.Helper()
	repoDir := t.TempDir()

	// Initialize mock git repo
	cmd := exec.Command("git", "init")
	cmd.Dir = repoDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to git init: %v", err)
	}

	// Create an initial commit so worktrees can be branched off legally
	cmd = exec.Command("git", "commit", "--allow-empty", "-m", "Initial mock commit")
	cmd.Dir = repoDir
	if err := cmd.Run(); err != nil {
		t.Fatalf("failed to create initial commit: %v", err)
	}

	return repoDir
}

func TestDifflessSandboxing(t *testing.T) {
	// 1. Setup isolated physical environment
	repoDir := setupMockRepo(t)

	// Change directory to the mock repo so the CLI executes relative to the test environment,
	// inherently preventing execution against github.com/SilverShaman/diffless
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	if err := os.Chdir(repoDir); err != nil {
		t.Fatalf("failed to chdir into mock repo: %v", err)
	}

	taskID := "mock-test-task"
	expectedWorktreePath, _ := filepath.Abs(filepath.Join("..", ".diffless-workspaces", taskID))

	// Clean up global temp environment afterwards
	defer os.RemoveAll(filepath.Dir(expectedWorktreePath))

	// === 2. Verify `diffless start` ===
	t.Run("StartSandboxCommand", func(t *testing.T) {
		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)
		rootCmd.SetErr(buf)
		rootCmd.SetArgs([]string{"start", taskID})
		
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("start command failed: %v", err)
		}

		if _, err := os.Stat(expectedWorktreePath); os.IsNotExist(err) {
			t.Fatalf("mathematically failed: expected worktree path %s was not physically created", expectedWorktreePath)
		}
	})

	// === 3. Verify `diffless lockdown` ===
	t.Run("LockdownSecurityCommand", func(t *testing.T) {
		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)
		rootCmd.SetErr(buf)
		rootCmd.SetArgs([]string{"lockdown", taskID})
		
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("lockdown command failed: %v", err)
		}

		// Algorithmically assert Path Permissions
		info, err := os.Stat(expectedWorktreePath)
		if err != nil {
			t.Fatalf("failed to stat worktree: %v", err)
		}
		if info.Mode().Perm() != 0700 {
			t.Errorf("critical security failure: expected permissions 0700, got %v", info.Mode().Perm())
		}

		// Algorithmically assert Ephemeral credential generation
		envPath := filepath.Join(expectedWorktreePath, ".env")
		envInfo, err := os.Stat(envPath)
		if os.IsNotExist(err) {
			t.Fatalf("critical security failure: expected ephemeral .env was not generated")
		}
		if envInfo.Mode().Perm() != 0600 {
			t.Errorf("critical security failure: expected .env permissions 0600, got %v", envInfo.Mode().Perm())
		}
	})

	// === 4. Verify `diffless clean` ===
	t.Run("CleanPruningCommand", func(t *testing.T) {
		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)
		rootCmd.SetErr(buf)
		rootCmd.SetArgs([]string{"clean", taskID})
		
		if err := rootCmd.Execute(); err != nil {
			t.Fatalf("clean command failed: %v", err)
		}

		// Assert immediate physical termination
		if _, err := os.Stat(expectedWorktreePath); !os.IsNotExist(err) {
			t.Errorf("expected worktree path %s to be deleted, but it dynamically persists", expectedWorktreePath)
		}

		// Assert internal git reference termination
		cmd := exec.Command("git", "worktree", "list")
		cmd.Dir = repoDir
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git worktree list failed: %v", err)
		}
		if strings.Contains(string(out), taskID) {
			t.Errorf("expected git worktree list to inherently lack %s, but reference persisted: %s", taskID, string(out))
		}
	})
}
