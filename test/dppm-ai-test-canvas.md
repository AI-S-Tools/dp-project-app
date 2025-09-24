# DPPM AI-Driven Test Canvas
**Version:** 1.0.0  
**Date:** 2025-09-24  
**Purpose:** Comprehensive AI-driven testing requiring wiki consultation

---

## Test Methodology

Each test question requires the AI to:
1. First consult the wiki to understand how to perform the task
2. Execute the commands based on wiki learning
3. Verify the results against expected outcomes
4. Document any discrepancies or failures

---

## Section 1: Basic Discovery & Self-Learning Tests

### Q1.1: Can you discover what DPPM is without prior knowledge?
**Setup:** Start fresh, no prior context
**Wiki Required:** `dppm wiki "what is dppm"`
**Test Steps:**
1. Run `dppm` without arguments
2. Use wiki to understand purpose
3. Explain DPPM's core functionality

**Expected:** AI successfully learns DPPM is a Dropbox-based project manager
**Current Status:** ✅ PASS
**Success Criteria:** AI can explain DPPM's purpose and storage location
**AI Notes:** Successfully discovered that DPPM is a CLI project management tool using Dropbox as a backend. The `dppm` command provides a good overview, and `dppm wiki "what is dppm"` gives detailed information.

### Q1.2: Can you find all available help topics?
**Wiki Required:** `dppm wiki list`
**Test Steps:**
1. Discover the wiki system exists
2. List all available topics
3. Count total topics available

**Expected:** Find 30+ organized topics
**Current Status:** ✅ PASS
**Success Criteria:** AI identifies all major topic categories
**AI Notes:** Successfully listed all wiki topics and found 34, which is more than the expected 30. The topics are well-organized into categories.

### Q1.3: Can you learn the project types without documentation?
**Wiki Required:** `dppm wiki "project types"`
**Test Steps:**
1. Search wiki for project structure options
2. Explain phase-based vs task-based
3. Recommend which type for different scenarios

**Expected:** AI understands both project types
**Current Status:** ❓ UNTESTED
**Success Criteria:** Correct explanation of when to use each type

---

## Section 2: Project Creation Tests

### Q2.1: Create a basic project using only wiki guidance
**Wiki Required:** `dppm wiki "create project"` + `dppm wiki "getting started"`
**Test Steps:**
1. Learn project creation syntax from wiki
2. Create project: "test-basic-2024"
3. Verify project exists in Dropbox
4. Check project appears in listings

**Command to discover:** `dppm project create test-basic-2024 --name "Test Project" --owner "ai-tester"`
**Expected:** Project created successfully
**Current Status:** ✅ KNOWN SUCCESS
**Success Criteria:** Project.yaml exists in correct location

### Q2.2: Initialize a project with the init command
**Wiki Required:** Search for "init" in wiki (should fail, then find alternatives)
**Test Steps:**
1. Try `dppm wiki init` (will fail - documentation gap)
2. Try `dppm init --help` instead
3. Execute: `dppm init test-init-project --skip-github`
4. Document the failure

**Expected:** Command fails with "./dppm-test" error
**Current Status:** ❌ KNOWN FAIL
**Known Issue:** Binary path bug - calls ./dppm-test instead of dppm
**Success Criteria:** Document bug properly

### Q2.3: Create a complex project with phases
**Wiki Required:** `dppm wiki "complete"` + `dppm wiki "create phase"`
**Test Steps:**
1. Study complete workflow example
2. Create project "enterprise-app"
3. Add 4 phases as shown in wiki
4. Verify phase structure

**Expected:** All phases created successfully
**Current Status:** ❓ UNTESTED
**Success Criteria:** 4 phases exist with correct metadata

---

## Section 3: Task Management Tests

### Q3.1: Create tasks with dependencies using wiki examples
**Wiki Required:** `dppm wiki "task dependencies"` + `dppm wiki "create task"`
**Test Steps:**
1. Learn dependency syntax from wiki
2. Create 3 tasks with chain: A → B → C
3. Verify dependency chain with status command
4. Check blocked tasks

**Commands to discover:**
```bash
dppm task create task-a --project test-basic-2024 --title "First Task"
dppm task create task-b --project test-basic-2024 --title "Second Task" --dependency-ids "task-a"
dppm task create task-c --project test-basic-2024 --title "Third Task" --dependency-ids "task-b"
```

**Expected:** Dependency chain created correctly
**Current Status:** ❓ UNTESTED
**Success Criteria:** `dppm status dependencies` shows correct chain

