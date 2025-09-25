package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Task struct {
	ID          string `yaml:"id"`
	Title       string `yaml:"title"`
	ProjectID   string `yaml:"project_id"`
	PhaseID     string `yaml:"phase_id,omitempty"`
	Status      string `yaml:"status"`
	Priority    string `yaml:"priority"`
	Assignee    string `yaml:"assignee,omitempty"`
	Reporter    string `yaml:"reporter,omitempty"`
	Created     string `yaml:"created"`
	Updated     string `yaml:"updated"`
	DueDate     string `yaml:"due_date,omitempty"`
	StoryPoints int    `yaml:"story_points,omitempty"`
	Description string `yaml:"description"`

	// Advanced features
	Components    []Component  `yaml:"components,omitempty"`
	Issues        []Issue      `yaml:"issues,omitempty"`
	DependencyIDs []string     `yaml:"dependency_ids,omitempty"`
	BlockedBy     []string     `yaml:"blocked_by,omitempty"`
	Blocking      []string     `yaml:"blocking,omitempty"`
	Labels        []string     `yaml:"labels,omitempty"`
	Attachments   []string     `yaml:"attachments,omitempty"`
	Comments      []Comment    `yaml:"comments,omitempty"`
	TimeTracking  TimeTracking `yaml:"time_tracking,omitempty"`
	Progress      Progress     `yaml:"progress,omitempty"`
}

type Component struct {
	ID          string `yaml:"id"`
	Title       string `yaml:"title"`
	Status      string `yaml:"status"`
	Type        string `yaml:"type"`
	Description string `yaml:"description,omitempty"`
	AssignedTo  string `yaml:"assigned_to,omitempty"`
	Created     string `yaml:"created"`
	Updated     string `yaml:"updated"`
}

type Issue struct {
	ID              string `yaml:"id"`
	Title           string `yaml:"title"`
	Type            string `yaml:"type"`
	Status          string `yaml:"status"`
	Severity        string `yaml:"severity,omitempty"`
	ParentComponent string `yaml:"parent_component,omitempty"`
	Description     string `yaml:"description"`
	ReportedBy      string `yaml:"reported_by,omitempty"`
	AssignedTo      string `yaml:"assigned_to,omitempty"`
	Created         string `yaml:"created"`
	Updated         string `yaml:"updated"`
}

type Comment struct {
	Timestamp string `yaml:"timestamp"`
	Author    string `yaml:"author"`
	Content   string `yaml:"content"`
	Type      string `yaml:"type"`
}

type TimeTracking struct {
	EstimatedHours int       `yaml:"estimated_hours,omitempty"`
	ActualHours    int       `yaml:"actual_hours,omitempty"`
	TimeLogs       []TimeLog `yaml:"time_logs,omitempty"`
}

type TimeLog struct {
	Date        string  `yaml:"date"`
	Hours       float32 `yaml:"hours"`
	Description string  `yaml:"description"`
	Author      string  `yaml:"author"`
}

type Progress struct {
	TotalComponents      int `yaml:"total_components"`
	CompletedComponents  int `yaml:"completed_components"`
	CompletionPercentage int `yaml:"completion_percentage"`
	TotalIssues          int `yaml:"total_issues"`
	ResolvedIssues       int `yaml:"resolved_issues"`
	OpenBugs             int `yaml:"open_bugs"`
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task management commands",
	Long: `Task Management Commands

Manage tasks within projects and sprints. Tasks are the core work units
that can be broken down into components and tracked with issues.

Available Commands:
  create       Create a new task
  show         Display detailed task information
  update       Update task properties
  list         List tasks (use 'dppm list tasks' instead)
  component    Manage task components
  issue        Manage task issues
  dependency   Manage task dependencies

Examples:
  dppm task create auth-system --project dash-lxd --title "Authentication System"
  dppm task show auth-system
  dppm task update auth-system --status in_progress

For more information about a specific command, use:
  dppm task [command] --help`,
}

