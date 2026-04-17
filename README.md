# Diffless: The AI-Augmented Git Workflow

![Diffless Concept](assets/infographic.webp)

Welcome to the **Diffless Workflow** repository. This project proposes a paradigm shift in version control and development lifecycles in the era of AI-augmented IDEs (such as Google Antigravity, Claude Code, Cursor, Codex, etc.).

## The Concept

For years, we've relied on **Trunk-Based Development (TBD)** as our defense mechanism against human limitations. Merging large branches textually in Git often creates a nightmare scenario called "merge hell," so TBD forced developers to commit thin, daily increments. 

However, with advanced LLM capabilities, an agent environment enables **Semantic Merging**. With AI handling code merging, the system doesn't just read Git diff text; it understands architectural intent, enabling synthetic code rewrites to resolve massive conflicts seamlessly. 

With these constraints broken, it’s time to move toward an **Agent-Augmented Gitflow**, powered by `git worktree` and managed via a dedicated CLI.

### Core Tenets

1. **Semantic Merging**: Eradicate "merge hell" by letting AI synthetically reconstruct conflicted code based on intent and functionality, replacing raw textual diffing.
2. **True Physical Sandboxes (via `git worktree`)**: If an AI agent operates in your main repository directory, it will overwrite your files and break your build. Using `git worktree`, AI agents instead check out into completely separate, hidden physical directories while sharing your exact `.git` database.
3. **Artifact-Driven PRs**: The "Monster PR" is dead. Agents replace raw 5,000-line diffs with rich Artifacts (videos, architectural diagrams, execution plans) that humans can quickly review conceptually.
4. **Gitflow CLI Abstraction**: Complex workspace management is abstracted into simple CLI commands that act natively as Agent Skills.

---

## The Diffless CLI

To make this workflow effortless for both developers and AI, we envision the lightweight **Diffless CLI** wrapping native `git worktree` commands:

- `diffless start <feature>`: Creates a physically isolated AI sandbox (e.g., `../.diffless-workspaces/feature`).
- `diffless sync`: Synthetically and semantically rebases agent code against the main trunk.
- `diffless propose`: Orchestrates the AI to generate markdown plans, architecture diagrams, and testing videos for a pull request.
- `diffless clean`: Prunes the external worktree safely once merged.

## Repository Structure

- **`plan.md`**  
  The step-by-step implementation plan for building the `diffless` CLI and enabling true physical agent sandboxing.
  
- **`docs/concept.md`**  
  The foundational theory detailing why TBD evolved for humans, and why `git worktree` sandboxes are necessary for AI.

- **`docs/Diffless.pdf`**  
  Detailed documentation showcasing the Diffless approach.

- **`assets`**  
  Contains visual assets like the infographic and the demo MP4 showcasing the workflow in action.

## Getting Started

1. Check out the infographic above for a clear visual representation of this shift.
2. Read the full problem statement and concept in [Concept](docs/concept.md).
3. Follow the CLI architectural steps detailed in the [Diffless Workflow Plan](plan.md).
4. Watch the demo video below to see the transition of standard diffing over to Artifacts.

## Demo Video

https://github.com/SilverShaman/diffless/raw/main/assets/demo.mp4
