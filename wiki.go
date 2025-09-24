/* ::GEMINI:09: Koden kan optimeres betydeligt ved at flytte de store tekstblokke i `show...` funktionerne til eksterne filer for at reducere binærstørrelsen og forbedre vedligeholdelsen.:: */
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// /* Definerer 'wiki' kommandoen for at søge i vidensbasen. */
var wikiCmd = &cobra.Command{
	Use:   "wiki [search-terms",
	Short: "Search DPPM knowledge base and examples",
	Long: `DPPM Wiki - AI-Friendly Knowledge Base

Search through comprehensive DPPM documentation, examples, and best practices.
Designed specifically for AI systems to quickly find relevant information.

Usage:
  dppm wiki [search-terms]        # Search knowledge base
  dppm --wiki "search terms"      # Alternative search syntax
  dppm wiki list                  # Show all available topics
  dppm wiki complete              # Show complete workflow examples

Examples:
  dppm wiki "create project"      # How to create projects
  dppm wiki "dependencies"       # Dependency management
  dppm wiki "phase workflow"     # Phase management
  dppm wiki "task blocking"      # Task blocking and dependencies
  dppm wiki "status commands"    # Status and reporting
  dppm wiki "project structure"  # Directory organization`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			showWikiIndex()
			return
		}

		searchTerm := strings.ToLower(strings.Join(args, " "))

		if searchTerm == "list" {
			showWikiTopics()
			return
		}

		if searchTerm == "complete" {
			showCompleteWorkflow()
			return
		}

		searchWiki(searchTerm)
	},
}

// /* Viser hovedindekset for wikien. */
func showWikiIndex() {
	fmt.Println(`DPPM Wiki - AI-Friendly Knowledge Base
=====================================

🚀 Quick Start:
  dppm wiki "create project"      # Project creation guide
  dppm wiki "complete"            # Complete workflow example
  dppm wiki list                  # All available topics

🔍 Common Searches:
  • Project Management: "create project", "project structure"
  • Phase Management: "phase workflow", "create phase"
  • Task Management: "create task", "task dependencies"
  • Status & Reporting: "status commands", "dependency chains"
  • Troubleshooting: "blocked tasks", "dependency issues"

💡 AI Usage:
This wiki is optimized for AI systems. Each topic includes:
  - Complete command examples with all parameters
  - Real-world workflow scenarios
  - Best practices and common patterns
  - Troubleshooting guides with solutions

Type 'dppm wiki list' to see all available topics.`)
}

// /* Viser en liste over alle tilgængelige wiki-emner. */
func showWikiTopics() {
	fmt.Println(`Available Wiki Topics:
=====================

📁 Project Management:
  • "what is dppm" - Introduction and overview
  • "getting started" - Quick start guide for beginners
  • "create project" - Complete project creation guide
  • "project types" - Phase-based vs Task-based projects
  • "project structure" - Directory organization
  • "project workflow" - End-to-end project management
  • "build project" - Complete project building workflow
  • "best practices" - Recommended patterns and tips

📋 Phase Management:
  • "create phase" - Phase creation and organization
  • "phase workflow" - Managing development phases
  • "phase structure" - Phase directory layout
  • "active phases" - List and monitor active phases

✅ Task Management:
  • "create task" - Task creation with all options
  • "task dependencies" - Dependency management
  • "task workflow" - Task lifecycle management
  • "task components" - Breaking tasks into components
  • "active tasks" - List all active/in-progress tasks
  • "dependency order" - View tasks in dependency order

📊 Status & Reporting:
  • "status commands" - All status reporting options
  • "dependency chains" - Understanding task relationships
  • "blocked tasks" - Identifying and resolving blocks
  • "find task" - Search for specific tasks
  • "list active" - Show all active work items

🔧 Advanced Features:
  • "time tracking" - Time logging and estimation
  • "issue tracking" - Bug and issue management
  • "project templates" - Using and creating templates
  • "collaboration" - Working with teams
  • "ai collaboration" - AI-to-AI task coordination with DSL markers
  • "automation" - Scripting and CI/CD integration
  • "reporting" - Generate progress reports
  • "troubleshooting" - Common issues and solutions

🚀 Complete Workflows:
  • "complete" - Full project creation to completion example
  • "real world" - Practical usage scenarios
  • "ai workflow" - AI-optimized project workflow`)
}

// /* Søger i wikien efter et bestemt emne. */
func searchWiki(searchTerm string) {
	// Convert to lowercase for case-insensitive search
	searchLower := strings.ToLower(searchTerm)

	switch {
	case strings.Contains(searchLower, "what is dppm") || strings.Contains(searchLower, "introduction"):
		showIntroductionGuide()
	case strings.Contains(searchLower, "getting started") || strings.Contains(searchLower, "quick start"):
		showGettingStartedGuide()
	case strings.Contains(searchLower, "create project") || strings.Contains(searchLower, "project creation"):
		showProjectCreationGuide()
	case strings.Contains(searchLower, "project types") || strings.Contains(searchLower, "phase-based") || strings.Contains(searchLower, "task-based"):
		showProjectTypesGuide()
	case strings.Contains(searchLower, "create phase") || strings.Contains(searchLower, "phase creation") || strings.Contains(searchLower, "phase management") || strings.Contains(searchLower, "phase workflow"):
		showPhaseCreationGuide()
	case strings.Contains(searchLower, "create task") || strings.Contains(searchLower, "task creation"):
		showTaskCreationGuide()
	case strings.Contains(searchLower, "dependencies") || strings.Contains(searchLower, "dependency"):
		showDependencyGuide()
	case strings.Contains(searchLower, "status") || strings.Contains(searchLower, "reporting"):
		showStatusGuide()
	case strings.Contains(searchLower, "blocked") || strings.Contains(searchLower, "blocking"):
		showBlockedTasksGuide()
	case strings.Contains(searchLower, "structure") || strings.Contains(searchLower, "organization"):
		showStructureGuide()
	case strings.Contains(searchLower, "workflow") || strings.Contains(searchLower, "process"):
		showWorkflowGuide()
	case strings.Contains(searchLower, "active tasks") || strings.Contains(searchLower, "in progress"):
		showActiveTasksGuide()
	case strings.Contains(searchLower, "active phases") || strings.Contains(searchLower, "current phases"):
		showActivePhasesGuide()
	case strings.Contains(searchLower, "dependency order") || strings.Contains(searchLower, "task order"):
		showDependencyOrderGuide()
	case strings.Contains(searchLower, "build project") || strings.Contains(searchLower, "project building"):
		showProjectBuildingGuide()
	case strings.Contains(searchLower, "list active") || strings.Contains(searchLower, "active work"):
		showListActiveGuide()
	case strings.Contains(searchLower, "find task") || strings.Contains(searchLower, "search task"):
		showFindTaskGuide()
	case strings.Contains(searchLower, "ai workflow") || strings.Contains(searchLower, "ai project"):
		showAIWorkflowGuide()
	case strings.Contains(searchLower, "ai collaboration") || strings.Contains(searchLower, "dsl markers"):
		showAICollaborationGuide()
	case strings.Contains(searchLower, "best practices") || strings.Contains(searchLower, "recommended patterns"):
		showBestPracticesGuide()
	case strings.Contains(searchLower, "task components") || strings.Contains(searchLower, "breaking tasks"):
		showTaskComponentsGuide()
	case strings.Contains(searchLower, "time tracking") || strings.Contains(searchLower, "time logging"):
		showTimeTrackingGuide()
	case strings.Contains(searchLower, "issue tracking") || strings.Contains(searchLower, "bug management"):
		showIssueTrackingGuide()
	case strings.Contains(searchLower, "project templates") || strings.Contains(searchLower, "templates"):
		showProjectTemplatesGuide()
	case strings.Contains(searchLower, "collaboration") && !strings.Contains(searchLower, "ai"):
		showCollaborationGuide()
	case strings.Contains(searchLower, "automation") || strings.Contains(searchLower, "scripting"):
		showAutomationGuide()
	case strings.Contains(searchLower, "troubleshooting") || strings.Contains(searchLower, "common issues"):
		showTroubleshootingGuide()
	case strings.Contains(searchLower, "real world") || strings.Contains(searchLower, "practical usage"):
		showRealWorldGuide()
	default:
		fmt.Printf("No specific guide found for '%s'\n\n", searchTerm)
		fmt.Println("Try one of these common searches:")
		fmt.Println("  dppm wiki \"project types\"")
		fmt.Println("  dppm wiki \"active tasks\"")
		fmt.Println("  dppm wiki \"dependency order\"")
		fmt.Println("  dppm wiki list")
	}
}

