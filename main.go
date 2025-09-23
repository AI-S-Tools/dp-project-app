package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

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

Storage Location: ~/Dropbox/project-management/

🚀 Quick Start Guide:
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
  dppm project create web-app --name "Web Application" --owner "dev-team"
  dppm phase create backend --project web-app --name "Backend Development"
  dppm task create auth --project web-app --phase backend --title "Authentication"
  dppm status project web-app
  dppm list projects

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

	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(phaseCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(wikiCmd)

	// Add --wiki flag for direct search
	rootCmd.Flags().String("wiki", "", "Search DPPM knowledge base (e.g. --wiki \"create task\")")
}

func main() {
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

🚀 CREATE NEW PROJECT:
  dppm project create my-project --name "My Project" --owner "your-name"
  dppm phase create phase-1 --project my-project --name "First Phase"
  dppm task create first-task --project my-project --phase phase-1

💡 COMMON WORKFLOWS:
  • New to DPPM? → dppm wiki "complete"
  • Creating tasks? → dppm --wiki "create task"
  • Managing dependencies? → dppm --wiki "dependencies"
  • Checking progress? → dppm status project PROJECT_NAME

🤖 AI TIP:
DPPM is designed for AI workflows. The wiki system contains comprehensive
guides for every feature. Use it to get detailed, actionable information.

Try: dppm wiki "project workflow" to see a complete example!`)
}