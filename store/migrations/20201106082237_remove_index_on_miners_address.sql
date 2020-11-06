-- +goose Up
DROP INDEX index_miners_on_address;

-- +goose Down
CREATE UNIQUE INDEX index_miners_on_address ON miners (address);
