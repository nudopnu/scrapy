-- +goose Up
CREATE TABLE images (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    ad_id TEXT NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    image_number INT NOT NULL,
    label TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE images;