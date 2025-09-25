# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Development Commands

```bash
# Build the application
go build -o dppm

# Build with version tag
go build -ldflags="-s -w -X main.version=v1.0.0" -o dppm

# Cross-platform builds (used in CI/CD)
GOOS=linux GOARCH=amd64 go build -o dppm-linux-amd64
GOOS=darwin GOARCH=arm64 go build -o dppm-macos-arm64
GOOS=windows GOARCH=amd64 go build -o dppm-windows-amd64.exe

# Format code before committing
gofmt -w *.go

# Update dependencies
go mod tidy

# Run type checking and linting
go vet ./...
```

## Testing Commands

```bash
# Test the binary after building
./dppm --version
./dppm --help
./dppm wiki list

# Test project creation flow
./dppm project create test-project --name "Test Project" --owner "user"
./dppm list projects
./dppm status project test-project

# Test complete workflow
./dppm init test-workflow --name "Test Workflow"
./dppm phase create phase-1 --project test-workflow
./dppm task create task-1 --project test-workflow --phase phase-1 --title "First Task"

# Test AI collaboration features
./dppm collab find .
./dppm wiki "complete"
```

## Project Architecture

DPPM is a CLI tool for managing projects using YAML files stored in Dropbox. The architecture emphasizes AI-friendly workflows and comprehensive documentation.

### Command Structure
- **Root Command (`main.go`)**: Entry point with startup guide, version handling, and wiki shortcut
- **Core Subcommands**: Each major feature area has its own file:
  - `init.go`: Complete project initialization workflow
  - `project.go`: Project CRUD operations
  - `task.go`: Advanced task management with components, issues, and dependencies
  - `sprint.go`: Phase (sprint) management (note: "phases" are the new term for sprints)
  - `status.go`: Multi-level status reporting across projects/phases/tasks
  - `list.go`: List entities at all levels with filtering
  - `wiki.go`: Built-in knowledge base with 20+ help topics
  - `collab.go`: AI collaboration features with DSL markers and task finding

### Data Model & Storage
- **Storage Path**: `~/Dropbox/project-management/projects/`
- **Hierarchical Structure**: Projects → Phases → Tasks → Components/Issues
- **YAML Schema**: All entities stored as YAML files with rich metadata
- **Template System**: Embedded templates for consistent entity creation
- **Dependency Management**: Tasks support dependency chains with blocking detection

### Key Design Patterns
1. **AI-First Design**: Verbose output, comprehensive help system, built-in wiki
2. **Command Pattern**: Every operation is a cobra.Command with consistent flags and extensive help
3. **File-based Storage**: No database; YAML files enable version control and easy inspection
4. **Hierarchical Organization**: Filesystem structure mirrors logical project organization
5. **Template-driven Creation**: Consistent entity structure via embedded templates
6. **Dropbox Sync**: Cross-platform synchronization for team collaboration

### Important Functions & Entry Points
- `main()`: CLI entry point with startup guide logic
- `showStartupGuide()`: Interactive help when run without commands
- `initProjectCmd()`: Complete project initialization workflow
- `createProjectCmd()`: Project creation with validation
- `createTaskCmd()`: Task creation with component/dependency support
- `getProjectPath()`: Helper to resolve project storage paths
- `readProjectYAML()`/`writeProjectYAML()`: Core YAML I/O functions

### YAML Schema Key Features
- **Tasks**: Support components, issues, dependencies, time tracking, progress
- **Projects**: Metadata, status tracking, owner information
- **Phases**: Date ranges, goals, completion tracking
- **Components**: Sub-task breakdown with individual status
- **Issues**: Bug tracking within tasks

## Development Workflow

When modifying DPPM:

1. **Adding New Commands**:
   - Create cobra.Command in appropriate file
   - Add to parent command in `init()` function
   - Follow existing pattern for flags, help text, and error handling
   - Consider updating wiki if adding major feature

2. **Modifying YAML Schemas**:
   - Update struct definitions with proper yaml tags
   - Ensure backward compatibility with existing files
   - Test with real YAML files
   - Update embedded templates if needed

3. **Testing Changes**:
   - Build locally with `go build -o dppm`
   - Test against real Dropbox project structure
   - Verify all command help text is accurate
   - Test edge cases and error conditions

4. **Code Quality**:
   - Run `gofmt -w *.go` before committing
   - Run `go vet ./...` to catch issues
   - Run `go mod tidy` after adding imports
   - Ensure all commands have comprehensive help text

## Release Process

Releases are automated via GitHub Actions when pushing version tags:

```bash
git tag v1.2.0
git push origin v1.2.0
```

This triggers `.github/workflows/release.yml` which:
1. Builds binaries for all platforms (Linux, macOS, Windows / amd64, arm64)
2. Creates GitHub release with detailed changelog
3. Uploads binaries with SHA256 checksums
4. Updates Homebrew tap for easy installation

## Key Dependencies

- **Cobra**: CLI framework providing command structure, flags, and help
- **YAML v3**: YAML parsing and generation for all data files
- **Standard Library**: File operations, path handling, time formatting

The project intentionally keeps dependencies minimal to ensure reliability and easy maintenance.
