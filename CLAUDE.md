# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands

```bash
# Build the application
go build -o dppm

# Build with version tag
go build -ldflags="-s -w -X main.version=v1.0.0" -o dppm

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o dppm-linux-amd64
GOOS=darwin GOARCH=arm64 go build -o dppm-macos-arm64
GOOS=windows GOARCH=amd64 go build -o dppm-windows-amd64.exe

# Format code
gofmt -w *.go

# Install dependencies
go mod tidy
```

## Testing Commands

```bash
# Test the binary after building
./dppm --version
./dppm --help
./dppm wiki list

# Test specific commands
./dppm project create test-project --name "Test" --owner "user"
./dppm list projects
./dppm status project test-project
```

## Project Architecture

DPPM is a CLI tool for managing projects using YAML files stored in Dropbox. The architecture consists of:

### Command Structure
- **Root Command (`main.go`)**: Entry point and root cobra command setup with global flags
- **Subcommands**: Each feature area has its own file with cobra commands:
  - `project.go`: Project CRUD operations
  - `task.go`: Task management with components, issues, and dependencies
  - `sprint.go`: Sprint planning and management (called "phases" in newer versions)
  - `status.go`: Status reporting across all levels
  - `list.go`: List projects, tasks, and other entities
  - `wiki.go`: Built-in knowledge base with 20+ help topics
  - `collab.go`: AI collaboration features with DSL markers
  - `init.go`: Complete project initialization workflow

### Data Model & Storage
- **Storage Path**: `~/Dropbox/project-management/projects/`
- **Hierarchical Structure**: Projects → Phases (Sprints) → Tasks → Components/Issues
- **YAML Schema**: All data stored as YAML files with specific schemas for each entity type
- **Template System**: Uses embedded templates for creating new entities

### Key Design Patterns
1. **Command Pattern**: Every operation is a cobra.Command with consistent flags and help
2. **File-based Storage**: No database; all data in YAML files for easy version control
3. **Hierarchical Organization**: Mirrors project structure in filesystem
4. **Verbose Output**: AI-friendly output with clear status messages and emojis
5. **Dependency Management**: Tasks can have dependencies with automatic blocking detection

### Important Functions & Entry Points
- `main()` in `main.go`: CLI entry point
- `initProjectCmd()` in `init.go`: Complete project initialization workflow
- `createProjectCmd()` in `project.go`: Project creation logic
- `createTaskCmd()` in `task.go`: Task creation with dependency validation
- `getProjectPath()`: Helper to resolve project storage paths
- `readProjectYAML()`/`writeProjectYAML()`: Core YAML I/O functions

### Cross-cutting Concerns
- **Error Handling**: Consistent error messages with context
- **Path Resolution**: All paths relative to `~/Dropbox/project-management/`
- **YAML Validation**: Schema validation on read/write operations
- **Help System**: Comprehensive help text for all commands with examples

## Development Workflow

When modifying DPPM:

1. **Adding New Commands**: Create cobra.Command in appropriate file, add to parent command
2. **Modifying YAML Schemas**: Update structs and ensure backward compatibility
3. **Testing Changes**: Build locally and test with real YAML files
4. **Formatting**: Run `gofmt -w *.go` before committing
5. **Dependencies**: Run `go mod tidy` after adding imports

## Release Process

Releases are automated via GitHub Actions when pushing tags:

```bash
git tag v1.0.0
git push origin v1.0.0
```

This triggers `.github/workflows/release.yml` which:
1. Builds binaries for all platforms (Linux, macOS, Windows / amd64, arm64)
2. Creates GitHub release with checksums
3. Uploads all binaries as release assets