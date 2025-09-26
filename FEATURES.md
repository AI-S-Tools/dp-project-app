# DPPM Features & API Reference

Complete feature list and command reference for DPPM - Dropbox Project Manager.

## 🎯 Core Features

### ✅ Project Management
- **Hierarchical Organization**: Projects → Phases → Tasks → Components → Issues
- **YAML-based Storage**: Human-readable, version-controllable project data
- **Dropbox Synchronization**: Cross-platform team collaboration via Dropbox
- **Template System**: Consistent project structure with embedded templates
- **Status Tracking**: Multi-level status management and reporting

### ✅ AI-First Design
- **Verbose Output**: Comprehensive, AI-friendly command responses
- **Built-in Knowledge Base**: 30+ wiki topics covering all functionality
- **Smart Setup**: Automatic Dropbox detection with manual fallback
- **DSL Markers**: AI-to-AI collaboration system with structured handoffs
- **Comprehensive Help**: Every command includes detailed help and examples

### ✅ Advanced Task Management
- **Component Breakdown**: Tasks split into manageable components
- **Dependency Management**: Task dependencies with blocking detection
- **Issue Tracking**: Bug tracking within tasks and components
- **Time Tracking**: Built-in time logging and estimation
- **Status Reporting**: Multi-level progress tracking and visualization

### ✅ Cross-Platform Support
- **Linux**: x86_64, ARM64
- **macOS**: Intel, Apple Silicon (with smart Dropbox detection)
- **Windows**: x86_64, ARM64
- **Dropbox Detection**: Supports multiple installation types (Personal, Business)

## 📚 Wiki Knowledge Base (30+ Topics)

### Project Management Topics
```bash
dppm wiki "what is dppm"          # Introduction and overview
dppm wiki "getting started"       # Quick start guide for beginners
dppm wiki "create project"        # Complete project creation guide
dppm wiki "project types"         # Phase-based vs Task-based projects
dppm wiki "project structure"     # Directory organization
dppm wiki "project workflow"      # End-to-end project management
dppm wiki "build project"         # Complete project building workflow
dppm wiki "best practices"        # Recommended patterns and tips
```

### Phase Management Topics
```bash
dppm wiki "create phase"          # Phase creation and organization
dppm wiki "phase workflow"        # Managing development phases
dppm wiki "phase structure"       # Phase directory layout
dppm wiki "active phases"         # List and monitor active phases
```

### Task Management Topics
```bash
dppm wiki "create task"           # Task creation with all options
dppm wiki "task dependencies"     # Dependency management
dppm wiki "task workflow"         # Task lifecycle management
dppm wiki "task components"       # Breaking tasks into components
dppm wiki "active tasks"          # List all active/in-progress tasks
dppm wiki "dependency order"      # View tasks in dependency order
```

### Status & Reporting Topics
```bash
dppm wiki "status commands"       # All status reporting options
dppm wiki "dependency chains"     # Understanding task relationships
dppm wiki "blocked tasks"         # Identifying and resolving blocks
dppm wiki "find task"             # Search for specific tasks
dppm wiki "list active"           # Show all active work items
```

### Advanced Features Topics
```bash
dppm wiki "time tracking"         # Time logging and estimation
dppm wiki "issue tracking"        # Bug and issue management
dppm wiki "project templates"     # Using and creating templates
dppm wiki "collaboration"         # Working with teams
dppm wiki "ai collaboration"      # AI-to-AI task coordination
dppm wiki "automation"            # Scripting and CI/CD integration
dppm wiki "reporting"             # Generate progress reports
dppm wiki "troubleshooting"       # Common issues and solutions
```

### Complete Workflows Topics
```bash
dppm wiki "complete"              # Full project to completion example
dppm wiki "real world"            # Practical usage scenarios
dppm wiki "ai workflow"           # AI-optimized project workflow
```

## 🛠️ Command API Reference

### Main Commands

#### Setup & Configuration
```bash
dppm --setup                       # First-time setup guide (REQUIRED)
dppm --version                     # Show version information
dppm --help                        # Show main help
```

#### Quick Actions
```bash
dppm                               # Smart startup guide
dppm wiki                          # Browse knowledge base
dppm wiki list                     # List all available topics
dppm --wiki "search term"          # Search specific help topic
```

### Project Commands

#### Project Creation & Management
```bash
# Create new project
dppm project create PROJECT_ID \
  --name "Project Name" \
  --description "Project description" \
  --owner "owner-name" \
  --type "standard"

# Initialize complete project (recommended)
dppm init PROJECT_ID \
  --name "Project Name" \
  --doc "./requirements.md" \
  --owner "owner-name"

# List projects
dppm list projects
dppm list projects --status active
dppm list projects --owner "username"

# Project status
dppm status project PROJECT_ID
dppm status projects
```