// /* Viser en introduktion til DPPM. */
func showIntroductionGuide() {
	fmt.Println(`What is DPPM?
============

DPPM (Dropbox Project Manager) is a comprehensive CLI tool for managing
projects, phases, and tasks using Dropbox as the storage backend.

🎯 PURPOSE:
DPPM was designed specifically for AI-driven development workflows,
providing verbose, structured output that AI systems can easily parse
and understand.

✨ KEY FEATURES:
  • YAML-based data storage
  • Hierarchical project organization
  • Phase-based development management
  • Comprehensive dependency tracking
  • Built-in knowledge base (wiki)
  • Cross-platform via Dropbox sync
  • AI-optimized verbose output

📁 STORAGE:
All data is stored in: ~/Dropbox/project-management/
This enables automatic sync across all your devices.

🤖 AI-FIRST DESIGN:
  • Self-documenting commands
  • Built-in wiki for self-service
  • Structured YAML output
  • Verbose help everywhere
  • Complete examples included

🚀 USE CASES:
  • Software development projects
  • Task and bug tracking
  • Sprint/phase planning
  • Dependency management
  • Team collaboration
  • Personal task management

🔍 Getting Help:
  dppm wiki "getting started"    # Quick start guide
  dppm wiki list                 # All available topics
  dppm --help                    # Command reference`)
}

// /* Viser en 'getting started' guide. */
func showGettingStartedGuide() {
	fmt.Println(`Getting Started Guide
====================

🚀 QUICK START IN 5 MINUTES:

1️⃣ CHECK INSTALLATION:
   dppm
   # Should show the startup guide

2️⃣ CREATE YOUR FIRST PROJECT:
   dppm project create my-project --name "My First Project" --owner "your-name"

3️⃣ ADD A PHASE (OPTIONAL):
   dppm phase create phase-1 --project my-project --name "Initial Development"

4️⃣ CREATE YOUR FIRST TASK:
   dppm task create first-task \
     --project my-project \
     --phase phase-1 \
     --title "Set up repository" \
     --priority high

5️⃣ CHECK PROJECT STATUS:
   dppm status project my-project

📚 LEARNING PATH:

Beginner:
  1. dppm wiki "what is dppm"      # Understand the tool
  2. dppm wiki "project types"     # Choose project structure
  3. dppm wiki "create project"    # Create first project
  4. dppm wiki "create task"       # Add tasks

Intermediate:
  1. dppm wiki "dependencies"      # Task relationships
  2. dppm wiki "phase workflow"    # Phase management
  3. dppm wiki "active tasks"      # Track progress
  4. dppm wiki "blocked tasks"     # Resolve blocks

Advanced:
  1. dppm wiki "build project"     # Complete workflows
  2. dppm wiki "ai workflow"       # AI automation
  3. dppm wiki "automation"        # CI/CD integration

💡 TIPS FOR SUCCESS:
  • Start simple with a task-based project
  • Use phases for projects > 10 tasks
  • Set dependencies thoughtfully
  • Check status daily
  • Use wiki for any questions

🆘 GETTING HELP:
  dppm wiki "topic"               # Search for help
  dppm wiki list                  # See all topics
  dppm wiki complete              # Full example

🔍 Next Steps:
  dppm wiki "project types"       # Understand options
  dppm wiki "create project"      # Start building`)
}

// /* Viser en guide til oprettelse af projekter. */
func showProjectCreationGuide() {
	fmt.Println(`Project Creation Guide
=====================

🎯 Basic Project Creation:
  dppm project create my-project --name "My Project" --owner "username"

📋 Complete Project Creation:
  dppm project create web-app \
    --name "Web Application" \
    --owner "dev-team" \
    --description "Modern web application with React frontend"

🗂️ What Gets Created:
  ~/Dropbox/project-management/projects/web-app/
  ├── project.yaml          # Project metadata
  └── phases/              # Empty phases directory

📄 project.yaml Structure:
  id: "web-app"
  name: "Web Application"
  description: "Modern web application..."
  status: "active"
  owner: "dev-team"
  created: "2025-09-23"
  updated: "2025-09-23"
  current_phase: ""
  phases: []

✅ Next Steps:
  1. dppm phase create phase-1 --project web-app --name "Setup Phase"
  2. dppm task create setup-repo --project web-app --phase phase-1
  3. dppm status project web-app

🔍 Related Commands:
  • dppm list projects              # List all projects
  • dppm project show web-app      # View project details
  • dppm wiki "create phase"       # Next step guide`)
}

// /* Viser en guide til oprettelse af faser. */
func showPhaseCreationGuide() {
	fmt.Println(`Phase Creation Guide
===================

🎯 Basic Phase Creation:
  dppm phase create phase-1 --project my-project --name "Setup Phase"

📋 Complete Phase Creation:
  dppm phase create backend-api \
    --project web-app \
    --name "Backend API Development" \
    --goal "Build REST API with authentication" \
    --start-date "2025-09-23" \
    --end-date "2025-10-07"

🗂️ What Gets Created:
  ~/Dropbox/project-management/projects/web-app/phases/backend-api/
  ├── phase.yaml           # Phase metadata
  └── tasks/              # Tasks directory

📄 phase.yaml Structure:
  id: "backend-api"
  name: "Backend API Development"
  project_id: "web-app"
  status: "planning"
  start_date: "2025-09-23"
  end_date: "2025-10-07"
  goal: "Build REST API with authentication"
  capacity: 10
  tasks: []

📅 Phase Status Values:
  • planning   - Phase is being planned (default)
  • active     - Currently working on this phase
  • completed  - Phase finished successfully
  • cancelled  - Phase abandoned

✅ Next Steps:
  1. dppm task create auth-system --project web-app --phase backend-api
  2. dppm task create user-mgmt --project web-app --phase backend-api
  3. dppm status project web-app

🔍 Related Commands:
  • dppm list phases --project web-app    # List project phases
  • dppm phase show backend-api --project web-app
  • dppm wiki "create task"               # Add tasks to phase`)
}

// /* Viser en guide til oprettelse af opgaver. */
func showTaskCreationGuide() {
	fmt.Println(`Task Creation Guide
==================

🎯 Basic Task Creation:
  dppm task create auth-system --project web-app --title "User Authentication"

📋 Complete Task Creation:
  dppm task create auth-system \
    --project web-app \
    --phase backend-api \
    --title "User Authentication System" \
    --description "Implement JWT-based authentication with login/logout" \
    --priority high \
    --assignee "john-doe"

🗂️ What Gets Created:
  ~/Dropbox/project-management/projects/web-app/phases/backend-api/tasks/
  └── auth-system.yaml

📄 auth-system.yaml Structure:
  id: "auth-system"
  title: "User Authentication System"
  project_id: "web-app"
  phase_id: "backend-api"
  status: "todo"
  priority: "high"
  assignee: "john-doe"
  description: "Implement JWT-based authentication..."
  created: "2025-09-23"
  updated: "2025-09-23"

🎯 Priority Levels:
  • low      - Nice to have features
  • medium   - Standard features (default)
  • high     - Important features
  • critical - Must-have features

📊 Status Values:
  • todo        - Not started (default)
  • in_progress - Currently working
  • done        - Completed
  • blocked     - Waiting for dependencies

✅ Advanced Features:
  # Add dependencies:
  --dependency-ids "setup-repo,database-setup"

  # Add story points:
  --story-points 8

  # Set due date:
  --due-date "2025-10-01"

🔍 Related Commands:
  • dppm list tasks --project web-app     # List all tasks
  • dppm task show auth-system            # View task details
  • dppm wiki "dependencies"              # Dependency management`)
}

// /* Viser en guide til afhængighedsstyring. */
func showDependencyGuide() {
	fmt.Println(`Dependency Management Guide
==========================

🔗 Task Dependencies:
Dependencies ensure tasks are completed in the correct order.
A task with dependencies cannot start until all dependency tasks are "done".

📋 Creating Tasks with Dependencies:
  # Task that depends on others:
  dppm task create frontend-auth \
    --project web-app \
    --phase frontend \
    --title "Frontend Authentication" \
    --dependency-ids "auth-system,user-api"

📄 Dependency Structure in YAML:
  dependency_ids: ["auth-system", "user-api"]
  blocked_by: []      # Auto-calculated
  blocking: []        # Auto-calculated

🚫 Understanding Blocking:
  • blocked_by: Tasks that must complete before this task can start
  • blocking: Tasks waiting for this task to complete
  • dependency_ids: Explicit dependencies you set

📊 Checking Dependencies:
  dppm status dependencies --project web-app
  dppm status blocked --project web-app
  dppm status project web-app

💡 Example Dependency Chain:
  1. setup-repo (no dependencies) ✅
  2. database-setup (depends on: setup-repo) 🔄
  3. auth-system (depends on: database-setup) ⏳ BLOCKED
  4. frontend-auth (depends on: auth-system) ⏳ BLOCKED

🔧 Dependency Best Practices:
  • Keep dependency chains short (max 3-4 levels)
  • Use phases to group related work
  • Check for circular dependencies
  • Mark tasks "done" promptly to unblock others

⚠️ Troubleshooting:
  # Find what's blocking a task:
  dppm status blocked --project web-app

  # See full dependency chain:
  dppm status dependencies --project web-app

🔍 Related Commands:
  • dppm wiki "blocked tasks"    # Resolving blocked tasks
  • dppm wiki "status commands"  # All status options`)
}

