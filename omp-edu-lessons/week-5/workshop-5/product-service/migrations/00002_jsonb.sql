-- +goose Up
ALTER TABLE products
    ADD info jsonb;


-- +goose Down
ALTER TABLE products
    DROP info;
