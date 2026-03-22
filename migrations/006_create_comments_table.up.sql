CREATE TABLE IF NOT EXISTS comments (
                                        id BIGSERIAL PRIMARY KEY,
                                        user_id BIGINT NOT NULL,
                                        post_id BIGINT NOT NULL,
                                        content TEXT NOT NULL,
                                        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_comments_user_id ON comments(user_id);
CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);

ALTER TABLE comments
    ADD CONSTRAINT fk_comments_user
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE comments
    ADD CONSTRAINT fk_comments_post
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE;