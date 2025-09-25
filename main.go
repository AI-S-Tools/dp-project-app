package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

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

For detailed command help, use: dppm [command] --help

🐛 Found a bug or missing feature?
Report issues: https://github.com/AI-S-Tools/dp-project-app/issues`,
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
	rootCmd.AddCommand(bindCmd)

	// Add --wiki flag for direct search
	rootCmd.Flags().String("wiki", "", "Search DPPM knowledge base (e.g. --wiki \"create task\")")

	// Add --link flag for project documentation linking
	rootCmd.Flags().String("link", "", "Create symlink to project docs folder (e.g. --link \"project-name\")")

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

	// Check for --link flag in args before executing
	for i, arg := range os.Args {
		if arg == "--link" && i+1 < len(os.Args) {
			// Execute project docs linking directly
			projectName := os.Args[i+1]
			createProjectDocsLink(projectName)
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
==============================`)

	// Check for local project binding and display context
	if context, err := getLocalProjectContext(); err == nil && context != nil {
		fmt.Println()
		if context.ProjectName != "" {
			fmt.Printf("🎯 Current DPPM Project: %s\n", context.ProjectName)
		} else {
			fmt.Printf("🎯 Current DPPM Project: %s\n", context.ProjectID)
		}
		fmt.Printf("📁 Project ID: %s\n", context.ProjectID)
		fmt.Printf("🔗 Dropbox Path: %s\n", context.DropboxPath)
		fmt.Println()
		fmt.Println("💡 Project-scoped commands will use this project automatically.")
		fmt.Println("   Use 'rm -rf .dppm' to unbind if needed.")
		fmt.Println()
	}

	fmt.Println(`
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

var bindCmd = &cobra.Command{
	Use:   "bind [project-id]",
	Short: "Bind local directory to a DPPM project",
	Long: `Bind Local Directory to DPPM Project

Creates a local project binding that allows you to use DPPM commands
without specifying the --project flag. The binding creates a .dppm/project.yaml
file in the current directory.

Arguments:
  project-id    ID of the existing project to bind to

Examples:
  dppm bind my-project           # Bind current directory to 'my-project'
  dppm task create new-task      # Now works without --project flag
  dppm phase create phase-1      # All commands use bound project

After binding, you can use any project-scoped command without the --project flag:
  dppm task create task-id --title "New Task"
  dppm phase create phase-id --name "New Phase"
  dppm status blocked

To unbind, simply remove the .dppm directory: rm -rf .dppm`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]

		// Check if project exists
		projectPath := filepath.Join(projectsPath, "projects", projectID)
		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Project '%s' does not exist.\n", projectID)
			fmt.Fprintf(os.Stderr, "Use 'dppm list projects' to see available projects.\n")
			return
		}

		// Create .dppm directory
		if err := os.MkdirAll(".dppm", 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating .dppm directory: %v\n", err)
			return
		}

		// Read project metadata from Dropbox if available
		projectFile := filepath.Join(projectPath, "project.yaml")
		var projectName string
		if data, err := os.ReadFile(projectFile); err == nil {
			var project map[string]interface{}
			if err := yaml.Unmarshal(data, &project); err == nil {
				if name, ok := project["name"].(string); ok {
					projectName = name
				}
			}
		}

		// Create binding
		binding := LocalProjectBinding{
			ProjectID:   projectID,
			ProjectName: projectName,
			DropboxPath: projectPath,
			Created:     time.Now().Format(time.RFC3339),
		}

		data, err := yaml.Marshal(binding)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating binding configuration: %v\n", err)
			return
		}

		// Write binding file
		if err := os.WriteFile(".dppm/project.yaml", data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing binding file: %v\n", err)
			return
		}

		fmt.Printf("✅ Successfully bound current directory to project '%s'\n", projectID)
		if projectName != "" {
			fmt.Printf("📁 Project Name: %s\n", projectName)
		}
		fmt.Printf("🔗 Dropbox Path: %s\n", projectPath)
		fmt.Println()
		fmt.Println("🎯 You can now use project-scoped commands without --project flag:")
		fmt.Println("   dppm task create new-task --title \"My Task\"")
		fmt.Println("   dppm phase create phase-1 --name \"New Phase\"")
		fmt.Println("   dppm status blocked")
		fmt.Println()
		fmt.Println("💡 To unbind: rm -rf .dppm")
	},
}

func createProjectDocsLink(projectID string) {
	// Check if project exists
	projectPath := filepath.Join(projectsPath, "projects", projectID)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Project '%s' does not exist.\n", projectID)
		fmt.Fprintf(os.Stderr, "Use 'dppm list projects' to see available projects.\n")
		return
	}

	// Create docs directory in project if it doesn't exist
	docsPath := filepath.Join(projectPath, "docs")
	if err := os.MkdirAll(docsPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating docs directory: %v\n", err)
		return
	}

	// Remove existing symlink if it exists
	symlinkPath := "./dbdocs"
	if _, err := os.Lstat(symlinkPath); err == nil {
		if err := os.Remove(symlinkPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error removing existing symlink: %v\n", err)
			return
		}
		fmt.Println("🗑️  Removed existing symlink 'dbdocs'")
	}

	// Create new symlink
	if err := os.Symlink(docsPath, symlinkPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating symlink: %v\n", err)
		return
	}

	// Add to .gitignore if it exists or create it
	gitignoreFile := "./.gitignore"

	// Read existing .gitignore if it exists
	if existingContent, err := os.ReadFile(gitignoreFile); err == nil {
		// Check if dbdocs is already in .gitignore
		if !strings.Contains(string(existingContent), "dbdocs") {
			// Append to existing .gitignore
			file, err := os.OpenFile(gitignoreFile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Warning: Could not update .gitignore: %v\n", err)
			} else {
				defer file.Close()
				if _, err := file.WriteString("\n# DPPM project docs symlink\ndbdocs\n"); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: Could not write to .gitignore: %v\n", err)
				} else {
					fmt.Println("📝 Added 'dbdocs' to .gitignore")
				}
			}
		} else {
			fmt.Println("📝 'dbdocs' already exists in .gitignore")
		}
	} else {
		// Create new .gitignore
		if err := os.WriteFile(gitignoreFile, []byte("# DPPM project docs symlink\ndbdocs\n"), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Could not create .gitignore: %v\n", err)
		} else {
			fmt.Println("📝 Created .gitignore with 'dbdocs' entry")
		}
	}

	fmt.Printf("✅ Created symlink 'dbdocs' → %s\n", docsPath)
	fmt.Printf("📁 Project: %s\n", projectID)
	fmt.Printf("🔗 Docs Path: %s\n", docsPath)
	fmt.Println()
	fmt.Println("📚 You can now access project documentation via the 'dbdocs' symlink:")
	fmt.Println("   ls dbdocs/")
	fmt.Println("   code dbdocs/")
	fmt.Println("   echo 'Project notes' > dbdocs/notes.md")
	fmt.Println()
	fmt.Println("💡 The symlink is automatically ignored by git via .gitignore")
}
