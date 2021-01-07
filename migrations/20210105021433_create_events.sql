-- +goose Up
CREATE TABLE events (
  id            BIGSERIAL NOT NULL PRIMARY KEY,
  height        INTEGER NOT NULL,
  miner_address TEXT NOT NULL,
  kind          TEXT NOT NULL,
  data          JSONB,

  created_at    TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at    TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE events;
