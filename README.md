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

```bash
cd ~/dp-project-app
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

## Fremtidige Features

- Sprint management commands
- Task management commands
- Search og filtering
- Status reports
- MCP server integration
- Web interface