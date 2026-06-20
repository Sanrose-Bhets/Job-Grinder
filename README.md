# Task Scheduler

A backend task scheduler built in Go, using PostgreSQL for storage and Docker for containerization. Built primarily as a learning project to practice Go concurrency (goroutines, channels, context cancellation) in a real backend context.

## Overview

The scheduler stores and manages tasks in a PostgreSQL database. Tasks aren't limited to one category — they can represent anything from work commitments to personal to-dos.

### Task Categories

- **Work** — meetings, deadlines, project milestones, interview tracking (see below)
- **Personal** — errands, reminders, daily routines
- **General** — anything that doesn't fit neatly into the above

## Interview Tracking

A dedicated section for work-related tasks that tracks job interview activity. This includes:

- Number of interviews scheduled / completed
- Interview stage (e.g. screening, technical, final)
- Company / role associated with each interview
- Outcome tracking (pending, passed, rejected, offer)

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Containerization:** Docker & Docker Compose
- **DB Admin UI:** pgAdmin

## Concurrency Concepts Practiced

- Goroutines for handling concurrent task execution
- Channels for communication between scheduler workers
- `context` package for cancellation and timeouts

## Running Locally

```bash
docker compose up -d
```

This spins up:
- `task-scheduler` — the Go application
- `postgres` — the database
- `pgadmin` — web UI for inspecting the database at `http://localhost:8090`

## Requirement
Docker is required to be able to run this application.
For Windows the Docker Desktop should be used.