package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var wikiCmd = &cobra.Command{
	Use:   "wiki [search-terms]",
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

func showWikiTopics() {
	fmt.Println(`Available Wiki Topics:
=====================

📁 Project Management:
  • "create project" - Complete project creation guide
  • "project structure" - Directory organization
  • "project workflow" - End-to-end project management

📋 Phase Management:
  • "create phase" - Phase creation and organization
  • "phase workflow" - Managing development phases
  • "phase structure" - Phase directory layout

✅ Task Management:
  • "create task" - Task creation with all options
  • "task dependencies" - Dependency management
  • "task workflow" - Task lifecycle management
  • "task components" - Breaking tasks into components

📊 Status & Reporting:
  • "status commands" - All status reporting options
  • "dependency chains" - Understanding task relationships
  • "blocked tasks" - Identifying and resolving blocks

🔧 Advanced Features:
  • "time tracking" - Time logging and estimation
  • "issue tracking" - Bug and issue management
  • "project templates" - Using and creating templates

🚀 Complete Workflows:
  • "complete" - Full project creation to completion example
  • "real world" - Practical usage scenarios`)
}

func searchWiki(searchTerm string) {
	switch {
	case strings.Contains(searchTerm, "create project") || strings.Contains(searchTerm, "project creation"):
		showProjectCreationGuide()
	case strings.Contains(searchTerm, "create phase") || strings.Contains(searchTerm, "phase creation"):
		showPhaseCreationGuide()
	case strings.Contains(searchTerm, "create task") || strings.Contains(searchTerm, "task creation"):
		showTaskCreationGuide()
	case strings.Contains(searchTerm, "dependencies") || strings.Contains(searchTerm, "dependency"):
		showDependencyGuide()
	case strings.Contains(searchTerm, "status") || strings.Contains(searchTerm, "reporting"):
		showStatusGuide()
	case strings.Contains(searchTerm, "blocked") || strings.Contains(searchTerm, "blocking"):
		showBlockedTasksGuide()
	case strings.Contains(searchTerm, "structure") || strings.Contains(searchTerm, "organization"):
		showStructureGuide()
	case strings.Contains(searchTerm, "workflow") || strings.Contains(searchTerm, "process"):
		showWorkflowGuide()
	default:
		fmt.Printf("No specific guide found for '%s'\n\n", searchTerm)
		fmt.Println("Try one of these common searches:")
		fmt.Println("  dppm wiki \"create project\"")
		fmt.Println("  dppm wiki \"dependencies\"")
		fmt.Println("  dppm wiki \"status commands\"")
		fmt.Println("  dppm wiki list")
	}
}

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

func init() {
	rootCmd.AddCommand(wikiCmd)
}