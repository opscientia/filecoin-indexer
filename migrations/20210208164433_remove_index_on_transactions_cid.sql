-- +goose Up
DROP INDEX index_transactions_on_cid;

-- +goose Down
CREATE UNIQUE INDEX index_transactions_on_cid ON transactions (cid);
