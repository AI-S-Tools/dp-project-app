package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Phase struct {
	ID        string       `yaml:"id"`
	Name      string       `yaml:"name"`
	ProjectID string       `yaml:"project_id"`
	Status    string       `yaml:"status"`
	StartDate string       `yaml:"start_date,omitempty"`
	EndDate   string       `yaml:"end_date,omitempty"`
	Created   string       `yaml:"created"`
	Updated   string       `yaml:"updated"`
	Goal      string       `yaml:"goal,omitempty"`
	Capacity  int          `yaml:"capacity,omitempty"`
	Tasks     []string     `yaml:"tasks,omitempty"`
	Metrics   PhaseMetrics `yaml:"metrics,omitempty"`
	Notes     string       `yaml:"notes,omitempty"`
}

type PhaseMetrics struct {
	CompletedTasks       int `yaml:"completed_tasks"`
	TotalTasks           int `yaml:"total_tasks"`
	StoryPointsCompleted int `yaml:"story_points_completed"`
	StoryPointsTotal     int `yaml:"story_points_total"`
}

var phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Phase management commands",
	Long: `Phase Management Commands

Manage project phases (development phases/sprints). Phases are time-boxed
periods where specific sets of tasks are completed toward project goals.

Project Structure with Phases:
  ~/Dropbox/project-management/projects/PROJECT_ID/
  ├── project.yaml           # Project metadata
  └── phases/                # All project phases
      ├── phase-1/           # Individual phase directory
      │   ├── phase.yaml     # Phase metadata
      │   └── tasks/         # Tasks for this phase
      └── phase-2/
          ├── phase.yaml
          └── tasks/

Available Commands:
  create    Create a new phase
  list      List all phases in a project
  show      Display detailed phase information
  update    Update phase metadata

Examples:
  dppm phase create phase-3 --project dash-lxd --name "File Integration Phase"
  dppm phase list --project dash-lxd
  dppm phase show phase-3 --project dash-lxd

For more information about a specific command, use:
  dppm phase [command] --help`,
}

var createPhaseCmd = &cobra.Command{
	Use:   "create [phase-id]",
	Short: "Create a new phase",
	Long: `Create a New Phase

Creates a new phase within a project. Phases are development periods
where specific tasks are grouped and completed together.

Arguments:
  phase-id    Unique identifier for the phase (required)
              Must be lowercase, alphanumeric with hyphens allowed
              Examples: phase-1, ui-enhancement, backend-v2

Phase Status Values:
  planning    Phase is being planned (default)
  active      Phase is currently being worked on
  completed   Phase has been finished
  cancelled   Phase has been abandoned

Examples:
  dppm phase create phase-3 --project dash-lxd --name "File Integration"
  dppm phase create backend-v2 --project web-app --name "Backend Version 2" --goal "Complete API redesign"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		phaseID := args[0]
		name, _ := cmd.Flags().GetString("name")
		projectID, _ := cmd.Flags().GetString("project")
		goal, _ := cmd.Flags().GetString("goal")
		startDate, _ := cmd.Flags().GetString("start-date")
		endDate, _ := cmd.Flags().GetString("end-date")

		// Validate phase ID for security
		if err := ValidatePhaseID(phaseID); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Validate project ID
		if err := ValidateProjectID(projectID); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid project ID: %v\n", err)
			os.Exit(1)
		}

		if name == "" {
			name = phaseID
		}

		// Warn about missing goal
		if goal == "" {
			fmt.Println("⚠️  Warning: Phase created without goal")
			fmt.Println("💡 Consider adding --goal for better phase context and planning")
			fmt.Println("   Example: --goal \"Deliver authentication system with complete user management\"")
			fmt.Println()
		}

		phase := Phase{
			ID:        phaseID,
			Name:      name,
			ProjectID: projectID,
			Status:    "planning",
			StartDate: startDate,
			EndDate:   endDate,
			Created:   time.Now().Format("2006-01-02"),
			Updated:   time.Now().Format("2006-01-02"),
			Goal:      goal,
			Tasks:     []string{},
		}

		// Create phase directory structure
		phaseDir := filepath.Join(projectsPath, "projects", projectID, "phases", phaseID)
		if err := os.MkdirAll(phaseDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating phase directory: %v\n", err)
			return
		}

		tasksDir := filepath.Join(phaseDir, "tasks")
		if err := os.MkdirAll(tasksDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating tasks directory: %v\n", err)
			return
		}

		phaseFile := filepath.Join(phaseDir, "phase.yaml")
		data, err := yaml.Marshal(phase)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling phase: %v\n", err)
			return
		}

		if err := os.WriteFile(phaseFile, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing phase file: %v\n", err)
			return
		}

		fmt.Printf("Phase '%s' created successfully in project '%s'\n", phaseID, projectID)
		fmt.Printf("Phase directory: %s\n", phaseDir)
	},
}

func init() {
	createPhaseCmd.Flags().StringP("name", "n", "", "Phase name")
	createPhaseCmd.Flags().StringP("project", "p", "", "Project ID (required)")
	createPhaseCmd.Flags().StringP("goal", "g", "", "Phase goal description")
	createPhaseCmd.Flags().String("start-date", "", "Phase start date (YYYY-MM-DD)")
	createPhaseCmd.Flags().String("end-date", "", "Phase end date (YYYY-MM-DD)")

	createPhaseCmd.MarkFlagRequired("project")

	phaseCmd.AddCommand(createPhaseCmd)
}
