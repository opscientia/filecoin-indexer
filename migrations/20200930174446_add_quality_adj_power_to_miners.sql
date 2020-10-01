-- +goose Up
ALTER TABLE miners
ADD quality_adj_power BIGINT;

-- +goose Down
ALTER TABLE miners
DROP COLUMN quality_adj_power;
