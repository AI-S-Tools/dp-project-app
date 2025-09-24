package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// /* Definerer 'list' kommandoen for at liste projekter, faser eller opgaver. */
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects, phases, or tasks",
	Long: `List Command

Display various entities in your Dropbox project management system.
Provides organized views of projects, phases, and tasks with detailed
information formatted for both human reading and AI processing.

Available Subcommands:
  projects    List all projects with status and metadata
  phases      List phases for a specific project
  tasks       List tasks for a project or phase

Output Format:
  Each listing provides comprehensive information including:
  - Unique identifiers
  - Names and descriptions
  - Status information
  - Ownership details
  - Creation and update timestamps
  - Hierarchical relationships

Examples:
  dppm list projects                    # Show all projects
  dppm list phases --project web-app   # Show phases for web-app project
  dppm list tasks --project web-app    # Show all tasks in web-app project
  dppm list tasks --phase phase-1      # Show tasks in specific phase

AI Usage:
  This command is designed to provide verbose, structured output that
  AI systems can easily parse and understand for project analysis,
  status reporting, and workflow automation.`,
}

// /* Definerer 'projects' underkommandoen for at liste alle projekter. */
var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List all projects",
	Long: `List All Projects

Display a comprehensive list of all projects in the Dropbox project
management system. Each project entry includes detailed metadata
to help with project overview and status tracking.

Information Displayed:
  • Project ID (unique identifier)
  • Project Name (human-readable name)
  • Status (active, completed, paused, cancelled)
  • Owner (responsible person or team)
  • Creation Date (when project was created)
  • Last Updated (most recent modification)
  • Description (if available)
  • Current Phase (if any)
  • Total Phases (count)

Output Format:
  Projects are displayed in a structured format with clear separators
  for easy reading by both humans and AI systems. Each project is
  separated by a divider line for visual clarity.

Examples:
  dppm list projects              # Show all projects
  dppm list projects | grep web   # Filter projects containing 'web'

AI Integration:
  The output is designed to be easily parseable by AI systems for:
  - Project status analysis
  - Resource allocation assessment
  - Progress tracking
  - Automated reporting`,
	Run: func(cmd *cobra.Command, args []string) {
		projectsDir := filepath.Join(projectsPath, "projects")

		entries, err := os.ReadDir(projectsDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading projects directory: %v\n", err)
			return
		}

		fmt.Println("Projects:")
		fmt.Println("=========")

		for _, entry := range entries {
			if entry.IsDir() {
				projectFile := filepath.Join(projectsDir, entry.Name(), "project.yaml")

				data, err := os.ReadFile(projectFile)
				if err != nil {
					continue
				}

				var project Project
				if err := yaml.Unmarshal(data, &project); err != nil {
					continue
				}

				fmt.Printf("ID: %s\n", project.ID)
				fmt.Printf("Name: %s\n", project.Name)
				fmt.Printf("Status: %s\n", project.Status)
				fmt.Printf("Owner: %s\n", project.Owner)
				fmt.Printf("Updated: %s\n", project.Updated)
				fmt.Println("---")
			}
		}
	},
}

// /* Definerer 'phases' underkommandoen for at liste faser for et projekt. */
var listPhasesCmd = &cobra.Command{
	Use:   "phases",
	Short: "List phases for a project",
	Long: `List Project Phases

Display all phases (sprints) for a specific project. Shows detailed
information about each phase including status, goals, and timeline.

Information Displayed:
  • Phase ID (unique identifier)
  • Phase Name (human-readable name)
  • Status (not_started, in_progress, completed, cancelled)
  • Goal (phase objective)
  • Start Date (if set)
  • Due Date (if set)
  • Task Count (number of tasks in phase)

Output Format:
  Phases are displayed in chronological order with clear separators
  for easy reading by both humans and AI systems.

Examples:
  dppm list phases --project web-app       # List phases for web-app project
  dppm list phases --project api-server    # List phases for api-server project

Flags:
  --project    Project ID to list phases for (required unless bound)`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")

		if projectID == "" {
			fmt.Fprintf(os.Stderr, "❌ Error: No project specified. Use --project flag or bind to a project.\n")
			return
		}

		projectDir := filepath.Join(projectsPath, "projects", projectID, "phases")

		entries, err := os.ReadDir(projectDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error reading phases directory for project %s: %v\n", projectID, err)
			return
		}

		fmt.Printf("Phases for project: %s\n", projectID)
		fmt.Println("========================")

		if len(entries) == 0 {
			fmt.Println("No phases found for this project.")
			return
		}

		for _, entry := range entries {
			if entry.IsDir() {
				phaseFile := filepath.Join(projectDir, entry.Name(), "phase.yaml")

				data, err := os.ReadFile(phaseFile)
				if err != nil {
					continue
				}

				var phase Phase
				if err := yaml.Unmarshal(data, &phase); err != nil {
					continue
				}

				fmt.Printf("ID: %s\n", phase.ID)
				fmt.Printf("Name: %s\n", phase.Name)
				fmt.Printf("Status: %s\n", phase.Status)
				fmt.Printf("Goal: %s\n", phase.Goal)
				if phase.StartDate != "" {
					fmt.Printf("Start Date: %s\n", phase.StartDate)
				}
				if phase.EndDate != "" {
					fmt.Printf("End Date: %s\n", phase.EndDate)
				}
				fmt.Printf("Updated: %s\n", phase.Updated)
				fmt.Println("---")
			}
		}
	},
}

// /* Definerer 'active' underkommandoen for at liste aktive opgaver. */
var listActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "List active tasks that can be worked on",
	Long: `List Active Tasks

Display all active tasks across projects that are not blocked by
dependencies and can be worked on immediately. This provides a
quick overview of available work items.

Active tasks are those that:
  • Have status 'todo' or 'in_progress'
  • Are not blocked by incomplete dependencies
  • Are ready for immediate work

Information Displayed:
  • Task ID and title
  • Project and phase information
  • Current status and priority
  • Assignee information (if any)
  • Last updated timestamp

Examples:
  dppm list active                    # List all active tasks across projects
  dppm list active --project web-app # List active tasks for specific project

AI Usage:
  Ideal for quick status checks and identifying immediately
  actionable work items across your project portfolio.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectID, _ := cmd.Flags().GetString("project")

		if projectID != "" {
			showActiveTasksForProject(projectID)
		} else {
			showAllActiveTasks()
		}
	},
}

// /* Initialiserer 'list' kommandoen og dens underkommandoer. */
func init() {
	listPhasesCmd.Flags().StringP("project", "p", "", "Project ID")
	listActiveCmd.Flags().StringP("project", "p", "", "Project ID")

	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listPhasesCmd)
	listCmd.AddCommand(listActiveCmd)
}