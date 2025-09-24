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

// /* Initialiserer 'list' kommandoen og dens underkommandoer. */
func init() {
	listCmd.AddCommand(listProjectsCmd)
}