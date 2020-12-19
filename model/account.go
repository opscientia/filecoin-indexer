package model

import (
	"github.com/shopspring/decimal"
)

// Account represents a blockchain account
type Account struct {
	Address string
	Balance decimal.Decimal
}
