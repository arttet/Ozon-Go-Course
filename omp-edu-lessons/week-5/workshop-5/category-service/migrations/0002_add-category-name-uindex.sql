-- +goose NO TRANSACTION
-- +goose Up
CREATE UNIQUE INDEX CONCURRENTLY category_name_uindex ON category (name);

-- +goose Down
DROP INDEX CONCURRENTLY category_name_uindex;
