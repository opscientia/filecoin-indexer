package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type transactionStore struct {
	db *gorm.DB
}

// FindOrCreate retrieves or stores a transaction record
func (ts *transactionStore) FindOrCreate(transaction *model.Transaction) error {
	err := ts.db.
		Where(model.Transaction{CID: transaction.CID}).
		FirstOrCreate(transaction).
		Error

	if err != nil {
		return err
	}

	return nil
}
