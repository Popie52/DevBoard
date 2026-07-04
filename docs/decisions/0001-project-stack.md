# ADR-0001: Technology Stack

## Status

Accepted

## Context

We need a production-like full-architecture that is easy to deploy on Linux VM and can evolve over time.

## Decision

Frontend:
- Next.js
- TypeScript
- Tailwind CSS

Backend:
- Golang
- net/http

Databases:
- PostgreSQL


Infrastructure
- Ubuntu Server VM
- Docker (later)
- Nginx (later)

## Consequences

- Clear separation between frontend and backend
- Easy migrations to GIN in the future for benchmarking
- Fronend remains a client of the Go API.
