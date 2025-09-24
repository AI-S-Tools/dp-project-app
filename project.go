package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// /* Definerer strukturen for et projekt. */
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

// /* Definerer 'project' kommandoen til projektstyring. */
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
  list      List all projects (use 'dppm list projects' instead)
  show      Display detailed project information
  update    Update project metadata
  delete    Remove a project and all its data

Examples:
  dppm project create web-app --name "Web Application" --owner "dev-team"
  dppm project show web-app
  dppm project update web-app --status completed

For more information about a specific command, use:
  dppm project [command] --help`,
}

// /* Definerer 'create' underkommandoen for at oprette et nyt projekt. */
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

		// Validate project ID
		if err := validateProjectID(projectID); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error: %v\n", err)
			return
		}

		// Check for duplicate project
		if projectExists(projectID) {
			fmt.Fprintf(os.Stderr, "❌ Error: Project '%s' already exists\n", projectID)
			fmt.Fprintf(os.Stderr, "Use a different project ID or run 'dppm status project %s' to see existing project\n", projectID)
			return
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

// /* Validerer projekt ID. */
func validateProjectID(projectID string) error {
	// Check for empty ID
	if strings.TrimSpace(projectID) == "" {
		return fmt.Errorf("Project ID cannot be empty")
	}

	// Check for valid characters (lowercase, alphanumeric, hyphens, underscores)
	validID := regexp.MustCompile(`^[a-z0-9][a-z0-9_-]*$`)
	if !validID.MatchString(projectID) {
		return fmt.Errorf("Invalid project ID '%s'. Project IDs must:\n  • Start with a lowercase letter or number\n  • Contain only lowercase letters, numbers, hyphens, and underscores\n  • Example: 'web-app', 'api_server', 'mobile2'", projectID)
	}

	// Check length limits
	if len(projectID) > 50 {
		return fmt.Errorf("Project ID '%s' is too long (max 50 characters)", projectID)
	}

	// Check for reserved names
	reserved := []string{"help", "version", "list", "status", "wiki", "collab", "bind", "init"}
	for _, reservedName := range reserved {
		if projectID == reservedName {
			return fmt.Errorf("Project ID '%s' is reserved. Please choose a different ID", projectID)
		}
	}

	return nil
}

// /* Tjekker om et projekt eksisterer. */
func projectExists(projectID string) bool {
	projectPath := filepath.Join(projectsPath, "projects", projectID, "project.yaml")
	_, err := os.Stat(projectPath)
	return !os.IsNotExist(err)
}

// /* Initialiserer 'project' kommandoen og dens underkommandoer. */
func init() {
	createProjectCmd.Flags().StringP("name", "n", "", "Project name")
	createProjectCmd.Flags().StringP("description", "d", "", "Project description")
	createProjectCmd.Flags().StringP("owner", "o", "", "Project owner")

	projectCmd.AddCommand(createProjectCmd)
}