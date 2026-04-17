# Diffless Workflow Implementation Plan

## Overview
This document outlines the transition plan to the **Diffless Workflow** via the creation of a lightweight **Diffless CLI**. The goal is to optimize environments for autonomous AI agents using `git worktree` for physical isolation, while leveraging semantic merging and artifact-driven reviews.

## Why a CLI?
To make this workflow effortless for both human developers and AI IDEs (like Google Antigravity or Claude Code), we will build the `diffless` CLI. Under the hood, the CLI heavily leverages `git worktree` to grant AI agents isolated, physical directories that share the same Git database footprint. This prevents AI agents from corrupting the human developer's active workspace state.

---

## Phase 1: The Diffless CLI & Worktree Sandboxing [COMPLETED]
**Goal:** Build a lightweight CLI that manages physical isolation via `git worktree`.
- **Command Implementation:**
  - `diffless start <task-id>`: Automatically runs `git worktree add -b <task-id> ../.diffless-workspaces/<task-id>`. This creates a physically isolated directory sharing the main `.git` database.
  - `diffless switch <task-id>`: Navigates the terminal or AI environment to the isolated worktree.
  - `diffless clean`: Prunes and removes completed worktrees using `git worktree remove` and `git worktree prune`.
- **Action for AI Agents:** Agents execute `diffless start` to safely experiment, install conflicting dependencies, or crash the local server in their own physical directory without impacting the developer's trunk.

## Phase 2: Autonomous Semantic Merging
**Goal:** Abstract away branch drift and complex merges inside the CLI.
- **Command Implementation:**
  - `diffless sync`: The CLI checks if the current worktree has drifted from `main`. If significant textual conflicts exist, the CLI pipes the conflicting patches to the AI.
- **Action for AI Agents:** The agent uses its semantic understanding to rewrite conflicting blocks based on intentionality, automatically resolving merges inside its isolated worktree.

## Phase 3: Artifact-Driven Reviews
**Goal:** Automatically compile rich PRs avoiding raw code-diff fatigue.
- **Command Implementation:**
  - `diffless propose`: Rather than a standard `git push`, the CLI orchestrates the AI to generate an **Artifact Package**.
- **Action for AI Agents:** When triggered, the AI must inspect its worktree changes and generate:
  1. A `markdown` execution plan detailing what was achieved.
  2. A `mermaid` architecture diagram highlighting system changes.
  3. Evidence of success (e.g., triggering a headless browser recording to output an `.mp4`).
  The CLI bundles these files into the PR description.

## Phase 4: Natural Language Execution (Agent Skills)
**Goal:** Make the CLI usable as native "tools" for LLMs.
- **Action for Developers:** Use intuitive language for complex repo operations (e.g., *"Antigravity, diffless start a new feature for the login portal, build the UI, and diffless propose it when you are done."*)
- **Implementation Considerations:**
  - Provide a machine-readable JSON schema for the `diffless` CLI so that AI agents can natively invoke `start`, `sync`, `propose`, and `clean` with zero hallucination.

---

## Antigravity IDE Optimizations
Because the primary target environment for the Diffless CLI is **Google Antigravity**, we will intentionally optimize the CLI hooks for its native architectural features:

- **Workspaces & Sandbox Mode**: When `diffless start` generates a `git worktree`, the CLI should use the Antigravity API to bind the agent purely to the new physical path. Engaging **Sandbox Mode** ensures the agent is hermetically sealed and cannot execute commands or read state from the human developer's parent `trunk` directory.
- **Browser Subagent Artifacts**: During `diffless propose`, the CLI should programmatically trigger Antigravity’s native `browser_subagent` tool. It will spin up the feature's dev server, navigate the frontend, and record the UI interactions, seamlessly injecting the `.mp4` or `.webp` artifacts into the PR.
- **Rules, Workflows & Task Groups**: Rather than using pure bash automation, the CLI will package standard operations into Antigravity **Task Groups** and **Workflows**. A dedicated ruleset (`.antigravity/diffless-workflow.yaml`) can safely orchestrate Phase 3: *Task 1: Semantic Merge*, *Task 2: Subagent Recording*, *Task 3: Diagram Generation*.
- **Strict Mode Commands**: By registering standard operations (`start`, `sync`, `propose`, `clean`) as native `/commands` inside the IDE, and engaging **Strict Mode**, we guarantee the AI doesn't hallucinate raw Git operations, strictly routing its actions through the safe Diffless APIs.

---

## Technical Architecture (Go)
To ensure the `diffless` CLI is lightweight, high-performance, and easily distributable across operating systems as a single compiled binary, the primary programming language will be **Go (Golang)**. 

### Proposed Project Structure
Following idiomatic Go project layouts, the repository acts as the host for the CLI source code:

```text
diffless/
├── cmd/
│   └── diffless/
│       └── main.go       # Entry point for the compiled CLI binary
├── internal/             # Encapsulated application logic
│   ├── cli/              # Command parsing (e.g., using spf13/cobra)
│   ├── git/              # Subprocess wrappers executing `git worktree`
│   ├── antigravity/      # API integrations for Antigravity Workspaces & Subagents
│   └── artifact/         # Generators for Markdown execution plans & diagram compilation
├── docs/                 # Diffless conceptual documentation
├── assets/               # Demonstration media
├── go.mod                # Go module definition
├── LICENSE               # GPL-3.0 Open Source License
└── README.md             
```

---

## Next Steps
- Begin implementation of Phase 2: Autonomous Semantic Merging (`diffless sync`).
- Hook up an LLM API client in Go to process textual merge conflicts synthetically.
- Register the newly built `diffless` CLI Sandbox commands as custom integrations in standard AI IDEs.
