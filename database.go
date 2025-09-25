package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// initDatabase initializes the SQLite database with ERD schema
func initDatabase() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	dbPath := filepath.Join(home, "Dropbox", "project-management", "dppm.db")

	// Create data directory if it doesn't exist
	dataDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("failed to create data directory: %v", err)
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Enable foreign key constraints
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return fmt.Errorf("failed to enable foreign keys: %v", err)
	}

	return createERDSchema()
}

// createERDSchema creates the ERD tables with AI-friendly constraints
func createERDSchema() error {
	schema := `
	-- Projects table with AI-friendly constraints
	CREATE TABLE IF NOT EXISTS projects (
		id TEXT PRIMARY KEY CHECK (
			id GLOB '[a-z0-9]*' AND
			id NOT GLOB '*_*' AND
			id NOT GLOB '*..*' AND
			LENGTH(id) BETWEEN 3 AND 50
		),
		name TEXT NOT NULL CHECK (LENGTH(name) > 0),
		owner TEXT NOT NULL,
		status TEXT DEFAULT 'active' CHECK (status IN ('active', 'completed', 'archived')),
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Phases table with P1/P2/P3/BUGS validation
	CREATE TABLE IF NOT EXISTS phases (
		id TEXT PRIMARY KEY CHECK (
			id GLOB 'P[1-9]*' OR id = 'BUGS'
		),
		project_id TEXT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
		name TEXT NOT NULL CHECK (LENGTH(name) > 0),
		phase_number INTEGER CHECK (phase_number > 0),
		status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'active', 'completed')),
		created DATETIME DEFAULT CURRENT_TIMESTAMP,

		UNIQUE(project_id, phase_number),
		CHECK (
			(id = 'BUGS' AND phase_number IS NULL) OR
			(id != 'BUGS' AND phase_number IS NOT NULL)
		)
	);

	-- Tasks table with T1.1/bug-name validation
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY CHECK (
			id GLOB 'T[1-9]*.[1-9]*' OR
			id GLOB 'bug-[a-z0-9]*'
		),
		phase_id TEXT NOT NULL REFERENCES phases(id) ON DELETE CASCADE,
		title TEXT NOT NULL CHECK (LENGTH(title) > 0),
		status TEXT DEFAULT 'pending' CHECK (status IN ('pending', 'in_progress', 'done')),

		-- Issue #36: Required structured descriptions
		project_description TEXT NOT NULL CHECK (LENGTH(project_description) >= 50),
		task_description TEXT NOT NULL CHECK (LENGTH(task_description) >= 30),
		expected_product_description TEXT NOT NULL CHECK (LENGTH(expected_product_description) >= 20),

		-- Optional fields
		assigned_to TEXT,
		priority TEXT DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high')),
		estimated_hours INTEGER CHECK (estimated_hours > 0),
		actual_hours INTEGER CHECK (actual_hours >= 0),
		due_date DATE,

		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Task dependencies with cycle prevention
	CREATE TABLE IF NOT EXISTS task_dependencies (
		task_id TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
		depends_on TEXT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (task_id, depends_on),
		CHECK (task_id != depends_on)
	);

	-- AI Guidance Views
	CREATE VIEW IF NOT EXISTS ai_ready_tasks AS
	SELECT
		t.id,
		t.title,
		'dppm task update ' || t.id || ' --status in_progress' as suggested_command
	FROM tasks t
	WHERE t.status = 'pending'
	AND NOT EXISTS (
		SELECT 1 FROM task_dependencies td
		JOIN tasks dep ON dep.id = td.depends_on
		WHERE td.task_id = t.id AND dep.status != 'done'
	);

	-- Next available task IDs
	CREATE VIEW IF NOT EXISTS next_available_tasks AS
	SELECT
		p.id as phase_id,
		'T' || SUBSTR(p.id, 2) || '.' || COALESCE(MAX(task_num), 0) + 1 as next_task_id
	FROM phases p
	LEFT JOIN (
		SELECT
			phase_id,
			CAST(SUBSTR(id, INSTR(id, '.') + 1) AS INTEGER) as task_num
		FROM tasks
		WHERE id GLOB 'T*.*'
	) t ON t.phase_id = p.id
	WHERE p.id != 'BUGS'
	GROUP BY p.id;

	-- AI Guidance dashboard
	CREATE VIEW IF NOT EXISTS ai_guidance AS
	SELECT
		'READY' as type,
		id as task_id,
		title,
		'✅ Can start immediately' as message,
		'dppm task update ' || id || ' --status in_progress' as command
	FROM tasks
	WHERE id IN (SELECT id FROM ai_ready_tasks)

	UNION ALL

	SELECT
		'BLOCKED' as type,
		t.id,
		t.title,
		'🚫 Waiting for: ' || GROUP_CONCAT(dep.title, ', ') as message,
		'Complete dependencies first' as command
	FROM tasks t
	JOIN task_dependencies td ON td.task_id = t.id
	JOIN tasks dep ON dep.id = td.depends_on
	WHERE t.status = 'pending' AND dep.status != 'done'
	GROUP BY t.id, t.title;
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %v", err)
	}

	return createERDTriggers()
}

// createERDTriggers creates status transition and dependency enforcement triggers
func createERDTriggers() error {
	triggers := `
	-- Status transition enforcement
	CREATE TRIGGER IF NOT EXISTS enforce_status_transitions
	BEFORE UPDATE ON tasks
	FOR EACH ROW
	WHEN NEW.status != OLD.status
	BEGIN
		-- Cannot go back from done
		SELECT CASE
			WHEN OLD.status = 'done' AND NEW.status != 'done' THEN
				RAISE(ABORT, '❌ Cannot change status from done. Task is completed.')
		END;

		-- Cannot skip in_progress
		SELECT CASE
			WHEN OLD.status = 'pending' AND NEW.status = 'done' THEN
				RAISE(ABORT, '❌ Cannot skip in_progress. Use --status in_progress first')
		END;
	END;

	-- Dependency enforcement
	CREATE TRIGGER IF NOT EXISTS enforce_dependencies
	BEFORE UPDATE ON tasks
	FOR EACH ROW
	WHEN NEW.status = 'in_progress' AND OLD.status = 'pending'
	BEGIN
		SELECT CASE
			WHEN EXISTS (
				SELECT 1 FROM task_dependencies td
				JOIN tasks dep ON dep.id = td.depends_on
				WHERE td.task_id = NEW.id AND dep.status != 'done'
			) THEN
				RAISE(ABORT, '❌ Cannot start task. Dependencies not completed. Use: dppm status dependencies ' || NEW.id)
		END;
	END;

	-- Update timestamp trigger
	CREATE TRIGGER IF NOT EXISTS update_task_timestamp
	AFTER UPDATE ON tasks
	FOR EACH ROW
	BEGIN
		UPDATE tasks SET updated = CURRENT_TIMESTAMP WHERE id = NEW.id;
	END;
	`

	_, err := db.Exec(triggers)
	return err
}

// validateTaskID checks if task ID follows ERD constraints
func validateTaskID(taskID string) error {
	if taskID == "" {
		return fmt.Errorf("❌ Task ID cannot be empty")
	}

	// Check T1.1 pattern
	if taskID[0] == 'T' {
		// Validate pattern with regex equivalent
		var phase, num int
		if n, err := fmt.Sscanf(taskID, "T%d.%d", &phase, &num); n != 2 || err != nil {
			return fmt.Errorf("❌ Task ID must match T{phase}.{number} format. Example: T1.1, T2.3")
		}
		if phase == 0 || num == 0 {
			return fmt.Errorf("❌ Phase and task numbers must start from 1. Example: T1.1 (not T0.1 or T1.0)")
		}
		return nil
	}

	// Check bug-* pattern
	if len(taskID) > 4 && taskID[:4] == "bug-" {
		if len(taskID) < 6 {
			return fmt.Errorf("❌ Bug task name too short. Example: bug-login-error")
		}
		return nil
	}

	return fmt.Errorf("❌ Task ID must be T{phase}.{number} or bug-{description}. Examples: T1.1, bug-login-error")
}

// validatePhaseID checks if phase ID follows ERD constraints
func validatePhaseID(phaseID string) error {
	if phaseID == "" {
		return fmt.Errorf("❌ Phase ID cannot be empty")
	}

	if phaseID == "BUGS" {
		return nil
	}

	if phaseID[0] == 'P' {
		var num int
		if n, err := fmt.Sscanf(phaseID, "P%d", &num); n != 1 || err != nil {
			return fmt.Errorf("❌ Phase ID must be P{number} or BUGS. Examples: P1, P2, P3, BUGS")
		}
		if num == 0 {
			return fmt.Errorf("❌ Phase numbers must start from 1. Example: P1 (not P0)")
		}
		return nil
	}

	return fmt.Errorf("❌ Phase ID must be P{number} or BUGS. Examples: P1, P2, P3, BUGS")
}

// suggestNextTaskID suggests the next available task ID for a phase
func suggestNextTaskID(phaseID string) (string, error) {
	if phaseID == "BUGS" {
		return "bug-describe-issue", nil
	}

	row := db.QueryRow(`
		SELECT next_task_id
		FROM next_available_tasks
		WHERE phase_id = ?
	`, phaseID)

	var nextID string
	if err := row.Scan(&nextID); err != nil {
		// Phase doesn't exist yet, suggest T{phase}.1
		var num int
		if n, err := fmt.Sscanf(phaseID, "P%d", &num); n == 1 && err == nil {
			return fmt.Sprintf("T%d.1", num), nil
		}
		return "", fmt.Errorf("invalid phase ID: %s", phaseID)
	}

	return nextID, nil
}

// getAIGuidance returns current AI guidance for ready/blocked tasks
func getAIGuidance() ([]map[string]string, error) {
	rows, err := db.Query(`
		SELECT type, task_id, title, message, command
		FROM ai_guidance
		ORDER BY
			CASE type
				WHEN 'READY' THEN 1
				WHEN 'BLOCKED' THEN 2
			END,
			task_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var guidance []map[string]string
	for rows.Next() {
		var g map[string]string = make(map[string]string)
		err := rows.Scan(&g["type"], &g["task_id"], &g["title"], &g["message"], &g["command"])
		if err != nil {
			return nil, err
		}
		guidance = append(guidance, g)
	}

	return guidance, nil
}

// closeDatabase closes the database connection
func closeDatabase() error {
	if db != nil {
		return db.Close()
	}
	return nil
}