// /* Viser en guide til status- og rapporteringskommandoer. */
func showStatusGuide() {
	fmt.Println(`Status & Reporting Guide
=======================

📊 Project Overview:
  dppm status project my-project

  Shows:
  • Total task count
  • Tasks by status (done, in_progress, todo, blocked)
  • List of ready-to-work tasks
  • List of blocked tasks with blocking reasons

🚫 Blocked Tasks Analysis:
  dppm status blocked --project my-project
  dppm status blocked    # All projects

  Shows:
  • Which tasks are blocked
  • What tasks are blocking them
  • Priority levels of blocked tasks

🔗 Dependency Chain Analysis:
  dppm status dependencies --project my-project
  dppm status dependencies    # All projects

  Shows:
  • Complete dependency relationships
  • Status of each dependency (✅ done, ❌ not done)
  • Full dependency chains

📋 Task Listing:
  dppm list projects                    # All projects
  dppm list phases --project my-project  # Project phases
  dppm list tasks --project my-project   # All project tasks
  dppm list tasks --phase phase-1        # Phase-specific tasks

📈 Example Status Output:
  Project Status: web-app
  =====================
  Total Tasks: 8
  ✅ Done: 2
  🔄 In Progress: 1
  📋 Ready to Start: 3
  🚫 Blocked: 2

  🚫 Blocked Tasks:
    • Frontend Auth (blocked by: Backend API)
    • User Dashboard (blocked by: Frontend Auth)

  📋 Ready to Work On:
    • Database Schema (high priority)
    • API Tests (medium priority)
    • Documentation (low priority)

💡 AI Usage Tips:
  • Use status commands to understand project health
  • Check blocked tasks daily to identify bottlenecks
  • Use dependency analysis to plan work order
  • Status output is structured for easy AI parsing

🔍 Related Commands:
  • dppm wiki "blocked tasks"      # Resolving blocks
  • dppm wiki "dependencies"       # Dependency management`)
}

// /* Viser en guide til løsning af blokerede opgaver. */
func showBlockedTasksGuide() {
	fmt.Println(`Blocked Tasks Resolution Guide
=============================

🚫 Understanding Blocked Tasks:
A task is "blocked" when it has dependencies that are not yet completed.
The task cannot start until ALL dependencies are marked "done".

🔍 Finding Blocked Tasks:
  dppm status blocked --project my-project
  dppm status project my-project    # Shows blocked count

📊 Example Blocked Task Output:
  🚫 Frontend Authentication
     Priority: high
     Blocked by: Backend API, User Database Schema

  This means Frontend Authentication cannot start until both
  "Backend API" AND "User Database Schema" are marked as "done".

✅ Resolving Blocked Tasks:
  1. Identify the blocking tasks
  2. Work on completing the blocking tasks first
  3. Mark blocking tasks as "done" when complete
  4. The blocked task will automatically become "ready to start"

🔧 Updating Task Status:
  dppm task update backend-api --status done
  dppm task update user-schema --status done
  # Now frontend-auth is automatically unblocked!

📈 Monitoring Block Resolution:
  Before:
  🚫 Blocked: 3 tasks
  📋 Ready: 2 tasks

  After completing blocking tasks:
  🚫 Blocked: 1 task
  📋 Ready: 4 tasks

💡 Prevention Strategies:
  • Plan dependencies carefully during task creation
  • Keep dependency chains short (2-3 levels max)
  • Use phases to group related work
  • Complete high-priority blocking tasks first
  • Check status daily to catch blocks early

⚠️ Common Issues:
  • Circular dependencies (Task A blocks B, B blocks A)
  • Long dependency chains (A→B→C→D→E)
  • Missing dependencies (forgot to mark prerequisite)
  • Wrong dependencies (dependency not actually needed)

🔍 Related Commands:
  • dppm wiki "dependencies"       # Dependency management
  • dppm wiki "task workflow"      # Task lifecycle
  • dppm wiki "status commands"    # Status monitoring`)
}

// /* Viser en guide til projektets mappestruktur. */
func showStructureGuide() {
	fmt.Println(`Project Structure Guide
======================

🗂️ DPPM Directory Organization:
  ~/Dropbox/project-management/
  ├── templates/               # Project templates
  │   ├── project.yaml        # Default project template
  │   └── phase.yaml          # Default phase template
  └── projects/               # All projects
      └── PROJECT_ID/         # Individual project
          ├── project.yaml    # Project metadata
          └── phases/         # All project phases
              └── PHASE_ID/   # Individual phase
                  ├── phase.yaml    # Phase metadata
                  └── tasks/        # Phase tasks
                      └── TASK_ID.yaml  # Individual task

📁 Example Real Structure:
  ~/Dropbox/project-management/projects/web-app/
  ├── project.yaml
  └── phases/
      ├── setup/
      │   ├── phase.yaml
      │   └── tasks/
      │       ├── repo-setup.yaml
      │       └── env-config.yaml
      ├── backend/
      │   ├── phase.yaml
      │   └── tasks/
      │       ├── auth-system.yaml
      │       ├── user-api.yaml
      │       └── database.yaml
      └── frontend/
          ├── phase.yaml
          └── tasks/
              ├── login-ui.yaml
              └── dashboard.yaml

🎯 Benefits of This Structure:
  • ✅ Clear separation of concerns
  • ✅ Easy navigation and organization
  • ✅ Scalable for large projects
  • ✅ AI-friendly hierarchical structure
  • ✅ Cross-platform via Dropbox sync
  • ✅ Version control friendly (YAML files)

📄 File Content Structure:

project.yaml:
  id: "web-app"
  name: "Web Application"
  status: "active"
  current_phase: "backend"
  phases: ["setup", "backend", "frontend"]

phase.yaml:
  id: "backend"
  name: "Backend Development"
  project_id: "web-app"
  status: "active"
  goal: "Build REST API"

task.yaml:
  id: "auth-system"
  title: "Authentication System"
  project_id: "web-app"
  phase_id: "backend"
  status: "in_progress"
  dependencies: ["database"]

🔍 Related Commands:
  • dppm list projects            # Browse all projects
  • dppm list phases --project X  # Browse project phases
  • dppm list tasks --phase Y     # Browse phase tasks`)
}

// /* Viser en komplet guide til projektets arbejdsgang. */
func showWorkflowGuide() {
	fmt.Println(`Complete Project Workflow Guide
==============================

🚀 End-to-End Project Creation:

1️⃣ CREATE PROJECT:
   dppm project create web-app \
     --name "Web Application" \
     --owner "dev-team" \
     --description "Modern React web app"

2️⃣ CREATE PHASES:
   dppm phase create setup --project web-app \
     --name "Project Setup" \
     --goal "Initialize project infrastructure"

   dppm phase create backend --project web-app \
     --name "Backend Development" \
     --goal "Build REST API with authentication"

   dppm phase create frontend --project web-app \
     --name "Frontend Development" \
     --goal "Build React user interface"

3️⃣ CREATE TASKS WITH DEPENDENCIES:
   # Phase 1: Setup (no dependencies)
   dppm task create repo-setup --project web-app --phase setup \
     --title "Repository Setup" --priority high

   dppm task create env-config --project web-app --phase setup \
     --title "Environment Configuration" \
     --dependency-ids "repo-setup"

   # Phase 2: Backend (depends on setup)
   dppm task create database --project web-app --phase backend \
     --title "Database Schema" \
     --dependency-ids "env-config"

   dppm task create auth-api --project web-app --phase backend \
     --title "Authentication API" \
     --dependency-ids "database"

   # Phase 3: Frontend (depends on backend)
   dppm task create login-ui --project web-app --phase frontend \
     --title "Login Interface" \
     --dependency-ids "auth-api"

4️⃣ MONITOR AND EXECUTE:
   dppm status project web-app    # Check project health
   dppm status blocked           # Find blocking issues

   # Work on ready tasks:
   dppm task update repo-setup --status in_progress
   dppm task update repo-setup --status done

   # Check what's now unblocked:
   dppm status project web-app

📊 Typical Workflow States:

Initial State:
   📋 Ready: repo-setup
   🚫 Blocked: env-config, database, auth-api, login-ui

After completing repo-setup:
   📋 Ready: env-config
   🚫 Blocked: database, auth-api, login-ui

After completing env-config:
   📋 Ready: database
   🚫 Blocked: auth-api, login-ui

And so on...

💡 Best Practices:
   • Plan all phases before creating tasks
   • Set up dependency chains thoughtfully
   • Use priority levels to guide work order
   • Check status daily to identify bottlenecks
   • Mark tasks "done" promptly to unblock others

🔍 Related Commands:
   • dppm wiki "create project"    # Detailed project creation
   • dppm wiki "dependencies"      # Dependency management
   • dppm wiki "status commands"   # Monitoring tools`)
}

