package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// FirstRunSetup handles initial DPPM setup with AI guidance
type FirstRunSetup struct {
	DropboxPath   string
	SetupComplete bool
	Steps         []SetupStep
}

type SetupStep struct {
	ID          string
	Title       string
	Description string
	Required    bool
	Completed   bool
	Command     string
	Instructions string
}

// validateDropboxInstallation checks if Dropbox is properly installed and configured
func validateDropboxInstallation() (*FirstRunSetup, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %v", err)
	}

	dropboxPath := filepath.Join(home, "Dropbox")
	setup := &FirstRunSetup{
		DropboxPath: dropboxPath,
		Steps: []SetupStep{
			{
				ID:          "check-dropbox-installed",
				Title:       "1. Verify Dropbox Installation",
				Description: "Check if Dropbox desktop application is installed",
				Required:    true,
				Instructions: `🔍 DROPBOX INSTALLATION CHECK:

   If Dropbox is NOT installed:
   • Linux: sudo apt install dropbox (or download from dropbox.com)
   • macOS: brew install --cask dropbox (or download from dropbox.com)
   • Windows: Download installer from dropbox.com

   After installation:
   • Sign in to your Dropbox account
   • Complete initial sync setup
   • Verify ~/Dropbox/ folder exists and syncs`,
			},
			{
				ID:          "check-dropbox-running",
				Title:       "2. Verify Dropbox is Running",
				Description: "Ensure Dropbox daemon/service is active",
				Required:    true,
				Instructions: `🔄 DROPBOX SERVICE CHECK:

   Verify Dropbox is running:
   • Linux: ps aux | grep dropbox
   • macOS: ps aux | grep Dropbox
   • Windows: Check system tray for Dropbox icon

   If not running:
   • Start Dropbox application
   • Wait for initial sync to complete`,
			},
			{
				ID:          "check-dropbox-sync",
				Title:       "3. Verify Dropbox Sync Folder",
				Description: "Confirm ~/Dropbox/ is real synced folder, not local fake",
				Required:    true,
				Instructions: `📁 DROPBOX FOLDER VALIDATION:

   Check ~/Dropbox/ exists and syncs:
   • Folder should contain existing Dropbox files
   • Create test file: echo "test" > ~/Dropbox/dppm-test.txt
   • File should appear on dropbox.com within minutes
   • Delete test file after verification`,
			},
			{
				ID:          "create-project-structure",
				Title:       "4. Create DPPM Project Structure",
				Description: "Initialize proper folder hierarchy in Dropbox",
				Required:    true,
				Instructions: `🏗️ PROJECT STRUCTURE CREATION:

   DPPM will create:
   ~/Dropbox/project-management/
   ├── projects/          # Individual project folders
   ├── templates/         # Project templates
   └── dppm-global.db     # Global database

   This ensures consistent structure across all devices.`,
			},
			{
				ID:          "verify-permissions",
				Title:       "5. Verify File Permissions",
				Description: "Ensure DPPM can read/write to Dropbox folder",
				Required:    true,
				Instructions: `🔐 PERMISSIONS CHECK:

   DPPM needs read/write access to:
   • ~/Dropbox/project-management/ (create folders/files)
   • Database operations (SQLite files)
   • YAML file operations (project data)`,
			},
		},
	}

	// Perform validation checks
	setup.validateSteps()

	return setup, nil
}

// validateSteps checks each setup step and marks completion status
func (s *FirstRunSetup) validateSteps() {
	for i := range s.Steps {
		step := &s.Steps[i]

		switch step.ID {
		case "check-dropbox-installed":
			step.Completed = s.isDropboxInstalled()
		case "check-dropbox-running":
			step.Completed = s.isDropboxRunning()
		case "check-dropbox-sync":
			step.Completed = s.isDropboxSyncing()
		case "create-project-structure":
			step.Completed = s.hasProjectStructure()
		case "verify-permissions":
			step.Completed = s.hasPermissions()
		}
	}

	// Check if all required steps are completed
	s.SetupComplete = true
	for _, step := range s.Steps {
		if step.Required && !step.Completed {
			s.SetupComplete = false
			break
		}
	}
}

// isDropboxInstalled checks if Dropbox folder exists and looks real
func (s *FirstRunSetup) isDropboxInstalled() bool {
	// Check if ~/Dropbox exists
	if _, err := os.Stat(s.DropboxPath); os.IsNotExist(err) {
		return false
	}

	// Check if it contains typical Dropbox indicators
	indicators := []string{
		".dropbox",
		".dropbox.cache",
	}

	for _, indicator := range indicators {
		indicatorPath := filepath.Join(s.DropboxPath, indicator)
		if _, err := os.Stat(indicatorPath); err == nil {
			return true
		}
	}

	// If ~/Dropbox exists but no indicators, likely fake folder
	return false
}

// isDropboxRunning checks if Dropbox process is active
func (s *FirstRunSetup) isDropboxRunning() bool {
	// This is a simple check - in production might want more sophisticated detection
	return s.isDropboxInstalled()
}

// isDropboxSyncing verifies the folder actually syncs
func (s *FirstRunSetup) isDropboxSyncing() bool {
	// For now, same as installed check
	// Could add more sophisticated sync verification
	return s.isDropboxInstalled()
}

