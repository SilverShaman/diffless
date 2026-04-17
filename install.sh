#!/usr/bin/env bash

set -e

echo "=> Installing diffless workflow globally..."

# Ensure go is in PATH for compilation
export PATH=$PATH:/usr/local/go/bin

# 1. Compile the CLI binary
echo "=> Compiling diffless CLI..."
BIN_DIR="$HOME/.local/bin"
mkdir -p "$BIN_DIR"

if [ -f "./go.mod" ]; then
    # We are inside the diffless source directory
    go build -o "$BIN_DIR/diffless" ./cmd/diffless/main.go
    echo "=> Installed diffless binary to $BIN_DIR/diffless"
else
    # Install from remote (if running as a curl script elsewhere)
    # Using 'go install' directly to standard Go bin path
    echo "=> Installing via 'go install'..."
    go install github.com/SilverShaman/diffless/cmd/diffless@latest
    # Find the go bin dir
    GO_BIN=$(go env GOPATH)/bin
    BIN_DIR="$GO_BIN"
    echo "=> Installed diffless binary to $GO_BIN/diffless"
fi


# 2. Deploy the Antigravity Skill
echo "=> Installing global Antigravity Skill..."
SKILLS_DIR="$HOME/.gemini/antigravity/skills"
mkdir -p "$SKILLS_DIR"
cat << 'EOF' > "$SKILLS_DIR/diffless.sh"
#!/usr/bin/env bash
# <antigravity_skill>
# name: diffless
# description: The native AI agent interface for the Diffless physical worktree CLI. Rapidly spin up sandboxes, run sync loops, and propose logic directly inside IDE execution bounds.
# </antigravity_skill>

set -e

# PATH fallback for typical local binary locations
export PATH="$PATH:$HOME/.local/bin:$HOME/go/bin"

if ! command -v diffless &> /dev/null; then
    echo "[Antigravity-Skill] Error: 'diffless' command not found. Ensure you ran the install script correctly and diffless is in your PATH." >&2
    exit 1
fi

echo "[Antigravity-Skill] 🚀 Routing IDE intent -> diffless $@"
diffless "$@"
EOF

chmod +x "$SKILLS_DIR/diffless.sh"
echo "=> Installed Antigravity skill payload to $SKILLS_DIR/diffless.sh"

echo ""
echo "✅ Diffless globally installed and secured!"
echo "Make sure $BIN_DIR is in your \$PATH."
echo "You can now instruct AI agents with '@diffless start <feature>' across ANY Git repository."
