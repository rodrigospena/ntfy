---
name: schedule
description: Schedule automated background tasks, periodic checks, or timers aligned with PXOS compact context rules.
---

# /schedule — PXOS Background Task & Timer Scheduling

When this workflow is invoked:

1. **Assess Task & Frequency**:
   - Determine whether a one-shot timer or recurring cron check is required.

2. **Execute Scheduler Tool**:
   - Use `schedule` with clear notification prompt.
   - Avoid manual sleep loops or blocking background commands.

3. **Log & Compact**:
   - Log active background schedules in `.ai/CURRENT_SPEC.md` if long-running.
