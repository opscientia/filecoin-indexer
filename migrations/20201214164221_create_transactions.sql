-- +goose Up
CREATE TABLE transactions (
  "id"         BIGSERIAL NOT NULL PRIMARY KEY,
  "cid"        TEXT NOT NULL,
  "height"     INTEGER NOT NULL,
  "from"       TEXT NOT NULL,
  "to"         TEXT NOT NULL,
  "value"      DECIMAL NOT NULL,
  "method"     TEXT NOT NULL,

  "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
  "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX index_transactions_on_cid ON transactions ("cid");

CREATE INDEX index_transactions_on_from ON transactions ("from");
CREATE INDEX index_transactions_on_to ON transactions ("to");

-- +goose Down
DROP TABLE transactions;
