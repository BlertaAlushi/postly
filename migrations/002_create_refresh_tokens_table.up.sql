CREATE TABLE IF NOT EXISTS refresh_tokens (
                                id SERIAL PRIMARY KEY,
                                user_id BIGINT NOT NULL,
                                token_hash TEXT NOT NULL,
                                expires_at TIMESTAMPTZ NOT NULL,
                                revoked BOOLEAN NOT NULL DEFAULT FALSE,
                                created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                                updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_refresh_token_hash
    ON refresh_tokens(token_hash);

CREATE INDEX IF NOT EXISTS idx_refresh_token_user
    ON refresh_tokens(user_id);