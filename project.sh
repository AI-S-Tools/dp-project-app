#!/bin/bash
# Project-scoped DPPM commands for: dp-project-app
# Auto-generated project wrapper to prevent cross-project task creation
# Usage: source project.sh or ./project.sh [command]

PROJECT_ID="dp-project-app"
DPPM_CMD="${DPPM_CMD:-dppm}"

# Project-scoped task management
task() {
    echo "🎯 [dp-project-app] Running: $DPPM_CMD task $* --project $PROJECT_ID"
    $DPPM_CMD task "$@" --project "$PROJECT_ID"
}

# Project-scoped phase management
phase() {
    echo "📋 [dp-project-app] Running: $DPPM_CMD phase $* --project $PROJECT_ID"
    $DPPM_CMD phase "$@" --project "$PROJECT_ID"
}

# Project-scoped status
project_status() {
    echo "📊 [dp-project-app] Project Status:"
    $DPPM_CMD status project "$PROJECT_ID" "$@"
}

# MCP server with project binding (future feature)
mcp_server() {
    echo "🤖 [dp-project-app] Starting MCP server bound to project: $PROJECT_ID"
    echo "Command would be: $DPPM_CMD mcp-server --project $PROJECT_ID $*"
    echo "(MCP server not yet implemented)"
}

# Show project context
context() {
    echo "🎯 Current DPPM Project: $PROJECT_ID"
    echo "📁 Working Directory: $(pwd)"
    echo "🔧 DPPM Command: $DPPM_CMD"
    echo ""
    project_status
}

# List project-specific items
list_items() {
    echo "📋 [dp-project-app] Listing project items:"
    if [ "$1" = "tasks" ]; then
        echo "Tasks for project $PROJECT_ID:"
        echo "(dppm task list --project command not yet available)"
    elif [ "$1" = "phases" ]; then
        echo "Phases for project $PROJECT_ID:"
        echo "(dppm phase list --project command not yet available)"
    else
        echo "Available list options: tasks, phases"
        echo "Usage: ./project.sh list_items tasks|phases"
    fi
}

# Create shortcuts for common operations
create_task() {
    if [ -z "$1" ]; then
        echo "Usage: create_task TASK_ID --title \"Task Title\" [options]"
        return 1
    fi
    echo "✨ [dp-project-app] Creating task: $1"
    task create "$@"
}

create_phase() {
    if [ -z "$1" ]; then
        echo "Usage: create_phase PHASE_ID --name \"Phase Name\" [options]"
        return 1
    fi
    echo "✨ [dp-project-app] Creating phase: $1"
    phase create "$@"
}

# Pass through other commands with project context awareness
dppm_global() {
    echo "🔄 [dp-project-app] Running global DPPM command: $DPPM_CMD $*"
    $DPPM_CMD "$@"
}

# Show available project commands
show_help() {
    echo "🎯 Project-scoped DPPM commands for: $PROJECT_ID"
    echo "=============================================="
    echo ""
    echo "📋 TASK MANAGEMENT:"
    echo "  task [cmd] [args]         - Task management (auto --project $PROJECT_ID)"
    echo "  create_task ID [opts]     - Quick task creation"
    echo ""
    echo "📁 PHASE MANAGEMENT:"
    echo "  phase [cmd] [args]        - Phase management (auto --project $PROJECT_ID)"
    echo "  create_phase ID [opts]    - Quick phase creation"
    echo ""
    echo "📊 PROJECT STATUS:"
    echo "  project_status            - Show project status"
    echo "  context                   - Show project context and status"
    echo "  list_items [tasks|phases] - List project items"
    echo ""
    echo "🤖 AI INTEGRATION:"
    echo "  mcp_server [opts]         - Start MCP server for this project"
    echo ""
    echo "🔄 GENERAL:"
    echo "  dppm_global [cmd] [args]  - Full DPPM command access"
    echo "  show_help                 - Show this help"
    echo ""
    echo "💡 EXAMPLES:"
    echo "  ./project.sh create_task auth-fix --title \"Fix authentication bug\" --priority high"
    echo "  ./project.sh create_phase testing --name \"Testing Phase\""
    echo "  ./project.sh project_status"
    echo "  ./project.sh context"
}

# Default action when called without arguments
if [ $# -eq 0 ]; then
    echo "🎯 DPPM Project: $PROJECT_ID"
    echo "📁 Directory: $(pwd)"
    echo ""
    project_status
    echo ""
    echo "💡 Run './project.sh show_help' for available commands"
else
    # Execute the requested command
    "$@"
fi