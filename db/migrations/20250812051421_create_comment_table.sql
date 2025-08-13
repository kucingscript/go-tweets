-- migrate:up
CREATE TABLE IF NOT EXISTS comments (
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL,
  user_id INT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (post_id) REFERENCES posts (id),
  FOREIGN KEY (user_id) REFERENCES users (id)
)

-- migrate:down
DROP TABLE IF EXISTS comments;