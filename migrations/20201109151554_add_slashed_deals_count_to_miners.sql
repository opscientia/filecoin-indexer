-- +goose Up
ALTER TABLE miners
ADD slashed_deals_count INTEGER NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN slashed_deals_count;
