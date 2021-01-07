package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type eventStore struct {
	db *gorm.DB
}

// Create stores an event record
func (es *eventStore) Create(event *model.Event) error {
	return es.db.Create(event).Error
}

// FindAll retrieves all events
func (es *eventStore) FindAll(height string) (*[]model.Event, error) {
	var events []model.Event

	tx := es.db

	if height != "" {
		tx = tx.Where("height = ?", height)
	}

	err := tx.Order("height DESC, kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// FindAllByMinerAddress retrieves all events for a given miner address
func (es *eventStore) FindAllByMinerAddress(address string, height string) (*[]model.Event, error) {
	var events []model.Event

	tx := es.db.Where("miner_address = ?", address)

	if height != "" {
		tx = tx.Where("height = ?", height)
	}

	err := tx.Order("height DESC, kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}
