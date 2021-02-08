-- +goose Up
CREATE INDEX index_events_on_miner_address ON events (miner_address);

-- +goose Down
DROP INDEX index_events_on_miner_address;
