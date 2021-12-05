-- +goose Up
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

ALTER TABLE products
    ADD CONSTRAINT fk_category_id
        FOREIGN KEY(category_id)
            REFERENCES categories(id);

-- +goose Down
DROP TABLE products;