var createTaskCmd = &cobra.Command{
	Use:   "create [task-id]",
	Short: "Create a new task",
	Long: `Create a New Task

Creates a new task within a project and optionally assigns it to a sprint.
Tasks are the primary work units that can be broken into components and
tracked with issues and dependencies.

Arguments:
  task-id    Unique identifier for the task (required)
             Must be lowercase, alphanumeric with hyphens allowed
             Examples: auth-system, file-browser, ui-enhancement

Examples:
  dppm task create auth-system --project dash-lxd --title "User Authentication System" --description "Implement JWT-based authentication with login/logout functionality"
  dppm task create file-ops --project web-app --title "File Operations" --phase phase-1 --description "Create file upload, download, and management features"
  dppm task create bug-fix --project api --title "Fix Login Bug" --priority high --description "Resolve authentication timeout issues reported by users"

💡 AI Best Practice:
  Always include descriptions for better task context and collaboration.
  Descriptions help other AI agents understand task requirements and scope.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set default project from local context if not explicitly provided
		setDefaultProjectFlag(cmd, "project")

		taskID := args[0]
		title, _ := cmd.Flags().GetString("title")
		projectID, _ := cmd.Flags().GetString("project")
		phaseID, _ := cmd.Flags().GetString("phase")
		description, _ := cmd.Flags().GetString("description")
		priority, _ := cmd.Flags().GetString("priority")
		assignee, _ := cmd.Flags().GetString("assignee")

		// Validate that project ID is available (either from flag or local binding)
		if projectID == "" {
			fmt.Fprintf(os.Stderr, "Error: project is required. Either use --project flag or run 'dppm bind PROJECT_ID' to set local project context.\n")
			return
		}

		if title == "" {
			title = taskID
		}
		if priority == "" {
			priority = "medium"
		}

		// Warn about missing description
		if description == "" {
			fmt.Println("⚠️  Warning: Task created without description")
			fmt.Println("💡 Consider adding --description for better AI collaboration")
			fmt.Println("   Example: --description \"Detailed explanation of what needs to be done\"")
			fmt.Println()
		}

		task := Task{
			ID:            taskID,
			Title:         title,
			ProjectID:     projectID,
			PhaseID:       phaseID,
			Status:        "todo",
			Priority:      priority,
			Assignee:      assignee,
			Reporter:      "dppm-user",
			Created:       time.Now().Format("2006-01-02"),
			Updated:       time.Now().Format("2006-01-02"),
			Description:   description,
			Components:    []Component{},
			Issues:        []Issue{},
			DependencyIDs: []string{},
			BlockedBy:     []string{},
			Blocking:      []string{},
			Labels:        []string{},
			Comments:      []Comment{},
		}

		// Create task directory structure
		var taskDir string
		if phaseID != "" {
			taskDir = filepath.Join(projectsPath, "projects", projectID, "phases", phaseID, "tasks")
		} else {
			taskDir = filepath.Join(projectsPath, "projects", projectID, "tasks")
		}

		if err := os.MkdirAll(taskDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating task directory: %v\n", err)
			return
		}

		taskFile := filepath.Join(taskDir, taskID+".yaml")
		data, err := yaml.Marshal(task)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling task: %v\n", err)
			return
		}

		if err := os.WriteFile(taskFile, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing task file: %v\n", err)
			return
		}

		fmt.Printf("Task '%s' created successfully in project '%s'\n", taskID, projectID)
		if phaseID != "" {
			fmt.Printf("Assigned to phase: %s\n", phaseID)
		}
	},
}

var showTaskCmd = &cobra.Command{
	Use:   "show [task-id]",
	Short: "Display detailed task information",
	Long: `Display Detailed Task Information

Shows comprehensive information about a specific task including its status,
dependencies, components, issues, and complete metadata.

Arguments:
  task-id    Task identifier to display

Examples:
  dppm task show auth-system
  dppm task show file-ops --project web-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set default project from local context if not explicitly provided
		setDefaultProjectFlag(cmd, "project")

		taskID := args[0]
		projectID, _ := cmd.Flags().GetString("project")

		// If no project specified, search all projects
		if projectID == "" {
			searchAndShowTask(taskID)
			return
		}

		showTask(projectID, "", taskID)
	},
}

