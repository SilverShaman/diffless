#!/usr/bin/env bash
set -e

VERSION="1.0.0"
ARCH=$(dpkg --print-architecture 2>/dev/null || echo "amd64")
PKG_NAME="diffless_${VERSION}_${ARCH}.deb"

# Ensure we're executing in this folder
cd "$(dirname "$0")"

if [ ! -f "$PKG_NAME" ]; then
    echo "=> Package $PKG_NAME not found! Building it first..."
    bash ./build.sh
fi

echo "=> Installing global diffless Linux Mint package (requires sudo)..."
sudo dpkg -i "$PKG_NAME"

echo "=> Deploying global Antigravity Skill payload..."
SKILLS_DIR="$HOME/.gemini/antigravity/skills"
mkdir -p "$SKILLS_DIR"
cat << 'EOF' > "$SKILLS_DIR/diffless.sh"
#!/usr/bin/env bash
# <antigravity_skill>
# name: diffless
# description: The native AI agent interface for the Diffless physical worktree CLI. Rapidly spin up sandboxes, run sync loops, and propose logic directly inside IDE execution bounds.
# </antigravity_skill>

set -e
export PATH="$PATH:/usr/local/bin"

if ! command -v diffless &> /dev/null; then
    echo "[Antigravity-Skill] Error: 'diffless' command not found." >&2
    exit 1
fi

echo "[Antigravity-Skill] 🚀 Routing IDE intent -> diffless $@"
diffless "$@"
EOF

chmod +x "$SKILLS_DIR/diffless.sh"

echo ""
echo "✅ Diffless Mint / APT installation complete!"
echo "The binary is installed globally to /usr/local/bin/diffless."
echo "The Antigravity skill payload is bound globally to $SKILLS_DIR/diffless.sh."
