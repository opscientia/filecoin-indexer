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
