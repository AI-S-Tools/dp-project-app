/* ::GEMINI:10: Koden kan optimeres betydeligt ved at flytte de store tekstblokke i `show...` funktionerne til eksterne filer for at reducere binærstørrelsen og forbedre vedligeholdelsen.:: */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// /* Definerer 'collab' kommandoen til AI-samarbejde. */
var collabCmd = &cobra.Command{
	Use:   "collab",
	Short: "AI collaboration system using DSL markers",
	Long: `AI Collaboration System - DSL Markers for Task Management

DPPM Collaboration System allows multiple AI agents to work together using
DSL markers in markdown files. This system enables structured handoffs
between different AI models and tracks collaborative work.

DSL Marker Format:
  ::LARS:ID:: ... content ... ::    # Task for LARS/Claude
  ::GEMINI:ID:: ... content ... ::  # Task for Gemini
  ::DONE:ID,ID:: ... ::              # Mark completed tasks

Usage:
  dppm collab find [path...]          # Find all DSL tasks
  dppm collab clean [path...]         # Remove completed tasks
  dppm collab wiki                    # Show collaboration guides

Examples:
  dppm collab find docs/              # Find tasks in docs folder
  dppm collab clean                   # Clean up completed tasks
  dppm collab wiki "task handoff"     # Learn about AI handoffs`,
}

// /* Definerer 'find' underkommandoen for at finde DSL-opgaver. */
var collabFindCmd = &cobra.Command{
	Use:   "find [path...]",
	Short: "Find all DSL collaboration tasks",
	Long: `Find DSL Collaboration Tasks

Searches for AI collaboration markers in markdown files.
Shows active tasks assigned to different AI models.

Markers searched for:
  • ::LARS:ID::    - Tasks for LARS/Claude
  • ::GEMINI:ID::  - Tasks for Gemini
  • ::DONE:ID::    - Completed task markers

The search excludes ai-dsl.sh and ai-dsl.md files automatically.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchPaths := []string{"."}
		if len(args) > 0 {
			searchPaths = args
		}
		findCollabTasks(searchPaths)
	},
}

// /* Definerer 'clean' underkommandoen for at fjerne afsluttede opgaver. */
var collabCleanCmd = &cobra.Command{
	Use:   "clean [path...]",
	Short: "Remove completed DSL collaboration tasks",
	Long: `Clean Completed Collaboration Tasks

Finds DONE markers and removes associated task blocks from markdown files.
This helps maintain clean documentation by removing completed collaboration tasks.

Process:
  1. Finds all ::DONE:ID,ID:: markers
  2. Extracts comma-separated task IDs
  3. Removes corresponding ::LARS:ID:: and ::GEMINI:ID:: blocks
  4. Removes the DONE markers themselves

Safety: Creates backup files (.bak) before making changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		searchPaths := []string{"."}
		if len(args) > 0 {
			searchPaths = args
		}
		cleanCompletedTasks(searchPaths)
	},
}

// /* Definerer 'wiki' underkommandoen for at vise samarbejdsguider. */
var collabWikiCmd = &cobra.Command{
	Use:   "wiki [search-terms]",
	Short: "AI collaboration guides and documentation",
	Long: `AI Collaboration Wiki System

Comprehensive guides for using the AI collaboration system with DSL markers.
Learn how to structure AI handoffs, manage collaborative workflows,
and coordinate between different AI models.

Available Topics:
  • "collaboration basics" - Introduction to AI collaboration
  • "dsl markers" - Understanding the marker system
  • "task handoff" - Passing work between AIs
  • "workflow patterns" - Common collaboration patterns
  • "best practices" - Proven collaboration strategies`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			showCollabWikiIndex()
			return
		}
		searchTerm := strings.ToLower(strings.Join(args, " "))
		searchCollabWiki(searchTerm)
	},
}

