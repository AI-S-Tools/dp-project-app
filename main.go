package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dp",
	Short: "Dropbox Project Management CLI",
	Long:  "A CLI tool for managing projects, sprints, and tasks in Dropbox",
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