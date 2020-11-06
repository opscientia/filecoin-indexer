-- +goose Up
CREATE UNIQUE INDEX index_miners_on_address ON miners (address);

-- +goose Down
DROP INDEX index_miners_on_address;
