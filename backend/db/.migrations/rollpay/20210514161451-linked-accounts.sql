
-- +migrate Up
CREATE TABLE IF NOT EXISTS linked_accounts (
    item_id text NOT NULL PRIMARY KEY,
    access_token text NOT NULL,
    alias TEXT NOT NULL,
    user_id bigint NOT NULL REFERENCES users (id),
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL
);

CREATE INDEX linked_accounts_created_at_idx ON linked_accounts (created_at DESC);
CREATE INDEX linked_accounts_deleted_at_idx ON linked_accounts (deleted_at DESC);

-- +migrate Down
DROP INDEX IF EXISTS linked_accounts_created_at_idx;
DROP INDEX IF EXISTS linked_accounts_deleted_at_idx;
DROP TABLE IF EXISTS linked_accounts;