// /* Viser et komplet eksempel på en arbejdsgang i DPPM. */
func showCompleteWorkflow() {
	fmt.Println(`Complete DPPM Workflow Example
=============================

🎯 SCENARIO: Building a Web Application
Let's walk through creating a complete project from scratch.

1️⃣ PROJECT CREATION:
   dppm project create web-app --name "Modern Web App" --owner "ai-team"

2️⃣ PHASE PLANNING:
   dppm phase create setup --project web-app --name "Project Setup"
   dppm phase create backend --project web-app --name "Backend API"
   dppm phase create frontend --project web-app --name "Frontend UI"
   dppm phase create deploy --project web-app --name "Deployment"

3️⃣ TASK CREATION WITH DEPENDENCIES:
   # Setup Phase
   dppm task create git-repo --project web-app --phase setup --title "Git Repository Setup"
   dppm task create docker-env --project web-app --phase setup --title "Docker Environment" --dependency-ids "git-repo"

   # Backend Phase
   dppm task create database --project web-app --phase backend --title "Database Schema" --dependency-ids "docker-env"
   dppm task create auth-api --project web-app --phase backend --title "Authentication API" --dependency-ids "database"
   dppm task create user-api --project web-app --phase backend --title "User Management API" --dependency-ids "auth-api"

   # Frontend Phase
   dppm task create react-setup --project web-app --phase frontend --title "React App Setup" --dependency-ids "auth-api"
   dppm task create login-ui --project web-app --phase frontend --title "Login Interface" --dependency-ids "react-setup"
   dppm task create dashboard --project web-app --phase frontend --title "User Dashboard" --dependency-ids "login-ui,user-api"

   # Deploy Phase
   dppm task create ci-cd --project web-app --phase deploy --title "CI/CD Pipeline" --dependency-ids "dashboard"

4️⃣ EXECUTION WORKFLOW:
   # Check initial status
   dppm status project web-app
   # Output: 1 ready task (git-repo), 7 blocked tasks

   # Start first task
   dppm task update git-repo --status in_progress
   # Work on it...
   dppm task update git-repo --status done

   # Check what's unblocked
   dppm status project web-app
   # Output: 1 ready task (docker-env), 6 blocked tasks

   # Continue the workflow...
   dppm task update docker-env --status in_progress
   dppm task update docker-env --status done

   # Now database becomes ready
   dppm status project web-app

5️⃣ MONITORING THROUGHOUT:
   # Check for any blocking issues
   dppm status blocked --project web-app

   # See full dependency chain
   dppm status dependencies --project web-app

   # List tasks by phase
   dppm list tasks --phase backend

📊 FINAL PROJECT STRUCTURE:
   web-app/
   ├── project.yaml
   └── phases/
       ├── setup/
       │   ├── phase.yaml
       │   └── tasks/
       │       ├── git-repo.yaml
       │       └── docker-env.yaml
       ├── backend/
       │   ├── phase.yaml
       │   └── tasks/
       │       ├── database.yaml
       │       ├── auth-api.yaml
       │       └── user-api.yaml
       ├── frontend/
       │   ├── phase.yaml
       │   └── tasks/
       │       ├── react-setup.yaml
       │       ├── login-ui.yaml
       │       └── dashboard.yaml
       └── deploy/
           ├── phase.yaml
           └── tasks/
               └── ci-cd.yaml

This example shows how DPPM manages complex projects with proper dependencies,
phase organization, and clear workflow progression!`)
}

// /* Viser en guide til forskellige projekttyper. */
func showProjectTypesGuide() {
	fmt.Println(`Project Types Guide
==================

DPPM supports two project organization strategies:

🗂️ PHASE-BASED PROJECTS (Recommended)
Best for: Larger projects with distinct development stages

Structure:
  project/
  └── phases/
      ├── phase-1/
      │   └── tasks/
      ├── phase-2/
      │   └── tasks/
      └── phase-3/
          └── tasks/

Benefits:
  ✅ Clear separation of development stages
  ✅ Better overview of project progress
  ✅ Easier to manage large numbers of tasks
  ✅ Natural grouping of related work

Example:
  dppm project create web-app --name "Web Application"
  dppm phase create backend --project web-app
  dppm phase create frontend --project web-app
  dppm task create api --project web-app --phase backend
  dppm task create ui --project web-app --phase frontend

📋 TASK-BASED PROJECTS (Simple)
Best for: Small projects with few tasks

Structure:
  project/
  └── tasks/
      ├── task-1.yaml
      ├── task-2.yaml
      └── task-3.yaml

Benefits:
  ✅ Simple and flat structure
  ✅ Quick to set up
  ✅ Good for maintenance tasks

Example:
  dppm project create bugfixes --name "Bug Fixes"
  dppm task create fix-login --project bugfixes --title "Fix login bug"
  dppm task create fix-api --project bugfixes --title "Fix API error"

💡 CHOOSING THE RIGHT TYPE:
Use Phase-Based When:
  • Project has > 10 tasks
  • Clear development stages exist
  • Multiple people working
  • Long-term project

Use Task-Based When:
  • Project has < 10 tasks
  • Simple maintenance work
  • Quick fixes needed
  • Personal todo list

🔍 Related Commands:
  • dppm wiki "create phase"      # Phase management
  • dppm wiki "project workflow"  # Complete examples`)
}

// /* Viser en guide til aktive opgaver. */
func showActiveTasksGuide() {
	fmt.Println(`Active Tasks Guide
=================

📋 FINDING ALL ACTIVE TASKS:

List all in-progress tasks across all projects:
  dppm status active

List active tasks in specific project:
  dppm status active --project web-app

List tasks by status:
  dppm list tasks --project web-app --status in_progress
  dppm list tasks --project web-app --status todo

📊 Example Output:
  Active Tasks (In Progress):
  ==========================

  Project: web-app
  • Authentication API (high priority)
    Phase: backend
    Assignee: john-doe
    Started: 2025-09-23

  • User Interface (medium priority)
    Phase: frontend
    Assignee: jane-smith
    Started: 2025-09-24

💡 MONITORING ACTIVE WORK:

Check who's working on what:
  dppm list tasks --assignee john-doe --status in_progress

Check phase progress:
  dppm list tasks --phase backend --status in_progress

Find overdue active tasks:
  dppm list tasks --status in_progress --overdue

🔧 UPDATING TASK STATUS:

Mark task as in progress:
  dppm task update AUTH-001 --status in_progress

Mark task as completed:
  dppm task update AUTH-001 --status done

⚠️ BEST PRACTICES:
  • Only have 1-2 tasks in_progress per person
  • Update status immediately when starting/stopping work
  • Review active tasks daily
  • Complete tasks before starting new ones

🔍 Related Commands:
  • dppm wiki "status commands"   # All status options
  • dppm wiki "task workflow"     # Task lifecycle`)
}

// /* Viser en guide til aktive faser. */
func showActivePhasesGuide() {
	fmt.Println(`Active Phases Guide
==================

📅 UNDERSTANDING PHASE STATUS:

Phase Status Values:
  • planning   - Phase being designed
  • active     - Currently working on this phase
  • completed  - Phase finished
  • cancelled  - Phase abandoned

📋 FINDING ACTIVE PHASES:

List all phases in a project:
  dppm list phases --project web-app

List only active phases:
  dppm list phases --project web-app --status active

Check phase details:
  dppm phase show backend --project web-app

📊 Example Phase Listing:
  Phases for project: web-app
  ==========================

  ✅ Phase: setup (completed)
     Tasks: 3/3 completed
     Duration: 2025-09-01 to 2025-09-07

  🔄 Phase: backend (active)
     Tasks: 5/8 completed, 2 in progress, 1 blocked
     Duration: 2025-09-08 to 2025-09-22

  📋 Phase: frontend (planning)
     Tasks: 0/5 completed
     Duration: 2025-09-23 to 2025-10-07

🔧 MANAGING PHASE TRANSITIONS:

Activate a phase:
  dppm phase update backend --project web-app --status active

Complete a phase:
  dppm phase update backend --project web-app --status completed

💡 PHASE WORKFLOW:
  1. Create phase in "planning" status
  2. Add all tasks to the phase
  3. Set phase to "active" when ready to start
  4. Work through tasks in dependency order
  5. Mark phase "completed" when all tasks done

⚠️ BEST PRACTICES:
  • Only one phase should be "active" at a time
  • Complete phases before starting new ones
  • Use phases to group related work
  • Plan all phases at project start

🔍 Related Commands:
  • dppm wiki "create phase"      # Phase creation
  • dppm wiki "phase workflow"    # Phase management`)
}

