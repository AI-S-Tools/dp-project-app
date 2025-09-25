package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var version = "dev" // Will be set during build

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
  dppm wiki "complete"                # Complete workflow example
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

func init() {
	home, _ := os.UserHomeDir()
	projectsPath = filepath.Join(home, "Dropbox", "project-management")

	rootCmd.AddCommand(initCmd)
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

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

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

// LocalProjectBinding represents the local project context
type LocalProjectBinding struct {
	ProjectID   string `yaml:"project_id"`
	ProjectName string `yaml:"project_name,omitempty"`
	DropboxPath string `yaml:"dropbox_path,omitempty"`
	Created     string `yaml:"created,omitempty"`
}

// getLocalProjectContext reads the local .dppm/project.yaml file if it exists
func getLocalProjectContext() (*LocalProjectBinding, error) {
	bindingFile := ".dppm/project.yaml"
	if _, err := os.Stat(bindingFile); os.IsNotExist(err) {
		return nil, nil // No local binding exists
	}

	data, err := os.ReadFile(bindingFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read binding file: %v", err)
	}

	// Try to parse as LocalProjectBinding first
	var binding LocalProjectBinding
	if err := yaml.Unmarshal(data, &binding); err != nil {
		// If that fails, try to extract project ID from regular project metadata
		var projectData map[string]interface{}
		if err := yaml.Unmarshal(data, &projectData); err != nil {
			return nil, fmt.Errorf("failed to parse binding file: %v", err)
		}

		// Extract project ID from the YAML structure
		if id, ok := projectData["id"].(string); ok {
			binding.ProjectID = id
		}
		if name, ok := projectData["name"].(string); ok {
			binding.ProjectName = name
		}
	}

	if binding.ProjectID == "" {
		return nil, fmt.Errorf("no project ID found in binding file")
	}

	return &binding, nil
}

// setDefaultProjectFlag sets the project flag to the local context if not explicitly provided
func setDefaultProjectFlag(cmd *cobra.Command, flagName string) {
	if cmd.Flags().Changed(flagName) {
		return // Flag was explicitly set by user
	}

	context, err := getLocalProjectContext()
	if err != nil || context == nil {
		return // No local context available
	}

	// Set the project flag to the local context value
	cmd.Flags().Set(flagName, context.ProjectID)
}
