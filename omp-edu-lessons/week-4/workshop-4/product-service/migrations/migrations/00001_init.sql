-- +goose Up
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          category_id BIGINT NOT NULL
);

-- +goose Down
DROP TABLE products;
