-- migrate:up
ALTER TABLE posts DROP CONSTRAINT posts_user_id_fkey;

ALTER TABLE posts
ADD CONSTRAINT posts_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE comments DROP CONSTRAINT comments_post_id_fkey;
ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey;

ALTER TABLE comments
ADD CONSTRAINT comments_post_id_fkey
FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE;

ALTER TABLE comments
ADD CONSTRAINT comments_user_id_fkey
FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE post_likes DROP CONSTRAINT post_likes_comment_id_fkey;
ALTER TABLE post_likes DROP CONSTRAINT post_likes_user_id_fkey;

ALTER TABLE post_likes 
ADD CONSTRAINT post_likes_comment_id_fkey 
FOREIGN KEY (comment_id) REFERENCES comments (id) ON DELETE CASCADE;

ALTER TABLE post_likes 
ADD CONSTRAINT post_likes_user_id_fkey 
FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

ALTER TABLE refresh_tokens DROP CONSTRAINT refresh_tokens_user_id_fkey;
ALTER TABLE refresh_tokens 
ADD CONSTRAINT refresh_tokens_user_id_fkey 
FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- migrate:down
ALTER TABLE posts DROP CONSTRAINT posts_user_id_fkey;
ALTER TABLE posts ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE comments DROP CONSTRAINT comments_post_id_fkey;
ALTER TABLE comments DROP CONSTRAINT comments_user_id_fkey;
ALTER TABLE comments ADD CONSTRAINT comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES posts (id);
ALTER TABLE comments ADD CONSTRAINT comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE post_likes DROP CONSTRAINT post_likes_comment_id_fkey;
ALTER TABLE post_likes DROP CONSTRAINT post_likes_user_id_fkey;
ALTER TABLE post_likes ADD CONSTRAINT post_likes_comment_id_fkey FOREIGN KEY (comment_id) REFERENCES comments (id);
ALTER TABLE post_likes ADD CONSTRAINT post_likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE refresh_tokens DROP CONSTRAINT refresh_tokens_user_id_fkey;
ALTER TABLE refresh_tokens ADD CONSTRAINT refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id);