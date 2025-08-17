-- migrate:up
ALTER TABLE users
ADD COLUMN password_reset_token VARCHAR(255),
ADD COLUMN password_reset_token_expires_at TIMESTAMPTZ;

-- migrate:down
ALTER TABLE users
DROP COLUMN password_reset_token,
DROP COLUMN password_reset_token_expires_at;