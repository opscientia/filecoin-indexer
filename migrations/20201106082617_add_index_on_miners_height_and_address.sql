-- +goose Up
CREATE UNIQUE INDEX index_miners_on_height_and_address ON miners (height, address);

-- +goose Down
DROP INDEX index_miners_on_height_and_address;
