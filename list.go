package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects, sprints, or tasks",
}

var listProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "List all projects",
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

func init() {
	listCmd.AddCommand(listProjectsCmd)
}