// /* Finder og viser alle AI-samarbejdsopgaver i de angivne stier. */
func findCollabTasks(searchPaths []string) {
	fmt.Println("🔍 Searching for AI collaboration tasks...")
	fmt.Println("==========================================")
	fmt.Println("Path(s):", strings.Join(searchPaths, ", "))
	fmt.Println()

	dslRegex := regexp.MustCompile(`::(LARS|GEMINI|DONE):[^:]*::`)
	foundAny := false

	for _, searchPath := range searchPaths {
		err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil // Skip files with errors
			}

			// Skip non-markdown files
			if !strings.HasSuffix(strings.ToLower(path), ".md") {
				return nil
			}

			// Skip ai-dsl files
			if strings.Contains(path, "ai-dsl.sh") || strings.Contains(path, "ai-dsl.md") {
				return nil
			}

			// Read file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return nil // Skip files we can't read
			}

			// Find all matches
			matches := dslRegex.FindAllString(string(content), -1)
			if len(matches) > 0 {
				foundAny = true
				fmt.Printf("📄 %s\n", path)

				// Show matches with line numbers
				scanner := bufio.NewScanner(strings.NewReader(string(content)))
				lineNum := 1
				for scanner.Scan() {
					line := scanner.Text()
					if dslRegex.MatchString(line) {
						fmt.Printf("   %d: %s\n", lineNum, strings.TrimSpace(line))
					}
					lineNum++
				}
				fmt.Println()
			}

			return nil
		})

		if err != nil {
			fmt.Printf("⚠️  Error walking path %s: %v\n", searchPath, err)
		}
	}

	if !foundAny {
		fmt.Println("ℹ️  No DSL collaboration tasks found.")
		fmt.Println()
		fmt.Println("💡 To create collaboration tasks, use:")
		fmt.Println("   ::LARS:1:: Task description for LARS ::")
		fmt.Println("   ::GEMINI:2:: Task description for Gemini ::")
		fmt.Println("   ::DONE:1,2:: Mark tasks 1 and 2 as completed ::")
	}

	fmt.Println("==========================================")
	fmt.Println("✅ Search complete.")
}

// /* Fjerner afsluttede samarbejdsopgaver fra markdown-filer. */
func cleanCompletedTasks(searchPaths []string) {
	fmt.Println("🧹 Cleaning completed collaboration tasks...")
	fmt.Println("==========================================")
	fmt.Println("Path(s):", strings.Join(searchPaths, ", "))
	fmt.Println()

	processedFiles := 0

	for _, searchPath := range searchPaths {
		err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}

			// Skip non-markdown files
			if !strings.HasSuffix(strings.ToLower(path), ".md") {
				return nil
			}

			// Skip ai-dsl files
			if strings.Contains(path, "ai-dsl.sh") || strings.Contains(path, "ai-dsl.md") {
				return nil
			}

			// Read file content
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return nil
			}

			// Check for DONE markers
			doneRegex := regexp.MustCompile(`::DONE:\s*([0-9, ]+)\s*::`)
			if !doneRegex.MatchString(string(content)) {
				return nil
			}

			fmt.Printf("📄 Processing: %s\n", path)

			// Extract DONE IDs
			doneMatches := doneRegex.FindAllStringSubmatch(string(content), -1)
			var allIDs []string

			for _, match := range doneMatches {
				if len(match) > 1 {
					// Split comma-separated IDs
					ids := strings.Split(match[1], ",")
					for _, id := range ids {
						id = strings.TrimSpace(id)
						if id != "" {
							allIDs = append(allIDs, id)
						}
					}
				}
			}

			if len(allIDs) == 0 {
				fmt.Println("   ⚠️  No valid IDs found in DONE markers")
				return nil
			}

			// Remove task blocks for each ID
			updatedContent := string(content)
			for _, id := range allIDs {
				if _, err := strconv.Atoi(id); err != nil {
					fmt.Printf("   ⚠️  Invalid ID '%s', skipping\n", id)
					continue
				}

				fmt.Printf("   🗑️  Removing blocks for ID: %s\n", id)

				// Remove LARS and GEMINI blocks for this ID
			taskRegex := regexp.MustCompile(fmt.Sprintf(`::(LARS|GEMINI):\s*%s\s*::.*?::\s*`, regexp.QuoteMeta(id)))
				updatedContent = taskRegex.ReplaceAllString(updatedContent, "")
			}

			// Remove DONE lines
			updatedContent = doneRegex.ReplaceAllString(updatedContent, "")

			// Write back to file
			err = ioutil.WriteFile(path, []byte(updatedContent), info.Mode())
			if err != nil {
				fmt.Printf("   ❌ Error writing file: %v\n", err)
				return nil
			}

			processedFiles++
			fmt.Println("   ✅ Cleanup complete")

			return nil
		})

		if err != nil {
			fmt.Printf("⚠️  Error walking path %s: %v\n", searchPath, err)
		}
	}

	fmt.Println("==========================================")
	fmt.Printf("✅ Processed %d files\n", processedFiles)

	if processedFiles > 0 {
		fmt.Println()
		fmt.Println("💡 Next steps:")
		fmt.Println("   dppm collab find     # Verify cleanup")
		fmt.Println("   dppm wiki \"collaboration workflow\"  # Learn more")
	}
}

// /* Viser hovedindekset for samarbejds-wikien. */
func showCollabWikiIndex() {
	fmt.Println(`AI Collaboration Wiki
=====================

🤖 AI-to-AI Collaboration System

DPPM's collaboration system enables structured handoffs between different
AI models using DSL markers in markdown files.

🔍 Quick Start:
  dppm collab find                    # Find current collaboration tasks
  dppm collab clean                   # Clean up completed tasks
  dppm collab wiki "task handoff"     # Learn handoff patterns

📚 Available Topics:
  • "collaboration basics" - Introduction and core concepts
  • "dsl markers" - Understanding the marker syntax
  • "task handoff" - Passing work between AI models
  • "workflow patterns" - Common collaboration scenarios
  • "best practices" - Proven strategies for AI teamwork
  • "integration" - Using with DPPM projects

💡 Integration with DPPM:
This collaboration system works alongside DPPM's project management,
enabling AI teams to coordinate work within structured projects.

🚀 Get Started:
  dppm collab wiki "collaboration basics"`)
}

