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

## Phase 2: Security Hardening & Zero-Trust Containment [COMPLETED]
**Goal:** Mathematically enforce the security boundaries of the generated agent sandboxes to prevent privilege escalation or credential harvesting.
- **Command Implementation:**
  - `diffless lockdown <task-id>`: Hardens the sandbox directory permissions (e.g. `chmod 700`). Automatically generates an ephemeral `.env` stripping out high-privilege production API keys in favor of safe development keys.
  - `diffless audit`: Scans the agent's worktree for unapproved binary executions or anomalous outbound network traffic patterns.
- **Action for AI Agents:** Agents operate under a Zero-Trust assumption. They must utilize the ephemeral keys provided in the sandbox `.env` and cannot route out to read the parent developer's global `~/.ssh` or `~/.aws` configurations.

## Phase 3: Test Harness for Core Sandboxing (Phases 1 & 2) [COMPLETED]
**Goal:** Build an automated execution suite in Go (`go test`) that mathematically proves the physical sandboxing and security constraints function safely against a real `.git` database.
- **Testing Implementation:**
  - **Isolated Test Data:** The test framework will dynamically generate a temporary local Git repository utilizing Go's `t.TempDir()`. This guarantees the mock `.git` repository safely exists in the OS temp directory, strictly ensuring the test environment never accidentally leaks or tracks into the main `diffless` repository.
  - **Command Verification:** It will programmatically run `diffless start mock-task`, verifying the `git worktree` spawns securely out of bounds.
  - **Security Assertion:** It will execute `diffless lockdown mock-task` and algorithmically assert that path permissions restrict to `0700` and the ephemeral `.env` generates securely.
  - **Detailed Reporting:** The framework will enforce verbose capturing (e.g., executing `go test -v -json`). This captures highly detailed, structured test results for every command execution, allowing CI/CD pipelines (and AI agents) to precisely parse why a native operation failed.
  - **Teardown:** Finally, it will invoke `diffless clean mock-task` to ensure native directory removal and Git graph pruning.
- **Action for CI/CD:** Ensures that future Diffless CLI builds cannot regress the isolation boundary, relying on the detailed test output logs as the authoritative truth for agents.

## Phase 4: Autonomous Semantic Merging [COMPLETED]
**Goal:** Abstract away branch drift and complex merges inside the CLI.
- **Command Implementation:**
  - `diffless sync`: The CLI checks if the current worktree has drifted from `main`. If significant textual conflicts exist, the CLI pipes the conflicting patches to the AI.
- **Testing Implementation:**
  - The `go test` harness invokes `diffless sync mock-task`, validating that the underlying `internal/ai` pipeline correctly processes and resolves branch conflict patches logically without failure, bypassing native Git text-merge failures.
- **Action for AI Agents:** The agent uses its semantic understanding to rewrite conflicting blocks based on intentionality, automatically resolving merges inside its isolated worktree.

## Phase 5: Artifact-Driven Reviews [COMPLETED]
**Goal:** Automatically compile rich PRs avoiding raw code-diff fatigue.
- **Command Implementation:**
  - `diffless propose`: Rather than a standard `git push`, the CLI orchestrates the AI to generate an **Artifact Package**.
- **Testing Implementation:**
  - The test framework logically ensures that `diffless propose` successfully spawns the internal directory structures (`.diffless-artifacts`) within the sandbox and algorithmically maps the PR validation (`.mp4`), trace logic (`.mermaid`), and notes (`.md`).
- **Action for AI Agents:** When triggered, the AI must inspect its worktree changes and generate:
  1. A `markdown` execution plan detailing what was achieved.
  2. A `mermaid` architecture diagram highlighting system changes.
  3. Evidence of success (e.g., triggering a headless browser recording to output an `.mp4`).
  The CLI bundles these files into the PR description.

## Phase 6: Natural Language Execution (Agent Skills) [COMPLETED]
**Goal:** Make the CLI usable as native "tools" for LLMs.
- **Action for Developers:** Use intuitive language for complex repo operations (e.g., *"Antigravity, diffless start a new feature for the login portal, build the UI, and diffless propose it when you are done."*)
- **Implementation Considerations:**
  - Provide a machine-readable JSON schema for the `diffless` CLI so that AI agents can natively invoke `start`, `sync`, `propose`, and `clean` with zero hallucination.

## Phase 7: Antigravity Skill Bindings [COMPLETED]
**Goal:** Integrate the CLI directly into Google Antigravity as natively callable "Skills".
- **Action for Developers:** Use the Antigravity interface or the chat input (e.g. `@diffless start feature`) to rapidly invoke CLI abstractions without ever touching a terminal.
- **Implementation Considerations:**
  - Author formal Antigravity Skill integration files (e.g., inside `.gemini/skills/`) that bind the `diffless` CLI commands to the Antigravity execution context.
  - Wire the IDE's Workspace API to automatically react to `diffless switch` directory transitions.

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
- The Diffless Workflow CLI architecture (Phases 1-7) is **100% COMPLETE**.
- Move towards marketing outreach, community distribution via Brew/Apt structures, and broad testing adoption.
