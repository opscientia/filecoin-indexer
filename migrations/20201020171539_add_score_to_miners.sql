-- +goose Up
ALTER TABLE miners
ADD score INTEGER;

-- +goose Down
ALTER TABLE miners
DROP COLUMN score;
