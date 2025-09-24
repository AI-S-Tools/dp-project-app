# GEMINI.md

This file provides guidance to Google Gemini when working with code in this repository.

## Project Overview

**DPPM - Dropbox Project Manager v1.1.0** is an AI-first CLI tool for project, phase, and task management using Dropbox as the storage backend. Perfect for AI-driven development workflows with verbose, AI-friendly output and local project binding.

### DPPM Project Tracking

This project is actively managed using DPPM itself. Current project status:

- **Project ID**: `dp-project-app`
- **Name**: DPPM - Dropbox Project Manager
- **Status**: ⚠️ **CRITICAL ISSUES - v1.1.1 RELEASE CANCELLED**
- **Total Tasks**: 44 (9 completed, 35 ready to start)
- **Health**: 🔴 **12 CRITICAL BUGS** blocking release (see GitHub issues #20-31)
- **Completion**: 20.5% done, but **13% functionality test success rate**

### Current Development Focus

**⚠️ CRITICAL UPDATE (2025-09-24T19:05:00Z):**

Testing revealed **12 critical bugs** that must be fixed before any release:

**🔴 COMPLETELY BROKEN COMMANDS:**
- `dppm init` - Binary path bug (Issue #20)
- `dppm bind` - Command doesn't exist (Issue #26)
- `dppm list active` - Command doesn't exist (Issue #27)
- `dppm status active` - Command doesn't exist (Issue #27)
- `dppm list phases` - Command doesn't exist (Issue #24)
- `dppm collab clean` - Not working correctly (Issue #28)

**🟠 PARTIALLY BROKEN FEATURES:**
- Local binding auto-scoping - Doesn't work (Issue #21)
- Task dependencies - Flag mismatch (Issue #22, #25)
- Wiki fuzzy search - Not working (Issue #30)
- Wiki completeness - 9 topics missing (Issue #29)
- Error handling - Fundamental flaws (Issue #31)

### Key Tasks Ready for Work

**🚨 MUST FIX BEFORE ANY RELEASE:**
1. Fix init command binary path bug (Issue #20)
2. Implement missing `dppm bind` command (Issue #26)
3. Add missing `dppm list active` command (Issue #27)
4. Add missing `dppm status active` command (Issue #27)
5. Add missing `dppm list phases` command (Issue #24)
6. Fix `dppm collab clean` (Issue #28)
7. Fix local binding auto-scoping (Issue #21)
8. Fix task dependency flags (Issues #22, #25)
9. Fix wiki fuzzy search (Issue #30)
10. Add 9 missing wiki topics (Issue #29)
11. Fix fundamental error handling (Issue #31)
12. Fix local binding project name sync (Issue #23)

**⚠️ DO NOT START NEW FEATURES** until critical bugs are fixed!

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

# ⚠️ WARNING: Many commands are broken! See issues #20-31

# Test with DPPM project management
./dppm-test status project dp-project-app  # ✅ Works
./dppm-test task create new-feature --project dp-project-app --phase core-development  # ✅ Works

# Test local binding functionality (new in v1.1.0)
go build -o dppm-test-local
./dppm-test-local bind dp-project-app  # ⚠️ Auto-scoping broken (Issue #21)
./dppm-test-local status project dp-project-app  # ❌ Still requires --project flag

# Always test help systems and wiki
./dppm-test --help  # ✅ Works
./dppm-test wiki "getting started"  # ✅ Works
./dppm-test wiki list  # ✅ Works

# Test comprehensive expectations (available in docs/test/)
# ⚠️ CRITICAL: See docs/test/expectation.yaml - Only 13% test success rate!
# See docs/test/wiki.md for all wiki texts and help content
```

## Release Process

**⚠️ RELEASE v1.1.1 CANCELLED DUE TO CRITICAL BUGS**

When bugs are fixed, releases are automated via GitHub Actions:
- Tag format: `v1.0.0`
- Multi-platform binaries: Linux, macOS, Windows (amd64, arm64)
- Homebrew tap: `AI-S-Tools/dp-project-app`

**DO NOT CREATE ANY RELEASE TAGS UNTIL ALL 12 ISSUES ARE RESOLVED!**

## Integration with MCP (High Priority)

The project includes active development of MCP (Model Context Protocol) server integration to enable seamless AI assistant access to project data. This will allow real-time project context sharing with AI development tools.

**Current Status:** Designed and ready for implementation as high priority task.

## New Features in v1.1.0

### Local Project Binding
- **Status:** ❌ **BROKEN** - Multiple critical issues
- **Feature:** `.dppm/project.yaml` binding system
- **Command:** `dppm bind PROJECT_ID` - **DOESN'T EXIST** (Issue #26)
- **Benefits:** Would eliminate need for `--project` flags if working
- **Known Issues:**
  - `dppm bind` command missing entirely (Issue #26)
  - Auto-scoping completely broken (Issue #21)
  - Project name not synced from Dropbox (Issue #23)

### Comprehensive Testing Framework
- **Status:** ✅ Test framework complete, ❌ Tests failing
- **Files:**
  - `docs/test/expectation.yaml` - 45+ comprehensive test cases (**13% pass rate**)
  - `docs/test/wiki.md` - Complete collection of all wiki texts and help texts
- **Test Results (2025-09-24):**
  - **6 tests passing** (13%)
  - **27 tests failing** (60%)
  - **12 GitHub issues created** for critical bugs
- **Coverage:** Init workflow, project lifecycle, stability tests, all wiki functions

### AI Collaboration System
- **Status:** ⚠️ Partially working
- **Features:** DSL markers, collaboration detection, wiki integration
- **Commands:**
  - `dppm collab find` - ✅ Works
  - `dppm collab wiki` - ✅ Works
  - `dppm collab clean` - ❌ **BROKEN** (Issue #28)

## Contributing Guidelines

**⚠️ PRIORITY: FIX THE 12 CRITICAL BUGS BEFORE ANY NEW DEVELOPMENT**

1. **Fix Bugs First**: Address GitHub issues #20-31 before new features
2. **Test Everything**: Run comprehensive tests from `docs/test/expectation.yaml`
3. **Follow Go Patterns**: Consistent with existing cobra command structure
4. **AI-First Design**: All features should enhance AI development workflows
5. **Comprehensive Help**: Every command needs detailed help and examples
6. **YAML-Based Storage**: Maintain the Dropbox-synchronized YAML approach

**Critical Test Commands Before Any PR:**
```bash
# These MUST work without errors:
./dppm-test init test-project --skip-github
./dppm-test bind test-project
./dppm-test list active
./dppm-test status active
./dppm-test list phases
./dppm-test collab clean
```

For detailed implementation guidance, see `CLAUDE.md` for Claude-specific instructions.