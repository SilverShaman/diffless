#!/usr/bin/env bash
# <antigravity_skill>
# name: diffless
# description: The native AI agent interface for the Diffless physical worktree CLI. Rapidly spin up sandboxes, run sync loops, and propose logic directly inside IDE execution bounds.
# </antigravity_skill>

set -e

# Dynamically intercept and compile the underlying Go binary if it's currently missing
if [ ! -f "diffless" ]; then
    echo "[Antigravity-Skill] Compiling native diffless Go execution payload..."
    go build -o diffless ./cmd/diffless/main.go
fi

echo "[Antigravity-Skill] 🚀 Routing IDE intent -> diffless $@"
./diffless "$@"
