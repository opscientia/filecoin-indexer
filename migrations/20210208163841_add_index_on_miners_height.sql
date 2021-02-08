-- +goose Up
CREATE INDEX index_miners_on_height ON miners (height);

-- +goose Down
DROP INDEX index_miners_on_height;
