# PROJECT CONTEXT — ntfy

## Product Overview
ntfy is a HTTP-based pub-sub notification service. It allows users to send push notifications to phones or desktop via simple HTTP PUT/POST requests.

## Stack
- **Backend**: Go (HTTP server, WebSockets, SSE, SQLite/Auth, Push Integrations)
- **Frontend**: React (Vite, Material UI, Dexie.js / IndexedDB, PWA Service Worker)
- **Deployment**: Binary executable (`ntfy-app.exe`), Docker, or Systemd service.

## Architecture & Structure
- `cmd/`: CLI commands and flags entrypoints.
- `server/`: Core server logic, HTTP router, background message bus, auth manager.
- `user/`: User and access control database management.
- `web/`: React frontend application source code.
- `.ai/`: Operational layer for AI assistant (PXOS framework).

## Key Conventions & Conventions
- All commit messages MUST follow Conventional Commits in English (`feat: ...`, `fix: ...`, `docs: ...`).
- Custom routes registered in `server/server.go` and `web/src/components/routes.js`.