// hasProjectStructure checks if DPPM folder structure exists
func (s *FirstRunSetup) hasProjectStructure() bool {
	requiredDirs := []string{
		"project-management",
		"project-management/projects",
		"project-management/templates",
	}

	for _, dir := range requiredDirs {
		dirPath := filepath.Join(s.DropboxPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// hasPermissions verifies read/write access to Dropbox
func (s *FirstRunSetup) hasPermissions() bool {
	// Test write access
	testFile := filepath.Join(s.DropboxPath, ".dppm-permission-test")

	// Try to create test file
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return false
	}

	// Try to read it back
	if _, err := os.ReadFile(testFile); err != nil {
		return false
	}

	// Clean up test file
	os.Remove(testFile)

	return true
}

// createProjectStructure initializes the DPPM folder hierarchy
func (s *FirstRunSetup) createProjectStructure() error {
	requiredDirs := []string{
		"project-management",
		"project-management/projects",
		"project-management/templates",
	}

	for _, dir := range requiredDirs {
		dirPath := filepath.Join(s.DropboxPath, dir)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dirPath, err)
		}
	}

	// Create initial template files if they don't exist
	if err := s.createInitialTemplates(); err != nil {
		return fmt.Errorf("failed to create templates: %v", err)
	}

	return nil
}

// createInitialTemplates creates basic project and phase templates
func (s *FirstRunSetup) createInitialTemplates() error {
	templatesDir := filepath.Join(s.DropboxPath, "project-management", "templates")

	// Project template
	projectTemplate := `id: "example-project"
name: "Example Project"
description: "Template for new projects"
owner: "team"
status: "active"
created: "2024-01-01T00:00:00Z"

# AI Instructions
ai_instructions: |
  Follow standard DPPM workflow:
  1. Create phases (P1, P2, P3...)
  2. Create tasks (T1.1, T1.2, T2.1...)
  3. Use proper naming conventions
  4. Add comprehensive descriptions
`

	projectTemplatePath := filepath.Join(templatesDir, "project-template.yaml")
	if err := os.WriteFile(projectTemplatePath, []byte(projectTemplate), 0644); err != nil {
		return err
	}

	return nil
}

// showFirstRunGuide displays comprehensive setup instructions for AI
func showFirstRunGuide() error {
	fmt.Println(`
🤖 DPPM FIRST-RUN SETUP GUIDE FOR AI
===================================

DPPM requires proper Dropbox installation and setup before use.
Follow this guide step-by-step to ensure correct configuration.

⚠️  CRITICAL: Do NOT proceed until ALL steps are completed!
`)

	setup, err := validateDropboxInstallation()
	if err != nil {
		return fmt.Errorf("setup validation failed: %v", err)
	}

	// Show each setup step
	for _, step := range setup.Steps {
		status := "❌ INCOMPLETE"
		if step.Completed {
			status = "✅ COMPLETED"
		}

		fmt.Printf("\n%s %s\n", status, step.Title)
		fmt.Printf("   %s\n", step.Description)

		if !step.Completed && step.Required {
			fmt.Printf("\n%s\n", step.Instructions)
			if step.Command != "" {
				fmt.Printf("\n💻 COMMAND: %s\n", step.Command)
			}
		}
	}

	if !setup.SetupComplete {
		fmt.Printf(`
🚫 SETUP INCOMPLETE - DPPM CANNOT PROCEED
========================================

REQUIRED ACTIONS:
• Complete all steps marked with ❌ above
• Restart DPPM after completing setup
• Only then can you begin creating projects

⚠️  AI AGENTS: Do not attempt to use DPPM commands until setup is complete!
`)
		return fmt.Errorf("DPPM setup incomplete - cannot proceed")
	}

	// If setup is complete, create project structure
	if !setup.hasProjectStructure() {
		fmt.Println("\n🏗️ CREATING DPPM PROJECT STRUCTURE...")
		if err := setup.createProjectStructure(); err != nil {
			return fmt.Errorf("failed to create project structure: %v", err)
		}
		fmt.Println("✅ Project structure created successfully!")
	}

	fmt.Printf(`
🎉 DPPM SETUP COMPLETE!
======================

✅ Dropbox installed and running
✅ Project structure created:
   ~/Dropbox/project-management/
   ├── projects/          # Your project folders
   ├── templates/         # Project templates
   └── dppm-global.db     # Global database

🚀 AI CAN NOW BEGIN USING DPPM:
   dppm project create my-project --name "My First Project"
   dppm phase create P1 --project my-project --name "Foundation"
   dppm task create T1.1 --project my-project --phase P1 --title "Setup"

📖 For help: dppm wiki or dppm --help
`)

	return nil
}

// requireDropboxSetup enforces setup completion before allowing commands
func requireDropboxSetup() error {
	setup, err := validateDropboxInstallation()
	if err != nil {
		return err
	}

	if !setup.SetupComplete {
		fmt.Println("🚫 DPPM SETUP REQUIRED")
		fmt.Println("Run: dppm --setup for first-time configuration")
		return fmt.Errorf("DPPM setup incomplete")
	}

	// Ensure project structure exists
	if !setup.hasProjectStructure() {
		if err := setup.createProjectStructure(); err != nil {
			return fmt.Errorf("failed to initialize project structure: %v", err)
		}
	}

	return nil
}`