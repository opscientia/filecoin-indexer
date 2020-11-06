-- +goose Up
ALTER TABLE miners
ADD raw_byte_power BIGINT NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN raw_byte_power;
