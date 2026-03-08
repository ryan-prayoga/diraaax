# diraaax

A private couple website for Ryan & Dira (Maboyyy & Magirll).

## Tech Stack

- **Backend:** Go, chi router, pgx/v5 (PostgreSQL)
- **Frontend:** SvelteKit, TypeScript, TailwindCSS v4
- **Auth:** PIN-based with secure HTTP-only cookie sessions
- **File Upload:** Local disk storage

## Prerequisites

- Go 1.22+
- Node.js 18+
- PostgreSQL (accessible via SSH tunnel or locally)
- SSH tunnel to VPS (if using remote DB)

## Project Structure

```
diraaax/
  backend/
    cmd/api/main.go          # Entry point
    internal/config/          # Environment config
    internal/db/              # PostgreSQL connection
    internal/models/          # Data models
    internal/repositories/    # DB queries
    internal/services/        # Business logic
    internal/handlers/        # HTTP handlers
    internal/middleware/       # Auth, logging
    internal/utils/           # Helpers
    uploads/gallery/          # Uploaded images
  frontend/
    src/lib/                  # API client, auth store
    src/routes/               # SvelteKit pages
```

## Setup

### 1. SSH Tunnel (if using remote PostgreSQL)

```bash
ssh -L 5433:localhost:5432 your-vps-user@your-vps-ip -N
```

Keep this running in a separate terminal.

### 2. Backend

```bash
cd backend

# Create .env from example
cp .env.example .env
# Edit .env with your actual DB password and settings

# Install dependencies
go mod tidy

# Run the server
go run cmd/api/main.go
```

Backend runs on `http://localhost:8080`.

### 3. Frontend

```bash
cd frontend

# Create .env from example
cp .env.example .env

# Install dependencies
npm install

# Run dev server
npm run dev
```

Frontend runs on `http://localhost:5173`.

## API Endpoints

| Method | Path                    | Description         |
| ------ | ----------------------- | ------------------- |
| GET    | `/health`               | Health check        |
| POST   | `/api/auth/verify-pin`  | Login with PIN      |
| POST   | `/api/auth/logout`      | Logout              |
| GET    | `/api/auth/me`          | Current user        |
| GET    | `/api/plans`            | List future plans   |
| POST   | `/api/plans`            | Create plan         |
| PATCH  | `/api/plans/:id/toggle` | Toggle plan status  |
| DELETE | `/api/plans/:id`        | Delete plan         |
| GET    | `/api/gallery`          | List gallery items  |
| POST   | `/api/gallery`          | Upload image        |
| DELETE | `/api/gallery/:id`      | Delete gallery item |
| GET    | `/api/secret-notes`     | List secret notes   |
| POST   | `/api/secret-notes`     | Create note         |

## Frontend Routes

| Route           | Description               |
| --------------- | ------------------------- |
| `/login`        | PIN login page            |
| `/dashboard`    | Home with quick links     |
| `/gallery`      | Photo gallery with upload |
| `/future-plans` | Shared plans list         |
| `/secret`       | Secret notes              |
| `/our-story`    | Relationship timeline     |

## Environment Variables

### Backend (.env)

```
APP_PORT=8080
DATABASE_URL=postgres://diraaax_user:PASSWORD@127.0.0.1:5433/diraaax_db?sslmode=disable
SESSION_COOKIE_NAME=diraaax_session
SESSION_SECRET=your_secret_here
UPLOAD_DIR=./uploads/gallery
CORS_ORIGIN=http://localhost:5173
```

### Frontend (.env)

```
PUBLIC_API_BASE_URL=http://localhost:8080
```
