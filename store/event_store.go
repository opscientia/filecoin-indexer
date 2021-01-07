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
func (es *eventStore) FindAll() (*[]model.Event, error) {
	var events []model.Event

	err := es.db.Order("height DESC, kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// FindAllByHeight retrieves all events for a given height
func (es *eventStore) FindAllByHeight(height string) (*[]model.Event, error) {
	var events []model.Event

	err := es.db.Where("height = ?", height).Order("kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// FindAllByMinerAddress retrieves all events for a given miner address
func (es *eventStore) FindAllByMinerAddress(address string) (*[]model.Event, error) {
	var events []model.Event

	err := es.db.
		Where("miner_address = ?", address).
		Order("height DESC, kind").
		Find(&events).
		Error

	if err != nil {
		return nil, err
	}

	return &events, nil
}

// FindAllByMinerAddressAndHeight retrieves all events for a given miner address and height
func (es *eventStore) FindAllByMinerAddressAndHeight(address string, height string) (*[]model.Event, error) {
	var events []model.Event

	err := es.db.
		Where("miner_address = ? AND height = ?", address, height).
		Order("kind").
		Find(&events).
		Error

	if err != nil {
		return nil, err
	}

	return &events, nil
}
