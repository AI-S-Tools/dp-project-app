/* ::GEMINI:06: Koden i 'loadProjectTasks' og 'loadTasksFromDir' er næsten identisk og kan refaktoreres til en enkelt, mere generisk funktion for at undgå gentagelse.:: */
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// /* Definerer 'status' kommandoen for at vise projektstatus. */
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

// /* Definerer 'project' underkommandoen for at vise en oversigt over projektstatus. */
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

// /* Definerer 'blocked' underkommandoen for at vise alle blokerede opgaver. */
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

// /* Definerer 'dependencies' underkommandoen for at vise alle afhængighedskæder. */
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

// /* Indlæser alle opgaver for et givet projekt. */
func loadProjectTasks(projectID string) ([]Task, error) {
	var tasks []Task

	// Load tasks from phases directories
	phasesDir := filepath.Join(projectsPath, "projects", projectID, "phases")
	phaseEntries, err := os.ReadDir(phasesDir)
	if err == nil {
		for _, phaseEntry := range phaseEntries {
			if phaseEntry.IsDir() {
				tasksDir := filepath.Join(phasesDir, phaseEntry.Name(), "tasks")
				if phaseTasks, err := loadTasksFromDir(tasksDir); err == nil {
					tasks = append(tasks, phaseTasks...)
				}
			}
		}
	}

	// Also check for tasks directly in project directory (for backwards compatibility)
	tasksDir := filepath.Join(projectsPath, "projects", projectID, "tasks")

	if directTasks, err := loadTasksFromDir(tasksDir); err == nil {
		tasks = append(tasks, directTasks...)
	}

	return tasks, nil
}