#### Project Flags
- `--name`: Human-readable project name
- `--description`: Detailed project description
- `--owner`: Project owner/maintainer
- `--type`: Project type (standard, template, etc.)
- `--doc`: Documentation file to include

### Phase Commands

#### Phase Creation & Management
```bash
# Create new phase
dppm phase create PHASE_ID \
  --project PROJECT_ID \
  --name "Phase Name" \
  --description "Phase description" \
  --start-date "2024-01-01" \
  --end-date "2024-01-31" \
  --goal "Phase objective"

# List phases
dppm list phases --project PROJECT_ID
dppm list phases --status active
dppm phase list --project PROJECT_ID

# Phase status
dppm status phase PHASE_ID --project PROJECT_ID
dppm status phases --project PROJECT_ID
```

#### Phase Flags
- `--project`: Parent project ID (required)
- `--name`: Human-readable phase name
- `--description`: Detailed phase description
- `--start-date`: Phase start date (YYYY-MM-DD)
- `--end-date`: Phase end date (YYYY-MM-DD)
- `--goal`: Phase objective/goal
- `--status`: Phase status (planning, active, completed, cancelled)

### Task Commands

#### Task Creation & Management
```bash
# Create new task
dppm task create TASK_ID \
  --project PROJECT_ID \
  --phase PHASE_ID \
  --title "Task Title" \
  --description "Task description" \
  --assignee "username" \
  --priority "medium" \
  --status "todo" \
  --estimated-hours 8 \
  --dependency-ids "task1,task2"

# Task updates
dppm task update TASK_ID \
  --status "in_progress" \
  --assignee "new-assignee" \
  --priority "high"

# List tasks
dppm list tasks --project PROJECT_ID
dppm list tasks --phase PHASE_ID
dppm list tasks --status "in_progress"
dppm list tasks --assignee "username"

# Task status and details
dppm task show TASK_ID --project PROJECT_ID
dppm status task TASK_ID --project PROJECT_ID
```

#### Task Flags
- `--project`: Parent project ID (required)
- `--phase`: Parent phase ID (required)
- `--title`: Task title/summary
- `--description`: Detailed task description
- `--assignee`: Person assigned to task
- `--priority`: Task priority (low, medium, high, critical)
- `--status`: Task status (todo, in_progress, review, blocked, done)
- `--estimated-hours`: Time estimation in hours
- `--dependency-ids`: Comma-separated list of dependency task IDs

#### Task Component Management
```bash
# Add component to task
dppm task component add TASK_ID \
  --title "Component Title" \
  --type "feature" \
  --status "todo" \
  --description "Component description"

# Update component
dppm task component update COMPONENT_ID \
  --status "done" \
  --notes "Completion notes"

# List components
dppm task component list TASK_ID
dppm list components --task TASK_ID
```

#### Task Issue Management
```bash
# Add issue to task
dppm task issue add TASK_ID \
  --title "Issue Title" \
  --type "bug" \
  --status "todo" \
  --component COMPONENT_ID \
  --description "Issue description"

# Update issue
dppm task issue update ISSUE_ID \
  --status "in_progress" \
  --assignee "username"

# List issues
dppm task issue list TASK_ID
dppm list issues --type "bug"
dppm list issues --status "todo"
```

#### Task Dependency Management
```bash
# Add dependency
dppm task dependency add TASK_ID --depends-on DEPENDENCY_TASK_ID

# Remove dependency
dppm task dependency remove TASK_ID DEPENDENCY_TASK_ID

# Check dependencies
dppm task dependency check TASK_ID
dppm task dependency show TASK_ID
```

### List Commands

#### Universal Listing
```bash
# List all entities
dppm list projects
dppm list phases --project PROJECT_ID
dppm list tasks --project PROJECT_ID
dppm list tasks --phase PHASE_ID

# Filtered listing
dppm list projects --status active
dppm list tasks --status "in_progress"
dppm list tasks --assignee "username"
dppm list tasks --priority "high"

# Active items across all levels
dppm list active                   # All active items
dppm list active --project PROJECT_ID
```

### Status Commands

#### Multi-level Status Reporting
```bash
# Project-level status
dppm status project PROJECT_ID
dppm status projects

# Phase-level status
dppm status phase PHASE_ID --project PROJECT_ID
dppm status phases --project PROJECT_ID

# Task-level status
dppm status task TASK_ID --project PROJECT_ID
dppm status tasks --project PROJECT_ID

# Cross-cutting status views
dppm status active                 # All active work
dppm status blocked                # All blocked tasks
dppm status dependencies          # All dependency chains
dppm status overdue               # Overdue items
```

### AI Collaboration Commands

