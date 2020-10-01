-- +goose Up
ALTER TABLE miners
ADD raw_byte_power BIGINT;

-- +goose Down
ALTER TABLE miners
DROP COLUMN raw_byte_power;
