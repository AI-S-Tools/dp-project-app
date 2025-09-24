/* ::GEMINI:03: Global variabel 'currentProjectConfig' kan føre til race conditions, hvis applikationen bliver multi-threaded i fremtiden. Overvej at overføre konfigurationen som en parameter.:: */
/* ::GEMINI:04: Fejlhåndtering i 'main' funktionen kan forbedres ved at undgå 'os.Exit(1)' og i stedet returnere fejl, hvilket gør koden mere testbar.:: */
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var version = "dev" // Will be set during build

// /* Strukturer til lokal projektbinding. */
type ProjectBinding struct {
	ProjectID    string    `yaml:"project_id"`
	DropboxPath  string    `yaml:"dropbox_path"`
	Created      time.Time `yaml:"created"`
	LastSync     time.Time `yaml:"last_sync"`
}

type LocalProjectConfig struct {
	Binding ProjectBinding `yaml:"binding"`
	Project struct {
		ID     string `yaml:"id"`
		Name   string `yaml:"name"`
		Status string `yaml:"status"`
		Owner  string `yaml:"owner"`
	} `yaml:"project"`
	Context struct {
		CurrentPhase     string `yaml:"current_phase,omitempty"`
		InProgressTask   string `yaml:"in_progress_task,omitempty"`
		LastTaskCreated  string `yaml:"last_task_created,omitempty"`
	} `yaml:"context"`
}

// Global variabel til at holde den aktuelle projektkonfiguration
var currentProjectConfig *LocalProjectConfig

// /* Definerer rodkommandoen for DPPM CLI-værktøjet. */
var rootCmd = &cobra.Command{
	Use:   "dppm",
	Short: "Dropbox Project Manager (DPPM)",
	Long: `DPPM - Dropbox Project Manager

A comprehensive CLI tool for managing projects, phases, and tasks using
Dropbox as the storage backend. Perfect for AI-driven development workflows.

Features:
  • YAML-based project, phase, and task management
  • Hierarchical project organization with phase folders
  • Cross-platform synchronization via Dropbox
  • AI-friendly verbose output and documentation
  • Built-in knowledge base and examples (dppm wiki)
  • Comprehensive dependency management
  • Template-based project creation
  • AI collaboration system with DSL markers

Storage Location: ~/Dropbox/project-management/

🚀 Quick Start Guide:
  dppm init my-project                # Complete project initialization
  dppm wiki                           # Show knowledge base
  dppm --wiki "create project"        # Search for help
  dppm project create my-project      # Create new project
  dppm phase create setup --project my-project
  dppm task create init --project my-project --phase setup

📖 Getting Help:
  dppm wiki list                      # All available topics
  dppm wiki complete                  # Complete workflow example
  dppm --help                         # Command reference

Examples:
  dppm init web-app --doc "./requirements.md" # Complete project setup
  dppm project create web-app --name "Web Application" --owner "dev-team"
  dppm phase create backend --project web-app --name "Backend Development"
  dppm task create auth --project web-app --phase backend --title "Authentication"
  dppm status project web-app
  dppm list projects
  dppm collab find docs/                # Find AI collaboration tasks
  dppm collab wiki "task handoff"       # Learn collaboration patterns

🤖 AI Usage:
DPPM is designed for AI-driven workflows. Use the wiki system for comprehensive
guidance on all features and best practices.

For detailed command help, use: dppm [command] --help`,
	Run: func(cmd *cobra.Command, args []string) {
		showStartupGuide()
	},
}

var projectsPath string

// /* Initialiserer stier og kommandoer. */
func init() {
	home, _ := os.UserHomeDir()
	projectsPath = filepath.Join(home, "Dropbox", "project-management")

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(bindCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(phaseCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(wikiCmd)
	rootCmd.AddCommand(collabCmd)

	// Add --wiki flag for direct search
	rootCmd.Flags().String("wiki", "", "Search DPPM knowledge base (e.g. --wiki \"create task\")")

	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")
}

// /* Finder projektbindingen ved at søge efter .dppm/project.yaml. */
func findProjectBinding() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Search upwards from current directory
	dir := currentDir
	for {
		dppmPath := filepath.Join(dir, ".dppm", "project.yaml")
		if _, err := os.Stat(dppmPath); err == nil {
			return dppmPath, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			// Reached root directory
			break
		}
		dir = parentDir
	}

	return "", fmt.Errorf("no .dppm/project.yaml found in current or parent directories")
}

// /* Indlæser den lokale projektkonfiguration fra en fil. */
func loadProjectConfig(configPath string) (*LocalProjectConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read project config: %v", err)
	}

	var config LocalProjectConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse project config: %v", err)
	}

	return &config, nil
}

