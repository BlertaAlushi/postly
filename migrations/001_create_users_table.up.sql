CREATE TABLE If NOT EXISTS users (
                                     id BIGSERIAL PRIMARY KEY,
                                     email TEXT NOT NULL UNIQUE,
                                     username TEXT NOT NULL UNIQUE,
                                     firstname TEXT NOT NULL,
                                     lastname TEXT NOT NULL,
                                     password TEXT NOT NULL,
                                     is_active BOOLEAN NOT NULL DEFAULT TRUE,
                                     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);