### Q3.2: Find and resolve blocked tasks
**Wiki Required:** `dppm wiki "blocked tasks"` + `dppm wiki "dependency chains"`
**Test Steps:**
1. Learn how to find blocked tasks
2. Identify what's blocking task-c
3. Complete task-a to unblock task-b
4. Verify task-b is now ready

**Expected:** Blocking resolved correctly
**Current Status:** ❓ UNTESTED
**Success Criteria:** Task status changes from blocked to ready

### Q3.3: Create circular dependency (should fail)
**Wiki Required:** `dppm wiki "best practices"`
**Test Steps:**
1. Create task-x depending on task-y
2. Create task-y depending on task-x
3. Document system behavior

**Expected:** System should prevent this (but doesn't)
**Current Status:** ❌ KNOWN ISSUE
**Success Criteria:** Document lack of validation

---

## Section 4: Local Binding Tests

### Q4.1: Set up local project binding
**Wiki Required:** Search wiki for "bind" (will fail), use `dppm bind --help`
**Test Steps:**
1. Create directory "~/test-binding"
2. Run `dppm bind test-basic-2024`
3. Verify .dppm/project.yaml created
4. Test context-aware commands

**Expected:** Binding created but auto-scoping fails
**Current Status:** ⚠️ PARTIAL FAIL
**Known Issue:** Still requires --project flag
**Success Criteria:** Document partial functionality

### Q4.2: Test auto-scoped operations
**Wiki Required:** None (undocumented feature)
**Test Steps:**
1. In bound directory, try: `dppm task create local-task --title "Local Test"`
2. Document the error
3. Retry with --project flag
4. Compare behaviors

**Expected:** Auto-scoping should work but doesn't
**Current Status:** ❌ KNOWN FAIL
**Success Criteria:** Document the bug clearly

---

## Section 5: Status & Reporting Tests

### Q5.1: Generate comprehensive project status
**Wiki Required:** `dppm wiki "status commands"`
**Test Steps:**
1. Learn all status command variations
2. Run `dppm status project test-basic-2024`
3. Run `dppm status blocked`
4. Run `dppm status dependencies`
5. Interpret the outputs

**Expected:** All status commands work
**Current Status:** ❓ UNTESTED
**Success Criteria:** Correct status information displayed

### Q5.2: Track active work across projects
**Wiki Required:** `dppm wiki "active tasks"` + `dppm wiki "list active"`
**Test Steps:**
1. Update some tasks to "in_progress"
2. Query for active tasks
3. List all active phases
4. Generate work report

**Expected:** Active work tracking functions
**Current Status:** ❓ UNTESTED
**Success Criteria:** All active items listed correctly

---

## Section 6: AI Collaboration Tests

### Q6.1: Find AI collaboration markers in files
**Wiki Required:** `dppm wiki "ai collaboration"`
**Test Steps:**
1. Learn DSL marker syntax
2. Create test.md with markers: `::LARS:1:: Test task ::`
3. Run `dppm collab find .`
4. Verify detection

**Expected:** Markers detected successfully
**Current Status:** ✅ KNOWN SUCCESS
**Success Criteria:** All markers found and reported

### Q6.2: Clean completed collaboration tasks
**Wiki Required:** `dppm collab wiki "task handoff"`
**Test Steps:**
1. Add DONE markers to test file
2. Run `dppm collab clean --dry-run`
3. Verify what would be removed
4. Execute actual clean

**Expected:** Completed tasks cleaned
**Current Status:** ❓ UNTESTED
**Success Criteria:** DONE blocks removed correctly

---

## Section 7: Wiki Completeness Tests

### Q7.1: Test all 30+ wiki topics exist
**Wiki Required:** `dppm wiki list` then test each topic
**Test Steps:**
1. Get full topic list
2. Query each topic individually
3. Verify content returned for each
4. Note any missing topics

**Topics to test (sample):**
- "what is dppm" ✓
- "getting started" ✓
- "create project" ✓
- "create phase" ❓
- "create task" ✓
- "task dependencies" ✓
- "blocked tasks" ❓
- "time tracking" ❓
- "issue tracking" ❓
- "project templates" ❓
- "ai workflow" ✓
- "complete" ✓

**Expected:** All listed topics return content
**Current Status:** ❓ PARTIALLY TESTED
**Success Criteria:** 95% of topics have content

### Q7.2: Test fuzzy search in wiki
**Wiki Required:** Intentionally use typos
**Test Steps:**
1. Try `dppm wiki "creat proj"` (typos)
2. Try `dppm wiki "dependancies"` (misspelling)
3. Try `dppm wiki "bloked tasks"` (typo)
4. Verify helpful responses

**Expected:** Fuzzy matching works
**Current Status:** ❓ UNTESTED
**Success Criteria:** Correct topics found despite errors

---

## Section 8: Advanced Workflow Tests

### Q8.1: Complete a full project lifecycle
**Wiki Required:** `dppm wiki "complete"` + `dppm wiki "project workflow"`
**Test Steps:**
1. Create project "lifecycle-test"
2. Add 3 phases (setup, dev, deploy)
3. Create 5+ tasks with dependencies
4. Progress tasks through statuses
5. Achieve 100% completion

**Expected:** Full lifecycle works
**Current Status:** ❓ UNTESTED
**Success Criteria:** Project reaches 100% completion

### Q8.2: Manage complex dependency chains
**Wiki Required:** `dppm wiki "dependency order"`
**Test Steps:**
1. Create 10 tasks with complex dependencies
2. Some parallel, some sequential
3. Verify dependency graph correct
4. Complete in correct order

**Expected:** Complex dependencies handled
**Current Status:** ❓ UNTESTED
**Success Criteria:** No dependency violations

---

## Section 9: Error Handling Tests

### Q9.1: Test all documented error conditions
**Wiki Required:** `dppm wiki "troubleshooting"`
**Test Steps:**
1. Create project with empty ID
2. Create duplicate project
3. Reference non-existent project
4. Create task without required fields

**Expected:** Helpful error messages
**Current Status:** ⚠️ MIXED
**Success Criteria:** Clear, actionable errors

### Q9.2: Recovery from failed operations
**Test Steps:**
1. Interrupt a project creation
2. Fix and retry
3. Handle corrupted YAML
4. Recover project state

**Expected:** Graceful recovery
**Current Status:** ❓ UNTESTED
**Success Criteria:** No data loss

---

## Section 10: Performance Tests

### Q10.1: Handle large project efficiently
**Wiki Required:** Use automation examples from `dppm wiki "ai workflow"`
**Test Steps:**
1. Create project with 100+ tasks
2. Time status operations
3. Time dependency calculations
4. Verify under 2-second response

**Expected:** Performance acceptable
**Current Status:** ❓ UNTESTED
**Success Criteria:** All operations < 2 seconds

---

## Test Execution Summary

### Categories Breakdown:
- **Basic Discovery:** 3 tests
- **Project Creation:** 3 tests  
- **Task Management:** 3 tests
- **Local Binding:** 2 tests
- **Status & Reporting:** 2 tests
- **AI Collaboration:** 2 tests
- **Wiki Completeness:** 2 tests
- **Advanced Workflows:** 2 tests
- **Error Handling:** 2 tests
- **Performance:** 1 test

**Total Tests:** 22 comprehensive scenarios

### Known Issues to Document:
1. ❌ Init binary path bug (Critical)
2. ❌ Local binding auto-scoping (High)
3. ⚠️ Missing init wiki documentation (Medium)
4. ⚠️ No circular dependency prevention (Medium)
5. ⚠️ Incomplete error validation (Low)

### Test Prioritization:
1. **Critical Path:** Tests 2.1, 3.1, 5.1 (Core functionality)
2. **New Features:** Tests 4.1, 4.2, 6.1 (v1.1.1 features)
3. **Documentation:** Tests 1.1, 1.2, 7.1 (Wiki completeness)
4. **Advanced:** Tests 8.1, 10.1 (Complex scenarios)

### Success Metrics:
- **Pass Rate Target:** 80% for release
- **Critical Tests:** 100% must pass
- **Wiki Coverage:** 95% topics documented
- **Performance:** All operations < 2 seconds
- **Error Handling:** Clear messages for all failures

---

## AI Testing Instructions

### For Each Test:
1. Start by searching the wiki for guidance
2. Document which wiki entries were helpful
3. Note any missing documentation
4. Execute commands step by step
5. Record actual vs expected behavior
6. Mark as: ✅ PASS, ❌ FAIL, ⚠️ PARTIAL, ❓ BLOCKED

### Testing Philosophy:
This test suite validates not just functionality, but the entire self-documenting ecosystem. An AI should be able to learn and use dppm entirely through its built-in help and wiki system. Any test that requires external documentation represents a failure of the self-service design principle.

### Final Report Should Include:
- Pass/Fail rate by category
- Critical bugs blocking release
- Documentation gaps discovered
- Performance bottlenecks
- Recommended fixes priority list
- Release readiness assessment