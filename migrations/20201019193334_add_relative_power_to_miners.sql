-- +goose Up
ALTER TABLE miners
ADD relative_power DOUBLE PRECISION;

-- +goose Down
ALTER TABLE miners
DROP COLUMN relative_power;