// /* Indlæser opgaver fra en bestemt mappe. */
func loadTasksFromDir(tasksDir string) ([]Task, error) {
	var tasks []Task

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

// /* Kontrollerer, om en opgave er blokeret af ufærdige afhængigheder. */
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

// /* Henter en liste over de opgaver, der blokerer en given opgave. */
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

// /* Viser alle blokerede opgaver for et specifikt projekt. */
func showBlockedTasksForProject(projectID string) {
	tasks, err := loadProjectTasks(projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		return
	}

	fmt.Printf("Blocked Tasks in %s:\n", projectID)
	fmt.Println("=======================")

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

// /* Viser alle blokerede opgaver på tværs af alle projekter. */
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

// /* Viser afhængighedskæderne for et specifikt projekt. */
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

// /* Viser alle afhængighedskæder på tværs af alle projekter. */
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

// /* Definerer 'active' underkommandoen for at vise aktive opgaver. */
var statusActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Show tasks that can be worked on now",
	Long: `Show Active Tasks

Display all tasks that are not blocked by dependencies and can be
worked on immediately. This shows tasks that are available for
team members to pick up and start working on.

Tasks are considered active when:
  • Status is 'todo' or 'in_progress'
  • No blocking dependencies exist
  • All dependency tasks are completed

Information Displayed:
  • Task ID and title
  • Project and phase information
  • Priority level
  • Assignee (if any)
  • Last updated timestamp

Examples:
  dppm status active                    # Show all active tasks across projects
  dppm status active --project web-app # Show active tasks for specific project

AI Usage:
  Perfect for daily standups and sprint planning to see what work
  is immediately available for team members.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")

		if projectID != "" {
			showActiveTasksForProject(projectID)
		} else {
			showAllActiveTasks()
		}
	},
}

// /* Viser aktive opgaver for et specifikt projekt. */
func showActiveTasksForProject(projectID string) {
	fmt.Printf("Active Tasks for project: %s\n", projectID)
	fmt.Println("==============================")

	projectDir := filepath.Join(projectsPath, "projects", projectID)
	if !dirExists(projectDir) {
		fmt.Printf("❌ Project '%s' not found\n", projectID)
		return
	}

	found := false
	// Check tasks in project root
	tasksDir := filepath.Join(projectDir, "tasks")
	if dirExists(tasksDir) {
		found = scanActiveTasksInDir(tasksDir, projectID, "") || found
	}

	// Check tasks in phases
	phasesDir := filepath.Join(projectDir, "phases")
	if dirExists(phasesDir) {
		phases, err := os.ReadDir(phasesDir)
		if err == nil {
			for _, phase := range phases {
				if phase.IsDir() {
					phaseTasksDir := filepath.Join(phasesDir, phase.Name(), "tasks")
					if dirExists(phaseTasksDir) {
						found = scanActiveTasksInDir(phaseTasksDir, projectID, phase.Name()) || found
					}
				}
			}
		}
	}

	if !found {
		fmt.Println("No active tasks found.")
	}
}

// /* Viser alle aktive opgaver på tværs af alle projekter. */
func showAllActiveTasks() {
	fmt.Println("All Active Tasks")
	fmt.Println("================")

	projectsDir := filepath.Join(projectsPath, "projects")
	projects, err := os.ReadDir(projectsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading projects directory: %v\n", err)
		return
	}

	found := false
	for _, project := range projects {
		if project.IsDir() {
			projectID := project.Name()
			fmt.Printf("\n📁 Project: %s\n", projectID)
			fmt.Println(strings.Repeat("-", len(projectID)+12))

			projectFound := false
			// Check tasks in project root
			tasksDir := filepath.Join(projectsDir, projectID, "tasks")
			if dirExists(tasksDir) {
				projectFound = scanActiveTasksInDir(tasksDir, projectID, "") || projectFound
			}

			// Check tasks in phases
			phasesDir := filepath.Join(projectsDir, projectID, "phases")
			if dirExists(phasesDir) {
				phases, phErr := os.ReadDir(phasesDir)
				if phErr == nil {
					for _, phase := range phases {
						if phase.IsDir() {
							phaseTasksDir := filepath.Join(phasesDir, phase.Name(), "tasks")
							if dirExists(phaseTasksDir) {
								projectFound = scanActiveTasksInDir(phaseTasksDir, projectID, phase.Name()) || projectFound
							}
						}
					}
				}
			}

			if !projectFound {
				fmt.Println("  No active tasks found")
			}
			found = found || projectFound
		}
	}

	if !found {
		fmt.Println("No active tasks found across all projects.")
	}
}

// /* Scanner aktive opgaver i en given mappe. */
func scanActiveTasksInDir(tasksDir, projectID, phaseID string) bool {
	entries, err := os.ReadDir(tasksDir)
	if err != nil {
		return false
	}

	found := false
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

			// Check if task is active (not blocked and in workable state)
			if isTaskActive(task) {
				displayActiveTask(task, phaseID)
				found = true
			}
		}
	}
	return found
}

// /* Kontrollerer om en opgave er aktiv (ikke blokeret og kan arbejdes på). */
func isTaskActive(task Task) bool {
	// Must be in a workable status
	if task.Status != "todo" && task.Status != "in_progress" {
		return false
	}

	// Must not be blocked by dependencies
	for _, depID := range task.DependencyIDs {
		if !isTaskCompleted(depID, task.ProjectID) {
			return false
		}
	}

	return true
}

// /* Kontrollerer om en opgave er færdig. */
func isTaskCompleted(taskID, projectID string) bool {
	// Check in project root tasks
	taskFile := filepath.Join(projectsPath, "projects", projectID, "tasks", taskID+".yaml")
	if fileExists(taskFile) {
		if task := loadTask(taskFile); task != nil {
			return task.Status == "done" || task.Status == "completed"
		}
	}

	// Check in phases
	phasesDir := filepath.Join(projectsPath, "projects", projectID, "phases")
	if dirExists(phasesDir) {
		phases, err := os.ReadDir(phasesDir)
		if err == nil {
			for _, phase := range phases {
				if phase.IsDir() {
					taskFile := filepath.Join(phasesDir, phase.Name(), "tasks", taskID+".yaml")
					if fileExists(taskFile) {
						if task := loadTask(taskFile); task != nil {
							return task.Status == "done" || task.Status == "completed"
						}
					}
				}
			}
		}
	}

	return false
}

// /* Indlæser en opgave fra en YAML-fil. */
func loadTask(taskFile string) *Task {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		return nil
	}

	var task Task
	if err := yaml.Unmarshal(data, &task); err != nil {
		return nil
	}
	return &task
}

// /* Viser en aktiv opgave. */
func displayActiveTask(task Task, phaseID string) {
	fmt.Printf("  ID: %s\n", task.ID)
	fmt.Printf("  Title: %s\n", task.Title)
	if phaseID != "" {
		fmt.Printf("  Phase: %s\n", phaseID)
	}
	fmt.Printf("  Status: %s\n", task.Status)
	fmt.Printf("  Priority: %s\n", task.Priority)
	if task.Assignee != "" {
		fmt.Printf("  Assignee: %s\n", task.Assignee)
	}
	fmt.Printf("  Updated: %s\n", task.Updated)
	fmt.Println("  ---")
}

// /* Kontrollerer om en mappe eksisterer. */
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// /* Initialiserer 'status' kommandoen og dens underkommandoer. */
func init() {
	statusBlockedCmd.Flags().StringP("project", "p", "", "Show blocked tasks for specific project")
	statusDependenciesCmd.Flags().StringP("project", "p", "", "Show dependencies for specific project")
	statusActiveCmd.Flags().StringP("project", "p", "", "Show active tasks for specific project")

	statusCmd.AddCommand(statusProjectCmd)
	statusCmd.AddCommand(statusBlockedCmd)
	statusCmd.AddCommand(statusDependenciesCmd)
	statusCmd.AddCommand(statusActiveCmd)
}