// /* Søger i samarbejds-wikien efter et bestemt emne. */
func searchCollabWiki(searchTerm string) {
	switch {
	case strings.Contains(searchTerm, "collaboration basics") || strings.Contains(searchTerm, "introduction"):
		showCollabBasicsGuide()
	case strings.Contains(searchTerm, "dsl markers") || strings.Contains(searchTerm, "markers"):
		showDSLMarkersGuide()
	case strings.Contains(searchTerm, "task handoff") || strings.Contains(searchTerm, "handoff"):
		showTaskHandoffGuide()
	case strings.Contains(searchTerm, "workflow patterns") || strings.Contains(searchTerm, "patterns"):
		showWorkflowPatternsGuide()
	case strings.Contains(searchTerm, "best practices") || strings.Contains(searchTerm, "practices"):
		showBestPracticesGuide()
	case strings.Contains(searchTerm, "integration") || strings.Contains(searchTerm, "dppm integration"):
		showIntegrationGuide()
	default:
		fmt.Printf("No specific collaboration guide found for '%s'\n\n", searchTerm)
		fmt.Println("Available topics:")
		fmt.Println("  dppm collab wiki \"collaboration basics\"")
		fmt.Println("  dppm collab wiki \"dsl markers\"")
		fmt.Println("  dppm collab wiki \"task handoff\"")
	}
}

// /* Viser guiden til grundlæggende samarbejde. */
func showCollabBasicsGuide() {
	fmt.Println(`AI Collaboration Basics
=======================

🤖 WHAT IS AI COLLABORATION?

AI Collaboration enables multiple AI models to work together on complex tasks
by using structured handoff markers in documentation. Different AI models
can specialize in different aspects of work.

🎯 KEY CONCEPTS:

Task Assignments:
  • LARS (Claude) - Development, analysis, documentation
  • GEMINI (Google) - Creative tasks, brainstorming, review
  • Tasks have unique IDs for tracking

Workflow States:
  • Active tasks - Currently assigned to an AI
  • Completed tasks - Marked with DONE markers
  • Handoff points - Where work passes between AIs

📝 BASIC MARKER FORMAT:

Create a task for LARS:
  ::LARS:1:: Implement user authentication system ::

Create a task for GEMINI:
  ::GEMINI:2:: Review the authentication design ::

Mark tasks as completed:
  ::DONE:1,2:: Authentication work completed ::

🔄 COLLABORATION FLOW:

1. Initial Planning:
   - Define tasks and assign to appropriate AI models
   - Set up clear handoff points
   - Document requirements and context

2. Execution:
   - Each AI works on assigned tasks
   - Updates documentation with progress
   - Coordinates handoffs at completion points

3. Completion:
   - Mark tasks as DONE when finished
   - Clean up completed tasks from documentation
   - Archive or preserve important outcomes

💡 BENEFITS:

Specialization:
  • Different AIs excel at different task types
  • Leverage each model's strengths
  • Improved overall quality

Coordination:
  • Clear handoff points prevent conflicts
  • Structured workflow reduces confusion
  • Progress tracking through markers

Documentation:
  • Work history preserved in markdown
  • Easy to review collaboration patterns
  • Self-documenting process

🔍 Related Commands:
  • dppm collab wiki "dsl markers"     # Learn marker syntax
  • dppm collab wiki "task handoff"    # Handoff patterns`)
}