// /* Gemmer den lokale projektkonfiguration til en fil. */
func saveProjectConfig(configPath string, config *LocalProjectConfig) error {
	config.Binding.LastSync = time.Now()

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal project config: %v", err)
	}

	// Ensure .dppm directory exists
	if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
		return fmt.Errorf("failed to create .dppm directory: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to save project config: %v", err)
	}

	return nil
}

// /* Returnerer det aktuelle projekt-ID fra den lokale binding. */
func getCurrentProjectID() string {
	if currentProjectConfig == nil {
		return ""
	}
	return currentProjectConfig.Binding.ProjectID
}

// /* Opretter en ny .dppm/project.yaml-fil i den aktuelle mappe. */
func createProjectBinding(projectID string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	dppmDir := filepath.Join(currentDir, ".dppm")
	configPath := filepath.Join(dppmDir, "project.yaml")

	// Check if .dppm already exists
	if _, err := os.Stat(configPath); err == nil {
		return fmt.Errorf(".dppm/project.yaml already exists in this directory")
	}

	// Create the config
	config := &LocalProjectConfig{
		Binding: ProjectBinding{
			ProjectID:   projectID,
			DropboxPath: filepath.Join(projectsPath, "projects", projectID),
			Created:     time.Now(),
			LastSync:    time.Now(),
		},
		Project: struct {
			ID     string `yaml:"id"`
			Name   string `yaml:"name"`
			Status string `yaml:"status"`
			Owner  string `yaml:"owner"`
		}{
			ID:     projectID,
			Name:   "", // Will be filled from Dropbox project
			Status: "active",
			Owner:  "",
		},
	}

	if err := saveProjectConfig(configPath, config); err != nil {
		return fmt.Errorf("failed to create project binding: %v", err)
	}

	fmt.Printf("✅ Created .dppm/project.yaml binding to project: %s\n", projectID)
	fmt.Printf("📁 Directory: %s\n", currentDir)
	fmt.Printf("🔗 Dropbox Path: %s\n", config.Binding.DropboxPath)

	return nil
}

// /* Definerer 'bind' kommandoen for at binde den aktuelle mappe til et projekt. */
var bindCmd = &cobra.Command{
	Use:   "bind [project-id]",
	Short: "Bind current directory to an existing DPPM project",
	Long: `Bind Current Directory to DPPM Project

Creates a local .dppm/project.yaml file that binds the current directory
to an existing DPPM project in your Dropbox. This enables automatic
project scoping for all DPPM commands.

Arguments:
  project-id    ID of existing DPPM project to bind to

Examples:
  dppm bind dp-project-app     # Bind to existing project
  dppm bind dash-terminal      # Bind to DASH Terminal project

After binding, all DPPM commands will automatically use this project:
  dppm task create feature     # Auto: --project dp-project-app
  dppm phase create testing    # Auto: --project dp-project-app`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

		// Verify project exists in Dropbox
		projectPath := filepath.Join(projectsPath, "projects", projectID, "project.yaml")
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "❌ Error: Project '%s' not found in Dropbox.\n", projectID)
			fmt.Fprintf(os.Stderr, "Available projects:\n")
			// TODO: List available projects
			os.Exit(1)
		}

		// Create binding
		if err := createProjectBinding(projectID); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n🎯 Project binding complete!")
		fmt.Println("All DPPM commands in this directory will now use this project.")
	},
}

// /* Kræver en projektbinding, ellers afsluttes programmet med en fejl. */
func requireProjectBinding() {
	if currentProjectConfig == nil {
		fmt.Fprintf(os.Stderr, "❌ Error: No DPPM project found in current directory.\n\n")
		fmt.Fprintf(os.Stderr, "To initialize a new project:\n")
		fmt.Fprintf(os.Stderr, "  dppm init project-name --name \"Project Name\"\n\n")
		fmt.Fprintf(os.Stderr, "To bind to an existing project:\n")
		fmt.Fprintf(os.Stderr, "  dppm bind existing-project-id\n\n")
		fmt.Fprintf(os.Stderr, "This prevents accidental cross-project task creation.\n")
		os.Exit(1)
	}
}

