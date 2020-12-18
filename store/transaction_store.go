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

// FindAll retrieves all transactions
func (ts *transactionStore) FindAll() (*[]model.Transaction, error) {
	transactions := []model.Transaction{}

	err := ts.db.Order("height DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// FindAllByAddress retrieves all transactions for a given address
func (ts *transactionStore) FindAllByAddress(address string) (*[]model.Transaction, error) {
	transactions := []model.Transaction{}

	err := ts.db.
		Where(`"from" = ? OR "to" = ?`, address, address).
		Order("height DESC").
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
