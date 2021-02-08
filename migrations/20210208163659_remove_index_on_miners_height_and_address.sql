-- +goose Up
DROP INDEX index_miners_on_height_and_address;

-- +goose Down
CREATE UNIQUE INDEX index_miners_on_height_and_address ON miners (height, address);
