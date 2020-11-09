-- +goose Up
ALTER TABLE miners
ADD faults_count INTEGER NOT NULL;

-- +goose Down
ALTER TABLE miners
DROP COLUMN faults_count;
