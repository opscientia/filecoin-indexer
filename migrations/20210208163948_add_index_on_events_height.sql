-- +goose Up
CREATE INDEX index_events_on_height ON events (height);

-- +goose Down
DROP INDEX index_events_on_height;
