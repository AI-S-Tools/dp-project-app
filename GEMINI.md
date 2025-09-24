# GEMINI.md

This file provides guidance to Google Gemini when working with code in this repository.

## Project Overview

**DPPM - Dropbox Project Manager v1.1.0** is an AI-first CLI tool for project, phase, and task management using Dropbox as the storage backend. Perfect for AI-driven development workflows with verbose, AI-friendly output and local project binding.

### DPPM Project Tracking

This project is actively managed using DPPM itself. Current project status:

- **Project ID**: `dp-project-app`
- **Name**: DPPM - Dropbox Project Manager
- **Status**: Active Development
- **Total Tasks**: 44 (9 completed, 35 ready to start)
- **Health**: ✅ No blocked tasks, excellent project health
- **Completion**: 20.5% done, 79.5% ready for implementation

### Current Development Focus

**6 Active Phases:**
1. **Core Development** - ✅ **COMPLETE** - All basic project/phase/task management functionality implemented
2. **AI Collaboration System** - ✅ **CORE COMPLETE** - DSL markers, collaboration detection, wiki integration done
3. **Feature Enhancements** - 📋 **PLANNED** - 20+ enhancement features ready for implementation
4. **Bug Fixes** - 🔴 **4 CRITICAL BUGS** - Init command, local binding issues identified
5. **Deployment** - ✅ **PRODUCTION READY** - Multi-platform binaries, GitHub Actions, Homebrew tap
6. **Milestones Extension** - 📋 **PLANNED** - Core milestone system designed

### Key Tasks Ready for Work

**🔴 Critical Priority:**
- Fix init command binary path bug (calls ./dppm-test instead of dppm)

**🟠 High Priority:**
- AI Rules System with hierarchical rule management
- DPPM MCP Server for AI integration
- Core Milestone System implementation
- Local binding auto-scoping fix
- Advanced dependency management system

**🟡 Medium Priority:**
- Missing list subcommands (phases, tasks)
- Missing phase subcommands (list, show, update)
- Missing task subcommands (list, component, issue, dependency)
- Enhanced status reporting and visualization
- Task templates and bulk operations
- Git integration workflow
- Phase numbering system (fase-1, fase-2, etc.)

## Build and Development

### Quick Commands

```bash
# Build the application
go build -o dppm

# Test basic functionality
./dppm --version
./dppm list projects
./dppm status project dp-project-app

# Format code
gofmt -w *.go

# Install dependencies
go mod tidy
```

### Project Structure

This is a Go-based CLI application using:
- `github.com/spf13/cobra` for CLI commands
- `gopkg.in/yaml.v3` for YAML parsing
- Standard library for file handling

**Key Files:**
- `main.go`: CLI entry point and root command
- `project.go`: Project CRUD operations
- `task.go`: Task management with dependencies
- `sprint.go`: Phase management (legacy naming)
- `status.go`: Status reporting across all levels
- `wiki.go`: Built-in knowledge base system
- `init.go`: Complete project initialization workflow

### Data Storage

All data is stored as YAML files in:
```
~/Dropbox/project-management/projects/
├── dp-project-app/
│   ├── project.yaml
│   └── phases/
│       ├── core-development/
│       ├── ai-collaboration/
│       ├── enhancements/
│       ├── deployment/
│       ├── bug-fixes/
│       └── milestones-extension/
```

## AI Development Guidelines

### When Working with DPPM Code

1. **Check Project Status First**: Always run `dppm status project dp-project-app` to understand current priorities
2. **Update Task Status**: When working on tasks, update their status in DPPM using `dppm task update`
3. **Follow Patterns**: Study existing commands in `project.go`, `task.go` etc. for consistent patterns
4. **Test with Real Data**: Build locally and test with actual DPPM projects
5. **AI-Friendly Output**: All commands should provide verbose, helpful output for AI workflows

### Common Development Patterns

```go
// Standard command structure
var createProjectCmd = &cobra.Command{
    Use:   "create PROJECT_ID",
    Short: "Create a new project",
    Long:  `Detailed description with examples...`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation with error handling
        // Verbose success/failure messages
    },
}
```

### Testing Approach

```bash
# Build test binary
go build -o dppm-test

# Test with DPPM project management
./dppm-test status project dp-project-app
./dppm-test task create new-feature --project dp-project-app --phase core-development

# Test local binding functionality (new in v1.1.0)
go build -o dppm-test-local
./dppm-test-local bind dp-project-app
./dppm-test-local status project dp-project-app

# Always test help systems and wiki
./dppm-test --help
./dppm-test wiki "getting started"
./dppm-test wiki list

# Test comprehensive expectations (available in docs/test/)
# See docs/test/expectation.yaml for complete test specification
# See docs/test/wiki.md for all wiki texts and help content
```

## Release Process

Releases are automated via GitHub Actions:
- Tag format: `v1.0.0`
- Multi-platform binaries: Linux, macOS, Windows (amd64, arm64)
- Homebrew tap: `AI-S-Tools/dp-project-app`

## Integration with MCP (High Priority)

The project includes active development of MCP (Model Context Protocol) server integration to enable seamless AI assistant access to project data. This will allow real-time project context sharing with AI development tools.

**Current Status:** Designed and ready for implementation as high priority task.

## New Features in v1.1.0

### Local Project Binding
- **Status:** ✅ Implemented, testing in progress
- **Feature:** `.dppm/project.yaml` binding system
- **Command:** `dppm bind PROJECT_ID`
- **Benefits:** Eliminates need for `--project` flags, prevents cross-project task creation
- **Known Issues:** Auto-scoping not fully integrated in all commands

### Comprehensive Testing Framework
- **Status:** ✅ Complete
- **Files:**
  - `docs/test/expectation.yaml` - 45+ comprehensive test cases
  - `docs/test/wiki.md` - Complete collection of all wiki texts and help texts
- **Coverage:** Init workflow, project lifecycle, stability tests, all wiki functions

### AI Collaboration System
- **Status:** ✅ Core features complete
- **Features:** DSL markers, collaboration detection, wiki integration
- **Commands:** `dppm collab find`, `dppm collab wiki`

## Contributing Guidelines

1. **Use DPPM**: Track your work in the DPPM project itself
2. **Follow Go Patterns**: Consistent with existing cobra command structure
3. **AI-First Design**: All features should enhance AI development workflows
4. **Comprehensive Help**: Every command needs detailed help and examples
5. **YAML-Based Storage**: Maintain the Dropbox-synchronized YAML approach

For detailed implementation guidance, see `CLAUDE.md` for Claude-specific instructions.