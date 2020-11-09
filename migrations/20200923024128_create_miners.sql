-- +goose Up
CREATE TABLE miners (
  id         BIGSERIAL NOT NULL PRIMARY KEY,
  address    TEXT NOT NULL,

  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- +goose Down
DROP TABLE miners;
