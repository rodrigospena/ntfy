---
name: grill-me
description: Interactively interview the user to resolve ambiguity, establish design trade-offs, and align strategic requirements under PXOS.
---

# /grill-me — PXOS Interactive Design & Discovery Interview

When this workflow is invoked:

1. **Analyze Context & Scope**:
   - Read `.ai/CURRENT_SPEC.md` and `.ai/PROJECT_CONTEXT.md`.
   - Identify unresolved strategic questions, high-risk architectural trade-offs, or UI/UX choices.

2. **Conduct Structured Interview**:
   - Present concise, numbered questions highlighting trade-offs (e.g., simplicity vs. future flexibility, performance vs. maintainability).
   - Flag any High-Risk PXOS changes (architectural rewrites, breaking API changes, auth changes) requiring explicit user authorization.

3. **Persist Strategic Decisions**:
   - Upon receiving user responses, record all agreed decisions into `.ai/DECISION_LOG.md`.
   - Update `.ai/CURRENT_SPEC.md` to reflect the refined spec.
