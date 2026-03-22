CREATE TABLE IF NOT EXISTS follows (
                                       id SERIAL PRIMARY KEY,
                                       user_id BIGINT NOT NULL,
                                       follow_id BIGINT NOT NULL,
                                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

                                       CONSTRAINT unique_follow UNIQUE (user_id, follow_id),
                                       CONSTRAINT no_self_follow CHECK (user_id <> follow_id)
);

CREATE INDEX IF NOT EXISTS idx_follows_user_id ON follows(user_id);
CREATE INDEX IF NOT EXISTS idx_follows_follow_id ON follows(follow_id);