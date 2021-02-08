package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type transactionStore struct {
	db *gorm.DB
}

// Create bulk-inserts the transaction records
func (ts *transactionStore) Create(transactions []*model.Transaction) error {
	return ts.db.Create(transactions).Error
}

// FindAll retrieves all transactions
func (ts *transactionStore) FindAll(height string, p Pagination) (*PaginatedResult, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	scope := ts.db.Table("transactions").Order("height DESC")

	if height != "" {
		scope = scope.Where("height = ?", height)
	}

	var count int64
	if err := scope.Count(&count).Error; err != nil {
		return nil, err
	}

	var transactions []model.Transaction

	err := scope.
		Offset(p.offset()).
		Limit(p.limit()).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalCount: count,
		Records:    transactions,
	}

	return result.update(), nil
}

// FindAllByAddress retrieves all transactions for the given addresses
func (ts *transactionStore) FindAllByAddresses(addresses []string, height string, p Pagination) (*PaginatedResult, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	scope := ts.db.
		Table("transactions").
		Where(`"from" IN ? OR "to" IN ?`, addresses, addresses).
		Order("height DESC")

	if height != "" {
		scope = scope.Where("height = ?", height)
	}

	var count int64
	if err := scope.Count(&count).Error; err != nil {
		return nil, err
	}

	var transactions []model.Transaction

	err := scope.
		Offset(p.offset()).
		Limit(p.limit()).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalCount: count,
		Records:    transactions,
	}

	return result.update(), nil
}

// CountSentByAddress retrieves sent transactions for the given addresses
func (ts *transactionStore) CountSentByAddresses(addresses []string) (int64, error) {
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

// CountReceivedByAddress retrieves received transactions for the given addresses
func (ts *transactionStore) CountReceivedByAddresses(addresses []string) (int64, error) {
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
