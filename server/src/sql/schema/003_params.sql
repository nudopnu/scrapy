-- +goose Up
CREATE TABLE params (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    keyword TEXT NOT NULL,
	location_id TEXT NOT NULL,
	distance INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    udpated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE params;