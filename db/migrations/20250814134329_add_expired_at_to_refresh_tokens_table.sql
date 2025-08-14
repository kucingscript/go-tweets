-- migrate:up
ALTER TABLE refresh_tokens
ADD COLUMN expired_at TIMESTAMPTZ;

-- migrate:down
ALTER TABLE refresh_tokens
DROP COLUMN expired_at;
