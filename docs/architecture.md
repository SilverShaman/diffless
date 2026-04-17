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
