CREATE TABLE IF NOT EXISTS likes (
                                     id BIGSERIAL PRIMARY KEY,
                                     user_id BIGINT NOT NULL,
                                     post_id BIGINT NOT NULL,
                                     CONSTRAINT unique_like UNIQUE (user_id, post_id)
);

CREATE INDEX IF NOT EXISTS idx_likes_user_id ON likes(user_id);
CREATE INDEX IF NOT EXISTS idx_likes_post_id ON likes(post_id);

ALTER TABLE likes
    ADD CONSTRAINT fk_likes_user
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE likes
    ADD CONSTRAINT fk_likes_post
        FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE;