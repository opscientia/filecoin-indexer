package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type transactionStore struct {
	db *gorm.DB
}

// Create stores a transaction record
func (ts *transactionStore) Create(transaction *model.Transaction) error {
	return ts.db.Create(transaction).Error
}
