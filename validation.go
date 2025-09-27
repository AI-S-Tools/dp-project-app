package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	// Valid ID pattern: starts with alphanumeric, contains only letters, numbers, hyphens, underscores
	validIDRegex = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9-_]*$`)

	// Reserved names that should not be used as IDs (OS-specific)
	reservedNames = []string{
		".", "..", "CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5",
		"COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5",
		"LPT6", "LPT7", "LPT8", "LPT9",
	}
)

// ValidateProjectID validates a project ID for security and correctness
func ValidateProjectID(id string) error {
	// Check empty
	if id == "" {
		return fmt.Errorf("project ID cannot be empty")
	}

	// Check length
	if len(id) > 255 {
		return fmt.Errorf("project ID too long (max 255 characters)")
	}

	// Check for path traversal attempts
	if strings.Contains(id, "..") ||
	   strings.Contains(id, "/") ||
	   strings.Contains(id, "\\") ||
	   strings.HasPrefix(id, "~") ||
	   strings.Contains(id, ":") {
		return fmt.Errorf("project ID contains invalid path characters: %s", id)
	}

	// Check for command injection attempts
	dangerousChars := []string{"|", "&", ";", "$", "`", "(", ")", "<", ">", "\"", "'", "\n", "\r", "\t"}
	for _, char := range dangerousChars {
		if strings.Contains(id, char) {
			return fmt.Errorf("project ID contains potentially dangerous character: %s", char)
		}
	}

	// Check reserved names (case-insensitive)
	upperID := strings.ToUpper(id)
	for _, reserved := range reservedNames {
		if upperID == reserved {
			return fmt.Errorf("'%s' is a reserved name and cannot be used as a project ID", id)
		}
	}

	// Check regex pattern for allowed characters
	if !validIDRegex.MatchString(id) {
		return fmt.Errorf("project ID must start with alphanumeric and contain only letters, numbers, hyphens, and underscores")
	}

	// Additional safety: ensure the resolved path stays within project directory
	// This is a secondary check after sanitization
	testPath := filepath.Clean(filepath.Join("projects", id))
	if !strings.HasPrefix(testPath, "projects") {
		return fmt.Errorf("project ID would create path outside project directory")
	}

	return nil
}

// ValidatePhaseID validates a phase ID must follow P1, P2, P3 format (with optional suffix)
func ValidatePhaseID(id string) error {
	// Check empty
	if id == "" {
		return fmt.Errorf("phase ID cannot be empty")
	}

	// Phase must start with P followed by a number
	// Valid formats: P1, P2, P1-backend, P2-frontend, etc.
	phaseRegex := regexp.MustCompile(`^P[1-9][0-9]*(-[a-zA-Z0-9][a-zA-Z0-9-_]*)?$`)
	if !phaseRegex.MatchString(id) {
		return fmt.Errorf("phase ID must follow format: P1, P2, P3... (optionally with suffix like P1-backend)")
	}

	// Check for dangerous characters (additional safety)
	if strings.Contains(id, "/") || strings.Contains(id, "\\") ||
	   strings.Contains(id, "..") || strings.Contains(id, "~") {
		return fmt.Errorf("phase ID contains invalid path characters")
	}

	return nil
}

// ValidateTaskID validates task ID format: T1.1, T2.1, subtasks T1.1.1, bugs T1.1.B1
func ValidateTaskID(id string, phaseID string) error {
	// Check empty
	if id == "" {
		return fmt.Errorf("task ID cannot be empty")
	}

	// Extract phase number from phase ID (e.g., P1 -> 1, P2-backend -> 2)
	var phaseNum int
	if n, err := fmt.Sscanf(phaseID, "P%d", &phaseNum); n != 1 || err != nil {
		// If phase ID is empty or invalid, allow any valid task format
		// This happens when validating without phase context
		if phaseID != "" {
			return fmt.Errorf("invalid phase ID format: %s", phaseID)
		}
		// Allow validation without phase context - just check format
		phaseNum = -1
	}

	// Task formats:
	// T1.1, T1.2 (regular tasks)
	// T1.1-auth, T1.2-login (with suffix)
	// T1.1.1, T1.1.2 (subtasks)
	// T1.1.B1, T1.1.B2 (bugs)
	taskRegex := regexp.MustCompile(`^T([1-9][0-9]*)\.([1-9][0-9]*)(\.[1-9][0-9]*|\.B[1-9][0-9]*)?(-[a-zA-Z0-9][a-zA-Z0-9-_]*)?$`)
	matches := taskRegex.FindStringSubmatch(id)
	if matches == nil {
		return fmt.Errorf("task ID must follow format: T1.1, T1.2 (subtasks: T1.1.1, bugs: T1.1.B1, with suffix: T1.1-auth)")
	}

	// If we have a phase context, validate the task belongs to that phase
	if phaseNum > 0 {
		taskPhaseNum, _ := strconv.Atoi(matches[1])
		if taskPhaseNum != phaseNum {
			return fmt.Errorf("task %s does not belong to phase %s (should start with T%d.)", id, phaseID, phaseNum)
		}
	}

	// Check for dangerous characters (additional safety)
	if strings.Contains(id, "/") || strings.Contains(id, "\\") ||
	   strings.Contains(id, "..") || strings.Contains(id, "~") {
		return fmt.Errorf("task ID contains invalid path characters")
	}

	return nil
}

// ValidateDescription validates a description or title field
func ValidateDescription(desc string) error {
	if len(desc) > 1000 {
		return fmt.Errorf("description too long (max 1000 characters)")
	}

	// Check for valid UTF-8
	if !utf8.ValidString(desc) {
		return fmt.Errorf("description contains invalid UTF-8 characters")
	}

	// Remove any control characters that could cause issues in YAML
	for i, r := range desc {
		if r < 32 && r != '\n' && r != '\t' {
			return fmt.Errorf("description contains invalid control character at position %d", i)
		}
	}

	return nil
}

// SanitizeForYAML escapes special characters that could cause YAML parsing issues
func SanitizeForYAML(s string) string {
	// Escape special YAML characters
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, "\"", "\\\"")
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

// CheckProjectExists checks if a project with the given ID already exists
func CheckProjectExists(projectID string) (bool, error) {
	projectPath := filepath.Join(projectsPath, "projects", projectID)
	if _, err := os.Stat(projectPath); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}