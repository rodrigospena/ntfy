---
name: goal
description: Run long-running autonomous goal workflow following PXOS 6-phase lifecycle (Discover -> Plan -> Execute -> Validate -> Review -> Compact).
---

# /goal — PXOS Autonomous Lifecycle Workflow

When this workflow is invoked, execute the task strictly adhering to the PXOS default 6-phase sequence:

1. **Discover**:
   - Inspect `.ai/AI_BASE.md`, `.ai/PROJECT_CONTEXT.md`, and `.ai/CURRENT_SPEC.md`.
   - Identify existing code patterns, boundaries, and risk tier (Low, Medium, High).
   - Do not modify source code in this phase.

2. **Plan**:
   - Write or update `implementation_plan.md`.
   - Outline affected components/files, main risks, and verification strategy.
   - For Medium/High risk changes, document explicit rationale and obtain user approval.

3. **Execute**:
   - Implement incremental, reversible changes.
   - Maintain consistency with surrounding architecture and existing codebase patterns.

4. **Validate**:
   - Run tests and build checks (`npm run build`, `go build`).
   - Confirm runtime behavior and edge cases without swallowing errors.

5. **Review**:
   - Audit for overengineering, scope creep, or unnecessary abstractions.

6. **Compact Context**:
   - Update `.ai/CURRENT_SPEC.md` and `.ai/DECISION_LOG.md`.
   - Summarize changes, architectural decisions, and remaining tasks.
