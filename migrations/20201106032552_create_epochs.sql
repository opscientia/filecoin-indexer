-- +goose Up
CREATE TABLE epochs (
  id           BIGSERIAL NOT NULL PRIMARY KEY,
  height       INTEGER NOT NULL,
  blocks_count SMALLINT NOT NULL,

  created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at   TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX index_epochs_on_height ON epochs (height);

-- +goose Down
DROP TABLE epochs;
