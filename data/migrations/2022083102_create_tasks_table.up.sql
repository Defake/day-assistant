BEGIN;

CREATE TABLE IF NOT EXISTS tasks (
    id         serial PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    body       jsonb
);

COMMIT;

