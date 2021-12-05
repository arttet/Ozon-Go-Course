-- +goose Up
CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    started_at TIMESTAMP
);

-- +goose Down
DROP TABLE task;
