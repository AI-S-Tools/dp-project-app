/* ::GEMINI:01: Der er hardcoded stier her ('./dppm-test'), der bør bruges variabler eller en konfigurationsmekanisme for at gøre det mere fleksibelt.:: */
/* ::GEMINI:02: Fejlhåndteringen kan forbedres ved at returnere mere specifikke fejltyper i stedet for generiske strenge.:: */
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// /* Definerer 'init' kommandoen for at initialisere et nyt projekt. */
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new project with complete setup",
	Long: `Project Initialization System

Creates a complete project setup including:
  • DPPM project in Dropbox
  • Local project directory with Git repository
  • Symlinked documentation between local and Dropbox
  • GitHub repository (optional)
  • AI-powered project structure analysis

This command automates the complete project setup workflow,
perfect for AI-driven development initialization.

Arguments:
  project-name    Name of the project to initialize

Examples:
  dppm init web-app --doc "/path/to/project/docs"
  dppm init api-server --org "my-org" --private
  dppm init mobile-app --doc "./requirements.md" --template "react-native"

AI Integration:
  The init system can analyze project documentation and automatically
  create appropriate phases, tasks, and project structure.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectID := strings.ToLower(strings.ReplaceAll(projectName, " ", "-"))

		docPath, _ := cmd.Flags().GetString("doc")
		org, _ := cmd.Flags().GetString("org")
		private, _ := cmd.Flags().GetBool("private")
		template, _ := cmd.Flags().GetString("template")
		skipGithub, _ := cmd.Flags().GetBool("skip-github")

		fmt.Printf("🚀 Initializing project '%s'\n", projectName)
		fmt.Printf("==========================================\n\n")

		// Step 1: Create DPPM project
		fmt.Printf("1️⃣ Creating DPPM project...\n")
		if err := createDPPMProject(projectID, projectName, docPath); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to create DPPM project: %v\n", err)
			return
		}
		fmt.Printf("✅ DPPM project created\n\n")

		// Step 2: Create local project directory
		fmt.Printf("2️⃣ Creating local project directory...\n")
		localDir := filepath.Join(".", projectID)
		if err := createLocalProject(localDir); err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to create local project: %v\n", err)
			return
		}
		fmt.Printf("✅ Local project directory created: %s\n\n", localDir)

		// Step 3: Create symlinked documentation
		fmt.Printf("3️⃣ Setting up documentation symlink...\n")
		if err := setupDocumentationLink(projectID, localDir, docPath); err != nil {
			fmt.Printf("⚠️  Warning: Could not create documentation symlink: %v\n", err)
		} else {
			fmt.Printf("✅ Documentation symlink created\n")
		}
		fmt.Println()

		// Step 4: Initialize Git repository
		fmt.Printf("4️⃣ Initializing Git repository...\n")
		if err := initializeGitRepo(localDir); err != nil {
			fmt.Printf("⚠️  Warning: Could not initialize Git: %v\n", err)
		} else {
			fmt.Printf("✅ Git repository initialized\n")
		}
		fmt.Println()

		// Step 5: Create GitHub repository (optional)
		if !skipGithub {
			fmt.Printf("5️⃣ Creating GitHub repository...\n")
			if err := createGithubRepo(projectID, projectName, org, private); err != nil {
				fmt.Printf("⚠️  Warning: Could not create GitHub repo: %v\n", err)
			} else {
				fmt.Printf("✅ GitHub repository created\n")
			}
			fmt.Println()
		}

		// Step 6: AI Analysis and Structure Creation
		fmt.Printf("6️⃣ AI Analysis and Project Structure...\n")
		if docPath != "" && fileExists(docPath) {
			analyzeAndCreateStructure(projectID, docPath, template)
		} else {
			createDefaultStructure(projectID, template)
		}
		fmt.Printf("✅ Project structure created\n\n")

		// Success summary
		fmt.Printf("🎉 Project initialization completed!\n")
		fmt.Printf("==========================================\n\n")

		fmt.Printf("📁 Project Details:\n")
		fmt.Printf("   • DPPM Project: %s\n", projectID)
		fmt.Printf("   • Local Directory: %s\n", localDir)
		fmt.Printf("   • Dropbox Storage: ~/Dropbox/project-management/projects/%s\n", projectID)
		if !skipGithub {
			repoUrl := fmt.Sprintf("https://github.com/%s/%s", getGithubUser(org), projectID)
			fmt.Printf("   • GitHub Repository: %s\n", repoUrl)
		}

		fmt.Printf("\n🚀 Next Steps:\n")
		fmt.Printf("   cd %s\n", localDir)
		fmt.Printf("   dppm status project %s\n", projectID)
		fmt.Printf("   dppm wiki \"ai workflow\"\n")

		fmt.Printf("\n💡 AI Tip:\n")
		fmt.Printf("   The project is ready for AI collaboration! Use:\n")
		fmt.Printf("   dppm collab wiki \"getting started\"\n")
	},
}

// /* Opretter et DPPM-projekt ved at kalde 'dppm-test' kommandoen. */
func createDPPMProject(projectID, projectName, docPath string) error {
	// Use the existing project create command
	cmd := exec.Command("./dppm-test", "project", "create", projectID, "--name", projectName)
	if docPath != "" {
		// If we have doc path, use it as description context
		cmd.Args = append(cmd.Args, "--description", fmt.Sprintf("Project initialized from %s", docPath))
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("project creation failed: %v\nOutput: %s", err, output)
	}
	return nil
}

// /* Opretter en lokal projektmappe. */
func createLocalProject(localDir string) error {
	return os.MkdirAll(localDir, 0755)
}

// /* Opretter et symbolsk link til dokumentationen i Dropbox. */
func setupDocumentationLink(projectID, localDir, docPath string) error {
	dropboxDocsDir := filepath.Join(os.Getenv("HOME"), "Dropbox", "project-management", "projects", projectID, "docs")
	localDocsDir := filepath.Join(localDir, "docs")

	// Create Dropbox docs directory
	if err := os.MkdirAll(dropboxDocsDir, 0755); err != nil {
		return fmt.Errorf("failed to create Dropbox docs directory: %v", err)
	}

	// Copy initial documentation if provided
	if docPath != "" && fileExists(docPath) {
		destPath := filepath.Join(dropboxDocsDir, filepath.Base(docPath))
		if err := copyFile(docPath, destPath); err != nil {
			return fmt.Errorf("failed to copy documentation: %v", err)
		}
	}

	// Create symlink from local to Dropbox
	return os.Symlink(dropboxDocsDir, localDocsDir)
}

// /* Initialiserer et Git-repository i den lokale projektmappe. */
func initializeGitRepo(localDir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = localDir
	return cmd.Run()
}

// /* Opretter et GitHub-repository ved hjælp af 'gh' CLI-værktøjet. */
func createGithubRepo(projectID, projectName, org string, private bool) error {
	args := []string{"repo", "create", projectID, "--description", projectName}

	if org != "" {
		args = append(args, "--org", org)
	}

	if private {
		args = append(args, "--private")
	} else {
		args = append(args, "--public")
	}

	cmd := exec.Command("gh", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("GitHub repo creation failed: %v\nOutput: %s", err, output)
	}
	return nil
}

// /* Analyserer projektdokumentation for at oprette en passende projektstruktur. */
func analyzeAndCreateStructure(projectID, docPath, template string) {
	fmt.Printf("📄 Analyzing project documentation: %s\n", docPath)

	// Read and analyze the documentation
	docContent, err := os.ReadFile(docPath)
	if err != nil {
		fmt.Printf("⚠️  Could not read documentation file: %v\n", err)
		createDefaultStructure(projectID, template)
		return
	}

	content := string(docContent)

	// Simple analysis - look for common project patterns
	phases := analyzeProjectPhases(content, template)

	fmt.Printf("🧠 AI Analysis Results:\n")
	fmt.Printf("   • Detected %d phases\n", len(phases))
	fmt.Printf("   • Template: %s\n", getTemplateOrDefault(template))

	// Create phases based on analysis
	for i, phase := range phases {
		phaseID := fmt.Sprintf("phase-%d", i+1)
		fmt.Printf("   Creating phase: %s\n", phase)

		cmd := exec.Command("./dppm-test", "phase", "create", phaseID,
			"--project", projectID,
			"--name", phase,
			"--goal", fmt.Sprintf("Complete %s development", phase))

		if output, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("     ⚠️  Warning: Could not create phase %s: %v\n", phase, err)
			fmt.Printf("     Output: %s\n", output)
		}
	}
}

// /* Opretter en standard projektstruktur baseret på en skabelon. */
func createDefaultStructure(projectID, template string) {
	fmt.Printf("📋 Creating default project structure for template: %s\n", getTemplateOrDefault(template))

	phases := getDefaultPhases(template)

	for i, phase := range phases {
		phaseID := fmt.Sprintf("phase-%d", i+1)
		fmt.Printf("   Creating phase: %s\n", phase)

		cmd := exec.Command("./dppm-test", "phase", "create", phaseID,
			"--project", projectID,
			"--name", phase,
			"--goal", fmt.Sprintf("Complete %s phase", phase))

		if output, err := cmd.CombinedOutput(); err != nil {
			fmt.Printf("     ⚠️  Warning: Could not create phase %s: %v\n", phase, err)
			fmt.Printf("     Output: %s\n", output)
		}
	}
}

// /* Analyserer projektindhold for at udlede relevante projektfaser. */
func analyzeProjectPhases(content, template string) []string {
	content = strings.ToLower(content)

	// Default phases based on common patterns
	phases := []string{"Planning", "Development", "Testing", "Deployment"}

	// Adjust based on content analysis
	if strings.Contains(content, "frontend") || strings.Contains(content, "ui") || strings.Contains(content, "react") || strings.Contains(content, "vue") || strings.Contains(content, "angular") {
		phases = []string{"Planning", "Backend", "Frontend", "Integration", "Deployment"}
	}

	if strings.Contains(content, "api") || strings.Contains(content, "backend") || strings.Contains(content, "database") {
		phases = []string{"Planning", "Database", "API", "Frontend", "Testing", "Deployment"}
	}

	if strings.Contains(content, "mobile") || strings.Contains(content, "ios") || strings.Contains(content, "android") {
		phases = []string{"Planning", "Backend API", "Mobile App", "Testing", "App Store"}
	}

	// Template-specific adjustments
	sswitch template {
	case "web":
		phases = []string{"Setup", "Backend", "Frontend", "Integration", "Deployment"}
	case "api":
		phases = []string{"Planning", "Database", "API Development", "Testing", "Documentation"}
	case "mobile":
		phases = []string{"Planning", "Backend", "Mobile UI", "Native Features", "Store Release"}
	}

	return phases
}

// /* Returnerer en liste over standardfaser for en given skabelon. */
func getDefaultPhases(template string) []string {
	sswitch template {
	case "web":
		return []string{"Setup", "Backend", "Frontend", "Integration"}
	case "api":
		return []string{"Planning", "Database", "API Development", "Testing"}
	case "mobile":
		return []string{"Planning", "Backend", "Mobile UI", "Testing"}
	default:
		return []string{"Planning", "Development", "Testing", "Deployment"}
	}
}

// /* Returnerer skabelonnavnet eller 'default', hvis det er tomt. */
func getTemplateOrDefault(template string) string {
	if template == "" {
		return "default"
	}
	return template
}

// /* Henter GitHub-brugernavnet fra 'gh' CLI-værktøjet. */
func getGithubUser(org string) string {
	if org != "" {
		return org
	}

	// Try to get current GitHub user
	cmd := exec.Command("gh", "api", "user", "--jq", ".login")
	output, err := cmd.Output()
	if err != nil {
		return "your-username"
	}

	return strings.TrimSpace(string(output))
}

// /* Kontrollerer, om en fil eksisterer. */
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// /* Kopierer en fil fra kilde til destination. */
func copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, input, 0644)
}

// /* Initialiserer 'init' kommandoen og dens flag. */
func init() {
	initCmd.Flags().StringP("doc", "d", "", "Path to project documentation file")
	initCmd.Flags().StringP("org", "o", "", "GitHub organization (optional)")
	initCmd.Flags().Bool("private", false, "Create private GitHub repository")
	initCmd.Flags().StringP("template", "t", "", "Project template (web, api, mobile)")
	initCmd.Flags().Bool("skip-github", false, "Skip GitHub repository creation")

	rootCmd.AddCommand(initCmd)
}