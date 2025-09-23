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
	ID          string            `yaml:"id"`
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Status      string            `yaml:"status"`
	Owner       string            `yaml:"owner"`
	Created     string            `yaml:"created"`
	Updated     string            `yaml:"updated"`
	Repository  string            `yaml:"repository,omitempty"`
	Tags        []string          `yaml:"tags,omitempty"`
	Metadata    map[string]interface{} `yaml:"metadata,omitempty"`
	Notes       string            `yaml:"notes,omitempty"`
	CurrentSprint string          `yaml:"current_sprint,omitempty"`
	Sprints     []string          `yaml:"sprints,omitempty"`
}

var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Project management commands",
}

var createProjectCmd = &cobra.Command{
	Use:   "create [project-id]",
	Short: "Create a new project",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectID := args[0]
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		owner, _ := cmd.Flags().GetString("owner")

		if name == "" {
			name = projectID
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
			Sprints:     []string{},
		}

		projectDir := filepath.Join(projectsPath, "projects", projectID)
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project directory: %v\n", err)
			return
		}

		sprintsDir := filepath.Join(projectDir, "sprints")
		if err := os.MkdirAll(sprintsDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating sprints directory: %v\n", err)
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

func init() {
	createProjectCmd.Flags().StringP("name", "n", "", "Project name")
	createProjectCmd.Flags().StringP("description", "d", "", "Project description")
	createProjectCmd.Flags().StringP("owner", "o", "", "Project owner")

	projectCmd.AddCommand(createProjectCmd)
}