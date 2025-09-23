package main

import (
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Task management commands",
}

// Task commands will be implemented here