// /* Viser en guide til afhængighedsorden. */
func showDependencyOrderGuide() {
	fmt.Println(`Dependency Order Guide
=====================

🔗 UNDERSTANDING DEPENDENCY ORDER:

Tasks must be completed in dependency order:
  1. Tasks with no dependencies (can start immediately)
  2. Tasks depending only on completed tasks
  3. Tasks with unmet dependencies (blocked)

📋 VIEWING DEPENDENCY ORDER:

Show dependency chain for project:
  dppm status dependencies --project web-app

Show tasks in workable order:
  dppm status project web-app
  # Shows: Ready tasks → Blocked tasks

📊 Example Dependency Chain:
  Dependency Order for web-app:
  ============================

  Level 1 (No dependencies):
  ✅ repo-setup (done)
  ✅ documentation (done)

  Level 2 (Depends on Level 1):
  ✅ docker-env (done) → depends on: repo-setup
  🔄 api-docs (in_progress) → depends on: documentation

  Level 3 (Depends on Level 2):
  📋 database (ready) → depends on: docker-env
  🚫 api-tests (blocked) → depends on: api-docs

  Level 4 (Depends on Level 3):
  🚫 auth-api (blocked) → depends on: database
  🚫 user-api (blocked) → depends on: database

🎯 FINDING NEXT TASK TO WORK ON:

Show ready tasks (no blocking dependencies):
  dppm status project web-app
  # Lists "Ready to Work On" section

Show blocked tasks and their blockers:
  dppm status blocked --project web-app

💡 DEPENDENCY BEST PRACTICES:

Good Dependencies:
  • auth-api depends on database (logical)
  • frontend depends on api (necessary)
  • deploy depends on tests (safe)

Bad Dependencies:
  • Circular: A→B→C→A (impossible)
  • Too deep: A→B→C→D→E→F (bottleneck)
  • Unnecessary: UI→Database (can work parallel)

🔧 MANAGING DEPENDENCIES:

Add dependency to existing task:
  dppm task update frontend --dependency-ids "api,auth"

Remove dependency:
  dppm task update frontend --dependency-ids ""

⚠️ TIPS:
  • Keep chains shallow (max 3-4 levels)
  • Allow parallel work where possible
  • Check for bottlenecks regularly
  • Complete blocking tasks first

🔍 Related Commands:
  • dppm wiki "dependencies"      # Dependency management
  • dppm wiki "blocked tasks"     # Resolving blocks`)
}

// /* Viser en guide til opbygning af projekter. */
func showProjectBuildingGuide() {
	fmt.Println(`Project Building Guide
=====================

🏗️ COMPLETE PROJECT BUILDING WORKFLOW:

Step-by-step guide to building a full project with DPPM.

1️⃣ ANALYZE REQUIREMENTS:
Before creating anything, understand:
  • What are the main deliverables?
  • What are the development phases?
  • What are the dependencies?
  • Who will work on what?

2️⃣ CREATE PROJECT STRUCTURE:
  # Create the project
  dppm project create ecommerce \
    --name "E-Commerce Platform" \
    --owner "dev-team" \
    --description "Full-stack e-commerce solution"

3️⃣ DEFINE PHASES:
  # Phase 1: Foundation
  dppm phase create foundation --project ecommerce \
    --name "Foundation Setup" \
    --goal "Set up development environment and infrastructure"

  # Phase 2: Backend
  dppm phase create backend --project ecommerce \
    --name "Backend Development" \
    --goal "Build API, database, and business logic"

  # Phase 3: Frontend
  dppm phase create frontend --project ecommerce \
    --name "Frontend Development" \
    --goal "Build user interface and experience"

  # Phase 4: Integration
  dppm phase create integration --project ecommerce \
    --name "Integration & Testing" \
    --goal "Connect all components and test"

4️⃣ CREATE TASKS WITH DEPENDENCIES:
  # Foundation tasks (no dependencies)
  dppm task create repo --project ecommerce --phase foundation \
    --title "Repository Setup" --priority high

  dppm task create docker --project ecommerce --phase foundation \
    --title "Docker Environment" --dependency-ids "repo"

  # Backend tasks
  dppm task create database --project ecommerce --phase backend \
    --title "Database Schema" --dependency-ids "docker"

  dppm task create auth --project ecommerce --phase backend \
    --title "Authentication System" --dependency-ids "database"

  dppm task create products-api --project ecommerce --phase backend \
    --title "Products API" --dependency-ids "database"

  dppm task create cart-api --project ecommerce --phase backend \
    --title "Shopping Cart API" --dependency-ids "products-api"

  # Frontend tasks
  dppm task create ui-setup --project ecommerce --phase frontend \
    --title "React Setup" --dependency-ids "docker"

  dppm task create product-list --project ecommerce --phase frontend \
    --title "Product Listing" --dependency-ids "products-api,ui-setup"

  dppm task create cart-ui --project ecommerce --phase frontend \
    --title "Shopping Cart UI" --dependency-ids "cart-api,ui-setup"

5️⃣ VERIFY PROJECT STRUCTURE:
  # Check overall status
  dppm status project ecommerce

  # View dependency order
  dppm status dependencies --project ecommerce

  # Find first tasks to work on
  dppm status project ecommerce | grep "Ready"

6️⃣ EXECUTE PROJECT:
  # Start with ready tasks
  dppm task update repo --status in_progress
  # ... work on task ...
  dppm task update repo --status done

  # Check what's unblocked
  dppm status project ecommerce

  # Continue with next ready task
  dppm task update docker --status in_progress

7️⃣ MONITOR PROGRESS:
  # Daily status check
  dppm status project ecommerce

  # Check active work
  dppm status active --project ecommerce

  # Find blockers
  dppm status blocked --project ecommerce

📊 PROJECT METRICS:
Track progress with:
  • Tasks completed vs total
  • Story points completed
  • Blocked task count
  • Phase completion status

💡 SUCCESS TIPS:
  ✅ Plan all phases before creating tasks
  ✅ Set realistic dependencies
  ✅ Keep 2-3 tasks ready at all times
  ✅ Review and update daily
  ✅ Mark tasks done promptly

🔍 Related Commands:
  • dppm wiki "complete"          # Full example
  • dppm wiki "project types"     # Choosing structure
  • dppm wiki "ai workflow"       # AI-optimized workflow`)
}

// /* Viser en guide til at liste aktivt arbejde. */
func showListActiveGuide() {
	fmt.Println(`List Active Work Guide
=====================

📋 COMMANDS TO LIST ACTIVE WORK:

All active work across projects:
  dppm status active

Active tasks in specific project:
  dppm list tasks --project web-app --status in_progress

Active phases:
  dppm list phases --status active

Ready to start tasks:
  dppm status project web-app
  # Shows "Ready to Work On" section

📊 COMPREHENSIVE ACTIVE WORK VIEW:

Get full picture of active work:
  # 1. Show all in-progress tasks
  dppm status active

  # 2. Show ready tasks
  dppm status project PROJECT_NAME

  # 3. Show blocked tasks
  dppm status blocked

Example Combined Output:
  🔄 IN PROGRESS (3 tasks):
  • web-app: Authentication API (john)
  • web-app: User Interface (jane)
  • mobile: Login Screen (alex)

  📋 READY TO START (5 tasks):
  • web-app: Database Backup
  • web-app: API Documentation
  • mobile: Settings Page
  • mobile: Profile View
  • backend: Cache Layer

  🚫 BLOCKED (2 tasks):
  • web-app: Deploy (waiting for: Tests)
  • mobile: API Integration (waiting for: API)

🔧 FILTERING ACTIVE WORK:

By assignee:
  dppm list tasks --assignee john --status in_progress

By priority:
  dppm list tasks --priority high --status in_progress

By phase:
  dppm list tasks --phase backend --status in_progress

By date:
  dppm list tasks --due-today --status in_progress

💡 DASHBOARD VIEW:
Create a project dashboard:
  echo "=== PROJECT DASHBOARD ==="
  dppm status project web-app
  echo ""
  echo "=== ACTIVE TASKS ==="
  dppm list tasks --project web-app --status in_progress
  echo ""
  echo "=== BLOCKED TASKS ==="
  dppm status blocked --project web-app

🔍 Related Commands:
  • dppm wiki "active tasks"      # Task-specific guide
  • dppm wiki "active phases"     # Phase-specific guide
  • dppm wiki "status commands"   # All status options`)
}

// /* Viser en guide til at finde opgaver. */
func showFindTaskGuide() {
	fmt.Println(`Find Task Guide
==============

🔍 SEARCHING FOR SPECIFIC TASKS:

Find task by ID:
  dppm task show AUTH-001

Find tasks by title (grep):
  dppm list tasks --project web-app | grep -i "auth"

Find tasks by status:
  dppm list tasks --project web-app --status todo
  dppm list tasks --project web-app --status in_progress
  dppm list tasks --project web-app --status done

Find tasks by assignee:
  dppm list tasks --assignee john-doe

Find tasks by priority:
  dppm list tasks --priority high
  dppm list tasks --priority critical

Find tasks in phase:
  dppm list tasks --phase backend

📊 ADVANCED SEARCH PATTERNS:

Find blocked tasks with specific dependency:
  dppm status dependencies --project web-app | grep "auth-api"

Find tasks modified today:
  dppm list tasks --updated-today

Find overdue tasks:
  dppm list tasks --overdue

Find tasks with specific labels:
  dppm list tasks --label "bug"
  dppm list tasks --label "security"

🔧 SEARCH COMBINATIONS:

High priority blocked tasks:
  dppm status blocked --project web-app | grep "high"

In-progress tasks by specific developer:
  dppm list tasks --assignee john --status in_progress

Tasks in backend phase that are ready:
  dppm list tasks --phase backend --status todo |
    grep -v "blocked"

💡 CREATE CUSTOM SEARCHES:

Alias for common searches:
  alias find-my-tasks='dppm list tasks --assignee $(whoami)'
  alias find-urgent='dppm list tasks --priority critical'
  alias find-blocked='dppm status blocked'

Script for task search:
  #!/bin/bash
  # find-task.sh
  dppm list tasks --project $1 | grep -i "$2"

  # Usage: ./find-task.sh web-app "auth"

📝 TASK INFORMATION:

Once found, get full details:
  dppm task show TASK-ID

View task file directly:
  cat ~/Dropbox/project-management/projects/PROJECT/phases/PHASE/tasks/TASK.yaml

🔍 Related Commands:
  • dppm wiki "list active"       # List active work
  • dppm wiki "status commands"   # Status queries
  • dppm wiki "task workflow"     # Task management`)
}