// /* Hovedfunktionen for DPPM-applikationen. */
func main() {
	// Check for version flag first
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-v" {
			fmt.Printf("DPPM (Dropbox Project Manager) %s\n", version)
			fmt.Println("AI-first CLI tool for project, phase, and task management")
			fmt.Println("Repository: https://github.com/AI-S-Tools/dp-project-app")
			return
		}
	}

	// Check for --wiki flag in args before executing
	for i, arg := range os.Args {
		if arg == "--wiki" && i+1 < len(os.Args) {
			// Execute wiki search directly
			wikiQuery := os.Args[i+1]
			wikiCmd.Run(wikiCmd, []string{wikiQuery})
			return
		}
	}

	// Special commands that don't require project binding
	if len(os.Args) > 1 {
		command := os.Args[1]
		if command == "init" || command == "bind" || command == "wiki" || command == "--help" || command == "help" {
			if err := rootCmd.Execute(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			return
		}
	}

	// Try to find and load project binding
	configPath, err := findProjectBinding()
	if err != nil {
		// No project binding found - show error and suggest init/bind
		if len(os.Args) == 1 {
			// No arguments - show startup guide which now suggests init
			showStartupGuide()
			return
		} else {
			// Command was provided but no project binding
			fmt.Fprintf(os.Stderr, "❌ Error: No DPPM project found in current directory.\n\n")
			fmt.Fprintf(os.Stderr, "To initialize a new project:\n")
			fmt.Fprintf(os.Stderr, "  dppm init project-name --name \"Project Name\"\n\n")
			fmt.Fprintf(os.Stderr, "To bind to an existing project:\n")
			fmt.Fprintf(os.Stderr, "  dppm bind existing-project-id\n\n")
			fmt.Fprintf(os.Stderr, "This prevents accidental cross-project task creation.\n")
			os.Exit(1)
		}
	}

	// Load project configuration
	currentProjectConfig, err = loadProjectConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error loading project config: %v\n", err)
		os.Exit(1)
	}

	// Show current project context
	if len(os.Args) == 1 {
		// No arguments - show project-aware status
		fmt.Printf("🎯 Current DPPM Project: %s\n", currentProjectConfig.Project.Name)
		fmt.Printf("📁 Project ID: %s\n", currentProjectConfig.Binding.ProjectID)
		fmt.Printf("📊 Directory: %s\n\n", filepath.Dir(filepath.Dir(configPath)))

		// Show project status
		if statusCmd.Run != nil {
			statusCmd.Run(statusCmd, []string{"project", currentProjectConfig.Binding.ProjectID})
		}
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// /* Viser en opstartsguide, når der ikke er angivet nogen kommandoer. */
func showStartupGuide() {
	fmt.Println(`DPPM - Dropbox Project Manager
==============================

🎯 You're running DPPM without any commands. Here's what you can do:

📚 GET HELP & LEARN:
  dppm wiki                    # Browse knowledge base
  dppm --wiki "create project" # Search for specific help
  dppm wiki list               # See all available topics
  dppm wiki complete           # Complete workflow example

📋 QUICK ACTIONS:
  dppm list projects           # See existing projects
  dppm status                  # Check overall status

🚀 CREATE NEW PROJECT (Two Options):

Option 1 - Complete Initialization (Recommended):
  dppm init my-project         # Interactive project setup wizard

Option 2 - Manual Creation:
  dppm project create my-project --name "My Project" --owner "your-name"
  dppm phase create phase-1 --project my-project --name "First Phase"
  dppm task create first-task --project my-project --phase phase-1

💡 COMMON WORKFLOWS:
  • New to DPPM? → dppm wiki "complete"
  • Creating tasks? → dppm --wiki "create task"
  • Managing dependencies? → dppm --wiki "dependencies"
  • Checking progress? → dppm status project PROJECT_NAME
  • AI collaboration? → dppm collab wiki

🤖 AI TIP:
DPPM is designed for AI workflows. The wiki system contains comprehensive
guides for every feature. Use it to get detailed, actionable information.

Try: dppm wiki "project workflow" to see a complete example!`)
}