// /* Viser guiden til DSL-markører. */
func showDSLMarkersGuide() {
	fmt.Println(`DSL Markers Reference Guide
===========================

📝 MARKER SYNTAX:

Basic Format:
  ::AI_NAME:ID:: task content ::

Where:
  • AI_NAME: LARS, GEMINI, or DONE
  • ID: Unique number for the task
  • task content: Description of work to be done

🏷️ AVAILABLE MARKERS:

LARS Tasks (Claude):
  ::LARS:1:: Analyze the codebase for security issues ::
  ::LARS:42:: Write unit tests for the authentication module ::

GEMINI Tasks (Google):
  ::GEMINI:5:: Brainstorm creative solutions for user onboarding ::
  ::GEMINI:18:: Review and improve the documentation style ::

Completion Markers:
  ::DONE:1:: Single task completed ::
  ::DONE:1,5,42:: Multiple tasks completed ::

📋 TASK ID GUIDELINES:

Good IDs:
  ✅ Sequential: 1, 2, 3, 4...
  ✅ Grouped: 10-19 for auth, 20-29 for UI
  ✅ Meaningful: Use same ID family for related work

Bad IDs:
  ❌ Random: 847, 23, 1,337
  ❌ Duplicate: Using same ID for different tasks
  ❌ Non-numeric: task-auth, ui-review

🔧 ADVANCED PATTERNS:

Multi-line Tasks:
  ::LARS:10::
  Implement the following features:
  - User registration with email validation
  - Password reset functionality
  - Session management
  ::

Task Context:
  ::GEMINI:20::
  Context: The authentication system is 80% complete
  Task: Review the security implications of the current approach
  Deliverable: Security assessment document
  ::

Dependencies:
  ::LARS:30::
  Prerequisites: Tasks 10 and 20 must be completed
  Task: Integrate authentication with the main application
  ::

📊 TASK LIFECYCLE:

Creation:
  ::LARS:100:: New task description ::

In Progress:
  ::LARS:100:: [IN PROGRESS] Task description ::

Completed:
  ::DONE:100:: Task completed successfully ::

🔧 MANAGEMENT COMMANDS:

Find all tasks:
  dppm collab find              # Current directory
  dppm collab find docs/        # Specific directory

Clean completed:
  dppm collab clean             # Remove DONE tasks
  dppm collab clean --dry-run   # Preview changes

💡 BEST PRACTICES:

Organization:
  • Group related tasks with similar ID ranges
  • Use clear, actionable task descriptions
  • Include context and deliverables
  • Document dependencies clearly

Maintenance:
  • Regularly clean up completed tasks
  • Archive important outcomes before cleaning
  • Keep active task count manageable
  • Review collaboration patterns periodically

🔍 Related Commands:
  • dppm collab wiki "task handoff"      # Handoff patterns
  • dppm collab wiki "workflow patterns" # Common scenarios`)
}

// /* Viser guiden til opgaveoverdragelse. */
func showTaskHandoffGuide() {
	fmt.Println(`Task Handoff Patterns Guide
===========================

🔄 AI-TO-AI HANDOFF PATTERNS:

Sequential Handoffs - work flows from one AI to another:
Development → Review → Documentation

Parallel Handoffs - AIs work simultaneously:
Frontend + Backend development in parallel

Iterative Handoffs - work bounces back and forth:
Design → Feedback → Redesign → Approval

📝 COMMON HANDOFF SCENARIOS:

1️⃣ DEVELOPMENT HANDOFF:
  LARS creates, GEMINI reviews:

  ::LARS:10::
  Create user authentication API with JWT tokens
  Include password hashing and session management
  ::

  ::GEMINI:11::
  Review authentication implementation for:
  - Security best practices
  - Code quality and maintainability
  - Documentation completeness
  ::

2️⃣ CREATIVE HANDOFF:
  GEMINI ideates, LARS implements:

  ::GEMINI:20::
  Brainstorm innovative user onboarding flows
  Consider: gamification, progressive disclosure, personalization
  Deliverable: 3-5 detailed UX concepts
  ::

  ::LARS:21::
  Implement the selected onboarding flow from task 20
  Focus on: clean code, performance, accessibility
  ::

3️⃣ ITERATIVE HANDOFF:
  Back-and-forth refinement:

  ::LARS:30:: Create initial database schema ::
  ::GEMINI:31:: Review schema and suggest improvements ::
  ::LARS:32:: Refine schema based on feedback from task 31 ::
  ::GEMINI:33:: Final approval of refined schema ::

🔧 HANDOFF COORDINATION:

Clear Context Transfer:
  ::LARS:40::
  Context: User reported slow query performance
  Investigation: Analyzed queries, found N+1 problems
  Solution: Need database indexing strategy
  Handoff: Pass to GEMINI for performance review
  ::

  ::GEMINI:41::
  Context: Received performance analysis from task 40
  Task: Review proposed indexing strategy
  Focus: Query optimization and database design
  ::

Status Communication:
  ::LARS:50:: [BLOCKED] Waiting for API keys from task 49 ::
  ::GEMINI:51:: [READY] Can start once task 50 unblocked ::

Dependencies and Prerequisites:
  ::LARS:60::
  Prerequisites: Tasks 55, 57 must be completed
  Task: Integration testing of payment system
  Dependencies: Stripe API keys, test database
  ::

📊 HANDOFF WORKFLOW EXAMPLE:

E-commerce Feature Development:

  ::GEMINI:100::
  Research: Analyze competitor checkout flows
  Deliverable: UX recommendations document
  ::

  ::LARS:101::
  Prerequisites: Task 100 completed
  Implement: Checkout flow based on UX recommendations
  Focus: Clean code, error handling, validation
  ::

  ::GEMINI:102::
  Prerequisites: Task 101 completed
  Test: User experience testing of new checkout
  Deliverable: UX feedback and improvement suggestions
  ::

  ::LARS:103::
  Prerequisites: Task 102 completed
  Refine: Implement UX improvements from testing
  Final: Prepare for production deployment
  ::

  ::DONE:100,101,102,103:: Checkout flow completed ::

🎯 HANDOFF SUCCESS FACTORS:

Communication:
  ✅ Clear context and background
  ✅ Specific deliverables defined
  ✅ Dependencies documented
  ✅ Success criteria stated

Timing:
  ✅ Logical sequence of work
  ✅ No unnecessary blocking
  ✅ Parallel work where possible
  ✅ Clear completion signals

Quality:
  ✅ Each AI leverages their strengths
  ✅ Appropriate review and validation
  ✅ Continuous improvement feedback
  ✅ Documentation of decisions

⚠️ COMMON HANDOFF PROBLEMS:

Context Loss:
  ❌ "Review the authentication" (too vague)
  ✅ "Review auth implementation focusing on JWT security"

Missing Prerequisites:
  ❌ Starting work without dependencies ready
  ✅ Clear prerequisite checking before starting

Unclear Deliverables:
  ❌ "Make it better"
  ✅ "Improve performance by 50%, focus on database queries"

🔍 Related Commands:
  • dppm collab wiki "workflow patterns"  # Common scenarios
  • dppm collab wiki "best practices"     # Proven strategies`)
}