var updateTaskCmd = &cobra.Command{
	Use:   "update [task-id]",
	Short: "Update task properties",
	Long: `Update Task Properties

Updates one or more properties of an existing task. Only specified properties
will be changed; others remain unchanged.

Arguments:
  task-id    Task identifier to update

Examples:
  dppm task update auth-system --status in_progress
  dppm task update file-ops --assignee john-doe --priority high
  dppm task update bug-fix --status done --description "Fixed login issue"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set default project from local context if not explicitly provided
		setDefaultProjectFlag(cmd, "project")

		taskID := args[0]
		projectID, _ := cmd.Flags().GetString("project")

		if projectID == "" {
			searchAndUpdateTask(taskID, cmd)
			return
		}

		updateTask(projectID, "", taskID, cmd)
	},
}

func searchAndShowTask(taskID string) {
	projectsDir := filepath.Join(projectsPath, "projects")

	err := filepath.Walk(projectsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.Name() == taskID+".yaml" {
			// Read and display task
			data, err := os.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading task file: %v\n", err)
				return nil
			}

			var task Task
			if err := yaml.Unmarshal(data, &task); err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing task file: %v\n", err)
				return nil
			}

			displayTask(task)
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Task '%s' not found\n", taskID)
	}
}

func showTask(projectID, phaseID, taskID string) {
	var taskFile string
	if phaseID != "" {
		taskFile = filepath.Join(projectsPath, "projects", projectID, "phases", phaseID, "tasks", taskID+".yaml")
	} else {
		// Try both locations
		taskFile = filepath.Join(projectsPath, "projects", projectID, "tasks", taskID+".yaml")
		if _, err := os.Stat(taskFile); os.IsNotExist(err) {
			// Search in phases
			phasesDir := filepath.Join(projectsPath, "projects", projectID, "phases")
			err := filepath.Walk(phasesDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				if info.Name() == taskID+".yaml" {
					taskFile = path
					return filepath.SkipDir
				}
				return nil
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Task '%s' not found in project '%s'\n", taskID, projectID)
				return
			}
		}
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading task file: %v\n", err)
		return
	}

	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing task file: %v\n", err)
		return
	}

	displayTask(task)
}

func displayTask(task Task) {
	fmt.Printf("Task: %s\n", task.ID)
	fmt.Printf("================\n\n")
	fmt.Printf("Title: %s\n", task.Title)
	fmt.Printf("Project: %s\n", task.ProjectID)
	if task.PhaseID != "" {
		fmt.Printf("Phase: %s\n", task.PhaseID)
	}
	fmt.Printf("Status: %s\n", task.Status)
	fmt.Printf("Priority: %s\n", task.Priority)

	if task.Assignee != "" {
		fmt.Printf("Assignee: %s\n", task.Assignee)
	}
	if task.Reporter != "" {
		fmt.Printf("Reporter: %s\n", task.Reporter)
	}

	fmt.Printf("Created: %s\n", task.Created)
	fmt.Printf("Updated: %s\n", task.Updated)

	if task.DueDate != "" {
		fmt.Printf("Due Date: %s\n", task.DueDate)
	}
	if task.StoryPoints > 0 {
		fmt.Printf("Story Points: %d\n", task.StoryPoints)
	}

	if task.Description != "" {
		fmt.Printf("\nDescription:\n%s\n", task.Description)
	}

	if len(task.DependencyIDs) > 0 {
		fmt.Printf("\nDependencies: %v\n", task.DependencyIDs)
	}
	if len(task.BlockedBy) > 0 {
		fmt.Printf("Blocked By: %v\n", task.BlockedBy)
	}
	if len(task.Blocking) > 0 {
		fmt.Printf("Blocking: %v\n", task.Blocking)
	}

	if len(task.Labels) > 0 {
		fmt.Printf("Labels: %v\n", task.Labels)
	}

	if len(task.Components) > 0 {
		fmt.Printf("\nComponents:\n")
		for _, comp := range task.Components {
			fmt.Printf("  - %s (%s): %s\n", comp.ID, comp.Status, comp.Title)
		}
	}

	if len(task.Issues) > 0 {
		fmt.Printf("\nIssues:\n")
		for _, issue := range task.Issues {
			fmt.Printf("  - %s (%s): %s\n", issue.ID, issue.Status, issue.Title)
		}
	}
}

func searchAndUpdateTask(taskID string, cmd *cobra.Command) {
	projectsDir := filepath.Join(projectsPath, "projects")
	updated := false

	err := filepath.Walk(projectsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.Name() == taskID+".yaml" {
			// Extract project ID from path
			rel, _ := filepath.Rel(projectsDir, path)
			parts := filepath.SplitList(rel)
			if len(parts) > 0 {
				projectID := parts[0]

				// Update task
				if updateTaskFile(path, cmd) {
					fmt.Printf("Task '%s' updated successfully in project '%s'\n", taskID, projectID)
					updated = true
				}
				return filepath.SkipDir
			}
		}
		return nil
	})

	if err != nil || !updated {
		fmt.Fprintf(os.Stderr, "Task '%s' not found\n", taskID)
	}
}

func updateTask(projectID, phaseID, taskID string, cmd *cobra.Command) {
	var taskFile string
	if phaseID != "" {
		taskFile = filepath.Join(projectsPath, "projects", projectID, "phases", phaseID, "tasks", taskID+".yaml")
	} else {
		// Try both locations
		taskFile = filepath.Join(projectsPath, "projects", projectID, "tasks", taskID+".yaml")
		if _, err := os.Stat(taskFile); os.IsNotExist(err) {
			// Search in phases
			phasesDir := filepath.Join(projectsPath, "projects", projectID, "phases")
			err := filepath.Walk(phasesDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				if info.Name() == taskID+".yaml" {
					taskFile = path
					return filepath.SkipDir
				}
				return nil
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Task '%s' not found in project '%s'\n", taskID, projectID)
				return
			}
		}
	}

	if updateTaskFile(taskFile, cmd) {
		fmt.Printf("Task '%s' updated successfully\n", taskID)
	}
}

func updateTaskFile(taskFile string, cmd *cobra.Command) bool {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading task file: %v\n", err)
		return false
	}

	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing task file: %v\n", err)
		return false
	}

	// Update fields if provided
	if cmd.Flags().Changed("status") {
		status, _ := cmd.Flags().GetString("status")
		task.Status = status
	}
	if cmd.Flags().Changed("priority") {
		priority, _ := cmd.Flags().GetString("priority")
		task.Priority = priority
	}
	if cmd.Flags().Changed("assignee") {
		assignee, _ := cmd.Flags().GetString("assignee")
		task.Assignee = assignee
	}
	if cmd.Flags().Changed("title") {
		title, _ := cmd.Flags().GetString("title")
		task.Title = title
	}
	if cmd.Flags().Changed("description") {
		description, _ := cmd.Flags().GetString("description")
		task.Description = description
	}
	if cmd.Flags().Changed("due-date") {
		dueDate, _ := cmd.Flags().GetString("due-date")
		task.DueDate = dueDate
	}
	if cmd.Flags().Changed("story-points") {
		storyPoints, _ := cmd.Flags().GetInt("story-points")
		task.StoryPoints = storyPoints
	}

	// Update timestamp
	task.Updated = time.Now().Format("2006-01-02")

	// Write back to file
	data, err = yaml.Marshal(task)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling task: %v\n", err)
		return false
	}

	if err := os.WriteFile(taskFile, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing task file: %v\n", err)
		return false
	}

	return true
}

var doneTaskCmd = &cobra.Command{
	Use:   "done [task-id]",
	Short: "Mark a task as completed",
	Long: `Mark Task as Completed

Quickly mark a task as done. This is a shortcut for 'dppm task update TASK_ID --status done'.

Arguments:
  task-id    Task identifier to mark as done

Examples:
  dppm task done auth-system        # Mark auth-system task as completed
  dppm task done bug-fix           # Mark bug-fix task as completed

  # In project-bound directory (after 'dppm bind PROJECT_ID'):
  dppm task done feature-x         # Auto-scoped to bound project

💡 Alternative: You can also use 'dppm task update TASK_ID --status done' for the same result.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set default project from local context if not explicitly provided
		setDefaultProjectFlag(cmd, "project")

		taskID := args[0]
		projectID, _ := cmd.Flags().GetString("project")

		if projectID == "" {
			// Search and update in all projects
			searchAndMarkTaskDone(taskID)
			return
		}

		// Mark task as done in specific project
		markTaskDone(projectID, "", taskID)
	},
}

