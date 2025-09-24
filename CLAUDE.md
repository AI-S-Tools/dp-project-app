# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

⚠️ **CRITICAL UPDATE (2025-09-24):** v1.1.1 RELEASE CANCELLED due to 12 critical bugs (GitHub issues #20-31). Only 13% of functionality tests pass. DO NOT start new features until all bugs are fixed!

## 🚨 Critical Bugs That Must Be Fixed

**Completely Broken Commands (6):**
1. Issue #20: `dppm init` - Binary path bug (calls ./dppm-test)
2. Issue #26: `dppm bind` - Command doesn't exist
3. Issue #27: `dppm list active` - Command doesn't exist
4. Issue #27: `dppm status active` - Command doesn't exist
5. Issue #24: `dppm list phases` - Command doesn't exist
6. Issue #28: `dppm collab clean` - Not working correctly

**Partially Broken Features (5):**
7. Issue #21: Local binding auto-scoping - Doesn't work
8. Issue #22: Task dependency flag mismatch - Wrong expectations
9. Issue #25: Task create --dependency-ids - Flag doesn't exist
10. Issue #30: Wiki fuzzy search - Not working
11. Issue #29: Wiki incomplete - 9 topics missing

**Low Priority Issues (1):**
12. Issue #23: Local binding project name sync - Cosmetic issue

**Test Results:**
- Success rate: 13% (6/45 tests passing)
- Failed tests: 27 (60%)
- Critical functionality broken: 6 core commands

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

⚠️ **WARNING: Many commands are broken! See GitHub issues #20-31**

```bash
# ✅ These work:
./dppm --version
./dppm --help
./dppm wiki list
./dppm project create test-project --name "Test" --owner "user"
./dppm list projects
./dppm status project test-project

# ❌ These are BROKEN:
./dppm init test-project --skip-github    # Issue #20 - Binary path bug
./dppm bind test-project                  # Issue #26 - Command missing
./dppm list active                        # Issue #27 - Command missing
./dppm status active                      # Issue #27 - Command missing
./dppm list phases                        # Issue #24 - Command missing
./dppm collab clean                       # Issue #28 - Not working

# ⚠️ These have issues:
./dppm task create test --dependency-ids 'other'  # Issue #22,25 - Flag missing

# 📋 Comprehensive test suite (13% pass rate):
# See docs/test/expectation.yaml for complete test results
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
  - `init.go`: Complete project initialization workflow ❌ **BROKEN** (Issue #20)

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
- **Error Handling**: Consistent error messages with context ❌ **FUNDAMENTAL FLAWS** (Issue #31)
- **Path Resolution**: All paths relative to `~/Dropbox/project-management/`
- **YAML Validation**: Schema validation on read/write operations
- **Help System**: Comprehensive help text for all commands with examples

## Development Workflow

⚠️ **CRITICAL: FIX BUGS FIRST BEFORE ANY NEW DEVELOPMENT**

When modifying DPPM:

1. **🚨 PRIORITY: Fix Critical Bugs**: Address GitHub issues #20-31 first
2. **Test Everything**: Run comprehensive tests from docs/test/expectation.yaml
3. **Critical Commands Must Work**: init, bind, list active, status active, list phases, collab clean
4. **Adding New Commands**: Create cobra.Command in appropriate file, add to parent command
5. **Modifying YAML Schemas**: Update structs and ensure backward compatibility
6. **Testing Changes**: Build locally and test with real YAML files
7. **Formatting**: Run `gofmt -w *.go` before committing
8. **Dependencies**: Run `go mod tidy` after adding imports

**❌ DO NOT:**
- Start new features until bugs are fixed
- Create any release tags
- Ignore the 12 critical GitHub issues

## Release Process

🚨 **RELEASE v1.1.1 CANCELLED DUE TO CRITICAL BUGS**

**DO NOT CREATE ANY RELEASE TAGS UNTIL ALL 12 GITHUB ISSUES (#20-31) ARE RESOLVED**

When bugs are fixed, releases are automated via GitHub Actions:

```bash
# ❌ DO NOT RUN UNTIL BUGS FIXED:
git tag v1.1.1
git push origin v1.1.1
```

This triggers `.github/workflows/release.yml` which:
1. Builds binaries for all platforms (Linux, macOS, Windows / amd64, arm64)
2. Creates GitHub release with checksums
3. Uploads all binaries as release assets

**Release Criteria:**
- All 12 GitHub issues resolved
- Test success rate above 90% (currently 13%)
- All critical commands working (init, bind, list active, etc.)
- Comprehensive testing completed