-- +goose Up
ALTER TABLE miners
ADD deals_count INTEGER NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN deals_count;