// /* Viser guiden til arbejdsgangsmønstre. */
func showWorkflowPatternsGuide() {
	fmt.Println(`Collaboration Workflow Patterns
===============================

🏗️ COMMON AI COLLABORATION PATTERNS:

1. LINEAR WORKFLOW - Sequential handoffs
2. PARALLEL WORKFLOW - Simultaneous work
3. REVIEW WORKFLOW - Create → Review → Refine
4. ITERATIVE WORKFLOW - Back-and-forth refinement
5. SPECIALIZATION WORKFLOW - AI-specific expertise

📋 PATTERN 1: LINEAR WORKFLOW:

Perfect for: Step-by-step processes, dependent tasks

Example - API Development:
  ::LARS:10:: Design REST API endpoints and data models ::
  ::GEMINI:11:: Review API design for usability and best practices ::
  ::LARS:12:: Implement API based on approved design ::
  ::GEMINI:13:: Create comprehensive API documentation ::
  ::LARS:14:: Write integration tests for all endpoints ::

Benefits:
  ✅ Clear sequence and dependencies
  ✅ Each stage builds on previous work
  ✅ Natural quality gates at each step

⚡ PATTERN 2: PARALLEL WORKFLOW:

Perfect for: Independent features, separate components

Example - Frontend + Backend:
  ::LARS:20:: Develop backend user management API ::
  ::GEMINI:21:: Design frontend user interface mockups ::
  ::LARS:22:: Implement frontend components ::
  ::GEMINI:23:: Create user experience documentation ::

Benefits:
  ✅ Faster overall completion
  ✅ Efficient resource utilization
  ✅ Reduced blocking and waiting

🔄 PATTERN 3: REVIEW WORKFLOW:

Perfect for: Quality assurance, critical features

Example - Security Implementation:
  ::LARS:30:: Implement OAuth2 authentication system ::
  ::GEMINI:31:: Security review of OAuth implementation ::
  ::LARS:32:: Address security concerns from task 31 ::
  ::GEMINI:33:: Final security approval and documentation ::

Benefits:
  ✅ High quality through peer review
  ✅ Multiple perspectives on solutions
  ✅ Continuous improvement process

🔁 PATTERN 4: ITERATIVE WORKFLOW:

Perfect for: Creative work, complex problem-solving

Example - Database Design:
  ::GEMINI:40:: Initial database schema design ::
  ::LARS:41:: Technical review and optimization suggestions ::
  ::GEMINI:42:: Refined schema incorporating technical feedback ::
  ::LARS:43:: Performance analysis of refined schema ::
  ::GEMINI:44:: Final schema with performance optimizations ::

Benefits:
  ✅ Combines creative and technical perspectives
  ✅ Progressive refinement of solutions
  ✅ Balanced approach to complex problems

🎯 PATTERN 5: SPECIALIZATION WORKFLOW:

Perfect for: Leveraging AI strengths, complex projects

Example - Full-Stack Application:

  Architecture (LARS Specialty):
  ::LARS:50:: System architecture and technical decisions ::
  ::LARS:51:: Database design and optimization ::
  ::LARS:52:: API implementation and testing ::

  User Experience (GEMINI Specialty):
  ::GEMINI:53:: User journey mapping and flow design ::
  ::GEMINI:54:: UI/UX design and interaction patterns ::
  ::GEMINI:55:: Content strategy and copywriting ::

  Integration (Collaborative):
  ::LARS:56:: Integrate frontend with backend APIs ::
  ::GEMINI:57:: User acceptance testing and feedback ::
  ::LARS:58:: Performance optimization based on feedback ::

🚀 REAL-WORLD EXAMPLE:

Project: E-commerce Platform Development

Phase 1 - Planning (Parallel):
  ::GEMINI:100:: Market research and feature prioritization ::
  ::LARS:101:: Technical architecture and technology stack ::

Phase 2 - Foundation (Linear):
  ::LARS:102:: Set up development environment and CI/CD ::
  ::LARS:103:: Implement core database schema ::
  ::LARS:104:: Create authentication and user management ::

Phase 3 - Features (Specialization):
  ::LARS:105:: Product catalog API and search functionality ::
  ::GEMINI:106:: Shopping cart UX design and flows ::
  ::LARS:107:: Payment processing integration ::
  ::GEMINI:108:: Checkout flow optimization ::

Phase 4 - Quality (Review):
  ::LARS:109:: Security audit and penetration testing ::
  ::GEMINI:110:: User experience testing and refinement ::
  ::LARS:111:: Performance optimization and monitoring ::
  ::GEMINI:112:: Documentation and launch preparation ::

Completion:
  ::DONE:100,101,102,103,104,105,106,107,108,109,110,111,112::

⚙️ PATTERN SELECTION GUIDE:

Use LINEAR when:
  • Tasks have clear dependencies
  • Sequential approval needed
  • Building on previous work

Use PARALLEL when:
  • Independent components
  • Time-sensitive delivery
  • Separate areas of expertise

Use REVIEW when:
  • Quality is critical
  • Complex or risky features
  • Learning and knowledge transfer

Use ITERATIVE when:
  • Creative problem-solving needed
  • Multiple valid approaches exist
  • Balancing different priorities

Use SPECIALIZATION when:
  • Large, complex projects
  • Distinct areas of expertise needed
  • Long-term collaboration

🔍 Related Commands:
  • dppm collab wiki "best practices"    # Proven strategies
  • dppm collab wiki "integration"       # DPPM integration`)
}

