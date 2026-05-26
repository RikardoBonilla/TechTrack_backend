# techtrack

> IT asset management and help desk API built in Go.

[![Go](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue.svg)](https://postgresql.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## What it does

REST API for managing IT assets and support tickets across multiple organizations. Multi-tenant SaaS-ready architecture — each tenant has fully isolated data.

**Core features:**
- Track hardware assets with QR codes and lifecycle statuses (ACTIVE, IN_REPAIR, RETIRED, LOST)
- Help desk ticketing with priorities (LOW → CRITICAL) and assignment to technicians
- Role-based access: ADMIN, TECHNICIAN, STAFF
- Full audit log — every change recorded with actor, entity and JSONB diff
- JWT authentication with per-tenant isolation

## Stack

- **Go** — chi router, golang-jwt, pgx/v5
- **PostgreSQL** — UUID primary keys, JSONB fields, partial indexes
- **Makefile** — run, build, test, lint

## Architecture

```
cmd/api/main.go           — entry point
internal/
├── domain/               — core models (Asset, Ticket, User, Tenant, AuditLog)
├── repository/           — data access layer
├── transport/http/       — HTTP handlers
├── config/               — env configuration
├── database/             — PostgreSQL connection pool
└── pkg/                  — shared utilities
migrations/               — SQL migrations (golang-migrate)
```

## Quick Start

```bash
git clone https://github.com/RikardoBonilla/techtrack.git
cd techtrack

cp .env.example .env
# Edit .env with your PostgreSQL credentials

make run
```

## Database Schema

Five tables: `tenants`, `users`, `assets`, `tickets`, `audit_logs`.

```bash
# Apply migrations
migrate -path migrations -database "$DATABASE_URL" up
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/auth/login` | Authenticate and get JWT |
| GET | `/assets` | List tenant assets |
| POST | `/assets` | Register new asset |
| GET | `/tickets` | List help desk tickets |
| POST | `/tickets` | Open new ticket |
| PATCH | `/tickets/:id` | Update ticket status |

## Development

```bash
make run    # start server
make build  # compile binary to bin/api
make test   # run tests
make lint   # run golangci-lint
```

---

*Built by [Ricardo Andres Bonilla Prada](https://github.com/RikardoBonilla)*
