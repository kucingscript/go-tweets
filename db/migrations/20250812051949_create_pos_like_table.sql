-- migrate:up
CREATE TABLE IF NOT EXISTS post_likes (
  id SERIAL PRIMARY KEY,
  comment_id INT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (comment_id) REFERENCES comments (id),
  FOREIGN KEY (user_id) REFERENCES users (id)
)


-- migrate:down
DROP TABLE IF EXISTS post_likes;