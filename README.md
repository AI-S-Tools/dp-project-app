# DPPM - Dropbox Project Manager

A comprehensive CLI tool for managing projects, sprints, and tasks using Dropbox as the storage backend. Perfect for AI-driven development workflows with verbose, AI-friendly output.

## Features

- ✅ YAML-based project, sprint, and task management
- ✅ Hierarchical project organization
- ✅ Cross-platform synchronization via Dropbox
- ✅ AI-friendly verbose output and comprehensive help
- ✅ Template-based project creation
- ✅ Extensive documentation and examples

## Installation

### Option 1: One-Shot Installation (Recommended)

#### Linux (x86_64)
```bash
curl -L https://github.com/AI-S-Tools/dppm/releases/download/v1.2.1/dppm-linux-amd64 -o dppm && chmod +x dppm && sudo mv dppm /usr/local/bin/dppm && dppm --version
```

#### Linux (ARM64)
```bash
curl -L https://github.com/AI-S-Tools/dppm/releases/download/v1.2.1/dppm-linux-arm64 -o dppm && chmod +x dppm && sudo mv dppm /usr/local/bin/dppm && dppm --version
```

#### macOS (Intel)
```bash
curl -L https://github.com/AI-S-Tools/dppm/releases/download/v1.2.1/dppm-macos-amd64 -o dppm && chmod +x dppm && sudo mv dppm /usr/local/bin/dppm && dppm --version
```

#### macOS (Apple Silicon)
```bash
curl -L https://github.com/AI-S-Tools/dppm/releases/download/v1.2.1/dppm-macos-arm64 -o dppm && chmod +x dppm && sudo mv dppm /usr/local/bin/dppm && dppm --version
```

#### Windows (PowerShell as Admin)
```powershell
Invoke-WebRequest -Uri "https://github.com/AI-S-Tools/dppm/releases/download/v1.2.1/dppm-windows-amd64.exe" -OutFile "dppm.exe" && Move-Item "dppm.exe" "C:\Windows\System32\dppm.exe" && dppm --version
```

### Option 2: Manual Download

