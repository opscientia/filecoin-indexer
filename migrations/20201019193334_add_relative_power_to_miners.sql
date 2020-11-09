-- +goose Up
ALTER TABLE miners
ADD relative_power REAL NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN relative_power;
