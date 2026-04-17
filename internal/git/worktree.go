package git

import (
	"fmt"
	"os/exec"
)

// AddWorktree creates a new git worktree with a new branch
func AddWorktree(taskID string, path string) error {
	cmd := exec.Command("git", "worktree", "add", "-b", taskID, path)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to add worktree: %v, out: %s", err, string(out))
	}
	return nil
}

// RemoveWorktree safely removes the given worktree from git
func RemoveWorktree(path string) error {
	cmd := exec.Command("git", "worktree", "remove", path)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to remove worktree: %v, out: %s", err, string(out))
	}
	return nil
}

// PruneWorktrees runs git worktree prune
func PruneWorktrees() error {
	cmd := exec.Command("git", "worktree", "prune")
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to prune worktrees: %v, out: %s", err, string(out))
	}
	return nil
}
