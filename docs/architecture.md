# Diffless Architecture

The **Diffless Workflow** system is composed of several interdependent layers that abstract physical Git environment constraints away from AI coding agents, providing safe, autonomous execution within "zero-trust" boundaries.

## 1. Native CLI Layer (Go)
The core of the system is the `diffless` CLI, written in **Go (Golang)**. The choice of Go provides a high-performance, universally compilable binary free from external scripting dependencies (like Python or Node.js). 

- **Environment Manager:** Manages underlying `git worktree` commands to orchestrate isolated physical sandboxes.
- **Security Boundary Enforcer:** Applies minimal permission models and ephemeral `.env` credentials dynamically whenever a sandbox is spun up.
- **Workflow Orchestrator:** Commands the system to trigger semantic merges and artifact PR generations.

## 2. Git Worktree Sub-System
Unlike typical branch checkouts, which replace the files in your current working directory, Diffless leverages `git worktree`.
- **Physical Isolation:** Each task requires a secondary directory (e.g., `../.diffless-workspaces/<feature>`).
- **Shared DB:** The `.git` database and object storage are centrally shared among all worktrees, making branching and caching nearly instantaneous with zero additional network overhead.
- **Agent Containment:** An AI cannot mistakenly destroy the original trunk environment because it is literally bounded to a completely separate OS path.

## 3. IDE Integration (Antigravity Skills)
The CLI operates under the hood while exposing itself cleanly up to the top level via IDE interceptors.
- **Global Deployment:** The `install.sh` script deploys a universal interceptor script (`diffless.sh`) into `~/.gemini/skills/`.
- **Chat Interception:** Developers or AI supervisor processes can trigger commands natively (e.g. `@diffless start feature`).
- **Context Routing:** The script natively catches the command, dynamically invokes the compiled Go binary in the user's `PATH`, and binds the IDE window natively into the new `git worktree` path.

## 4. Semantic Merge & Artifact Pipeline
Diffless re-architects how Pull Requests are formulated.
- **Diff Resolution Layer:** Resolves massive refactors textually by passing broken patches through an LLM to interpret logical intent rather than simple line matching.
- **Generative Artifact Assembly:** When a feature is completed, the agent triggers `diffless propose`. Natively compiled outputs (like Subagent UI testing `.mp4` videos, Markdown execution logs, and Mermaid architecture diagrams) securely track to `.diffless-artifacts/`. 
- **PR Aggregation:** The finalized PR displays the summarized artifacts immediately to human reviewers rather than thousands of lines of unreadable code drift.

## 5. Cross-Platform OS Jailing (Advanced Hardening)
While physical `git worktree` isolation prevents repository drift, Diffless also supports deep, OS-level cryptographic containment via the `diffless run` command and its optional `--jail` flag.
- **Unified Security Interface:** The core CLI exposes a single, OS-agnostic `JailProvider` interface to minimize code divergence.
- **Native Primitives:** Depending on the compiled host, it dynamically enforces boundaries using the OS's deepest primitives:
  - **Linux:** Kernel Namespaces (`CLONE_NEWNS`, `CLONE_NEWNET`) and eBPF syscall filtering.
  - **Windows:** Host Compute System (HCS), AppContainer capability profiles, and Windows Filtering Platform (WFP).
  - **macOS:** App Sandbox (`sandbox-exec` profiles) and Endpoint Security (ES).
  - **NixOS:** Declarative application sandboxing using `bubblewrap` and system-level isolation via `systemd-nspawn`, leveraging the reproducible Nix ecosystem.
- **Containerized Isolation Backends:** Beyond native OS primitives, the `JailProvider` also supports standard container ecosystems:
  - **Docker (OCI-Based):** Standard, portable container jailing via the Docker engine for environments heavily reliant on containerized dev flows.
  - **Bubblewrap (Super Lightweight):** Extremely fast, unprivileged, daemon-less process sandboxing that boots instantly without the overhead of full OCI containers.
- **Opt-In Execution:** Because full OS-level isolation can complicate basic local development workflows, it is strictly opt-in (`--jail`). If omitted, executions safely fall back to the standard Phase 2 user-space boundary (0700 file permissions and `.env` stripping).