// /* Viser en guide til AI-optimeret arbejdsgang. */
func showAIWorkflowGuide() {
	fmt.Println(`AI-Optimized Workflow Guide
===========================

🤖 DPPM IS DESIGNED FOR AI WORKFLOWS:

DPPM provides structured, verbose output perfect for AI parsing and
comprehensive wiki system for self-service learning.

📚 AI SELF-DISCOVERY:

1. AI runs dppm without args:
   dppm
   # Shows smart guide with next steps

2. AI searches for help:
   dppm --wiki "how to create project"
   dppm --wiki "task dependencies"
   dppm --wiki "find blocked tasks"

3. AI gets complete examples:
   dppm wiki complete
   # Shows full workflow with all commands

🎯 AI PROJECT WORKFLOW:

Step 1: Understand project scope
  dppm wiki "project types"
  # AI learns about phase-based vs task-based

Step 2: Create project structure
  dppm project create ai-app --name "AI Application"
  dppm wiki "create phase"
  dppm phase create ml --project ai-app --name "ML Development"

Step 3: Build task dependency graph
  dppm wiki "dependencies"
  dppm task create data-prep --project ai-app --phase ml
  dppm task create training --project ai-app --phase ml --dependency-ids "data-prep"

Step 4: Monitor and execute
  dppm status project ai-app
  # AI sees ready tasks, blocked tasks, progress

Step 5: Find specific information
  dppm --wiki "active tasks"
  dppm status active --project ai-app
  dppm status dependencies --project ai-app

📊 AI-FRIENDLY FEATURES:

Structured Output:
  • YAML format for easy parsing
  • Consistent status values
  • Clear dependency chains
  • Verbose help text

Self-Service Documentation:
  • Built-in wiki system
  • Complete examples
  • All parameters documented
  • Common patterns included

Intelligent Defaults:
  • Smart suggestions in output
  • Helpful error messages
  • Next-step guidance
  • Context-aware help

🔧 AI AUTOMATION EXAMPLES:

Daily status report:
  dppm status project $PROJECT > daily-report.txt
  dppm status blocked --project $PROJECT >> daily-report.txt
  dppm status active --project $PROJECT >> daily-report.txt

Find next task to work on:
  dppm status project $PROJECT | grep "Ready to Work On" -A 10

Check for blockers:
  dppm status blocked --project $PROJECT
  if [ $? -eq 0 ]; then
    echo "Found blocking issues"
  fi

💡 AI BEST PRACTICES:

1. Always start with wiki:
   dppm wiki "topic"

2. Use structured commands:
   dppm COMMAND SUBCOMMAND --flag value

3. Parse YAML output:
   dppm task show TASK-ID
   # Returns structured YAML

4. Follow dependency order:
   dppm status dependencies --project PROJECT

5. Check before acting:
   dppm status project PROJECT
   # Before updating tasks

🚀 COMPLETE AI EXAMPLE:

AI receives request: "Set up a new web project with auth"

AI executes:
  # Learn how
  dppm wiki "build project"

  # Create structure
  dppm project create web --name "Web App"
  dppm phase create setup --project web
  dppm phase create auth --project web

  # Create tasks
  dppm task create repo --project web --phase setup
  dppm task create docker --project web --phase setup --dependency-ids "repo"
  dppm task create auth-api --project web --phase auth --dependency-ids "docker"

  # Verify
  dppm status project web
  dppm status dependencies --project web

This workflow is fully self-documented and AI-discoverable!

🔍 Related Commands:
  • dppm wiki list               # All topics
  • dppm wiki complete           # Full example
  • dppm --wiki "any question"   # Direct search`)
}

// /* Viser en guide til AI-samarbejdssystemet. */
func showAICollaborationGuide() {
	fmt.Println(`AI Collaboration System
=======================

🤖 AI-TO-AI TASK COORDINATION:

DPPM includes a comprehensive AI collaboration system using DSL markers
for structured handoffs between different AI models. This enables teams
of AI agents to work together on complex projects.

🏷️ DSL MARKER SYSTEM:

Task Assignment:
  ::LARS:1:: Implement user authentication system ::
  ::GEMINI:2:: Design the user interface for login ::

Completion Tracking:
  ::DONE:1,2:: Authentication and UI completed ::

🔄 COLLABORATION WORKFLOW:

1. Create collaboration workspace in project docs
2. Assign tasks to appropriate AI models using DSL markers
3. AIs work on assigned tasks and update progress
4. Mark completed tasks with DONE markers
5. Clean up completed tasks to maintain workspace

📋 AVAILABLE COMMANDS:

Find Active Tasks:
  dppm collab find                # Find all DSL tasks
  dppm collab find docs/          # Search specific directory

Clean Completed Work:
  dppm collab clean               # Remove completed tasks
  dppm collab clean docs/         # Clean specific directory

Learn Collaboration:
  dppm collab wiki                # Collaboration wiki index
  dppm collab wiki "task handoff" # Learn handoff patterns

🎯 AI SPECIALIZATION:

LARS (Claude) - Best for:
  • Code implementation and debugging
  • Technical analysis and architecture
  • Documentation and structured writing
  • Security and best practices review

GEMINI (Google) - Best for:
  • Creative problem-solving and brainstorming
  • User experience and design thinking
  • Content creation and copywriting
  • Alternative approaches and innovation

🚀 INTEGRATION WITH DPPM:

The collaboration system works seamlessly with DPPM projects:
  • Store collaboration tasks in project documentation
  • Link AI tasks to DPPM phases and milestones
  • Use project structure to organize AI workflows
  • Archive completed collaborative work

📊 EXAMPLE COLLABORATIVE PROJECT:

Web Application Development:
  ::LARS:10:: Design system architecture and data models ::
  ::GEMINI:11:: Create user journey maps and wireframes ::
  ::LARS:12:: Implement backend API based on architecture ::
  ::GEMINI:13:: Design responsive frontend components ::
  ::LARS:14:: Write comprehensive tests and documentation ::

Completion:
  ::DONE:10,11,12,13,14:: Web application completed ::

💡 BEST PRACTICES:

Task Design:
  ✅ Clear, actionable task descriptions
  ✅ Specific deliverables and success criteria
  ✅ Appropriate AI assignment based on strengths
  ✅ Logical dependency ordering

Workflow Management:
  ✅ Regular cleanup of completed tasks
  ✅ Archive important decisions and outcomes
  ✅ Use sequential numbering for task IDs
  ✅ Document handoff context clearly

🔍 Getting Started:

Quick Setup:
  # Create collaboration workspace
  mkdir -p docs/collaboration
  echo "# Active AI Tasks" > docs/collaboration/active-tasks.md
  echo "::LARS:1:: Plan the project structure ::" >> docs/collaboration/active-tasks.md

  # Find and manage tasks
  dppm collab find docs/collaboration/
  dppm collab clean docs/collaboration/

🔍 Related Commands:
  • dppm collab wiki "collaboration basics"  # Detailed introduction
  • dppm collab wiki "workflow patterns"     # Common patterns
  • dppm collab wiki "integration"           # DPPM integration
  • dppm wiki "ai workflow"                  # AI-optimized DPPM usage`)
}


// /* Viser task components guiden. */
func showTaskComponentsGuide() {
	fmt.Println(`Task Components Guide
====================

🧩 BREAKING TASKS INTO MANAGEABLE COMPONENTS

Components are sub-tasks that represent specific deliverables within a larger task.
They help track detailed progress and enable parallel work on complex features.

Component Structure:
  • Each component has a unique ID within the task
  • Components include title, description, and status
  • Components can have their own time estimates
  • Components support assignee tracking

Creating Components:
  # Components are defined in the task YAML structure
  components:
    - id: frontend-ui
      title: "User Interface Components"
      description: "Create login form, dashboard, and navigation"
      status: todo
      assignee: frontend-dev
      estimated_hours: 8

    - id: backend-api
      title: "API Endpoints"
      description: "Implement authentication endpoints"
      status: in_progress
      assignee: backend-dev
      estimated_hours: 6

Best Practices:
  ✅ Break complex tasks into 2-8 components
  ✅ Each component should be completable in 1-2 days
  ✅ Use consistent naming across similar tasks
  ✅ Include time estimates for resource planning
  ✅ Assign components to specific team members

Common Component Types:
  • UI/Frontend work
  • Backend/API development
  • Database schema changes
  • Testing and validation
  • Documentation
  • Configuration and deployment

Examples:
  Authentication System Components:
    → Login form implementation
    → Password reset functionality
    → JWT token management
    → Session handling
    → Security testing

  File Management Components:
    → Upload interface
    → File validation
    → Storage integration
    → Download functionality
    → Access permissions`)
}

