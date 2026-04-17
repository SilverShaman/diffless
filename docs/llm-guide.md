# LLM & Agent Context Guide (`diffless`)

<system_context>
## Core Objective
You are an AI coding assistant. You are currently operating inside the `github.com/SilverShaman/diffless` repository.
This repository is building the **Diffless CLI** in **Go**. 

The fundamental purpose of this CLI is to wrap `git worktree` commands to provide true, physical sandbox isolation for AI agents working in codebases. Do NOT assume this is a standard web application or script. It is a strictly compiled Go binary.

## Domain Model
- **Diffless**: The conceptual workflow replacing Trunk-Based Development with AI-Augmented Branching.
- **Git Worktree**: The underlying native Git mechanism that allows parallel physical directories attached to the same `.git` database.
- **Sandbox**: The physically isolated directory (e.g., `../.diffless-workspaces/<feature-id>`) where AI agents safely write code without breaking the human developer's IDE in the main trunk.
- **Semantic Merging**: The process by which an AI agent intentionally merges code by understanding the *architectural intent* rather than just blindly merging textual Git diffs.
- **Artifacts**: Pull requests in this workflow NEVER consist of raw code diffs. They consist of Execution Plans (`markdown`), Architecture Diagrams (`mermaid`), and Validation Videos (`mp4`/`webp`).
</system_context>

<architecture_rules>
## Go Project Constraints
- **Language**: Go 1.21+
- **Entrypoint**: `cmd/diffless/main.go`
- **Application Logic**: All domain logic must sit inside `internal/`.
  - `internal/cli`: Framework definitions for terminal commands (e.g., using `spf13/cobra`).
  - `internal/git`: Subprocess execution wrappers directly calling `git worktree`, `git add`, `git status`, etc.
  - `internal/antigravity`: Handlers for integrating with the Antigravity IDE APIs (Workspaces, Sandbox boundaries, `browser_subagent` execution).
  - `internal/artifact`: Generation of the PR markdown, mermaid diagrams, and video embeddings.
- **Dependencies**: Keep external dependencies to an absolute bare minimum. Rely strictly on the Go standard library (e.g., `os/exec` for git shell execution) wherever possible.
</architecture_rules>

<contribution_protocol>
## Agent Interaction Rules
1. **Never alter the GPL-3.0 License**. 
2. **Sequential Thinking**: When constructing the CLI, implement one subsystem completely before moving to the next (e.g., build and test the `git worktree` bindings entirely before moving to cobra CLI scaffolding).
3. **No Placeholders**: When instructed to write a Go file, generate the complete, production-ready implementation. Do not use `// TODO: implement logic here` responses. 
4. **Physical Data Safety**: If you are asked to test the CLI within the host repository, be highly aware that it mutates `.git` states. Do not execute destructive repository purge commands.
</contribution_protocol>

<antigravity_specific_hooks>
## System Instructions for Antigravity Agents
If you are Google Antigravity (or another advanced IDE agent) reading this context file:
- **Sandbox Mode Awareness**: Understand that the CLI you are helping to build is mathematically designed to trigger your native "Sandbox Mode" when `diffless start` runs. We are building the host infrastructure for *you*.
- **Browser Subagent**: Be prepared to embed code that securely triggers your `@browser_subagent` for UI validation testing during the `diffless propose` pipeline. 
- **Precision Edits**: When the human USER asks you to edit existing `.go` files across this repository, strictly adhere to your internal specific string-replacement tools (e.g., `multi_replace_file_content`) to prevent exhausting your context window on raw rewrites. 
</antigravity_specific_hooks>

<end_of_system_prompt>
