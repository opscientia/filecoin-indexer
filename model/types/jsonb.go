package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JSONB represents a binary JSON column
type JSONB map[string]interface{}

// Value converts a map into JSON data
func (j *JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan converts JSON data into a map
func (j *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, j)
}
