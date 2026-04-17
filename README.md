# Diffless: The AI-Augmented Git Workflow

![Diffless Concept](assets/infographic.webp)

Welcome to the **Diffless Workflow** repository. This project proposes a paradigm shift in version control and development lifecycle in the era of AI-augmented IDEs (such as Google Antigravity, Claude Code, Cursor, Codex, etc.).

## The Concept

For years, we've relied on **Trunk-Based Development (TBD)** as our defense mechanism against human limitations. Merging large branches textually in Git often creates a nightmare scenario called "merge hell," so TBD forced developers to commit thin, daily increments. 

However, with advanced LLM capabilities, an agent environment enables **Semantic Merging**. When AI handles merging natively, it doesn't just look at textual code lines; it understands architectural intent, enabling synthetic code rewrites to resolve massive conflicts automatically. 

With these constraints broken, it’s time to move toward an **Agent-Augmented Gitflow**.

### Core Tenets

1. **Semantic Merging Over Textual Diffing**: Eliminate "merge hell" by letting AI synthetically reconstruct conflicted code based on intent.
2. **True Isolated Sandboxes**: Instead of risking a chaotic `trunk` build, async AI agents should operate in deep, long-lived feature branches or forks.
3. **Artifact-Driven PRs**: The "Monster PR" is dead. Agents replace raw 5,000-line diffs with rich Artifacts (videos, architectural diagrams, execution plans) that are conceptually reviewed by humans.
4. **Gitflow as a "Skill"**: Advanced Git maneuvers (cutting releases, cherry-picking, back-merging) simply become conversational prompts.

---

## Repository Structure

- **`plan.md`**  
  The formal, step-by-step implementation plan turning the Diffless concept into practice for human teams and AI tooling alike.
  
- **`docs/concept.md`**  
  The foundational theory file introducing why TBD was built for humans, and why sandboxes are built for AI.

- **`docs/Diffless.pdf`**  
  Detailed documentation showcasing the Diffless approach.

- **`assets/infographic.webp`**  
  Visual breakdown comparing the Human-Centric Era against the new Agent-Augmented Era of code workflows.

- **`assets/demo.mp4`**  
  A video recording showcasing an artifact-driven review process in action.

## Getting Started

1. Check out the infographic above for a clear visual representation of this shift.
2. Read the full problem statement and concept in [Concept](docs/concept.md).
3. Follow the rollout steps detailed in the [Diffless Workflow Plan](plan.md).
4. Watch the demo video below to see the transition of standard diffing over to Artifacts.

## Demo Video

https://github.com/SilverShaman/diffless/raw/main/assets/demo.mp4
