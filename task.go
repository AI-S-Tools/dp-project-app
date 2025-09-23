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
	ID          string            `yaml:"id"`
	Title       string            `yaml:"title"`
	ProjectID   string            `yaml:"project_id"`
	PhaseID     string            `yaml:"phase_id,omitempty"`
	Status      string            `yaml:"status"`
	Priority    string            `yaml:"priority"`
	Assignee    string            `yaml:"assignee,omitempty"`
	Reporter    string            `yaml:"reporter,omitempty"`
	Created     string            `yaml:"created"`
	Updated     string            `yaml:"updated"`
	DueDate     string            `yaml:"due_date,omitempty"`
	StoryPoints int               `yaml:"story_points,omitempty"`
	Description string            `yaml:"description"`

	// Advanced features
	Components    []Component       `yaml:"components,omitempty"`
	Issues        []Issue           `yaml:"issues,omitempty"`
	DependencyIDs []string          `yaml:"dependency_ids,omitempty"`
	BlockedBy     []string          `yaml:"blocked_by,omitempty"`
	Blocking      []string          `yaml:"blocking,omitempty"`
	Labels        []string          `yaml:"labels,omitempty"`
	Attachments   []string          `yaml:"attachments,omitempty"`
	Comments      []Comment         `yaml:"comments,omitempty"`
	TimeTracking  TimeTracking      `yaml:"time_tracking,omitempty"`
	Progress      Progress          `yaml:"progress,omitempty"`
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
	Date        string `yaml:"date"`
	Hours       float32 `yaml:"hours"`
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
}

type Progress struct {
	TotalComponents       int `yaml:"total_components"`
	CompletedComponents   int `yaml:"completed_components"`
	CompletionPercentage  int `yaml:"completion_percentage"`
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
  dppm task create auth-system --project dash-lxd --title "User Authentication System"
  dppm task create file-ops --project web-app --title "File Operations" --phase phase-1
  dppm task create bug-fix --project api --title "Fix Login Bug" --priority high`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]
		title, _ := cmd.Flags().GetString("title")
		projectID, _ := cmd.Flags().GetString("project")
		phaseID, _ := cmd.Flags().GetString("phase")
		description, _ := cmd.Flags().GetString("description")
		priority, _ := cmd.Flags().GetString("priority")
		assignee, _ := cmd.Flags().GetString("assignee")

		if title == "" {
			title = taskID
		}
		if priority == "" {
			priority = "medium"
		}

		task := Task{
			ID:          taskID,
			Title:       title,
			ProjectID:   projectID,
			PhaseID:     phaseID,
			Status:      "todo",
			Priority:    priority,
			Assignee:    assignee,
			Reporter:    "dppm-user",
			Created:     time.Now().Format("2006-01-02"),
			Updated:     time.Now().Format("2006-01-02"),
			Description: description,
			Components:  []Component{},
			Issues:      []Issue{},
			DependencyIDs: []string{},
			BlockedBy:   []string{},
			Blocking:    []string{},
			Labels:      []string{},
			Comments:    []Comment{},
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

func init() {
	createTaskCmd.Flags().StringP("title", "t", "", "Task title")
	createTaskCmd.Flags().StringP("project", "p", "", "Project ID (required)")
	createTaskCmd.Flags().StringP("phase", "s", "", "Phase ID (optional)")
	createTaskCmd.Flags().StringP("description", "d", "", "Task description")
	createTaskCmd.Flags().String("priority", "medium", "Task priority (low, medium, high, critical)")
	createTaskCmd.Flags().StringP("assignee", "a", "", "Task assignee")

	createTaskCmd.MarkFlagRequired("project")

	taskCmd.AddCommand(createTaskCmd)
}