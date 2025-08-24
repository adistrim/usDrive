CREATE TYPE file_status AS ENUM ('pending_upload', 'active', 'error');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    google_id VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    full_name VARCHAR(255),
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE files (
    id BIGSERIAL PRIMARY KEY,
    parent_id BIGINT,
    name TEXT NOT NULL,
    is_folder BOOLEAN NOT NULL DEFAULT FALSE,
    mime_type VARCHAR(255),
    size_bytes BIGINT,
    status file_status NOT NULL DEFAULT 'pending_upload',
    storage_path TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_parent FOREIGN KEY(parent_id) REFERENCES files(id) ON DELETE CASCADE
);

CREATE INDEX idx_files_parent_id ON files(parent_id);