// /* Viser time tracking guiden. */
func showTimeTrackingGuide() {
	fmt.Println(`Time Tracking Guide
==================

⏱️  TRACKING TIME AND ESTIMATES IN DPPM

DPPM includes built-in time tracking capabilities for tasks and components.
Track estimates, actual time spent, and generate accurate project timelines.

Time Tracking Structure:
  time_tracking:
    estimated_hours: 16
    actual_hours: 18.5
    start_date: "2023-10-01"
    completion_date: "2023-10-03"
    time_logs:
      - date: "2023-10-01"
        hours: 6.5
        description: "Initial setup and research"
        contributor: "dev-team"
      - date: "2023-10-02"
        hours: 8.0
        description: "Core implementation"
        contributor: "dev-team"

Time Logging Commands:
  # Log time for a specific task
  dppm task update auth-system --log-time 4.5 --description "Frontend development"

  # View time summary for project
  dppm status project web-app --include-time

  # Generate time report for phase
  dppm phase show backend --time-report

Best Practices:
  ✅ Estimate time during task creation
  ✅ Log time daily for accuracy
  ✅ Include meaningful descriptions
  ✅ Track time at component level for detail
  ✅ Review estimates vs actual regularly

Time Categories:
  • Development (coding, implementation)
  • Research (investigation, learning)
  • Testing (validation, bug fixes)
  • Documentation (writing, updates)
  • Meetings (planning, reviews)
  • Debugging (issue investigation)

Reporting Features:
  • Project-level time summaries
  • Phase burn-down charts
  • Individual contributor reports
  • Estimation accuracy analysis
  • Overtime and capacity planning

Example Workflow:
  1. Estimate task during creation: --estimated-hours 12
  2. Log daily progress: dppm task update task-id --log-time 3.5
  3. Review weekly: dppm status project --time-summary
  4. Adjust estimates: dppm task update task-id --estimated-hours 15`)
}

// /* Viser issue tracking guiden. */
func showIssueTrackingGuide() {
	fmt.Println(`Issue Tracking Guide
===================

🐛 BUG AND ISSUE MANAGEMENT IN DPPM

Track bugs, improvements, and issues directly within your DPPM tasks.
Integrated issue tracking keeps everything in one place.

Issue Structure:
  issues:
    - id: bug-001
      type: bug
      title: "Login form validation error"
      description: "Email validation fails for certain domains"
      status: open
      priority: high
      reporter: qa-team
      assignee: frontend-dev
      created: "2023-10-01"
      labels: ["frontend", "validation", "critical"]

Issue Types:
  • bug - Defects and problems
  • enhancement - Feature improvements
  • documentation - Doc issues
  • performance - Speed/efficiency problems
  • security - Security vulnerabilities
  • technical-debt - Code quality issues

Issue Lifecycle:
  open → in_progress → testing → resolved → closed

Creating Issues:
  # Add issue to existing task
  dppm task update auth-system --add-issue "Login timeout" --issue-type bug --priority high

  # Create task specifically for issue
  dppm task create bug-fix-001 --title "Fix login validation" --description "Address email domain validation issue"

Issue Management:
  # List all open issues
  dppm status issues --status open

  # Show issues by priority
  dppm status issues --priority high

  # Update issue status
  dppm task update auth-system --issue-id bug-001 --issue-status resolved

Best Practices:
  ✅ Create issues for all identified problems
  ✅ Use consistent labeling system
  ✅ Assign priority levels (low, medium, high, critical)
  ✅ Include reproduction steps for bugs
  ✅ Link issues to related tasks
  ✅ Close issues when verified fixed

Reporting:
  • Open issue counts by project/phase
  • Issue resolution time tracking
  • Bug discovery rates
  • Quality metrics and trends

Integration:
  • Issues can reference task components
  • Link to external bug trackers (GitHub, Jira)
  • Export issues for stakeholder reports
  • Automated issue creation from test failures`)
}

// /* Viser project templates guiden. */
func showProjectTemplatesGuide() {
	fmt.Println(`Project Templates Guide
======================

📋 USING AND CREATING DPPM PROJECT TEMPLATES

Templates provide pre-configured project structures for common development patterns.
Speed up project initialization and ensure consistency across teams.

Available Templates:
  • web - Web application (frontend + backend)
  • api - REST API service
  • mobile - Mobile application
  • library - Software library/package
  • data - Data processing pipeline
  • microservice - Microservice architecture

Using Templates:
  # Initialize project with template
  dppm init my-web-app --template web --name "E-commerce Platform"

  # Create project with template
  dppm project create api-service --template api --name "User Management API"

Template Structure:
  templates/web/
  ├── project.yaml          # Project metadata
  ├── phases/              # Pre-defined phases
  │   ├── planning/        # Requirements gathering
  │   ├── backend/         # API development
  │   ├── frontend/        # UI development
  │   ├── integration/     # System integration
  │   └── deployment/      # Go-live activities
  └── tasks/               # Common task templates

Web Application Template:
  Phases:
    1. Planning & Design
    2. Backend API Development
    3. Frontend Development
    4. Integration & Testing
    5. Deployment & Monitoring

  Common Tasks:
    • Database schema design
    • Authentication system
    • API endpoint implementation
    • Frontend component development
    • Integration testing
    • Deployment pipeline setup

Creating Custom Templates:
  1. Create template directory structure
  2. Define project.yaml with metadata
  3. Add phase and task templates
  4. Configure default dependencies
  5. Include documentation templates

Template Configuration:
  # templates/custom-web/project.yaml
  name: "{{ .ProjectName }}"
  template: "custom-web"
  description: "{{ .Description }}"
  phases:
    - id: setup
      name: "Project Setup"
      tasks:
        - environment-setup
        - dependency-management
        - ci-cd-pipeline

Best Practices:
  ✅ Use templates for repeated project types
  ✅ Customize templates for your organization
  ✅ Include comprehensive task descriptions
  ✅ Define realistic time estimates
  ✅ Document template usage guidelines
  ✅ Version control template definitions`)
}

// /* Viser collaboration guiden. */
func showCollaborationGuide() {
	fmt.Println(`Team Collaboration Guide
=======================

👥 WORKING WITH TEAMS IN DPPM

DPPM supports multi-person teams with role-based access, assignment tracking,
and collaborative workflows for effective project management.

Team Structure:
  team:
    - name: "Alice Developer"
      role: "lead-developer"
      email: "alice@company.com"
      skills: ["backend", "database", "devops"]
    - name: "Bob Frontend"
      role: "frontend-developer"
      email: "bob@company.com"
      skills: ["react", "javascript", "ui-design"]

Task Assignment:
  # Assign tasks to team members
  dppm task create user-profile --assignee "alice@company.com"
  dppm task update existing-task --assignee "bob@company.com"

  # List tasks by assignee
  dppm list tasks --assignee "alice@company.com"
  dppm status active --assignee "bob@company.com"

Collaborative Workflows:
  1. Project Lead creates project and phases
  2. Tasks assigned to appropriate team members
  3. Dependencies defined to coordinate work
  4. Regular status reviews using dppm status commands
  5. Blocked tasks escalated and resolved quickly

Communication Patterns:
  # Daily standups
  dppm status active                    # What can we work on today?
  dppm status blocked                   # What's blocking progress?
  dppm list active --assignee me       # My current workload

  # Weekly reviews
  dppm status project --include-team    # Project health check
  dppm phase show current --team-view   # Phase completion status

Role-Based Access:
  • Project Owner: Full project control
  • Lead Developer: Task creation and assignment
  • Developer: Task updates and time logging
  • Reviewer: Status viewing and reporting
  • Stakeholder: Read-only project visibility

Best Practices:
  ✅ Assign tasks to specific individuals
  ✅ Use meaningful task descriptions for context
  ✅ Update status regularly (daily standups)
  ✅ Communicate dependencies clearly
  ✅ Review blocked tasks in team meetings
  ✅ Celebrate completed milestones

Integration:
  • Sync with external calendar systems
  • Export to team communication tools (Slack, Teams)
  • Generate progress reports for stakeholders
  • Connect to code repositories for commit tracking

Conflict Resolution:
  • Clear ownership of tasks prevents overlap
  • Dependency management coordinates work order
  • Status visibility identifies conflicts early
  • Regular team sync prevents misalignment`)
}