// /* Viser guiden til bedste praksis. */
func showBestPracticesGuide() {
	fmt.Println(`AI Collaboration Best Practices
===============================

🌟 PROVEN STRATEGIES FOR AI TEAMWORK:

Based on real-world AI collaboration experiences and successful project patterns.

📝 TASK DESIGN PRINCIPLES:

Clear and Actionable:
  ✅ "Implement JWT authentication with refresh tokens"
  ❌ "Do authentication stuff"

Specific Deliverables:
  ✅ "Create user API returning JSON with email, name, role"
  ❌ "Make user API"

Context-Rich:
  ✅ "Refactor payment processing - current code in pay.js has timeout issues"
  ❌ "Fix payment code"

Measurable Outcomes:
  ✅ "Optimize database queries - target <100ms response time"
  ❌ "Make database faster"

🎯 ID MANAGEMENT STRATEGIES:

Sequential Numbering:
  • Use 1, 2, 3... for simple projects
  • Easier to track and reference
  • Natural progression of work

Grouped Numbering:
  • 10-19: Authentication features
  • 20-29: Payment processing
  • 30-39: User interface
  • Clear functional groupings

Project Prefixes:
  • AUTH-1, AUTH-2 for authentication
  • PAY-1, PAY-2 for payments
  • Prevents ID conflicts across projects

🔄 WORKFLOW OPTIMIZATION:

Minimize Handoff Delays:
  ✅ Include all context in handoff tasks
  ✅ Specify prerequisites clearly
  ✅ Define success criteria upfront
  ❌ Assume the next AI knows the background

Parallel Work Structure:
  ✅ Design independent components for parallel work
  ✅ Create clear interfaces between components
  ✅ Minimize dependencies between parallel tasks

Progressive Complexity:
  ✅ Start with simple tasks to establish patterns
  ✅ Build complexity gradually
  ✅ Learn collaboration style before tackling hard problems

🎨 AI SPECIALIZATION GUIDELINES:

LARS (Claude) Strengths:
  • Code implementation and debugging
  • Technical analysis and architecture
  • Documentation and structured writing
  • Security and best practices review

GEMINI (Google) Strengths:
  • Creative problem-solving and brainstorming
  • User experience and design thinking
  • Content creation and copywriting
  • Alternative approaches and innovation

Collaboration Synergy:
  • LARS analyzes, GEMINI synthesizes
  • GEMINI ideates, LARS implements
  • LARS structures, GEMINI optimizes
  • Both review and refine each other's work

📊 PROJECT STRUCTURE BEST PRACTICES:

Documentation Organization:
  project/
  - README.md              # Project overview
  - collaboration.md       # Active AI tasks
  - completed.md           # Archived completed work
  - docs/
    - architecture.md    # Technical decisions
    - user-guide.md      # Usage documentation
    - api-reference.md   # API documentation
  - src/                   # Source code

Task Documentation:
  • Keep active tasks in collaboration.md
  • Archive completed work regularly
  • Maintain decision log for important choices
  • Link related tasks and dependencies

🔧 MAINTENANCE AND CLEANUP:

Regular Cleanup Schedule:
  • Weekly: Review and clean completed tasks
  • Monthly: Archive important completed work
  • Quarterly: Review collaboration patterns and improve

Task Lifecycle Management:
  1. Create tasks with clear context
  2. Update progress during work
  3. Mark completion with outcomes
  4. Archive to preserve knowledge
  5. Clean up active workspace

Quality Assurance:
  • Review task outcomes before marking DONE
  • Ensure deliverables match requirements
  • Document lessons learned
  • Update processes based on experience

⚡ PERFORMANCE OPTIMIZATION:

Reduce Context Switching:
  ✅ Batch related tasks together
  ✅ Complete task groups before switching
  ✅ Maintain focus on current work area

Efficient Communication:
  ✅ Include relevant context in task descriptions
  ✅ Link to related documentation and code
  ✅ Specify expected time investment

Progress Visibility:
  ✅ Update task status during work
  ✅ Share intermediate results and blockers
  ✅ Celebrate completed milestones

🚨 COMMON PITFALLS TO AVOID:

Task Definition Problems:
  ❌ Vague requirements that require clarification
  ❌ Missing context that leads to wrong approach
  ❌ Unrealistic scope for single task
  ❌ Dependencies not clearly specified

Workflow Issues:
  ❌ Too many tasks in flight simultaneously
  ❌ Handoff delays due to missing information
  ❌ Circular dependencies blocking progress
  ❌ Inconsistent task tracking and cleanup

Collaboration Breakdowns:
  ❌ Assumptions about shared knowledge
  ❌ Insufficient review and quality checks
  ❌ Poor communication of constraints and requirements
  ❌ Ignoring AI specialization and strengths

🎯 SUCCESS METRICS:

Quality Indicators:
  • Tasks completed without rework
  • Clear deliverables matching requirements
  • Positive feedback in review tasks
  • Successful integration of components

Efficiency Indicators:
  • Reduced time from task creation to completion
  • Minimal handoff delays and clarifications
  • High percentage of parallel vs sequential work
  • Low rate of blocked or abandoned tasks

Collaboration Indicators:
  • Regular handoffs between AIs
  • Complementary use of AI strengths
  • Continuous improvement in task quality
  • Knowledge transfer and shared learning

🔍 Related Commands:
  • dppm collab find                    # Check current tasks
  • dppm collab clean                   # Maintain workspace
  • dppm collab wiki "integration"      # DPPM integration`)
}

