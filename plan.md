# Diffless Workflow Implementation Plan

## Overview
This document outlines the transition plan from Trunk-Based Development (TBD) to the **Diffless Workflow** (AI-Augmented Feature Branching & Gitflow). The goal is to optimize our repository environments for long-running, autonomous AI agents (like Google Antigravity, Claude Code, and Codex) while maintaining high visibility and control for human reviewers.

## Why This Plan?
As outlined in our initial concepts (see `docs/concept.md`), traditional Git workflows struggle with long-lived branches because merges are strictly *textual*. This leads to "merge hell," which forced the industry to adopt TBD as a defense mechanism to limit cognitive bandwidth exhaustion. 

The **Diffless Workflow** leverages the *semantic merging* capabilities of modern AI. AI can understand the architectural intent of code changes and resolve conflicts seamlessly—enabling agents to work for weeks in isolated, deep sandboxes without generating catastrophic merge conflicts.

---

## Phase 1: Isolated Agent Sandboxing
**Goal:** Prevent agent-driven changes from disrupting the main build.
- **Action for Developers:** Avoid committing complex work directly to `main` if it can be offloaded.
- **Action for AI Agents:** When an agent receives a task, it must **fork** the repository or create an isolated, long-lived **feature branch**.
- **Implementation Considerations:**
  - Define branch naming conventions, e.g., `agent/<agent-id>-<feature-name>`.
  - Configure CI/CD pipelines to isolate agent branches, giving the AI a true sandbox to safely experiment and iterate.

## Phase 2: Autonomous Semantic Merging
**Goal:** Eradicate textual merge conflicts.
- **Action for Developers:** Rely on AI IDEs to handle branch drift.
- **Action for AI Agents:** The agent must safely handle conflicts. If a long-lived agent branch drifts from `main`, the AI should periodically review the upstream changes, understand their intent, and synthetically re-implement them or resolve conflicts inside its sandbox autonomously.
- **Implementation Considerations:**
  - Introduce regular, automated rebasing triggers for active agent sandboxes.

## Phase 3: Artifact-Driven Reviews (Solving the "Monster PR")
**Goal:** Replace line-by-line review fatigue with conceptual "Diffless" artifact checks.
- **Action for Developers:** Stop reviewing raw 5,000-line code diffs. Review the rich artifacts instead, providing high-level feedback in natural language.
- **Action for AI Agents:** Before merging into `main`, agents must generate an **Artifact Package**.
- **Implementation Considerations:**
  - Agents must generate a `markdown` execution plan detailing what was achieved.
  - Agents must build `mermaid` architecture diagrams to map out large changes.
  - Agents must attach functionality validation in the PRs (e.g., MP4s of an agent navigating the UI to prove it works).

## Phase 4: Formal Gitflow via Agent "Skills"
**Goal:** Eliminate Gitflow overhead via conversational prompts.
- **Action for Developers:** Use intuitive language for complex repo operations (e.g., *"Cut a release branch, run integration tests, and generate a changelog artifact if tests pass"*).
- **Action for AI Agents:** Safely execute complex chains of Git operations (hotfixes, squashing, merging, cherry-picking) with precision.
- **Implementation Considerations:**
  - Ensure operations like `cherry-pick` and `rebase` are formalized as executable Agent "Skills."

---

## Next Steps
- Verify the local configuration of Antigravity or other AI IDE setups for semantic merges.
- Create an automated PR/Merge-Request template that forces the generation of execution plans and execution recordings in place of standard diffs.
- Distribute the Diffless infographic (`assets/infographic.webp`) to the engineering team.
