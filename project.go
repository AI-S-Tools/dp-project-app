package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Project struct {
	ID           string                 `yaml:"id"`
	Name         string                 `yaml:"name"`
	Description  string                 `yaml:"description"`
	Status       string                 `yaml:"status"`
	Owner        string                 `yaml:"owner"`
	Created      string                 `yaml:"created"`
	Updated      string                 `yaml:"updated"`
	Repository   string                 `yaml:"repository,omitempty"`
	Tags         []string               `yaml:"tags,omitempty"`
	Metadata     map[string]interface{} `yaml:"metadata,omitempty"`
	Notes        string                 `yaml:"notes,omitempty"`
	CurrentPhase string                 `yaml:"current_phase,omitempty"`
	Phases       []string               `yaml:"phases,omitempty"`
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project management commands",
	Long: `Project Management Commands

Manage projects in your Dropbox project management system. Projects are
the top-level organizational unit that contain phases and tasks.

Project Structure:
  ~/Dropbox/project-management/projects/PROJECT_ID/
  ├── project.yaml           # Project metadata and configuration
  └── phases/              # Directory containing all phases
      └── PHASE_ID/        # Individual phase directories

Available Commands:
  create    Create a new project with specified parameters
  update    Update project metadata and properties

Examples:
  dppm project create web-app --name "Web Application" --owner "dev-team"
  dppm project show web-app
  dppm project update web-app --status completed

For more information about a specific command, use:
  dppm project [command] --help`,
}

var createProjectCmd = &cobra.Command{
	Use:   "create [project-id]",
	Short: "Create a new project",
	Long: `Create a New Project

Creates a new project in the Dropbox project management system with the
specified project ID and metadata. This will create the directory structure
and initial project.yaml file.

Arguments:
  project-id    Unique identifier for the project (required)
                Must be lowercase, alphanumeric with hyphens allowed
                Examples: web-app, mobile-project, ai-tool

Directory Structure Created:
  ~/Dropbox/project-management/projects/PROJECT_ID/
  ├── project.yaml          # Project metadata file
  └── phases/             # Empty directory for future phases

Project Status Values:
  active      Project is currently being worked on (default)
  completed   Project has been finished
  paused      Project is temporarily stopped
  cancelled   Project has been abandoned

Examples:
  dppm project create web-app --name "Web Application"
  dppm project create ai-tool --name "AI Development Tool" --owner "john-doe" --description "Advanced AI automation tool"
  dppm project create mobile --name "Mobile App" --owner "dev-team" --description "Cross-platform mobile application"

AI Usage Tips:
  - Use descriptive project names for better organization
  - Include clear descriptions for AI context understanding
  - Set appropriate owners for team collaboration`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		owner, _ := cmd.Flags().GetString("owner")
		template, _ := cmd.Flags().GetString("template")

		// Validate project ID for security
		if err := ValidateProjectID(projectID); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Check if project already exists
		exists, err := CheckProjectExists(projectID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error checking project existence: %v\n", err)
			os.Exit(1)
		}
		if exists {
			fmt.Fprintf(os.Stderr, "Error: Project '%s' already exists\n", projectID)
			os.Exit(1)
		}

		// Validate description if provided
		if description != "" {
			if err := ValidateDescription(description); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}

		if name == "" {
			name = projectID
		}

		// Warn about missing description
		if description == "" {
			fmt.Println("⚠️  Warning: Project created without description")
			fmt.Println("💡 Consider adding --description for better project context and AI collaboration")
			fmt.Println("   Example: --description \"Detailed explanation of project goals and scope\"")
			fmt.Println()
		}

		// Handle template if specified
		if template != "" {
			fmt.Printf("📋 Using template: %s\n", template)
			// TODO: Implement template loading and application
			// For now, just acknowledge the template parameter
		}

		project := Project{
			ID:          projectID,
			Name:        name,
			Description: description,
			Status:      "active",
			Owner:       owner,
			Created:     time.Now().Format("2006-01-02"),
			Updated:     time.Now().Format("2006-01-02"),
			Tags:        []string{},
			Phases:      []string{},
		}

		projectDir := filepath.Join(projectsPath, "projects", projectID)
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project directory: %v\n", err)
			return
		}

		phasesDir := filepath.Join(projectDir, "phases")
		if err := os.MkdirAll(phasesDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating phases directory: %v\n", err)
			return
		}

		projectFile := filepath.Join(projectDir, "project.yaml")
		data, err := yaml.Marshal(project)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling project: %v\n", err)
			return
		}

		if err := os.WriteFile(projectFile, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing project file: %v\n", err)
			return
		}

		fmt.Printf("Project '%s' created successfully\n", projectID)
	},
}