#### DSL Marker System
```bash
# Find collaboration markers in code/docs
dppm collab find PATH
dppm collab find . --type "TODO"
dppm collab find . --assignee "ai-agent"

# Collaboration wiki
dppm collab wiki "handoff patterns"
dppm collab wiki "task coordination"
dppm collab wiki "ai workflows"
```

## 🏗️ Project Structure

### Storage Layout
```
[Your Dropbox]/project-management/
├── dppm-global.db              # Global database
├── templates/                  # YAML templates
│   ├── project.yaml
│   ├── phase.yaml
│   └── task.yaml
└── projects/
    └── PROJECT_ID/
        ├── project.yaml        # Project metadata
        ├── phases/
        │   └── PHASE_ID/
        │       ├── phase.yaml  # Phase metadata
        │       └── tasks/
        │           ├── TASK_ID.yaml
        │           └── ...
        └── data/              # Project-specific database
            └── dppm.db
```

### Configuration
```
~/.dppm/
└── dropbox.conf               # Saved Dropbox path
```

## 📋 YAML Schemas

### Project Schema
```yaml
id: "project-id"
name: "Project Name"
description: "Detailed project description"
status: "active"  # active, completed, paused, cancelled
owner: "owner-name"
type: "standard"  # standard, template, research
created: "2024-01-01T10:00:00Z"
updated: "2024-01-01T10:00:00Z"
metadata:
  repository: "https://github.com/org/repo"
  documentation: "path/to/docs"
  tags: ["web", "api", "backend"]
```

### Phase Schema
```yaml
id: "phase-id"
name: "Phase Name"
project_id: "parent-project-id"
description: "Phase description"
status: "active"  # planning, active, completed, cancelled
start_date: "2024-01-01"
end_date: "2024-01-31"
goal: "Phase objective"
created: "2024-01-01T10:00:00Z"
updated: "2024-01-01T10:00:00Z"
```

### Task Schema
```yaml
id: "task-id"
title: "Task Title"
project_id: "parent-project-id"
phase_id: "parent-phase-id"
description: "Detailed task description"
status: "todo"  # todo, in_progress, review, blocked, done
priority: "medium"  # low, medium, high, critical
assignee: "username"
estimated_hours: 8
actual_hours: 0
dependency_ids: ["task1", "task2"]
blocked_by: []  # Auto-calculated
created: "2024-01-01T10:00:00Z"
updated: "2024-01-01T10:00:00Z"
due_date: "2024-01-15"

# Components (subtasks)
components:
  - id: "comp-1"
    title: "Component Title"
    type: "feature"  # feature, bug, enhancement, documentation, testing
    status: "todo"
    description: "Component description"
    estimated_hours: 4

# Issues (bugs, problems)
issues:
  - id: "issue-1"
    title: "Issue Title"
    type: "bug"  # bug, enhancement, question
    status: "todo"
    parent_component: "comp-1"
    description: "Issue description"
```

## 🔄 Status Values

### Project Status
- `active`: Currently being worked on
- `completed`: Project finished successfully
- `paused`: Temporarily stopped
- `cancelled`: Project cancelled

### Phase Status
- `planning`: Design and planning phase
- `active`: Currently executing
- `completed`: Phase finished
- `cancelled`: Phase cancelled

### Task Status
- `todo`: Not started
- `in_progress`: Currently being worked on
- `review`: Ready for review
- `blocked`: Cannot proceed (dependencies)
- `done`: Completed successfully

### Priority Levels
- `low`: Nice-to-have features
- `medium`: Standard priority
- `high`: Important, time-sensitive
- `critical`: Must be done immediately

## 🎯 Exit Codes

- `0`: Success
- `1`: General error
- `2`: Configuration error (Dropbox not found/configured)
- `3`: File/project not found
- `4`: Invalid command syntax
- `5`: Dependency conflict

## 🚀 Best Practices

### Project Workflow
1. **Initialize**: Use `dppm init` for complete setup
2. **Plan**: Create phases with clear goals and timelines
3. **Break Down**: Split phases into manageable tasks
4. **Dependencies**: Set up task dependencies early
5. **Track**: Use status commands to monitor progress
6. **Collaborate**: Leverage DSL markers for AI coordination

### Task Management
- Keep tasks small and focused (< 8 hours)
- Use components for complex tasks
- Set dependencies to prevent blocking
- Regular status updates
- Document decisions and changes

### AI Collaboration
- Use `dppm collab find` to discover handoff points
- Follow DSL marker conventions
- Document AI workflow patterns
- Regular `dppm wiki` consultation for best practices

---

**Total Commands**: 50+ commands and subcommands
**Wiki Topics**: 30+ comprehensive help topics
**Platforms**: Linux, macOS, Windows (all architectures)
**Storage**: Dropbox-synchronized YAML files with SQLite optimization
**AI-Ready**: Designed for automated workflows and AI collaboration