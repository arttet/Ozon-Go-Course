-- +goose Up
CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    exec_duration bigint NOT NULL,
    started_at TIMESTAMP
);

-- +goose Down
DROP TABLE task;
