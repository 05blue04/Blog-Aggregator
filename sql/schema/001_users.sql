-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;