// /* Viser automation guiden. */
func showAutomationGuide() {
	fmt.Println(`Automation Guide
================

🤖 SCRIPTING AND CI/CD INTEGRATION WITH DPPM

Automate project workflows, integrate with CI/CD pipelines, and create
custom scripts for repetitive DPPM operations.

Command Line Scripting:
  #!/bin/bash
  # Daily status report script
  echo "=== Daily Project Status ==="
  dppm list active
  echo -e "\n=== Blocked Tasks ==="
  dppm status blocked
  echo -e "\n=== Team Workload ==="
  dppm status project --team-summary

CI/CD Integration:
  # In your CI pipeline (GitHub Actions, Jenkins, etc.)
  steps:
    - name: Update Task Status
      run: |
        dppm task update build-pipeline --status in_progress
        # Run build/test commands
        if [ $? -eq 0 ]; then
          dppm task update build-pipeline --status done
        else
          dppm task update build-pipeline --status blocked
        fi

Automated Task Creation:
  # Create tasks from issue tracker
  curl -s "https://api.github.com/repos/owner/repo/issues" | \
  jq -r '.[] | "dppm task create issue-\(.number) --title \"\(.title)\" --description \"\(.body)\""' | \
  bash

Project Templates Automation:
  # Bulk project creation
  for project in web-app mobile-app admin-panel; do
    dppm init $project --template web --skip-github
    dppm bind $project
    dppm task create setup --title "Initial Setup" --phase planning
  done

Monitoring and Alerting:
  #!/bin/bash
  # Alert on blocked tasks
  BLOCKED_COUNT=$(dppm status blocked | grep -c "ID:")
  if [ $BLOCKED_COUNT -gt 5 ]; then
    echo "WARNING: $BLOCKED_COUNT tasks are blocked!"
    # Send alert to team (Slack, email, etc.)
  fi

API Integration:
  # Export DPPM data to external systems
  dppm status project --format json | curl -X POST \
    -H "Content-Type: application/json" \
    -d @- https://your-dashboard.com/api/projects

Batch Operations:
  # Bulk task updates
  dppm list tasks --status in_progress --format json | \
  jq -r '.[].id' | \
  xargs -I {} dppm task update {} --add-label "sprint-2"

Custom Commands:
  # Create alias for common operations
  alias daily-standup='dppm list active && dppm status blocked'
  alias my-tasks='dppm list active --assignee $(git config user.email)'

Reporting Automation:
  # Weekly progress report
  #!/bin/bash
  DATE=$(date +%Y-%m-%d)
  echo "# Weekly Report - $DATE" > report.md
  echo "## Active Tasks" >> report.md
  dppm list active >> report.md
  echo "## Completed This Week" >> report.md
  dppm list tasks --status done --updated-since "7 days ago" >> report.md

Best Practices:
  ✅ Use dppm commands in build scripts
  ✅ Automate status updates based on CI results
  ✅ Create monitoring for blocked tasks
  ✅ Generate regular progress reports
  ✅ Integrate with team communication tools
  ✅ Version control your automation scripts`)
}

// /* Viser troubleshooting guiden. */
func showTroubleshootingGuide() {
	fmt.Println(`Troubleshooting Guide
====================

🔧 COMMON ISSUES AND SOLUTIONS

Problem: "No DPPM project found in current directory"
Solution:
  • Run: dppm bind <project-id> to bind current directory
  • Or: cd to a directory with existing DPPM binding
  • Or: Use --project flag: dppm task create --project <project-id>

Problem: "Task creation fails with dependency error"
Solution:
  • Verify dependency task exists: dppm task show <dependency-id>
  • Check project scope: dependencies must be in same project
  • Use correct task ID format (lowercase, no spaces)

Problem: "Wiki topics return 'No specific guide found'"
Solution:
  • Use: dppm wiki list to see all available topics
  • Try alternative search terms
  • Use quotes for multi-word searches: "project types"

Problem: "Collab clean not removing DSL markers"
Solution:
  • Verify DONE marker format: ::DONE:01:: content ::
  • Ensure task ID matches: ::LARS:01:: and ::DONE:01::
  • Check file has .md extension

Problem: "Commands run slowly or timeout"
Solution:
  • Check Dropbox sync status
  • Verify ~/Dropbox/project-management exists
  • Restart Dropbox client if needed
  • Use local binding instead of network paths

Problem: "Project binding lost after directory changes"
Solution:
  • Re-run: dppm bind <project-id>
  • Check .dppm/project.yaml file exists
  • Verify project still exists in Dropbox

Problem: "Task dependencies create circular references"
Solution:
  • Use: dppm status dependencies to visualize
  • Remove problematic dependency: dppm task update --remove-dependency
  • Redesign task breakdown to avoid cycles

Problem: "Phase creation fails with permission error"
Solution:
  • Check Dropbox folder permissions
  • Verify project directory exists
  • Run dppm with proper user permissions
  • Check disk space availability

Problem: "Time tracking data not saving"
Solution:
  • Use correct time format: --log-time 4.5
  • Verify task exists before logging time
  • Check YAML file is writable
  • Ensure proper date format in logs

Problem: "Wiki search returns wrong results"
Solution:
  • Use specific terms: "task creation" not "tasks"
  • Try alternative phrasings
  • Check spelling of search terms
  • Use dppm wiki list to browse available topics

Diagnostic Commands:
  dppm --version                # Check DPPM version
  dppm status project --debug   # Debug project status
  dppm wiki list               # Verify wiki availability
  ls ~/.dppm/                  # Check local configuration

Getting Help:
  • dppm --help for command reference
  • dppm <command> --help for specific help
  • dppm wiki list for available guides
  • Check project documentation
  • Review error messages carefully

Performance Tips:
  • Use local binding to avoid network calls
  • Minimize large YAML files
  • Regular cleanup of completed tasks
  • Use filters to limit output size`)
}

// /* Viser real world guiden. */
func showRealWorldGuide() {
	fmt.Println(`Real World Usage Examples
========================

🌍 PRACTICAL SCENARIOS AND IMPLEMENTATION PATTERNS

Scenario 1: E-commerce Web Application
======================================

Project Structure:
  dppm project create ecommerce --name "Online Store Platform"

  Phases:
    1. Requirements & Planning (2 weeks)
    2. Backend API Development (4 weeks)
    3. Frontend Development (4 weeks)
    4. Payment Integration (2 weeks)
    5. Testing & QA (3 weeks)
    6. Deployment & Launch (1 week)

Sample Tasks with Dependencies:
  # Planning phase
  dppm task create market-research --phase planning --estimated-hours 40
  dppm task create tech-stack --phase planning --dependency-ids market-research

  # Backend development
  dppm task create database-schema --phase backend --dependency-ids tech-stack
  dppm task create user-auth --phase backend --dependency-ids database-schema
  dppm task create product-catalog --phase backend --dependency-ids database-schema
  dppm task create order-system --phase backend --dependency-ids user-auth,product-catalog

Team Assignment:
  Alice (Lead): database-schema, user-auth
  Bob (Backend): product-catalog, order-system
  Carol (Frontend): user-interface, shopping-cart
  Dave (DevOps): deployment, monitoring

Scenario 2: API Microservice Development
========================================

Project Structure:
  dppm project create user-service --template api

  Components per Task:
  dppm task create authentication --components api-endpoints,database-layer,security-tests
  dppm task create user-profile --components crud-operations,validation,caching
  dppm task create admin-panel --components user-management,permissions,audit-log

CI/CD Integration:
  # In deployment pipeline
  dppm task update user-service --status in_progress
  # Run tests
  if tests_pass; then
    dppm task update user-service --status done --log-time 2.5
  else
    dppm task update user-service --status blocked --add-issue "Test failures"
  fi

Scenario 3: Mobile App with Backend
===================================

Cross-Platform Coordination:
  # Backend team
  dppm project create mobile-backend --template api
  dppm task create auth-api --assignee backend-team
  dppm task create user-data-api --dependency-ids auth-api

  # Mobile team
  dppm project create mobile-app --template mobile
  dppm task create login-screen --dependency-ids auth-api --assignee mobile-team
  dppm task create profile-screen --dependency-ids user-data-api

Scenario 4: Data Processing Pipeline
====================================

ETL Workflow:
  dppm project create data-pipeline --template data

  Sequential Tasks:
  dppm task create data-extraction --estimated-hours 16
  dppm task create data-transformation --dependency-ids data-extraction
  dppm task create data-loading --dependency-ids data-transformation
  dppm task create monitoring-setup --dependency-ids data-loading

Scenario 5: Legacy System Migration
===================================

Phased Migration:
  Phase 1: Assessment (dppm phase create assessment)
    → Inventory existing systems
    → Identify migration challenges
    → Plan migration strategy

  Phase 2: Infrastructure (dppm phase create infrastructure)
    → Set up new environment
    → Configure databases
    → Establish security protocols

  Phase 3: Data Migration (dppm phase create data-migration)
    → Export legacy data
    → Transform data format
    → Import to new system
    → Validate data integrity

Real-World Tips:
================

Team Coordination:
  • Daily: dppm list active --assignee me
  • Weekly: dppm status project --include-team
  • Monthly: dppm phase show --completion-report

Issue Management:
  • Link DPPM tasks to GitHub issues
  • Use consistent labeling across projects
  • Track issue resolution time

Time Tracking:
  • Log time daily for accuracy
  • Use consistent time categories
  • Review estimates vs actual monthly

Automation Integration:
  • Update task status from CI/CD results
  • Generate reports for stakeholders
  • Monitor blocked tasks automatically

Quality Assurance:
  • Include testing in every phase
  • Define clear acceptance criteria
  • Track defect discovery rates`)
}

// /* Initialiserer 'wiki' kommandoen. */
func init() {
	rootCmd.AddCommand(wikiCmd)
}