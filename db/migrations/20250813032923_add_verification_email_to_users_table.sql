-- migrate:up
ALTER TABLE users 
ADD COLUMN is_verified BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN verification_token VARCHAR(255),
ADD COLUMN verified_at TIMESTAMPTZ;

-- migrate:down
ALTER TABLE users
DROP COLUMN is_verified,
DROP COLUMN verification_token,
DROP COLUMN verified_at;
