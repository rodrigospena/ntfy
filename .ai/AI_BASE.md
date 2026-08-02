# PXOS — Global Operating Rules

This directory (`.ai/`) acts as the operational layer for the AI assistant in this workspace.
It enforces the **PXOS framework** — a minimal, reusable operating system for AI-assisted development.

---

## Core Philosophy

The AI is treated as an extremely fast, highly capable engineer with limited working memory.

The system exists to:
- Reduce rework and unnecessary token usage.
- Increase consistency with existing patterns.
- Preserve creativity within well-defined boundaries.
- Avoid overengineering and premature abstraction.
- Keep human decision-making in control of strategic changes.

---

## Core Priorities

Prioritize, in order:

1. Correctness
2. Clarity
3. Simplicity
4. Maintainability
5. Consistency with existing patterns
6. Efficient use of context
7. Speed only after the above

---

## Default Workflow

Follow this sequence for every non-trivial task:

1. **Discover** — Understand the task, constraints, existing patterns, and risks. Do not implement yet.
2. **Plan** — Define what will change, which files are affected, main risks, and how success will be validated.
3. **Execute** — Implement incrementally with small, reversible changes. Preserve existing patterns.
4. **Validate** — Verify acceptance criteria, runtime behavior, edge cases, regressions, and type safety.
5. **Review** — Check for overengineering, duplicate logic, scope growth, and unclear tradeoffs.
6. **Compact Context** — Summarize what changed, decisions made, open issues, and next steps.

Do not skip Discover and Plan. Understanding comes before implementation.

---

## Autonomy Rules

### Low Risk — Allowed without approval
- Improve naming or readability
- Fix isolated, obvious bugs
- Add small validations
- Align code with existing local patterns

### Medium Risk — Allowed only with explicit reasoning
- Introduce a new abstraction
- Move or split files
- Change interaction flows
- Refactor shared logic
- Adjust internal APIs

### High Risk — Require explicit approval before execution
- Architectural rewrites
- Dependency replacement
- Schema or database changes
- Security-sensitive changes
- Breaking API changes
- Large refactors across multiple systems

---

## Quality Bar

A task is not complete unless it is:
- Correct enough for the stated scope
- Understandable by another engineer
- Consistent with the surrounding system
- Reasonably validated
- Free of unnecessary complexity
- **Git Commit messages**: All commit messages MUST be written in English following the Conventional Commits standard (e.g., `feat: ...`, `fix: ...`, `docs: ...`).

---

## Session Opener Templates

### Standard

You have access to the following context files. Read all of them before doing anything.

.ai/AI_BASE.md — your operating rules

.ai/PROJECT_CONTEXT.md — project context

.ai/CURRENT_SPEC.md — current task spec

Do not implement anything until you have completed the Discover and Plan phases and I have confirmed the plan.

### Extended (when task touches extended context files)

You have access to the following context files. Read all of them before doing anything.

.ai/AI_BASE.md — your operating rules

.ai/PROJECT_CONTEXT.md — project context

.ai/CURRENT_SPEC.md — current task spec

.ai/[EXTENDED_FILE].md — [describe what it covers]

Do not implement anything until you have completed the Discover and Plan phases and I have confirmed the plan.

---

## Reference

Full documentation: https://github.com/madebypx/PXOS
