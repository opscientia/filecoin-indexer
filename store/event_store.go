package store

import (
	"github.com/figment-networks/indexing-engine/metrics"
	m "github.com/figment-networks/indexing-engine/pipeline/metrics"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/model/types"
)

type eventStore struct {
	db *gorm.DB
}

// Create bulk-inserts the event records
func (es *eventStore) Create(event []*model.Event) error {
	observer := m.DatabaseQueryDuration.WithLabels("eventStore_Create")

	timer := metrics.NewTimer(observer)
	defer timer.ObserveDuration()

	return es.db.Create(event).Error
}

// FindAll retrieves all events
func (es *eventStore) FindAll(height string, kind string, p Pagination) (*PaginatedResult, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	scope := es.db.Table("events").Order("height DESC, kind")

	if height != "" {
		scope = scope.Where("height = ?", height)
	}

	if kind != "" {
		scope = scope.Where("kind = ?", kind)
	}

	var count int64
	if err := scope.Count(&count).Error; err != nil {
		return nil, err
	}

	var events []model.Event

	err := scope.
		Offset(p.offset()).
		Limit(p.limit()).
		Find(&events).
		Error

	if err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalCount: count,
		Records:    events,
	}

	return result.update(), nil
}

// FindAllByMinerAddress retrieves all events for a given miner address
func (es *eventStore) FindAllByMinerAddress(address string, height string, kind string, p Pagination) (*PaginatedResult, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	scope := es.db.
		Table("events").
		Where("miner_address = ?", address).
		Order("height DESC, kind")

	if height != "" {
		scope = scope.Where("height = ?", height)
	}

	if kind != "" {
		scope = scope.Where("kind = ?", kind)
	}

	var count int64
	if err := scope.Count(&count).Error; err != nil {
		return nil, err
	}

	var events []model.Event

	err := scope.
		Offset(p.offset()).
		Limit(p.limit()).
		Find(&events).
		Error

	if err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalCount: count,
		Records:    events,
	}

	return result.update(), nil
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
