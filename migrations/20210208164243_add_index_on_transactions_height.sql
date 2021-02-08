-- +goose Up
CREATE INDEX index_transactions_on_height ON transactions (height);

-- +goose Down
DROP INDEX index_transactions_on_height;
