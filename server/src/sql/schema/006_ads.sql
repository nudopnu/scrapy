-- +goose Up
CREATE TABLE ads (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    price TEXT,
    description TEXT,
    location TEXT,
    postal_code TEXT,
    category_id TEXT,
    posted_at TEXT,
    link TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE ads;