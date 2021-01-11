package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/model/types"
)

type eventStore struct {
	db *gorm.DB
}

// Create stores an event record
func (es *eventStore) Create(event *model.Event) error {
	return es.db.Create(event).Error
}

// FindAll retrieves all events
func (es *eventStore) FindAll(height string, kind string) (*[]model.Event, error) {
	var events []model.Event

	tx := es.db

	if height != "" {
		tx = tx.Where("height = ?", height)
	}

	if kind != "" {
		tx = tx.Where("kind = ?", kind)
	}

	err := tx.Order("height DESC, kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// FindAllByMinerAddress retrieves all events for a given miner address
func (es *eventStore) FindAllByMinerAddress(address string, height string, kind string) (*[]model.Event, error) {
	var events []model.Event

	tx := es.db.Where("miner_address = ?", address)

	if height != "" {
		tx = tx.Where("height = ?", height)
	}

	if kind != "" {
		tx = tx.Where("kind = ?", kind)
	}

	err := tx.Order("height DESC, kind").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// DealIDsByKind returns deal IDs for a given event kind
func (es *eventStore) DealIDsByKind(kind types.EventKind) ([]string, error) {
	var result []string

	err := es.db.
		Table("events").
		Select("data->>'deal_id'").
		Where("kind = ?", kind).
		Scan(&result).
		Error

	if err != nil {
		return nil, err
	}

	return result, nil
}
