package store

import (
	"github.com/figment-networks/indexing-engine/metrics"
	m "github.com/figment-networks/indexing-engine/pipeline/metrics"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type epochStore struct {
	db *gorm.DB
}

// Create inserts the epoch record
func (es *epochStore) Create(epoch *model.Epoch) error {
	observer := m.DatabaseQueryDuration.WithLabels("epochStore_Create")

	timer := metrics.NewTimer(observer)
	defer timer.ObserveDuration()

	return es.db.Create(epoch).Error
}

// LastHeight returns the last indexed height
func (es *epochStore) LastHeight() (int64, error) {
	var result int64

	err := es.db.Table("epochs").Select("MAX(height)").Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
