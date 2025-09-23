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

A comprehensive CLI tool for managing projects, sprints, and tasks using
Dropbox as the storage backend. Perfect for AI-driven development workflows.

Features:
  • YAML-based project, sprint, and task management
  • Hierarchical project organization
  • Cross-platform synchronization via Dropbox
  • AI-friendly verbose output and documentation
  • Template-based project creation
  • Comprehensive help system

Storage Location: ~/Dropbox/project-management/

Examples:
  dppm project create my-project --name "My Project" --owner "username"
  dppm list projects
  dppm project show my-project
  dppm --help

For detailed command help, use: dppm [command] --help`,
}

var projectsPath string

func init() {
	home, _ := os.UserHomeDir()
	projectsPath = filepath.Join(home, "Dropbox", "project-management")

	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(sprintCmd)
	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(listCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}