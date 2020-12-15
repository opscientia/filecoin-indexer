package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	ID        uint64          `json:"-"`
	CID       string          `json:"cid" gorm:"column:cid"`
	Height    *int64          `json:"height"`
	From      string          `json:"from"`
	To        string          `json:"to"`
	Value     decimal.Decimal `json:"value"`
	Method    string          `json:"method"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
}
