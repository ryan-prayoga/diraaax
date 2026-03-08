-- cleanup_duplicates.sql
-- Run this to clean up any duplicate seed data.
-- Safe to run multiple times (idempotent).

-- Remove duplicate access_codes, keeping the one with the lowest id
DELETE FROM access_codes
WHERE id NOT IN (
    SELECT MIN(id)
    FROM access_codes
    GROUP BY code
);

-- Remove duplicate future_plans (same title + created_by), keeping lowest id
DELETE FROM future_plans
WHERE id NOT IN (
    SELECT MIN(id)
    FROM future_plans
    GROUP BY title, created_by
);

-- Remove duplicate secret_notes (same title + content + created_by), keeping lowest id
DELETE FROM secret_notes
WHERE id NOT IN (
    SELECT MIN(id)
    FROM secret_notes
    GROUP BY title, content, created_by
);

-- Remove expired sessions
DELETE FROM sessions
WHERE expires_at < NOW();

-- Verify counts after cleanup
SELECT 'access_codes' AS table_name, COUNT(*) AS count FROM access_codes
UNION ALL
SELECT 'future_plans', COUNT(*) FROM future_plans
UNION ALL
SELECT 'secret_notes', COUNT(*) FROM secret_notes
UNION ALL
SELECT 'sessions', COUNT(*) FROM sessions
UNION ALL
SELECT 'gallery_items', COUNT(*) FROM gallery_items
UNION ALL
SELECT 'users', COUNT(*) FROM users;
