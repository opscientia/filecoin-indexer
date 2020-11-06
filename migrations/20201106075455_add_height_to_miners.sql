-- +goose Up
ALTER TABLE miners
ADD height INTEGER NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN height;
