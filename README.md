# DP Project App

Dropbox Project Management CLI - Et kommandolinje værktøj til at administrere projekter, sprints og tasks i Dropbox.

## Installation

```bash
cd ~/dp-project-app
go build -o dp
sudo cp dp /usr/local/bin/
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
dp project create my-project --name "Mit Projekt" --description "Beskrivelse" --owner "username"

# List alle projekter
dp list projects
```

### Sprints (kommer snart)

```bash
dp sprint create sprint-1 --project my-project
dp sprint list --project my-project
```

### Tasks (kommer snart)

```bash
dp task create task-1 --project my-project --sprint sprint-1 --title "Min task"
dp task list --project my-project
dp task update task-1 --status in_progress
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