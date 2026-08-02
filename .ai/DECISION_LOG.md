# DECISION LOG

## [2026-08-02] Adopt PXOS Framework Structure
- **Decision**: Added `.ai/` operational layer containing `AI_BASE.md`, `PROJECT_CONTEXT.md`, `CURRENT_SPEC.md`, and `DECISION_LOG.md`.
- **Rationale**: Align repository operations with minimal, reusable AI operating standards (PXOS).

## [2026-08-02] Topic Landing Page (`/sub/:topic`)
- **Decision**: Implemented standalone landing page component `TopicLandingPage.jsx` and backend route `webAppTopicSubRegex`.
- **Rationale**: Allow non-technical end-users to subscribe to a single topic via a clean UI without exposing the full ntfy management console.
