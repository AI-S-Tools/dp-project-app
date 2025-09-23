package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show project status and dependency information",
	Long: `Status Command

Display comprehensive status information about projects, dependencies,
and task relationships. Shows what tasks are blocked by dependencies
and provides overview of project health.

Available Subcommands:
  project      Show overall project status
  dependencies Show all dependency chains
  blocked      Show tasks blocked by dependencies
  active       Show tasks that can be worked on now

Examples:
  dppm status project dash-lxd
  dppm status dependencies
  dppm status blocked
  dppm status active --project dash-lxd`,
}

var statusProjectCmd = &cobra.Command{
	Use:   "project [project-id]",
	Short: "Show project status overview",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

		fmt.Printf("Project Status: %s\n", projectID)
		fmt.Println("=====================")

		// Load all tasks for project
		tasks, err := loadProjectTasks(projectID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
			return
		}

		// Calculate statistics
		todoCount := 0
		inProgressCount := 0
		doneCount := 0
		blockedCount := 0

		for _, task := range tasks {
			switch task.Status {
			case "todo":
				if isTaskBlocked(task, tasks) {
					blockedCount++
				} else {
					todoCount++
				}
			case "in_progress":
				inProgressCount++
			case "done":
				doneCount++
			}
		}

		fmt.Printf("Total Tasks: %d\n", len(tasks))
		fmt.Printf("✅ Done: %d\n", doneCount)
		fmt.Printf("🔄 In Progress: %d\n", inProgressCount)
		fmt.Printf("📋 Ready to Start: %d\n", todoCount)
		fmt.Printf("🚫 Blocked: %d\n", blockedCount)

		if blockedCount > 0 {
			fmt.Println("\n🚫 Blocked Tasks:")
			for _, task := range tasks {
				if task.Status == "todo" && isTaskBlocked(task, tasks) {
					blockers := getBlockingTasks(task, tasks)
					fmt.Printf("  • %s (blocked by: %s)\n", task.Title, strings.Join(blockers, ", "))
				}
			}
		}

		if todoCount > 0 {
			fmt.Println("\n📋 Ready to Work On:")
			for _, task := range tasks {
				if task.Status == "todo" && !isTaskBlocked(task, tasks) {
					fmt.Printf("  • %s (%s priority)\n", task.Title, task.Priority)
				}
			}
		}
	},
}

var statusBlockedCmd = &cobra.Command{
	Use:   "blocked",
	Short: "Show all blocked tasks",
	Long: `Show Blocked Tasks

Display all tasks that are currently blocked by dependencies.
Shows which tasks are blocking each blocked task.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")

		if projectID != "" {
			showBlockedTasksForProject(projectID)
		} else {
			showAllBlockedTasks()
		}
	},
}

var statusDependenciesCmd = &cobra.Command{
	Use:   "dependencies",
	Short: "Show all dependency chains",
	Long: `Show Dependency Chains

Display comprehensive view of all task dependencies across projects.
Shows the dependency graph and highlights potential issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")

		if projectID != "" {
			showDependenciesForProject(projectID)
		} else {
			showAllDependencies()
		}
	},
}

func loadProjectTasks(projectID string) ([]Task, error) {
	var tasks []Task

	tasksDir := filepath.Join(projectsPath, "projects", projectID, "tasks")

	entries, err := os.ReadDir(tasksDir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") {
			taskFile := filepath.Join(tasksDir, entry.Name())

			data, err := os.ReadFile(taskFile)
			if err != nil {
				continue
			}

			var task Task
			if err := yaml.Unmarshal(data, &task); err != nil {
				continue
			}

			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}

func isTaskBlocked(task Task, allTasks []Task) bool {
	if len(task.DependencyIDs) == 0 {
		return false
	}

	taskMap := make(map[string]Task)
	for _, t := range allTasks {
		taskMap[t.ID] = t
	}

	for _, depID := range task.DependencyIDs {
		if depTask, exists := taskMap[depID]; exists {
			if depTask.Status != "done" {
				return true
			}
		}
	}

	return false
}

func getBlockingTasks(task Task, allTasks []Task) []string {
	var blockers []string

	taskMap := make(map[string]Task)
	for _, t := range allTasks {
		taskMap[t.ID] = t
	}

	for _, depID := range task.DependencyIDs {
		if depTask, exists := taskMap[depID]; exists {
			if depTask.Status != "done" {
				blockers = append(blockers, depTask.Title)
			}
		}
	}

	return blockers
}

func showBlockedTasksForProject(projectID string) {
	tasks, err := loadProjectTasks(projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		return
	}

	fmt.Printf("Blocked Tasks in %s:\n", projectID)
	fmt.Println("========================")

	hasBlocked := false
	for _, task := range tasks {
		if task.Status == "todo" && isTaskBlocked(task, tasks) {
			hasBlocked = true
			blockers := getBlockingTasks(task, tasks)
			fmt.Printf("🚫 %s\n", task.Title)
			fmt.Printf("   Priority: %s\n", task.Priority)
			fmt.Printf("   Blocked by: %s\n", strings.Join(blockers, ", "))
			fmt.Println()
		}
	}

	if !hasBlocked {
		fmt.Println("✅ No blocked tasks! All tasks are ready to work on.")
	}
}

func showAllBlockedTasks() {
	projectsDir := filepath.Join(projectsPath, "projects")

	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading projects: %v\n", err)
		return
	}

	fmt.Println("All Blocked Tasks:")
	fmt.Println("==================")

	for _, entry := range entries {
		if entry.IsDir() {
			showBlockedTasksForProject(entry.Name())
		}
	}
}

func showDependenciesForProject(projectID string) {
	tasks, err := loadProjectTasks(projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		return
	}

	fmt.Printf("Dependency Chain for %s:\n", projectID)
	fmt.Println("===========================")

	for _, task := range tasks {
		if len(task.DependencyIDs) > 0 {
			fmt.Printf("📋 %s\n", task.Title)
			fmt.Printf("   Status: %s\n", task.Status)
			fmt.Printf("   Depends on:\n")

			for _, depID := range task.DependencyIDs {
				for _, depTask := range tasks {
					if depTask.ID == depID {
						status := "✅"
						if depTask.Status != "done" {
							status = "❌"
						}
						fmt.Printf("     %s %s (%s)\n", status, depTask.Title, depTask.Status)
					}
				}
			}
			fmt.Println()
		}
	}
}

func showAllDependencies() {
	projectsDir := filepath.Join(projectsPath, "projects")

	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading projects: %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			showDependenciesForProject(entry.Name())
		}
	}
}

func init() {
	statusBlockedCmd.Flags().StringP("project", "p", "", "Show blocked tasks for specific project")
	statusDependenciesCmd.Flags().StringP("project", "p", "", "Show dependencies for specific project")

	statusCmd.AddCommand(statusProjectCmd)
	statusCmd.AddCommand(statusBlockedCmd)
	statusCmd.AddCommand(statusDependenciesCmd)
}