# Time to Switch your Git Workflow?

For years, we’ve relied on Trunk-Based Development (TBD) as our gold standard. Why? Because as humans, we have a limited cognitive bandwidth. TBD forces us to merge small, daily increments to avoid the "merge hell" and monster Pull Requests that come with Gitflow or long-lived feature branches. TBD was a necessary defense mechanism against our own limits.  

However, looking at the new LLM capabilities rolling out in AI IDEs today, **I have a theory: it’s time to move away from TBD and return to isolated, long-lived feature branches—powered by Git Worktrees and managed via a lightweight CLI.**

Because **Antigravity, Claude Code, Cursor, and Codex** can autonomously handle complex semantic merges, articulate massive changes via Artifacts, and embed complex version-control operations as executable "skills," the very constraints that made TBD necessary are gone.

Here are the four reasons why I believe **The Diffless Workflow**, powered by Git Worktrees, is our next move:

**1. The End of "Merge Hell" (Semantic vs. Textual Merging)**
Git handles merges *textually*. Untangling a 3-week branch divergence is a nightmare for humans. AI doesn't just read text; it understands semantic intent. If a long-lived agent branch drifts from `main`, the AI can autonomously read the conflicting changes, understand the architectural goals of both, and synthetically rewrite the code to resolve the conflict. When AI does the heavy lifting, the "penalty" of long-lived branches disappears.

**2. Agents Need True Physical Sandboxes (Enter Git Worktree)**  
If an AI agent operates in your main repository directory, it will constantly overwrite your files, ruin your IDE state, and break your local build. But wholesale repository cloning wastes disk space and breaks shared context. 

The solution is `git worktree`. A worktree allows multiple branches to be checked out simultaneously in separate physical directories, all linked to the same `.git` database. By wrapping `git worktree` in a lightweight **Diffless CLI**, an AI agent can spin up `diffless start feature-x`, executing complex, build-breaking code in a completely isolated directory (e.g., `../.diffless/feature-x`), while the human developer works uninterrupted in the main trunk.

**3. "Artifacts" Solve the Monster PR Problem**  
Reviewing massive PRs is a bottleneck. Diffless fundamentally changes this by generating rich **Artifacts**. When an agent finishes a worktree branch, it doesn't just hand us a raw 5,000-line code diff. Using `diffless propose`, the agent outputs:  
- Architecture diagrams  
- Markdown execution plans  
- Browser recordings of the feature working  

We review Artifacts, leave intuitive feedback, and let the agent incorporate it async. Strict review gates become an advantage rather than a bottleneck.  

**4. Gitflow Overhead is Abstracted by the CLI**  
Complex Gitflow operations—cherry-picking, branching, worktree cleanup, and back-merging—require significant overhead. By creating a lightweight **Diffless CLI**, we encapsulate these bare-metal Git operations into standard Agent "Skills." 

**The Takeaway**  
Trunk-Based Development is an optimization designed for human limits. As we shift to agent-driven development, our `trunk` should revert to being a pristine vault for perfectly verified code. The messy, long-term construction work should be offloaded to AI agents operating in deep, physically isolated Git Worktrees, managed seamlessly by a lightweight CLI.
