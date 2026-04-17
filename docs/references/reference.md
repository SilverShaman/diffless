# Antigravity Reference Guide

This guide aggregates the core documentation for Google Antigravity, including agent behavior, workspaces, and models.

<!-- Document: get-started.md -->
## Getting Started

### Download

Please visit [antigravity.google/download](https://antigravity.google/download) to download Google Antigravity.

*   **macOS**: macOS versions with Apple security update support. This is typically the current and two previous versions. Min Version 12 (Monterey), X86 is not supported
*   **Windows**: Windows 10 (64 bit)
*   **Linux**: glibc >= 2.28, glibcxx >= 3.4.25 (e.g. Ubuntu 20. Debian 10, Fedora 36, RHEL 8)

The application will prompt when updates are available:

![Update Available](https://antigravity.google/assets/image/docs/restart-to-update.png)

### Basic Navigation

The Agent Manager can be opened from the Editor via the button on the top bar or via keyboard shortcut `Cmd + E`:

![Editor Open Agent Manager](https://antigravity.google/assets/image/docs/editor-open-agent-manager.png)

Similarly, from the Agent Manager, the Editor can be opened from any workspace via the “Focus Editor” option in the workspace’s drop down. When focused on a workspace, the Editor can be opened from any of the “Open Editor” buttons, or via the keyboard shortcut `Cmd + E`.

![Agent Manager Open Editor](https://antigravity.google/assets/image/docs/agent-manager-open-editor.png)

---

<!-- Document: workspaces.md -->
## Workspaces

In the Agent Manager, you can work across multiple workspaces simultaneously. In order to open a new workspace, just select the button in the left sidebar and select a starting folder. At any point, you can switch between conversations across workspaces through the left sidebar.

![Switch Workspaces](https://antigravity.google/assets/image/docs/workspaces/switch_workspace.png)

To start a new conversation within a workspace, either select the desired workspace from the Start Conversation tab or hit the Plus button next to the workspace name in the sidebar.

![Start Conversation Within Workspace](https://antigravity.google/assets/image/docs/workspaces/start_within_workspace.png)

---

<!-- Document: sandbox-mode.md -->
## Sandboxing Terminal Commands

Sandboxing provides kernel-level isolation for terminal commands executed by the Agent. When enabled, commands run in a restricted environment with limited file system and network access, protecting your system from unintended modifications.

Sandboxing is currently disabled by default, but this may change in future releases. It is supported on macOS and Linux. On macOS, it leverages Seatbelt (`sandbox-exec`), Apple's kernel-level sandboxing mechanism. On Linux, it uses `nsjail` for process isolation.

### Enabling Sandboxing

You can enable or disable sandboxing in Antigravity User Settings. Toggle "Enable Terminal Sandboxing" to turn sandboxing on or off. When enabled, you can also control network access separately using the "Sandbox Allow Network" toggle.

![Sandbox settings toggles](https://antigravity.google/assets/image/docs/sandbox-settings-toggle.png)

### Restrictions

When sandboxing is enabled, the Agent's terminal commands are subject to the following restrictions:

*   **File System**: Commands can only write to your designated workspace directory and essential system locations. This prevents the Agent from accidentally deleting or modifying files outside your project.

![File system operation blocked by sandbox](https://antigravity.google/assets/image/docs/sandbox-filesystem-denied.png)

*   **Network Access**: Network connectivity can be independently controlled. Use the "Sandbox Network Access" toggle in Antigravity User Settings to allow or deny network access while maintaining file system restrictions.

Here's an example of a command being blocked due to network restrictions:

![Sandbox network denial example](https://antigravity.google/assets/image/docs/sandbox-network-denied.png)

### Handling Sandbox Violations

If a command fails due to sandbox restrictions, you'll see a message indicating the failure may be sandbox-related. To resolve this, you can:

*   **Disable Sandbox Permanently**: Turn off sandboxing entirely in Antigravity User Settings.
*   **Bypass for a Single Command**: When using "Request Review" mode, you can choose to run an individual command with or without sandbox restrictions.

In "Request Review" mode, you'll see a "Bypass Sandbox" option when prompted to run a command:

![Bypass Sandbox option in Request Review mode](https://antigravity.google/assets/image/docs/sandbox-bypass-option.png)

### Interaction with Strict Mode

When strict mode is enabled, sandboxing is automatically activated with network access denied. This ensures maximum protection when operating in a strict environment.

![Sandbox settings in strict mode](https://antigravity.google/assets/image/docs/sandbox-secure-mode-settings.png)

---

<!-- Document: strict-mode.md -->
## Strict Mode

Strict mode provides enhanced security controls for the Agent, allowing you to restrict its access to external resources and sensitive operations. When strict mode is enabled, several security measures are enforced to protect your environment.

### Features

#### Browser URL Allowlist/Denylist

In strict mode, the Agent's ability to interact with external websites is governed by the browser's Allowlist and Denylist. This applies to:

*   **External Markdown Images**: The Agent will only render images from URLs that are allowed.
*   **Read URL Tool**: The Read URL tool will only auto-execute for allowed URLs.

#### Terminal, Browser, and Artifact Review Policies

Strict mode enforces the following behavior for terminal, browser, and artifact interactions:

*   **Terminal Auto Execution**: Set to "Request Review". The Agent will always prompt for permission before executing any terminal command. The terminal allowlist is ignored when strict mode is enabled.
*   **Browser Javascript Execution**: Set to "Request Review". The Agent will always prompt for permission before executing Javascript in the browser.
*   **Artifact Review**: Set to "Request Review". The Agent will always prompt for confirmation before acting on plans laid out in artifacts.

#### File System Access

Strict mode restricts the Agent's access to the file system to ensure it only interacts with authorized files:

*   **Respect .gitignore**: The Agent will respect `.gitignore` rules, preventing it from accessing ignored files.
*   **Workspace Isolation**: Access to files outside the workspace is disabled. The Agent can only view and edit files within the designated workspace.

---

<!-- Document: models.md -->
## Models

### Reasoning Model

For the core reasoning model, Antigravity offers leading frontier models from the Google Vertex Model Garden:

*   Gemini 3.1 Pro (high)
*   Gemini 3.1 Pro (low)
*   Gemini 3 Flash
*   Claude Sonnet 4.6 (thinking)
*   Claude Opus 4.6 (thinking)
*   GPT-OSS-120b

Users can select which reasoning model they want to use within the model selector drop down under the conversation prompt box:

![Model Selector Drop Down](https://antigravity.google/assets/image/docs/model-selector.png)

The choice of reasoning model is sticky between user messages within a conversation, so if you change the reasoning model while the Agent is running, it will continue to use the previously selected reasoning model until it has completed its steps for that user turn (or until the user cancels the current execution).

Learn more about reasoning model rate limits in [our plans page](https://antigravity.google/docs/plans).

### Additional Models

Antigravity uses a number of other models for various parts of the stack that are not customizable:

*   **Nano Banana Pro 2**: Used by the generative image tool when the Agent wants to produce a UI mockup, needs images to populate a web page or application, generate system or architecture diagrams, or other generative image tasks.
*   **Gemini 2.5 Pro UI Checkpoint**: Used by the [browser subagent](https://antigravity.google/docs/browser-subagent) to actuate the browser, such as clicking, scrolling, or filling in input.
*   **Gemini 2.5 Flash**: Used in the background for checkpointing and context summarization.
*   **Gemini 2.5 Flash Lite**: Used by the codebase semantic search tool.

---

<!-- Document: agent-modes-settings.md -->
## Agent Modes / Settings

### Conversation-Level

When starting a new Agent conversation, users can choose between multiple modes:

*   Planning: Agent can plan before executing tasks. Use for deep research, complex tasks, or collaborative work. In this mode, the Agent organizes its work in [task groups](https://antigravity.google/docs/task-groups), produces [Artifacts](https://antigravity.google/docs/artifacts), and takes other steps to thoroughly research, think through, and plan its work for optimal quality.
*   Fast: Agent will execute tasks directly. Use for simple tasks that can be completed faster, such as renaming variables, kicking off a few bash commands, or other smaller, localized tasks. This is helpful for when speed is an important factor, and the task is simple enough that there is low worry of worse quality.

### Overall Settings

Settings across every Agent conversation can be found in the “Agent” tab of the Settings pane. Some of the major ones include:

#### Artifact Review Policy

These are the possible options for Artifact Review Policy:

*   Always Proceed: Agent never asks for review
*   Request Review: Agent always asks for review

When Agent decides to request review from the user for implementation plans, this policy determines what the agent does. When set to “Request Review”, the agent will always terminate after notifying, allowing the user to spend time reviewing the plan and adding comments to augment proposed changes.

![Settings Review Policy Proceed](https://antigravity.google/assets/image/docs/agent/settings-review-policy-proceed.png)

If you do not need to manually review Agent’s plan before making changes, set this to “Always Proceed”, in which case every time the agent decides to request review from the user, it will then immediately continue with executing the plan.

![Settings Review Policy Manual](https://antigravity.google/assets/image/docs/agent/settings-review-policy-manual.png)

#### Terminal Command Auto Execution

For the terminal command generation tool:

*   Request Review: Never auto-execute terminal commands (except those in a configurable Allow list)
*   Always Proceed: Agent never asks for review (except those in a configurable Deny list)

#### Agent Non-Workspace File Access

Allow Agent to view and edit files outside of the current workspace. By default, the Agent only has access to the files in the workspace and in the application’s root folder `~/.antigravity/`, which contains [Artifacts](https://antigravity.google/docs/artifacts), [Knowledge Items](https://antigravity.google/docs/knowledge), and other Antigravity-specific data.

Use with caution, as this could expose local secret or sensitive data to the Agent.

---

<!-- Document: agent-permissions.md -->
*   Agent

## Agent Permissions

Antigravity uses a unified permission system to control what actions the Agent can perform on your behalf. Every action is represented as a **permission resource** that can be placed into one of three lists:

*   **Allow**: The action is auto-approved without prompting.
*   **Deny**: The action is blocked immediately.
*   **Ask**: The Agent pauses and asks for your approval before proceeding.

### Resource String Format

Every entry in the Allow, Deny, or Ask list follows the format:

            `action(target)`
        

Where `action` is one of the supported action types and `target` is a pattern describing what the permission covers.

#### Supported Actions

Action

Target Format

Matching

`command`

`command(prefix)` or `command(*)`

Matches commands by prefix. `command(git)` matches `git add`, `git commit`, etc.

`read_file`

`read_file(/path)`

Matches the file or everything under the directory. Paths must be literal and absolute. Globs (`*.go`), regex, and `~` are not supported.

`write_file`

`write_file(/path)`

Same as `read_file`. Also implicitly covers `read_file` for the same path.

`read_url`

`read_url(domain)` or `read_url(*)`

Matches the domain and all subdomains. Does not match URL paths.

`mcp`

`mcp(server/tool)`, `mcp(server/*)`, or `mcp(*)`

Matches by exact server name. `server/*` covers all tools on that server.

#### Examples

**Allow list** — these actions run without prompting:

```
command(git)                       # Any git command
command(npm run build)             # Build commands
read_file(/home/user/project)      # Read anything in the project
write_file(/home/user/project/src) # Edit files under src/
read_url(example.com)              # Fetch any public page
mcp(workspace/*)                   # All workspace tools
```
        

**Deny list** — these actions are always blocked:

```
command(rm)                        # No rm commands
command(sudo)                      # No sudo
write_file(/home/user/.ssh)        # No writes to .ssh
```
        

**Ask list** — the Agent pauses and asks before proceeding:

```
command(*)                         # Prompt for every command
mcp(*)                             # Prompt for every MCP tool call
```

---

<!-- Document: rules-workflows.md -->
## Rules

Rules are manually defined constraints for the Agent to follow, at both the local and global levels. Rules allow users to guide the agent to follow behaviors particular to their own use cases and style.

To get started with Rules:

1.  Open the Customizations panel via the "..." dropdown at the top of the editor's agent panel.
2.  Navigate to the Rules panel.
3.  Click **\+ Global** to create new Global Rules, or **\+ Workspace** to create new Workspace-specific rules.

A Rule itself is simply a Markdown file, where you can input the constraints to guide the Agent to your tasks, stack, and style.

Rules files are limited to 12,000 characters each.

### Global Rules

Global rules live in ~/.gemini/GEMINI.md and are applied across all workspaces.

### Workspace Rules

Workspace rules live in the .agents/rules folder of your workspace or git root.

At the rule level you can define how a rule should be activated:

*   Manual: The rule is manually activated via at mention in Agent’s input box.
*   Always On: The rule is always applied.
*   Model Decision: Based on a natural language description of the rule, the model decides whether to apply the rule.
*   Glob: Based on the glob pattern you define (e.g., _.js, src/\*\*/_.ts), the rule will be applied to all files that match the pattern.

Note: Antigravity now defaults to .agents/rules, but still maintains backward support for .agent/rules.

### @ Mentions

You can reference other files using @filename in a Rules file. If filename is a relative path, it will be interpreted relative to the location of the Rules file. If filename is an absolute path, it will be resolved as a true absolute path, otherwise it will be resolved relative to the repository. For example, @/path/to/file.md will first attempt to be resolved to /path/to/file.md, and if that file does not exist, it will be resolved to workspace/path/to/file.md.

## Workflows

Workflows enable you to define a series of steps to guide the Agent through a repetitive set of tasks, such as deploying a service or responding to PR comments. These Workflows are saved as markdown files, allowing you to have an easy repeatable way to run key processes. Once saved, Workflows can be invoked in Agent via a slash command with the format /workflow-name.

While Rules provide models with guidance by providing persistent, reusable context at the prompt level, Workflows provide a structured sequence of steps or prompts at the trajectory level, guiding the model through a series of interconnected tasks or actions.

To create a workflow:

1.  Open the Customizations panel via the "..." dropdown at the top of the editor's agent panel.
2.  Navigate to the Workflows panel.
3.  Click the **\+ Global** button to create a new global workflow that can be accessed across all your workspaces, or click the **\+ Workspace** button to create a workflow specific to your current workspace.

To execute a workflow, simply invoke it in Agent using the /workflow-name command. You can call other Workflows from within a workflow! For example, /workflow-1 can include instructions like “Call /workflow-2” and “Call /workflow-3”. Upon invocation, Agent sequentially processes each step defined in the workflow, performing actions or generating responses as specified.

Workflows are saved as markdown files and contain a title, a description and a series of steps with specific instructions for Agent to follow. Workflow files are limited to 12,000 characters each.

### Agent-Generated Workflows

You can also ask Agent to generate Workflows for you! This works particularly well after manually working with Agent through a series of steps since it can use the conversation history to create the Workflow.

---

<!-- Document: task-groups.md -->
## Task Groups

When Agent is in planning mode, large and complex tasks are handled with Task Groups, which break down these problems into smaller, more approachable units of work. Oftentimes, Agent will work on multiple parts of the greater task at the same time, and task sections are how these changes are presented to the user. Here is an example task group.

![Task Group](https://antigravity.google/assets/image/docs/agent/task-group.png)

The top component of the task group specifies the overarching goal of this task as well as summarizes the changes made within this unit of work. There is also a section of edited files for quick user audit of changes: click on the file pill and you will view the current state of the changed files.

![Task Group Clicked Pill](https://antigravity.google/assets/image/docs/agent/task-group-clicked-pill.png)

Within a task group, Agent identifies subtasks that help modularize necessary changes, and all work done by the Agent is viewable within these progress update sections. By default, the details in each subtask are not directly exposed to the user, but if you are interested, there is a toggle that will expand on the exact steps that Agent made.

![Task Group Expanded](https://antigravity.google/assets/image/docs/agent/task-group-expanded.png)

Sometimes, there are pending steps, such as browser setup or terminal commands requiring approval, that are created inside these progress updates. In this case, instead of expanding all of the updates, Agent provides a special section at the end of the task group where you can review these pending steps accordingly.

![Task Group Pending](https://antigravity.google/assets/image/docs/agent/task-group-pending.png)

---

<!-- Document: browser-subagent.md -->
## Browser Subagent

When the agent wants to interact with the browser, it invokes a browser subagent to handle the task at hand. The browser subagent runs a model specialized to operate on the pages that are open within the Antigravity-managed browser, which is different from the model you selected for the main agent.

This subagent has access to a variety of tools that are necessary to control your browser, including clicking, scrolling, typing, reading console logs, and more. It can also read your open pages through DOM capture, screenshots, or markdown parsing, as well as taking videos.

While the agent is controlling a page, it will show an overlay on the page with a blue border and a small panel with short descriptions of the actions being taken. When this is shown, you will not be allowed to interact with the page to ensure it doesn’t get confused by your actions.

The browser subagent can act on tabs that are not focused, so you are free to open other tabs and use them uninterrupted as it works.

---

<!-- Document: command.md -->
## Antigravity Editor: Command

The **Command** feature brings the power of natural language directly into your workflow, allowing you to request specific inline completions or terminal commands on the fly.

### How it Works

1.  **Trigger**: Press `Command + I` (Mac) or `Ctrl + I` (Windows/Linux).
2.  **Prompt**: A text input box will appear at your current cursor position.
3.  **Instruction**: Type your request in natural language (e.g., "Create a function to sort this list" or "Add error handling to this block").
4.  **Execution**: Antigravity generates the code or command directly inline for you to review and accept.

### Where to Use It

#### In the Editor

![Command in Editor](https://antigravity.google/assets/image/docs/editor/command_editor.png)

Use Command to generate boilerplate code, refactor complex functions, or write documentation without breaking your coding flow.

*   _Example_: "Create a React component for a login form."

#### In the Terminal

![Command in Terminal](https://antigravity.google/assets/image/docs/editor/command_terminal.png)

Use Command within the integrated Antigravity terminal to generate complex shell commands without needing to memorize syntax.

*   _Example_: "Find all processes listening on port 3000 and kill them."

---

<!-- Document: skills.md -->
## Agent Skills

Skills are an [open standard](https://agentskills.io/home) for extending agent capabilities. A skill is a folder containing a `SKILL.md` file with instructions that the agent can follow when working on specific tasks.

### What are skills?

Skills are reusable packages of knowledge that extend what the agent can do. Each skill contains:

*   **Instructions** for how to approach a specific type of task
*   **Best practices** and conventions to follow
*   **Optional scripts and resources** the agent can use

When you start a conversation, the agent sees a list of available skills with their names and descriptions. If a skill looks relevant to your task, the agent reads the full instructions and follows them.

### Where skills live

Antigravity supports two types of skills:

Location

Scope

`<workspace-root>/.agents/skills/<skill-folder>/`

Workspace-specific

`~/.gemini/antigravity/skills/<skill-folder>/`

Global (all workspaces)

**Workspace skills** are great for project-specific workflows, like your team's deployment process or testing conventions.

**Global skills** work across all your projects. Use these for personal utilities or general-purpose tools you want everywhere.

Note: Antigravity now defaults to .agents/skills, but still maintains backward support for .agent/skills.

### Creating a skill

To create a skill:

1.  Create a folder for your skill in one of the skill directories
2.  Add a `SKILL.md` file inside that folder

```
.agents/skills/
└─── my-skill/
    └─── SKILL.md
```

Every skill needs a `SKILL.md` file with YAML frontmatter at the top:

```yaml
---
name: my-skill
description: Helps with a specific task. Use when you need to do X or Y.
---

# My Skill

Detailed instructions for the agent go here.

## When to use this skill

- Use this when...
- This is helpful for...

## How to use it

Step-by-step guidance, conventions, and patterns the agent should follow.
```

#### Frontmatter fields

Field

Required

Description

`name`

No

A unique identifier for the skill (lowercase, hyphens for spaces). Defaults to the folder name if not provided.

`description`

Yes

A clear description of what the skill does and when to use it. This is what the agent sees when deciding whether to apply the skill.

Tip: Write your description in third person and include keywords that help the agent recognize when the skill is relevant. For example: "Generates unit tests for Python code using pytest conventions."

### Skill folder structure

While `SKILL.md` is the only required file, you can include additional resources:

```
.agents/skills/my-skill/
├─── SKILL.md       # Main instructions (required)
├─── scripts/       # Helper scripts (optional)
├─── examples/      # Reference implementations (optional)
└─── resources/     # Templates and other assets (optional)
```

The agent can read these files when following your skill's instructions.

### How the agent uses skills

Skills follow a **progressive disclosure** pattern:

1.  **Discovery**: When a conversation starts, the agent sees a list of available skills with their names and descriptions
2.  **Activation**: If a skill looks relevant to your task, the agent reads the full `SKILL.md` content
3.  **Execution**: The agent follows the skill's instructions while working on your task

You don't need to explicitly tell the agent to use a skill—it decides based on context. However, you can mention a skill by name if you want to ensure it's used.

### Best practices

#### Keep skills focused

Each skill should do one thing well. Instead of a "do everything" skill, create separate skills for distinct tasks.

#### Write clear descriptions

The description is how the agent decides whether to use your skill. Make it specific about what the skill does and when it's useful.

#### Use scripts as black boxes

If your skill includes scripts, encourage the agent to run them with `--help` first rather than reading the entire source code. This keeps the agent's context focused on the task.

#### Include decision trees

For complex skills, add a section that helps the agent choose the right approach based on the situation.

### Example: A code review skill

Here's a simple skill that helps the agent review code:

```yaml
---
name: code-review
description: Reviews code changes for bugs, style issues, and best practices. Use when reviewing PRs or checking code quality.
---

# Code Review Skill

When reviewing code, follow these steps:

## Review checklist

1. **Correctness**: Does the code do what it's supposed to?
2. **Edge cases**: Are error conditions handled?
3. **Style**: Does it follow project conventions?
4. **Performance**: Are there obvious inefficiencies?

## How to provide feedback

- Be specific about what needs to change
- Explain why, not just what
- Suggest alternatives when possible
```

---

# Git Worktree Reference

### NAME

git-worktree - Manage multiple working trees

### SYNOPSIS

`git` `worktree` `add` \[`-f`\] \[`--detach`\] \[`--checkout`\] \[`--lock` \[`--reason` _<string>_\]\]
		 \[`--orphan`\] \[(`-b` | `-B`) _<new-branch>_\] _<path>_ \[_<commit-ish>_\]
`git` `worktree` `list` \[`-v` | `--porcelain` \[`-z`\]\]
`git` `worktree` `lock` \[`--reason` _<string>_\] _<worktree>_
`git` `worktree` `move` _<worktree>_ _<new-path>_
`git` `worktree` `prune` \[`-n`\] \[`-v`\] \[`--expire` _<expire>_\]
`git` `worktree` `remove` \[`-f`\] _<worktree>_
`git` `worktree` `repair` \[_<path>_…​\]
`git` `worktree` `unlock` _<worktree>_

### DESCRIPTION

Manage multiple working trees attached to the same repository.

A git repository can support multiple working trees, allowing you to check out more than one branch at a time. With `git` `worktree` `add` a new working tree is associated with the repository, along with additional metadata that differentiates that working tree from others in the same repository. The working tree, along with this metadata, is called a "worktree".

This new worktree is called a "linked worktree" as opposed to the "main worktree" prepared by [git-init\[1\]](https://antigravity.google/docs/git-init) or [git-clone\[1\]](https://antigravity.google/docs/git-clone). A repository has one main worktree (if it’s not a bare repository) and zero or more linked worktrees. When you are done with a linked worktree, remove it with `git` `worktree` `remove`.

In its simplest form, `git` `worktree` `add` _<path>_ automatically creates a new branch whose name is the final component of _<path>_, which is convenient if you plan to work on a new topic. For instance, `git` `worktree` `add` `../hotfix` creates new branch `hotfix` and checks it out at path `../hotfix`. To instead work on an existing branch in a new worktree, use `git` `worktree` `add` _<path>_ _<branch>_. On the other hand, if you just plan to make some experimental changes or do testing without disturbing existing development, it is often convenient to create a _throwaway_ worktree not associated with any branch. For instance, `git` `worktree` `add` `-d` _<path>_ creates a new worktree with a detached `HEAD` at the same commit as the current branch.

If a working tree is deleted without using `git` `worktree` `remove`, then its associated administrative files, which reside in the repository (see "DETAILS" below), will eventually be removed automatically (see `gc.worktreePruneExpire` in [git-config\[1\]](https://antigravity.google/docs/git-config)), or you can run `git` `worktree` `prune` in the main or any linked worktree to clean up any stale administrative files.

If the working tree for a linked worktree is stored on a portable device or network share which is not always mounted, you can prevent its administrative files from being pruned by issuing the `git` `worktree` `lock` command, optionally specifying `--reason` to explain why the worktree is locked.

### COMMANDS

`add` _<path>_ \[_<commit-ish>_\]

Create a worktree at _<path>_ and checkout _<commit-ish>_ into it. The new worktree is linked to the current repository, sharing everything except per-worktree files such as `HEAD`, `index`, etc. As a convenience, _<commit-ish>_ may be a bare "`-`", which is synonymous with `@{-1}`.

If _<commit-ish>_ is a branch name (call it _<branch>_) and is not found, and neither `-b` nor `-B` nor `--detach` are used, but there does exist a tracking branch in exactly one remote (call it _<remote>_) with a matching name, treat as equivalent to:

$ git worktree add --track -b <branch> <path> <remote>/<branch>

If the branch exists in multiple remotes and one of them is named by the `checkout.defaultRemote` configuration variable, we’ll use that one for the purposes of disambiguation, even if the _<branch>_ isn’t unique across all remotes. Set it to e.g. `checkout.defaultRemote=origin` to always checkout remote branches from there if _<branch>_ is ambiguous but exists on the `origin` remote. See also `checkout.defaultRemote` in [git-config\[1\]](https://antigravity.google/docs/git-config).

If _<commit-ish>_ is omitted and neither `-b` nor `-B` nor `--detach` used, then, as a convenience, the new worktree is associated with a branch (call it _<branch>_) named after `$`(`basename` _<path>_). If _<branch>_ doesn’t exist, a new branch based on `HEAD` is automatically created as if `-b` _<branch>_ was given. If _<branch>_ does exist, it will be checked out in the new worktree, if it’s not checked out anywhere else, otherwise the command will refuse to create the worktree (unless `--force` is used).

If _<commit-ish>_ is omitted, neither `--detach`, or `--orphan` is used, and there are no valid local branches (or remote branches if `--guess-remote` is specified) then, as a convenience, the new worktree is associated with a new unborn branch named _<branch>_ (after `$`(`basename` _<path>_) if neither `-b` or `-B` is used) as if `--orphan` was passed to the command. In the event the repository has a remote and `--guess-remote` is used, but no remote or local branches exist, then the command fails with a warning reminding the user to fetch from their remote first (or override by using `-f`/`--force`).

`list`

List details of each worktree. The main worktree is listed first, followed by each of the linked worktrees. The output details include whether the worktree is bare, the revision currently checked out, the branch currently checked out (or "detached HEAD" if none), "locked" if the worktree is locked, "prunable" if the worktree can be pruned by the `prune` command.

`lock`

If a worktree is on a portable device or network share which is not always mounted, lock it to prevent its administrative files from being pruned automatically. This also prevents it from being moved or deleted. Optionally, specify a reason for the lock with `--reason`.

`move`

Move a worktree to a new location. Note that the main worktree or linked worktrees containing submodules cannot be moved with this command. (The `git` `worktree` `repair` command, however, can reestablish the connection with linked worktrees if you move the main worktree manually.)

`prune`

Prune worktree information in `$GIT_DIR/worktrees`.

`remove`

Remove a worktree. Only clean worktrees (no untracked files and no modification in tracked files) can be removed. Unclean worktrees or ones with submodules can be removed with `--force`. The main worktree cannot be removed.

`repair` \[_<path>_...\]

Repair worktree administrative files, if possible, if they have become corrupted or outdated due to external factors.

For instance, if the main worktree (or bare repository) is moved, linked worktrees will be unable to locate it. Running `repair` in the main worktree will reestablish the connection from linked worktrees back to the main worktree.

Similarly, if the working tree for a linked worktree is moved without using `git` `worktree` `move`, the main worktree (or bare repository) will be unable to locate it. Running `repair` within the recently-moved worktree will reestablish the connection. If multiple linked worktrees are moved, running `repair` from any worktree with each tree’s new _<path>_ as an argument, will reestablish the connection to all the specified paths.

If both the main worktree and linked worktrees have been moved or copied manually, then running `repair` in the main worktree and specifying the new _<path>_ of each linked worktree will reestablish all connections in both directions.

`unlock`

Unlock a worktree, allowing it to be pruned, moved or deleted.

### OPTIONS

`-f`

`--force`

By default, `add` refuses to create a new worktree when _<commit-ish>_ is a branch name and is already checked out by another worktree, or if _<path>_ is already assigned to some worktree but is missing (for instance, if _<path>_ was deleted manually). This option overrides these safeguards. To add a missing but locked worktree path, specify `--force` twice.

`move` refuses to move a locked worktree unless `--force` is specified twice. If the destination is already assigned to some other worktree but is missing (for instance, if _<new-path>_ was deleted manually), then `--force` allows the move to proceed; use `--force` twice if the destination is locked.

`remove` refuses to remove an unclean worktree unless `--force` is used. To remove a locked worktree, specify `--force` twice.

`-b` _<new-branch>_

`-B` _<new-branch>_

With `add`, create a new branch named _<new-branch>_ starting at _<commit-ish>_, and check out _<new-branch>_ into the new worktree. If _<commit-ish>_ is omitted, it defaults to `HEAD`. By default, `-b` refuses to create a new branch if it already exists. `-B` overrides this safeguard, resetting _<new-branch>_ to _<commit-ish>_.

`-d`

`--detach`

With `add`, detach `HEAD` in the new worktree. See "DETACHED HEAD" in [git-checkout\[1\]](https://antigravity.google/docs/git-checkout).

`--checkout`

`--no-checkout`

By default, `add` checks out _<commit-ish>_, however, `--no-checkout` can be used to suppress checkout in order to make customizations, such as configuring sparse-checkout. See "Sparse checkout" in [git-read-tree\[1\]](https://antigravity.google/docs/git-read-tree).

`--guess-remote`

`--no-guess-remote`

With `worktree` `add` _<path>_, without _<commit-ish>_, instead of creating a new branch from `HEAD`, if there exists a tracking branch in exactly one remote matching the basename of _<path>_, base the new branch on the remote-tracking branch, and mark the remote-tracking branch as "upstream" from the new branch.

This can also be set up as the default behaviour by using the `worktree.guessRemote` config option.

`--relative-paths`

`--no-relative-paths`

Link worktrees using relative paths or absolute paths (default). Overrides the `worktree.useRelativePaths` config option, see [git-config\[1\]](https://antigravity.google/docs/git-config).

With `repair`, the linking files will be updated if there’s an absolute/relative mismatch, even if the links are correct.

`--track`

`--no-track`

When creating a new branch, if _<commit-ish>_ is a branch, mark it as "upstream" from the new branch. This is the default if _<commit-ish>_ is a remote-tracking branch. See `--track` in [git-branch\[1\]](https://antigravity.google/docs/git-branch) for details.

`--lock`

Keep the worktree locked after creation. This is the equivalent of `git` `worktree` `lock` after `git` `worktree` `add`, but without a race condition.

`-n`

`--dry-run`

With `prune`, do not remove anything; just report what it would remove.

`--orphan`

With `add`, make the new worktree and index empty, associating the worktree with a new unborn branch named _<new-branch>_.

`--porcelain`

With `list`, output in an easy-to-parse format for scripts. This format will remain stable across Git versions and regardless of user configuration. It is recommended to combine this with `-z`. See below for details.

`-z`

Terminate each line with a _NUL_ rather than a newline when `--porcelain` is specified with `list`. This makes it possible to parse the output when a worktree path contains a newline character.

`-q`

`--quiet`

With `add`, suppress feedback messages.

`-v`

`--verbose`

With `prune`, report all removals.

With `list`, output additional information about worktrees (see below).

`--expire` _<time>_

With `prune`, only expire unused worktrees older than _<time>_.

With `list`, annotate missing worktrees as prunable if they are older than _<time>_.

`--reason` _<string>_

With `lock` or with `add` `--lock`, an explanation why the worktree is locked.

_<worktree>_

Worktrees can be identified by path, either relative or absolute.

If the last path components in the worktree’s path is unique among worktrees, it can be used to identify a worktree. For example if you only have two worktrees, at `/abc/def/ghi` and `/abc/def/ggg`, then `ghi` or `def/ghi` is enough to point to the former worktree.

### REFS

When using multiple worktrees, some refs are shared between all worktrees, but others are specific to an individual worktree. One example is `HEAD`, which is different for each worktree. This section is about the sharing rules and how to access refs of one worktree from another.

In general, all pseudo refs are per-worktree and all refs starting with `refs/` are shared. Pseudo refs are ones like `HEAD` which are directly under `$GIT_DIR` instead of inside `$GIT_DIR/refs`. There are exceptions, however: refs inside `refs/bisect`, `refs/worktree` and `refs/rewritten` are not shared.

Refs that are per-worktree can still be accessed from another worktree via two special paths, `main-worktree` and `worktrees`. The former gives access to per-worktree refs of the main worktree, while the latter to all linked worktrees.

For example, `main-worktree/HEAD` or `main-worktree/refs/bisect/good` resolve to the same value as the main worktree’s `HEAD` and `refs/bisect/good` respectively. Similarly, `worktrees/foo/HEAD` or `worktrees/bar/refs/bisect/bad` are the same as `$GIT_COMMON_DIR/worktrees/foo/HEAD` and `$GIT_COMMON_DIR/worktrees/bar/refs/bisect/bad`.

To access refs, it’s best not to look inside `$GIT_DIR` directly. Instead use commands such as [git-rev-parse\[1\]](https://antigravity.google/docs/git-rev-parse) or [git-update-ref\[1\]](https://antigravity.google/docs/git-update-ref) which will handle refs correctly.

### CONFIGURATION FILE

By default, the repository `config` file is shared across all worktrees. If the config variables `core.bare` or `core.worktree` are present in the common config file and `extensions.worktreeConfig` is disabled, then they will be applied to the main worktree only.

In order to have worktree-specific configuration, you can turn on the `worktreeConfig` extension, e.g.:

$ git config extensions.worktreeConfig true

In this mode, specific configuration stays in the path pointed by `git` `rev-parse` `--git-path` `config.worktree`. You can add or update configuration in this file with `git` `config` `--worktree`. Older Git versions will refuse to access repositories with this extension.

Note that in this file, the exception for `core.bare` and `core.worktree` is gone. If they exist in `$GIT_DIR/config`, you must move them to the `config.worktree` of the main worktree. You may also take this opportunity to review and move other configuration that you do not want to share to all worktrees:

*   `core.worktree` should never be shared.
    
*   `core.bare` should not be shared if the value is `core.bare=true`.
    
*   `core.sparseCheckout` should not be shared, unless you are sure you always use sparse checkout for all worktrees.
    

See the documentation of `extensions.worktreeConfig` in [git-config\[1\]](https://antigravity.google/docs/git-config) for more details.

### DETAILS

Each linked worktree has a private sub-directory in the repository’s `$GIT_DIR/worktrees` directory. The private sub-directory’s name is usually the base name of the linked worktree’s path, possibly appended with a number to make it unique. For example, when `$GIT_DIR=/path/main/.git` the command `git` `worktree` `add` `/path/other/test-next` `next` creates the linked worktree in `/path/other/test-next` and also creates a `$GIT_DIR/worktrees/test-next` directory (or `$GIT_DIR/worktrees/test-next1` if `test-next` is already taken).

Within a linked worktree, `$GIT_DIR` is set to point to this private directory (e.g. `/path/main/.git/worktrees/test-next` in the example) and `$GIT_COMMON_DIR` is set to point back to the main worktree’s `$GIT_DIR` (e.g. `/path/main/.git`). These settings are made in a `.git` file located at the top directory of the linked worktree.

Path resolution via `git` `rev-parse` `--git-path` uses either `$GIT_DIR` or `$GIT_COMMON_DIR` depending on the path. For example, in the linked worktree `git` `rev-parse` `--git-path` `HEAD` returns `/path/main/.git/worktrees/test-next/HEAD` (not `/path/other/test-next/.git/HEAD` or `/path/main/.git/HEAD`) while `git` `rev-parse` `--git-path` `refs/heads/master` uses `$GIT_COMMON_DIR` and returns `/path/main/.git/refs/heads/master`, since refs are shared across all worktrees, except `refs/bisect`, `refs/worktree` and `refs/rewritten`.

See [gitrepository-layout\[5\]](https://antigravity.google/docs/gitrepository-layout) for more information. The rule of thumb is do not make any assumption about whether a path belongs to `$GIT_DIR` or `$GIT_COMMON_DIR` when you need to directly access something inside `$GIT_DIR`. Use `git` `rev-parse` `--git-path` to get the final path.

If you manually move a linked worktree, you need to update the `gitdir` file in the entry’s directory. For example, if a linked worktree is moved to `/newpath/test-next` and its `.git` file points to `/path/main/.git/worktrees/test-next`, then update `/path/main/.git/worktrees/test-next/gitdir` to reference `/newpath/test-next` instead. Better yet, run `git` `worktree` `repair` to reestablish the connection automatically.

To prevent a `$GIT_DIR/worktrees` entry from being pruned (which can be useful in some situations, such as when the entry’s worktree is stored on a portable device), use the `git` `worktree` `lock` command, which adds a file named `locked` to the entry’s directory. The file contains the reason in plain text. For example, if a linked worktree’s `.git` file points to `/path/main/.git/worktrees/test-next` then a file named `/path/main/.git/worktrees/test-next/locked` will prevent the `test-next` entry from being pruned. See [gitrepository-layout\[5\]](https://antigravity.google/docs/gitrepository-layout) for details.

When `extensions.worktreeConfig` is enabled, the config file `.git/worktrees/`_<id>_`/config.worktree` is read after `.git/config` is.

### LIST OUTPUT FORMAT

The `worktree` `list` command has two output formats. The default format shows the details on a single line with columns. For example:

$ git worktree list
/path/to/bare-source            (bare)
/path/to/linked-worktree        abcd1234 \[master\]
/path/to/other-linked-worktree  1234abc  (detached HEAD)

The command also shows annotations for each worktree, according to its state. These annotations are:

*   `locked`, if the worktree is locked.
    
*   `prunable`, if the worktree can be pruned via `git` `worktree` `prune`.
    

$ git worktree list
/path/to/linked-worktree    abcd1234 \[master\]
/path/to/locked-worktree    acbd5678 (brancha) locked
/path/to/prunable-worktree  5678abc  (detached HEAD) prunable

For these annotations, a reason might also be available and this can be seen using the verbose mode. The annotation is then moved to the next line indented followed by the additional information.

$ git worktree list --verbose
/path/to/linked-worktree              abcd1234 \[master\]
/path/to/locked-worktree-no-reason    abcd5678 (detached HEAD) locked
/path/to/locked-worktree-with-reason  1234abcd (brancha)
	locked: worktree path is mounted on a portable device
/path/to/prunable-worktree            5678abc1 (detached HEAD)
	prunable: gitdir file points to non-existent location

Note that the annotation is moved to the next line if the additional information is available, otherwise it stays on the same line as the worktree itself.

#### [](#_porcelain_format)Porcelain Format

The porcelain format has a line per attribute. If `-z` is given then the lines are terminated with NUL rather than a newline. Attributes are listed with a label and value separated by a single space. Boolean attributes (like `bare` and `detached`) are listed as a label only, and are present only if the value is true. Some attributes (like `locked`) can be listed as a label only or with a value depending upon whether a reason is available. The first attribute of a worktree is always `worktree`, an empty line indicates the end of the record. For example:

$ git worktree list --porcelain
worktree /path/to/bare-source
bare

worktree /path/to/linked-worktree
HEAD abcd1234abcd1234abcd1234abcd1234abcd1234
branch refs/heads/master

worktree /path/to/other-linked-worktree
HEAD 1234abc1234abc1234abc1234abc1234abc1234a
detached

worktree /path/to/linked-worktree-locked-no-reason
HEAD 5678abc5678abc5678abc5678abc5678abc5678c
branch refs/heads/locked-no-reason
locked

worktree /path/to/linked-worktree-locked-with-reason
HEAD 3456def3456def3456def3456def3456def3456b
branch refs/heads/locked-with-reason
locked reason why is locked

worktree /path/to/linked-worktree-prunable
HEAD 1233def1234def1234def1234def1234def1234b
detached
prunable gitdir file points to non-existent location

Unless `-z` is used any "unusual" characters in the lock reason such as newlines are escaped and the entire reason is quoted as explained for the configuration variable `core.quotePath` (see [git-config\[1\]](https://antigravity.google/docs/git-config)). For Example:

$ git worktree list --porcelain
...
locked "reason\\nwhy is locked"
...

### EXAMPLES

You are in the middle of a refactoring session and your boss comes in and demands that you fix something immediately. You might typically use [git-stash\[1\]](https://antigravity.google/docs/git-stash) to store your changes away temporarily, however, your working tree is in such a state of disarray (with new, moved, and removed files, and other bits and pieces strewn around) that you don’t want to risk disturbing any of it. Instead, you create a temporary linked worktree to make the emergency fix, remove it when done, and then resume your earlier refactoring session.

$ git worktree add -b emergency-fix ../temp master
$ pushd ../temp
## ... hack hack hack ...
$ git commit -a -m 'emergency fix for boss'
$ popd
$ git worktree remove ../temp

### CONFIGURATION

Everything below this line in this section is selectively included from the [git-config\[1\]](https://antigravity.google/docs/git-config) documentation. The content is the same as what’s found there:

`worktree.guessRemote`

If no branch is specified and neither `-b` nor `-B` nor `--detach` is used, then `git` `worktree` `add` defaults to creating a new branch from HEAD. If `worktree.guessRemote` is set to true, `worktree` `add` tries to find a remote-tracking branch whose name uniquely matches the new branch name. If such a branch exists, it is checked out and set as "upstream" for the new branch. If no such match can be found, it falls back to creating a new branch from the current `HEAD`.

`worktree.useRelativePaths`

Link worktrees using relative paths (when "`true`") or absolute paths (when "`false`"). This is particularly useful for setups where the repository and worktrees may be moved between different locations or environments. Defaults to "`false`".

Note that setting `worktree.useRelativePaths` to "`true`" implies enabling the `extensions.relativeWorktrees` config (see [git-config\[1\]](https://antigravity.google/docs/git-config)), thus making it incompatible with older versions of Git.

### BUGS

Multiple checkout in general is still experimental, and the support for submodules is incomplete. It is NOT recommended to make multiple checkouts of a superproject.

### GIT

Part of the [git\[1\]](https://antigravity.google/docs/git) suite

#### worktree

