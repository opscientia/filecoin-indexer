package model

import "time"

// Epoch represents a blockchain height
type Epoch struct {
	ID          uint64    `json:"-"`
	Height      *int64    `json:"height"`
	BlocksCount *uint16   `json:"blocks_count"`
	CreatedAt   time.Time `json:"processed_at"`
	UpdatedAt   time.Time `json:"-"`
}

// TableName returns the custom table name
func (Epoch) TableName() string {
	return "epochs"
}