func searchAndMarkTaskDone(taskID string) {
	projectsDir := filepath.Join(projectsPath, "projects")
	projects, err := os.ReadDir(projectsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading projects directory: %v\n", err)
		return
	}

	var found bool
	for _, project := range projects {
		if project.IsDir() {
			projectPath := filepath.Join(projectsDir, project.Name())

			// Check in project root tasks
			taskFile := filepath.Join(projectPath, "tasks", taskID+".yaml")
			if _, err := os.Stat(taskFile); err == nil {
				markTaskDone(project.Name(), "", taskID)
				found = true
				return
			}

			// Check in phase tasks
			phasesPath := filepath.Join(projectPath, "phases")
			if phases, err := os.ReadDir(phasesPath); err == nil {
				for _, phase := range phases {
					if phase.IsDir() {
						taskFile = filepath.Join(phasesPath, phase.Name(), "tasks", taskID+".yaml")
						if _, err := os.Stat(taskFile); err == nil {
							markTaskDone(project.Name(), phase.Name(), taskID)
							found = true
							return
						}
					}
				}
			}
		}
	}

	if !found {
		fmt.Fprintf(os.Stderr, "Error: Task '%s' not found in any project\n", taskID)
		fmt.Fprintf(os.Stderr, "Use 'dppm list tasks' to see available tasks\n")
	}
}

