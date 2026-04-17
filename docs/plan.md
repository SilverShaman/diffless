# Diffless Workflow Implementation Plan

## Overview
This document outlines the transition plan to the **Diffless Workflow** via the creation of a lightweight **Diffless CLI**. The goal is to optimize environments for autonomous AI agents using `git worktree` for physical isolation, while leveraging semantic merging and artifact-driven reviews.

## Why a CLI?
To make this workflow effortless for both human developers and AI IDEs (like Google Antigravity or Claude Code), we will build the `diffless` CLI. Under the hood, the CLI heavily leverages `git worktree` to grant AI agents isolated, physical directories that share the same Git database footprint. This prevents AI agents from corrupting the human developer's active workspace state.

---

## Historical Information: Completed Phases (1-10)
To see the full architectural progression of the Diffless CLI, review the completed foundational phases below:

<details>
<summary>Click to expand historical phases (1-10)</summary>

1. **Phase 1: The Diffless CLI & Worktree Sandboxing** - Built the lightweight CLI that manages physical isolation via `git worktree` (`start`, `switch`, `clean`).
2. **Phase 2: Security Hardening & Zero-Trust Containment** - Mathematically enforced security boundaries (`chmod 700`) and ephemeral `.env` stripping (`lockdown`, `audit`).
3. **Phase 3: Test Harness for Core Sandboxing** - Built an automated execution suite in Go (`go test`) proving physical sandboxing constraints.
4. **Phase 4: Autonomous Semantic Merging** - Abstracted branch drift and complex merges inside the CLI (`sync`).
5. **Phase 5: Artifact-Driven Reviews** - Automatically compiled rich PRs avoiding raw code-diff fatigue (`propose`).
6. **Phase 6: Natural Language Execution (Agent Skills)** - Made the CLI usable as native "tools" for LLMs via JSON schemas.
7. **Phase 7: Antigravity Skill Bindings** - Integrated the CLI directly into Google Antigravity as natively callable "Skills".
8. **Phase 8: Linux Mint Packaging & Deployment** - Packaged and deployed the `diffless` CLI globally for Linux Mint.
9. **Phase 9: Cross-Platform Packaging & Distribution** - Expanded package support to macOS (Homebrew) and Windows.
10. **Phase 10: Advanced Sandbox Hardening** - Unified cross-platform OS-level containment natively (Linux Namespaces, macOS Sandbox, Windows AppContainer) via an opt-in `--jail` execution flag.

</details>

## Phase 11: Runtime Context & Documentation [PLANNED]
**Goal:** Improve the help and documentation during runtime for `diffless` to drastically reduce developer onboarding friction.
- **Context-Aware Help:** Overhaul the CLI's `-h` / `--help` flag to dynamically adjust its output based on the user's current environment (e.g., prioritizing `propose` and `sync` if invoked inside a sandbox, or `start` if invoked in the trunk).
- **Interactive Terminal Guides:** Implement a `diffless guide` command that prints beautiful, rich-text walkthroughs of the Agent-Augmented Gitflow directly in the terminal, negating the need to constantly reference web documentation.
- **Action for Developers & Agents:** Provides a completely self-documenting CLI experience, ensuring humans and AI agents always know what native actions are available and safe to execute in their current context.

---

## Phase 12: Expanded Sandboxing Ecosystems [PLANNED]
**Goal:** Broaden the `diffless run --jail` capabilities beyond native kernel primitives to support diverse container and reproducible environments.
- **NixOS Integration:** Implement native `systemd-nspawn` and `bubblewrap` (`bwrap`) providers to allow developers to leverage strict, declarative Nix sandbox configurations.
- **Docker-Based Isolation:** Add an OCI-compliant provider targeting the Docker Engine. This allows teams already using DevContainers or Dockerized workflows to jail AI agents inside standard, heavy-duty Docker containers.
- **Lightweight Containerization (Bubblewrap):** Provide a daemon-less, super lightweight isolation backend using `bwrap` for Linux users who want microsecond startup times, unprivileged execution, and extremely low resource overhead compared to a full Docker daemon.

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
- The Diffless Workflow CLI architecture (Phases 1-10) is **100% COMPLETE**.
- **ACTIVE:** Implement Phase 11 (Runtime Context & Documentation), move towards marketing outreach, community distribution via Brew/curl setups, and broad testing adoption.
