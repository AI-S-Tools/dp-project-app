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

### Option 1: Homebrew (macOS/Linux)

```bash
# Add the tap (once)
brew tap AI-S-Tools/dp-project-app https://github.com/AI-S-Tools/dp-project-app

# Install DPPM
brew install dppm

# Verify installation
dppm --version
dppm wiki "getting started"
```

### Option 2: One-Line Install (All Platforms)

```bash
curl -fsSL https://raw.githubusercontent.com/AI-S-Tools/dp-project-app/master/install.sh | bash
```

This script automatically:
- ✅ Detects your platform (Linux, macOS, Windows)
- ✅ Downloads the appropriate binary
- ✅ Installs to the correct location
- ✅ Makes it executable

### Option 3: Manual Download

Download the latest binary for your platform from [Releases](https://github.com/AI-S-Tools/dp-project-app/releases):

#### Linux / macOS
```bash
# Download binary (replace with your platform)
curl -L -o dppm-linux-amd64 https://github.com/AI-S-Tools/dp-project-app/releases/latest/download/dppm-linux-amd64

# Make executable and install
chmod +x dppm-linux-amd64
sudo mv dppm-linux-amd64 /usr/local/bin/dppm

# Verify installation
dppm --version
dppm wiki "getting started"
```

#### Windows
1. Download `dppm-windows-amd64.exe` from [Releases](https://github.com/AI-S-Tools/dp-project-app/releases)
2. Move to a directory in your PATH
3. Run: `dppm wiki "getting started"`

#### Available Platforms
- **Linux**: amd64, arm64
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

### Option 3: Build from Source

```bash
git clone https://github.com/AI-S-Tools/dp-project-app.git
cd dp-project-app
go build -o dppm
sudo cp dppm /usr/local/bin/
```

## Struktur

Projekter gemmes i `~/Dropbox/project-management/` med følgende struktur:

```
~/Dropbox/project-management/
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

### Projekter

```bash
# Opret nyt projekt
dppm project create my-project --name "Mit Projekt" --description "Beskrivelse" --owner "username"

# List alle projekter
dppm list projects
```

### Sprints (kommer snart)

```bash
dppm sprint create sprint-1 --project my-project
dppm sprint list --project my-project
```

### Tasks (kommer snart)

```bash
dppm task create task-1 --project my-project --sprint sprint-1 --title "Min task"
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
goal: "Sprint mål"
```

### Task
```yaml
id: "task-id"
title: "Task titel"
project_id: "parent-project-id"
sprint_id: "parent-sprint-id"
status: "todo"  # todo, in_progress, review, blocked, done
priority: "medium"  # low, medium, high, critical
assignee: "username"
description: |
  Detaljeret task beskrivelse
```

## Udvikling

Appen er bygget i Go med:
- `github.com/spf13/cobra` til CLI commands
- `gopkg.in/yaml.v3` til YAML parsing
- Standard library til filhåndtering

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