Download the latest binary for your platform from [Releases](https://github.com/AI-S-Tools/dppm/releases):

#### Linux
```bash
# Download binary (choose your architecture: amd64 or arm64)
wget https://github.com/AI-S-Tools/dppm/releases/latest/download/dppm-linux-amd64
# OR for ARM:
# wget https://github.com/AI-S-Tools/dppm/releases/latest/download/dppm-linux-arm64

# Make executable and install
chmod +x dppm-linux-amd64
sudo mv dppm-linux-amd64 /usr/local/bin/dppm

# Verify installation
dppm --version
```

#### macOS
```bash
# For Intel Macs:
curl -L -o dppm https://github.com/AI-S-Tools/dppm/releases/latest/download/dppm-macos-amd64
# For Apple Silicon (M1/M2/M3):
# curl -L -o dppm https://github.com/AI-S-Tools/dppm/releases/latest/download/dppm-macos-arm64

# Make executable and install
chmod +x dppm
sudo mv dppm /usr/local/bin/dppm

# Verify installation
dppm --version
```

#### Windows
1. Download `dppm-windows-amd64.exe` from [Releases](https://github.com/AI-S-Tools/dppm/releases)
2. Move to a directory in your PATH
3. Run: `dppm --version`

#### Available Platforms
- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

### Option 3: Build from Source

```bash
git clone https://github.com/AI-S-Tools/dppm.git
cd dppm
go build -o dppm
sudo cp dppm /usr/local/bin/
```

## First Run Setup

DPPM requires Dropbox for project synchronization. On first run:

1. **Automatic Detection**: DPPM will try to find your Dropbox folder automatically
2. **Multiple Locations**: Supports `~/Dropbox`, `~/Dropbox (Personal)`, `~/Dropbox (Business)`
3. **Manual Path**: If not found automatically, you'll be prompted to enter the path
4. **Persistent Config**: Your Dropbox location is saved in `~/.dppm/dropbox.conf`

```bash
# First run - DPPM will guide you through setup
dppm --setup

# After setup, normal usage
dppm project create my-project --name "My Project"
```

## Struktur

Projects are stored in `[Your Dropbox]/project-management/` with the following structure:

```
[Your Dropbox]/project-management/
├── templates/              # YAML templates
│   ├── project.yaml
│   ├── sprint.yaml
│   └── task.yaml
└── projects/
    └── project-id/
        ├── project.yaml     # Project metadata
        └── sprints/
            └── sprint-id/
                ├── sprint.yaml
                └── tasks/
                    ├── task-1.yaml
                    ├── task-2.yaml
                    └── ...
```

## Kommandoer

### Projects

```bash
# Create new project
dppm project create my-project --name "My Project" --description "Description" --owner "username"

# List all projects
dppm list projects
```

### Sprints (coming soon)

```bash
dppm sprint create sprint-1 --project my-project
dppm sprint list --project my-project
```

### Tasks (coming soon)

```bash
dppm task create task-1 --project my-project --sprint sprint-1 --title "My task"
dppm task list --project my-project
dppm task update task-1 --status in_progress
```

## Comprehensive Help System

DPPM features extensive help documentation for every command:

```bash
# Main help
dppm --help

# Project command help
dppm project --help

# Specific command help
dppm project create --help

# List command help
dppm list --help
dppm list projects --help
```

## YAML Format

### Project
```yaml
id: "project-id"
name: "Project Name"
description: "Project description"
status: "active"  # active, completed, paused, cancelled
owner: "owner-name"
created: "2025-09-23"
updated: "2025-09-23"
```

### Sprint
```yaml
id: "sprint-id"
name: "Sprint Name"
project_id: "parent-project-id"
status: "planning"  # planning, active, completed, cancelled
start_date: "2025-09-23"
end_date: "2025-10-07"
goal: "Sprint goal"
```

### Task
```yaml
id: "task-id"
title: "Task title"
project_id: "parent-project-id"
sprint_id: "parent-sprint-id"
status: "todo"  # todo, in_progress, review, blocked, done
priority: "medium"  # low, medium, high, critical
assignee: "username"
description: |
  Detailed task description
```

## Udvikling

The app is built in Go with:
- `github.com/spf13/cobra` for CLI commands
- `gopkg.in/yaml.v3` for YAML parsing
- Standard library for file handling

## Advanced Task Management Features

### Task Components & Subtasks
- **Multi-part tasks**: Tasks can be broken into multiple components
- **Subtask types**: bug, enhancement, feature, documentation
- **Component status tracking**: Each part can be todo/in_progress/done
- **Bug issues**: Special subtasks for tracking bugs and fixes
- **Change requests**: Track modifications and improvements

### Dependency Management
- **Task dependencies**: Tasks can depend on other tasks (dependency_ids)
- **Automatic blocking**: Tasks cannot start until dependencies are completed
- **Dependency validation**: System prevents circular dependencies
- **Visual dependency chains**: Clear display of task relationships

### Enhanced Status & Reporting
- **Multi-level status**: Project → Sprint → Task → Components → Issues
- **Smart filtering**: Show active/pending/completed across all levels
- **Dependency visualization**: See what's blocking what
- **Progress tracking**: Automatic calculation of completion percentages
- **Status summaries**: Quick overview of project health

### Expanded YAML Schema

#### Task with Components
```yaml
id: "task-001"
title: "User Authentication System"
components:
  - id: "auth-backend"
    title: "Backend API"
    status: "done"
    type: "feature"
  - id: "auth-frontend"
    title: "Frontend UI"
    status: "in_progress"
    type: "feature"
  - id: "auth-tests"
    title: "Unit Tests"
    status: "todo"
    type: "testing"
issues:
  - id: "bug-001"
    title: "Login fails on mobile"
    type: "bug"
    status: "todo"
    parent_component: "auth-frontend"
dependency_ids: ["task-database", "task-security"]
blocked_by: ["task-database"]  # Auto-calculated
```

### Command Extensions
```bash
# Component management
dppm task component add task-001 --title "API Documentation" --type documentation
dppm task component update auth-backend --status done
dppm task component list task-001

# Issue tracking
dppm task issue add task-001 --title "Fix mobile login" --type bug --component auth-frontend
dppm task issue update bug-001 --status in_progress
dppm task issue list --type bug --status todo

# Dependency management
dppm task dependency add task-002 --depends-on task-001
dppm task dependency remove task-002 task-001
dppm task dependency check task-002  # Show what's blocking this task

# Status reporting
dppm status project web-app        # Overall project status
dppm status dependencies          # Show all dependency chains
dppm status blocked               # Show all blocked tasks
dppm status active               # Show all active work
```

## Fremtidige Features

- Advanced sprint planning with capacity
- Automated progress reporting
- Gantt chart export
- Time tracking integration
- Team collaboration features
- MCP server integration
- Web interface with dependency graphs