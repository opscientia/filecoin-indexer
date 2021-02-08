-- +goose Up
CREATE INDEX index_events_on_kind ON events (kind);

-- +goose Down
DROP INDEX index_events_on_kind;
