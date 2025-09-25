# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

<<<<<<< HEAD
## Build Commands

```bash
# Build the application
=======
## Project Setup Check

**IMPORTANT:** Always check for `.dppm/project.yaml` at startup to understand the local project binding and configuration. This file contains critical project metadata including GitHub repository, Dropbox path, and current issues.

## Current Open GitHub Issues

### High Priority Architecture Issues
- **#36**: Project tree management - Phases need numbering, tasks need structure, bugs should be in BUGS phase
- **#35**: Task descriptions need 100-400 lines for AI development (anti-DRY for project management)
- **#34**: Building projects with tasks/milestones is iterative
- **#33**: Missing task numbering and sequential ordering best practices
- **#32**: GitHub issues integration needed

### Medium Priority Issues
- **#23**: Local binding doesn't load project name from Dropbox
- **#22**: Task dependency flag mismatch in expectations
- **#21**: Local binding auto-scoping doesn't work

### Recently Closed Critical Bugs (Verify fixes)
- **#20**: `dppm init` - Binary path bug (was calling ./dppm-test)
- **#26**: `dppm bind` - Command was missing
- **#27**: `dppm list active` & `dppm status active` - Commands were missing
- **#24**: `dppm list phases` - Command was missing
- **#28**: `dppm collab clean` - Was not working correctly
- **#25**: Task dependencies - Flag `--dependency-ids` was missing
- **#31**: Fundamental error handling flaws
- **#30**: Wiki fuzzy search not working
- **#29**: Wiki incomplete (9 topics missing)

## Build & Development Commands

```bash
# Build the binary
>>>>>>> b1654786e04937704748bf855514fc428e4dd828
go build -o dppm

# Build with version tag
go build -ldflags="-s -w -X main.version=v1.0.0" -o dppm

<<<<<<< HEAD
# Build for different platforms
=======
# Cross-platform builds
>>>>>>> b1654786e04937704748bf855514fc428e4dd828
GOOS=linux GOARCH=amd64 go build -o dppm-linux-amd64
GOOS=darwin GOARCH=arm64 go build -o dppm-macos-arm64
GOOS=windows GOARCH=amd64 go build -o dppm-windows-amd64.exe

<<<<<<< HEAD
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
=======
# Format code before committing
gofmt -w *.go

# Update dependencies
go mod tidy

# Run type checking and linting
go vet ./...
```

## Architecture & Command Structure

### Core Architecture
DPPM uses a hierarchical command structure built with Cobra. Each major feature area has its own file:

- **Entry Point**: `main.go` - Root command setup with version handling
- **Command Files**: Each file implements a subcommand tree
  - `init.go`: Project initialization workflow (BROKEN - #20)
  - `project.go`: Project CRUD operations
  - `task.go`: Task management with components/issues
  - `phase.go`: Phase (sprint) management
  - `status.go`: Multi-level status reporting
  - `list.go`: List entities at all levels
  - `wiki.go`: Built-in documentation system
  - `collab.go`: AI collaboration features

### Adding New Commands
1. Create cobra.Command in appropriate file or new file
2. Add to parent command using `parentCmd.AddCommand(newCmd)`
3. Follow existing pattern for flags and help text
4. Update wiki topics if adding major feature

### Data Storage Pattern
- **Base Path**: `~/Dropbox/project-management/projects/`
- **Structure**: Projects → Phases → Tasks → Components/Issues
- **Format**: YAML files with specific schemas per entity type
- **Templates**: Embedded templates for entity creation
- **Path Resolution**: All paths relative to base, use `getProjectPath()` helper

### YAML Schema Modifications
When changing schemas:
1. Update struct definitions in relevant file
2. Ensure backward compatibility
3. Test with existing YAML files
4. Update template files if needed

## Testing Approach

```bash
# Comprehensive test (see docs/test/expectation.yaml)
./dppm-test --version
./dppm-test project create test --name "Test" --owner "user"
./dppm-test list projects

# Test broken commands to verify fixes
./dppm-test init test-project  # Should NOT call ./dppm-test
./dppm-test bind test-project  # Should work
./dppm-test list active  # Should work
./dppm-test collab clean  # Should clean properly
```

## Release Process

**DO NOT CREATE RELEASES UNTIL ALL GITHUB ISSUES #20-31 ARE RESOLVED**

Once fixed:
```bash
git tag v1.1.1
git push origin v1.1.1
# GitHub Actions will build and release all platform binaries
```

## Key Functions & Entry Points

- `initProjectCmd()`: Complete project initialization (fix binary path)
- `createTaskCmd()`: Task creation with dependency validation
- `getProjectPath()`: Resolves project storage paths
- `readProjectYAML()`/`writeProjectYAML()`: Core YAML I/O
- `bindCmd()`: Local binding command (implement missing)
- `listActiveCmd()`: Active items listing (implement missing)

## Development Priorities

1. Fix all 12 critical bugs first
2. Ensure 90%+ test pass rate
3. Verify all core commands work
4. Then consider new features
>>>>>>> b1654786e04937704748bf855514fc428e4dd828
