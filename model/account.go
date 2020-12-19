package model

import (
	"github.com/shopspring/decimal"
)

// Account represents a blockchain account
type Account struct {
	ID                   string          `json:"id"`
	PublicKey            string          `json:"public_key"`
	Balance              decimal.Decimal `json:"balance"`
	Nonce                uint64          `json:"nonce"`
	TransactionsSent     int64           `json:"transactions_sent"`
	TransactionsReceived int64           `json:"transactions_received"`
}
