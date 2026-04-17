Write-Host "=> Installing diffless workflow globally..."

# 1. Compile the CLI binary
Write-Host "=> Compiling diffless CLI..."
$BinDir = "$env:USERPROFILE\.local\bin"
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

if (Test-Path ".\go.mod") {
    go build -o "$BinDir\diffless.exe" .\cmd\diffless\main.go
    Write-Host "=> Installed diffless binary to $BinDir\diffless.exe"
} else {
    Write-Host "=> Installing via 'go install'..."
    go install github.com/SilverShaman/diffless/cmd/diffless@latest
    $GoBin = go env GOPATH
    $BinDir = "$GoBin\bin"
    Write-Host "=> Installed diffless binary to $BinDir\diffless.exe"
}

# 2. Deploy the Antigravity Skill
Write-Host "=> Installing global Antigravity Skill..."
$SkillsDir = "$env:USERPROFILE\.gemini\antigravity\skills\diffless"
$ScriptsDir = "$SkillsDir\scripts"
New-Item -ItemType Directory -Force -Path $ScriptsDir | Out-Null

$SkillMdContent = @"
---
name: diffless
description: The native AI agent interface for the Diffless physical worktree CLI. Rapidly spin up sandboxes, run sync loops, and propose logic directly inside IDE execution bounds.
---

# Diffless Skill

This skill provides the native AI agent interface for the Diffless physical worktree CLI.

## When to use this skill
- When the user asks to start a new feature using diffless (diffless start).
- When the user asks to sync or propose changes using diffless (diffless sync, diffless propose).
- To work in an isolated physical Git worktree instead of the main trunk.

## How to use it
You can execute the diffless CLI directly if it is in your PATH.
"@

Set-Content -Path "$SkillsDir\SKILL.md" -Value $SkillMdContent

$ScriptContent = @"
@echo off
setlocal
set PATH=%PATH%;%USERPROFILE%\.local\bin;%USERPROFILE%\go\bin

where diffless >nul 2>&1
if %ERRORLEVEL% neq 0 (
    echo [Antigravity-Skill] Error: 'diffless' command not found. Ensure you ran the install script correctly and diffless is in your PATH. >&2
    exit /b 1
)

echo [Antigravity-Skill] 🚀 Routing IDE intent -^> diffless %*
diffless %*
"@

Set-Content -Path "$ScriptsDir\diffless.bat" -Value $ScriptContent

Write-Host "=> Installed Antigravity skill payload to $SkillsDir"

Write-Host ""
Write-Host "✅ Diffless globally installed and secured!"
Write-Host "Make sure $BinDir is in your PATH environment variable."
