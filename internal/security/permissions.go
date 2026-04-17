package security

import (
	"os"
)

// HardenDirectory locks down a directory to owner-only permissions
func HardenDirectory(path string) error {
	// 0700 grants Read/Write/Execute for owner ONLY, enforcing Zero-Trust containment
	// against other users or lower privilege processes sharing the filesystem.
	return os.Chmod(path, 0700)
}
