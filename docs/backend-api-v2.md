# diraaax Backend API v2

This document describes the current diraaax v2 backend foundation.

Base path: `/api`

All responses use this envelope:

```json
{
  "success": true,
  "data": {}
}
```

Error responses use:

```json
{
  "success": false,
  "error": {
    "code": "invalid_request",
    "message": "title is required"
  }
}
```

## Authentication

- Auth method: PIN verification against `access_codes`
- Session storage: `sessions`
- Transport: HTTP-only cookie using `SESSION_COOKIE_NAME`
- Public routes:
  - `GET /health`
  - `POST /api/auth/verify-pin`
- Every other `/api/*` route requires a valid session cookie

## Routes

### Auth

- `POST /api/auth/verify-pin`
- `POST /api/auth/logout`
- `GET /api/auth/me`

### Timeline

- `GET /api/timeline`
- `POST /api/timeline`
- `DELETE /api/timeline/:id`

### Memories

- `GET /api/memories`
- `GET /api/memories/random`
- `POST /api/memories`
- `DELETE /api/memories/:id`

### Bucket List

- `GET /api/bucket-list`
- `POST /api/bucket-list`
- `PATCH /api/bucket-list/:id/toggle`
- `DELETE /api/bucket-list/:id`

### Love Capsules

- `GET /api/capsules`
- `GET /api/capsules/:id`
- `POST /api/capsules`
- `POST /api/capsules/:id/open`
- `GET /api/capsules/:id/scenes`
- `POST /api/capsules/:id/scenes`

### Daily Moods

- `GET /api/moods`
- `POST /api/moods`

### Love Reasons

- `GET /api/love-reasons`
- `POST /api/love-reasons`
- `DELETE /api/love-reasons/:id`

### Scaffolded Endpoints

- `GET /api/voice-notes`
- `POST /api/voice-notes`
- `GET /api/memory-locations`
- `POST /api/memory-locations`

## Notes

- The backend reads configuration from `.env`.
- PostgreSQL is treated as the source of truth. No schema generation or migrations are included in this layer.
- Capsules can only be opened once their `open_date` has passed unless they were already opened.