var updateProjectCmd = &cobra.Command{
	Use:   "update [project-id]",
	Short: "Update an existing project",
	Long: `Update Project Information

Updates metadata for an existing project. You can update the name, description,
owner, status, and other project properties.

Arguments:
  project-id    ID of the project to update (required)

Available Status Values:
  active      Project is currently being worked on
  completed   Project has been finished
  paused      Project is temporarily stopped
  archived    Project is archived but preserved
  cancelled   Project has been abandoned

Examples:
  dppm project update web-app --description "Updated project description"
  dppm project update ai-tool --owner "new-team" --status completed
  dppm project update mobile --name "Mobile App v2.0"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

		// Validate project ID
		if err := ValidateProjectID(projectID); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Check if project exists
		projectPath := filepath.Join(projectsPath, "projects", projectID, "project.yaml")
		data, err := os.ReadFile(projectPath)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "Error: Project '%s' not found\n", projectID)
			} else {
				fmt.Fprintf(os.Stderr, "Error reading project: %v\n", err)
			}
			os.Exit(1)
		}

		var project Project
		if err := yaml.Unmarshal(data, &project); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing project file: %v\n", err)
			os.Exit(1)
		}

		// Get flags and update if provided
		updated := false

		if name, _ := cmd.Flags().GetString("name"); name != "" {
			project.Name = name
			updated = true
		}

		if description, _ := cmd.Flags().GetString("description"); description != "" {
			if err := ValidateDescription(description); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			project.Description = description
			updated = true
		}

		if owner, _ := cmd.Flags().GetString("owner"); owner != "" {
			project.Owner = owner
			updated = true
		}

		if status, _ := cmd.Flags().GetString("status"); status != "" {
			validStatuses := []string{"active", "completed", "paused", "archived", "cancelled"}
			isValid := false
			for _, valid := range validStatuses {
				if status == valid {
					isValid = true
					break
				}
			}
			if !isValid {
				fmt.Fprintf(os.Stderr, "Error: Invalid status '%s'. Must be one of: %v\n", status, validStatuses)
				os.Exit(1)
			}
			project.Status = status
			updated = true
		}

		if !updated {
			fmt.Println("No updates specified. Use --help to see available flags.")
			return
		}

		// Update timestamp
		project.Updated = time.Now().Format("2006-01-02")

		// Write updated project back
		updatedData, err := yaml.Marshal(project)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling updated project: %v\n", err)
			os.Exit(1)
		}

		if err := os.WriteFile(projectPath, updatedData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing updated project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Project '%s' updated successfully\n", projectID)
	},
}

func init() {
	createProjectCmd.Flags().StringP("name", "n", "", "Project name")
	createProjectCmd.Flags().StringP("description", "d", "", "Project description")
	createProjectCmd.Flags().StringP("owner", "o", "", "Project owner")
	createProjectCmd.Flags().StringP("template", "t", "", "Project template (yaml, web-app, api, or path to custom template)")

	updateProjectCmd.Flags().StringP("name", "n", "", "New project name")
	updateProjectCmd.Flags().StringP("description", "d", "", "New project description")
	updateProjectCmd.Flags().StringP("owner", "o", "", "New project owner")
	updateProjectCmd.Flags().StringP("status", "s", "", "New project status (active, completed, paused, archived, cancelled)")

	projectCmd.AddCommand(createProjectCmd)
	projectCmd.AddCommand(updateProjectCmd)
}