func markTaskDone(projectID, phaseID, taskID string) {
	var taskPath string
	if phaseID != "" {
		taskPath = filepath.Join(projectsPath, "projects", projectID, "phases", phaseID, "tasks", taskID+".yaml")
	} else {
		taskPath = filepath.Join(projectsPath, "projects", projectID, "tasks", taskID+".yaml")
	}

	// Read existing task
	data, err := os.ReadFile(taskPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading task file: %v\n", err)
		return
	}

	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing task file: %v\n", err)
		return
	}

	// Update status and timestamp
	oldStatus := task.Status
	task.Status = "done"
	task.Updated = time.Now().Format("2006-01-02")

	// Write updated task
	updatedData, err := yaml.Marshal(task)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling task: %v\n", err)
		return
	}

	if err := os.WriteFile(taskPath, updatedData, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing task file: %v\n", err)
		return
	}

	fmt.Printf("✅ Task '%s' marked as done\n", taskID)
	fmt.Printf("📁 Project: %s", projectID)
	if phaseID != "" {
		fmt.Printf(" → Phase: %s", phaseID)
	}
	fmt.Println()
	if oldStatus != "done" {
		fmt.Printf("📊 Status: %s → done\n", oldStatus)
	}
	fmt.Printf("⏰ Updated: %s\n", time.Now().Format("2006-01-02"))
}

func init() {
	createTaskCmd.Flags().StringP("title", "t", "", "Task title")
	createTaskCmd.Flags().StringP("project", "p", "", "Project ID (required)")
	createTaskCmd.Flags().StringP("phase", "s", "", "Phase ID (optional)")
	createTaskCmd.Flags().StringP("description", "d", "", "Task description")
	createTaskCmd.Flags().String("priority", "medium", "Task priority (low, medium, high, critical)")
	createTaskCmd.Flags().StringP("assignee", "a", "", "Task assignee")

	// Project flag will be handled manually in Run function to allow auto-scoping from local binding
	// createTaskCmd.MarkFlagRequired("project")

	// Show command flags
	showTaskCmd.Flags().StringP("project", "p", "", "Project ID (if not specified, searches all projects)")

	// Update command flags
	updateTaskCmd.Flags().StringP("project", "p", "", "Project ID (if not specified, searches all projects)")
	updateTaskCmd.Flags().String("status", "", "Task status (todo, in_progress, review, blocked, done)")
	updateTaskCmd.Flags().String("priority", "", "Task priority (low, medium, high, critical)")
	updateTaskCmd.Flags().StringP("assignee", "a", "", "Task assignee")
	updateTaskCmd.Flags().StringP("title", "t", "", "Task title")
	updateTaskCmd.Flags().StringP("description", "d", "", "Task description")
	updateTaskCmd.Flags().String("due-date", "", "Due date (YYYY-MM-DD)")
	updateTaskCmd.Flags().Int("story-points", 0, "Story points")

	// Done command flags
	doneTaskCmd.Flags().StringP("project", "p", "", "Project ID (if not specified, searches all projects)")

	taskCmd.AddCommand(createTaskCmd)
	taskCmd.AddCommand(showTaskCmd)
	taskCmd.AddCommand(updateTaskCmd)
	taskCmd.AddCommand(doneTaskCmd)
}
