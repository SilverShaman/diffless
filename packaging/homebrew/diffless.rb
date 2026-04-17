class Diffless < Formula
  desc "AI-augmented git workflow CLI builder"
  homepage "https://github.com/SilverShaman/diffless"
  url "https://github.com/SilverShaman/diffless/archive/refs/tags/v0.0.1.tar.gz"
  sha256 "0000000000000000000000000000000000000000000000000000000000000000"
  license "GPL-3.0-only"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"diffless", "./cmd/diffless/main.go"
    
    # We install the Antigravity skill by providing an install script
    # because Homebrew sandboxes writes outside of the Cellar.
    
    pkgshare.install Dir["*"]
    
    (bin/"diffless-install-skill").write <<~EOS
      #!/usr/bin/env bash
      echo "=> Installing global Antigravity Skill for macOS..."
      SKILLS_DIR="$HOME/.gemini/antigravity/skills/diffless"
      mkdir -p "$SKILLS_DIR/scripts"

      cat << 'EOF2' > "$SKILLS_DIR/SKILL.md"
      ---
      name: diffless
      description: The native AI agent interface for the Diffless physical worktree CLI. Rapidly spin up sandboxes, run sync loops, and propose logic directly inside IDE execution bounds.
      ---

      # Diffless Skill
      This skill provides the native AI agent interface for the Diffless physical worktree CLI.
      EOF2

      cat << 'EOF2' > "$SKILLS_DIR/scripts/diffless.sh"
      #!/usr/bin/env bash
      set -e
      export PATH="$PATH:$HOME/.local/bin:$HOME/go/bin:/opt/homebrew/bin"
      
      if ! command -v diffless &> /dev/null; then
          echo "[Antigravity-Skill] Error: 'diffless' command not found." >&2
          exit 1
      fi
      
      echo "[Antigravity-Skill] 🚀 Routing IDE intent -> diffless $@"
      diffless "$@"
      EOF2

      chmod +x "$SKILLS_DIR/scripts/diffless.sh"
      echo "=> Installed Antigravity skill payload to $SKILLS_DIR"
    EOS
  end

  def caveats
    <<~EOS
      Diffless has been installed!
      
      To complete the Google Antigravity integration, please run:
        diffless-install-skill
    EOS
  end

  test do
    system "#{bin}/diffless", "--version"
  end
end
