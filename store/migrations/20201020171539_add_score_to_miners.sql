-- +goose Up
ALTER TABLE miners
ADD score INTEGER NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN score;
