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
	var transactions []model.Transaction

	err := ts.db.Order("height DESC").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// FindAllByHeight retrieves all transactions for a given height
func (ts *transactionStore) FindAllByHeight(height string) (*[]model.Transaction, error) {
	var transactions []model.Transaction

	err := ts.db.Where("height = ?", height).Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// FindAllByAddress retrieves all transactions for given addresses
func (ts *transactionStore) FindAllByAddress(addresses ...string) (*[]model.Transaction, error) {
	var transactions []model.Transaction

	err := ts.db.
		Where(`"from" IN ? OR "to" IN ?`, addresses, addresses).
		Order("height DESC").
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// FindAllByAddressAndHeight retrieves all transactions for given addresses and height
func (ts *transactionStore) FindAllByAddressAndHeight(height string, addresses ...string) (*[]model.Transaction, error) {
	var transactions []model.Transaction

	err := ts.db.
		Where(`"from" IN ? OR "to" IN ?`, addresses, addresses).
		Where("height = ?", height).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// CountSentByAddress retrieves sent transactions for given addresses
func (ts *transactionStore) CountSentByAddress(addresses ...string) (int64, error) {
	var count int64

	err := ts.db.
		Table("transactions").
		Where(`"from" IN ?`, addresses).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// CountReceivedByAddress retrieves received transactions for given addresses
func (ts *transactionStore) CountReceivedByAddress(addresses ...string) (int64, error) {
	var count int64

	err := ts.db.
		Table("transactions").
		Where(`"to" IN ?`, addresses).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