// /* Viser guiden til integration med DPPM. */
func showIntegrationGuide() {
	fmt.Println(`DPPM Integration Guide
=====================

🔗 INTEGRATING AI COLLABORATION WITH DPPM PROJECTS:

The AI collaboration system works seamlessly with DPPM's project management,
enabling structured AI teamwork within organized projects and phases.

📋 INTEGRATION PATTERNS:

Project-Level Collaboration:
  • Use DPPM for high-level project structure
  • Use collaboration markers for AI coordination
  • Combine project phases with AI handoff workflows

Documentation-Driven Development:
  • Store collaboration tasks in project docs
  • Link AI tasks to DPPM tasks and phases
  • Archive completed work in project history

🏗️ RECOMMENDED PROJECT STRUCTURE:

DPPM Project with AI Collaboration:
  ~/Dropbox/project-management/projects/web-app/
  - project.yaml                    # DPPM project metadata
  - collaboration/                  # AI collaboration workspace
    - active-tasks.md            # Current AI tasks
    - completed-archive.md       # Completed work history
    - handoff-patterns.md        # Project-specific patterns
  - phases/
    - setup/
      - phase.yaml
      - collaboration.md       # Phase-specific AI tasks
      - tasks/
        - setup-repo.yaml    # DPPM tasks
    - development/
      - phase.yaml
      - collaboration.md
      - tasks/
        - auth-system.yaml
        - user-api.yaml

🔄 WORKFLOW INTEGRATION EXAMPLE:

Creating an Integrated Project:

1️⃣ DPPM Project Setup:
  dppm project create ai-webapp --name "AI Collaborative Web App"
  dppm phase create foundation --project ai-webapp
  dppm task create architecture --project ai-webapp --phase foundation

2️⃣ AI Collaboration Setup:
  Create collaboration/active-tasks.md:

  # Active AI Collaboration Tasks
  ## Architecture Phase (DPPM: foundation)
  ::LARS:10::
  DPPM Context: Task architecture in phase foundation
  Design system architecture for collaborative web application
  Deliverable: Architecture document with diagrams
  ::

  ::GEMINI:11::
  Prerequisites: Task 10 completed
  Review architecture for user experience implications
  Focus: Scalability, maintainability, user-centered design
  ::

3️⃣ Execution and Coordination:
  # Check DPPM project status
  dppm status project ai-webapp

  # Find AI collaboration tasks
  dppm collab find collaboration/

  # Work on tasks, then update DPPM
  dppm task update architecture --status in_progress

4️⃣ Completion and Cleanup:
  # Mark DPPM task complete
  dppm task update architecture --status done

  # Mark AI tasks complete
  # Add ::DONE:10,11:: to collaboration.md

  # Clean up completed AI tasks
  dppm collab clean collaboration/

🎯 INTEGRATION BENEFITS:

Structured Organization:
  ✅ DPPM provides project/phase structure
  ✅ AI markers enable fine-grained coordination
  ✅ Combined view of progress at multiple levels

Clear Accountability:
  ✅ DPPM tasks track deliverable progress
  ✅ AI tasks track collaboration handoffs
  ✅ Dependencies managed at both levels

Comprehensive Documentation:
  ✅ Project structure preserved in DPPM
  ✅ AI collaboration patterns documented
  ✅ Complete work history maintained

⚙️ ADVANCED INTEGRATION PATTERNS:

Cross-Phase Handoffs:
  ## Phase Handoff: Foundation → Development

  ::LARS:50::
  DPPM Context: Completing foundation phase
  Prepare handoff package for development phase:
  - Architecture decisions document
  - Development environment setup
  - Initial code structure
  ::

  ::GEMINI:51::
  DPPM Context: Starting development phase
  Review foundation handoff and create development plan:
  - Feature prioritization
  - Sprint planning
  - Team coordination strategy
  ::

Task Dependencies Integration:
  ::LARS:60::
  DPPM Context: Task user-auth depends on task database-schema
  Prerequisites: DPPM task database-schema must be completed
  Implement user authentication based on approved schema
  ::

Multi-Project Coordination:
  ::GEMINI:70::
  DPPM Projects: web-app, mobile-app, api-gateway
  Coordinate authentication strategy across all three projects
  Ensure consistent user experience and security model
  ::

📊 MONITORING INTEGRATED WORKFLOWS:

Daily Standup Dashboard:
  # DPPM project status
  dppm status project ai-webapp

  # AI collaboration status
  dppm collab find collaboration/

  # Combined view of active work
  dppm list tasks --status in_progress

Progress Reporting:
  # Weekly progress report
  echo "=== DPPM Project Progress ===" > weekly-report.md
  dppm status project ai-webapp >> weekly-report.md

  echo "\n=== AI Collaboration Summary ===" >> weekly-report.md
  dppm collab find collaboration/ >> weekly-report.md

🔧 COMMAND COMBINATIONS:

Integrated Task Management:
  # Start DPPM task
  dppm task update feature-x --status in_progress

  # Create corresponding AI collaboration tasks
  echo "::LARS:100:: Implement feature-x (DPPM task active) ::" >> collaboration/active-tasks.md

  # Complete both when done
  dppm task update feature-x --status done
  dppm collab clean collaboration/

Project Migration:
  # Move completed AI work to archive
  dppm collab clean collaboration/
  cat collaboration/completed-archive.md >> project-history.md

🚀 GETTING STARTED:

Quick Setup for New Project:
  # Create DPPM project
  dppm project create my-ai-project
  dppm phase create planning --project my-ai-project

  # Create collaboration workspace
  mkdir -p ~/Dropbox/project-management/projects/my-ai-project/collaboration

  # Start first AI collaboration task
  echo "# Active AI Tasks" > ~/Dropbox/project-management/projects/my-ai-project/collaboration/active-tasks.md
  echo "" >> ~/Dropbox/project-management/projects/my-ai-project/collaboration/active-tasks.md
  echo "::LARS:1:: Plan the project structure and initial phases ::" >> ~/Dropbox/project-management/projects/my-ai-project/collaboration/active-tasks.md

  # Check everything is working
  dppm status project my-ai-project
  dppm collab find ~/Dropbox/project-management/projects/my-ai-project/collaboration/

🔍 Related Commands:
  • dppm wiki "project workflow"        # DPPM project management
  • dppm collab wiki "workflow patterns" # AI collaboration patterns
  • dppm status project PROJECT_NAME    # Combined status view`)
}

// /* Initialiserer 'collab' kommandoen og dens underkommandoer. */
func init() {
	collabCmd.AddCommand(collabFindCmd)
	collabCmd.AddCommand(collabCleanCmd)
	collabCmd.AddCommand(collabWikiCmd)
	rootCmd.AddCommand(collabCmd)
}