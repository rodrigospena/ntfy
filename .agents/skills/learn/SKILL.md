---
name: learn
description: Capture workspace patterns, corrections, and architectural decisions into PXOS context files (.ai/).
---

# /learn — PXOS Context Compaction & Pattern Capture

When this workflow is invoked:

1. **Extract Insights**:
   - Identify new architectural patterns, user corrections, environment configurations, or domain rules discovered during the session.

2. **Update Core PXOS Files**:
   - Update `.ai/PROJECT_CONTEXT.md` for permanent project rules, tech stack details, or directory structure updates.
   - Append to `.ai/DECISION_LOG.md` for durable architectural, product, or design decisions.
   - If domain-specific or data-model invariants were introduced, create or update `.ai/DOMAIN_RULES.md` or `.ai/DATA_MODEL.md`.

3. **Verify Conventions**:
   - Ensure commit history aligns with English Conventional Commits (`feat: ...`, `fix: ...`, `docs: ...`).
