# diraaax Agent Development Guide

This file provides instructions for coding agents working on the diraaax project.

diraaax is a **private romantic web platform** created for Ryan and Dira.

The goal of diraaax is to build a **digital love universe** where memories,
photos, notes, dreams, and stories live together.

---

# Project Structure

frontend/

- SvelteKit
- TailwindCSS
- Motion animations

backend/

- Golang
- PostgreSQL
- REST API

docs/

- roadmap-v2.md
- ui-design-v2.md
- database-schema-v2.md
- animation-design-v2.md

---

# Development Rules

Agents must follow these rules when implementing features.

1. Always follow the design system from `docs/ui-design-v2.md`
2. Database schema must follow `docs/database-schema-v2.md`
3. New features must match the roadmap in `docs/roadmap-v2.md`
4. Animations must follow `docs/animation-design-v2.md`
5. Do not modify API routes without updating documentation
6. Mobile-first UI is required
7. Maintain emotional and romantic design tone

---

# Backend Guidelines

Language: Go

Structure:

handlers/
services/
repositories/
models/

API should follow REST conventions.

Example:

GET /api/timeline
POST /api/timeline
DELETE /api/timeline/:id

Authentication:

- PIN based authentication
- session cookie

---

# Frontend Guidelines

Framework: SvelteKit

Styling:
TailwindCSS

Animation:
Motion One or Lottie

Design priorities:

- soft
- romantic
- cute
- minimal
- emotional

---

# UI Principles

The website must feel like:

- a digital love diary
- a scrapbook of memories
- a romantic private world

Avoid:

- corporate UI
- overly dark themes
- aggressive colors
