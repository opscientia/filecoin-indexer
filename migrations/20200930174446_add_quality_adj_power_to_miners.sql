-- +goose Up
ALTER TABLE miners
ADD quality_adj_power BIGINT NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN quality_adj_power;
