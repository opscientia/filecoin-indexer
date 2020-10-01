-- +goose Up
ALTER TABLE miners
ADD sector_size BIGINT;

-- +goose Down
ALTER TABLE miners
DROP COLUMN sector_size;
