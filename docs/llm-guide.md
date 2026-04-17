# Antigravity Agent Context & Operations Guide (`diffless`)

<system_context>
## Welcome to the Diffless Project, Antigravity
You are currently operating inside the `github.com/SilverShaman/diffless` repository. 

**What is Diffless?** 
We are building a highly specialized **CLI tool in Go** designed specifically to empower *you* (and other AI agents). Crucially, the primary focus of the `diffless` tool is to manage and sandbox agent operations across **any external repository** a user throws at you, not just the `github.com/SilverShaman/diffless` codebase you are currently inside! Its goal is to replace chaotic, conflict-prone trunk-based development with **AI-Augmented Branching** universally. 

By wrapping native `git worktree` mechanisms, this global CLI automatically grants you a physically isolated sandbox directory (`../.diffless-workspaces/<feature>`) in whatever target repo you are working on. Inside this sandbox, you can freely modify code, install conflicting dependencies, and run tests without ever polluting the human developer's primary workspace.

You are not building a standard web app; you are building the global infrastructure that optimizes your own workflow across all future projects. 
</system_context>

<architecture_and_domain>
## Core Domain Concepts
- **Sandbox Mode Integration**: The `diffless start` command is mathematically designed to transition you into a hermetic "Sandbox Mode", bounding your execution context entirely to the `git worktree`.
- **Zero-Trust Hardening**: The CLI forces strict 0700 permissions and ephemeral `.env` stripping to guarantee secure containment.
- **Semantic Merging**: We do not do raw text merges. You analyze intent and intelligently resolve logic conflicts.
- **Artifact-Driven PRs**: Instead of proposing a standard diff, the CLI orchestrates you to generate an **Artifact Package**: Execution Markdown, Mermaid Arch Diagrams, and Browser Subagent Validation Videos (using your `@browser_subagent` capability).

## Go Project Architecture Constraints
* **Language & Build**: Go 1.21+. The binary is compiled from `cmd/diffless/main.go`.
* **Standard Library First**: Keep dependencies slim. Favor `os/exec` and standard library patterns over massive 3rd-party modules. Only use modules like `spf13/cobra` where absolutely necessary for robust CLI semantics.
* **Internal Structure**:
  * `internal/cli`: Framework definitions for all CLI commands.
  * `internal/git`: Low-level subprocess drivers handling `git worktree` routing.
  * `internal/antigravity`: The bridge layer. This is where we write Go routines that invoke *your* IDE APIs, task groups, and subagent executors.
  * `internal/artifact`: Handlers for generating validation media, markdown, and diagrams for PRs.
</architecture_and_domain>

<agent_interaction_rules>
## How You Should Operate Here
As an Antigravity Agent, strictly adhere to these developmental practices:

1. **Be Precision-Driven:** Use specific tooling (like `replace_file_content` or `multi_replace_file_content`) to inject changes. Do not perform full-file rewrites manually via shell scripts.
2. **Be Production-Ready:** There are no placeholders. When writing or modifying Go code, output fully functional logic, including error handling. Do not leave `// TODO` comments.
3. **Validate With Tests:** When modifying CLI components or Antigravity API boundaries, always run `go test -v ./...` to guarantee you have not fractured the isolation constraints or sandbox lock-down capabilities.
4. **License Preservation:** Never mutate or remove the GPL-3.0 License.
5. **Protect Global State:** Remember that the CLI impacts `.git` databases and global IDE state (like `~/.gemini/antigravity/skills/diffless.sh`). Be delicate with repository-level file removals or deep tree modifications.
</agent_interaction_rules>

<reference_documentation>
## External Context Links

> **NOTE:** The content for all of the links below has been downloaded and aggregated locally into simple markdown format. You **do not** need to fetch these from the web going forward. Please reference `docs/references/reference.md` instead.

For deeper architectural alignment, query or reference the underlying system documentation directly:

**Git Native:**
- Git Worktree: https://git-scm.com/docs/git-worktree

**Antigravity Native APIs:**
- Get Started: https://antigravity.google/docs/get-started
- Models: https://antigravity.google/docs/models
- Agent Modes & Settings: https://antigravity.google/docs/agent-modes-settings
- Agent Permissions: https://antigravity.google/docs/agent-permissions
- Rules & Workflows: https://antigravity.google/docs/rules-workflows
- Task Groups: https://antigravity.google/docs/task-groups
- Browser Subagent: https://antigravity.google/docs/browser-subagent
- Strict Mode: https://antigravity.google/docs/strict-mode
- Sandbox Mode: https://antigravity.google/docs/sandbox-mode
- Workspace Commands: https://antigravity.google/docs/command
- Workspaces: https://antigravity.google/docs/workspaces
- Skills: https://antigravity.google/docs/skills
</reference_documentation>

<end_of_system_prompt>
