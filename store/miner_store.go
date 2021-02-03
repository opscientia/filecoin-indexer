package store

import (
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
)

type minerStore struct {
	db *gorm.DB
}

// CreateOrUpdate stores or updates a miner record
func (ms *minerStore) CreateOrUpdate(miner *model.Miner) (*model.Miner, error) {
	result := model.Miner{}

	err := ms.db.
		Where(model.Miner{
			Height:  miner.Height,
			Address: miner.Address,
		}).
		Assign(model.Miner{
			SectorSize:        miner.SectorSize,
			RawBytePower:      miner.RawBytePower,
			QualityAdjPower:   miner.QualityAdjPower,
			RelativePower:     miner.RelativePower,
			FaultsCount:       miner.FaultsCount,
			DealsCount:        miner.DealsCount,
			SlashedDealsCount: miner.SlashedDealsCount,
			Score:             miner.Score,
		}).
		FirstOrCreate(&result).
		Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FindByHeight retrieves a miner record for a given height
func (ms *minerStore) FindByHeight(address string, height int64) (*model.Miner, error) {
	var miner model.Miner

	err := ms.db.
		Where("address = ? AND height = ?", address, height).
		Take(&miner).
		Error

	if err != nil {
		return nil, err
	}

	return &miner, nil
}

// FindAllAtPreviousHeight retrieves all miners at a height lower than the given height
func (ms *minerStore) FindAllAtPreviousHeight(height int64) ([]model.Miner, error) {
	var miners []model.Miner

	err := ms.db.
		Distinct("ON(address) *").
		Where("height < ?", height).
		Order("address, height DESC").
		Find(&miners).
		Error

	if err != nil {
		return nil, err
	}

	return miners, nil
}

// FindAllByHeight retrieves all miners for a given height
func (ms *minerStore) FindAllByHeight(height int64, p Pagination) (*PaginatedResult, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}

	scope := ms.db.Table("miners").Where("height = ?", height)

	var count int64
	if err := scope.Count(&count).Error; err != nil {
		return nil, err
	}

	var miners []model.Miner

	err := scope.
		Offset(p.offset()).
		Limit(p.limit()).
		Find(&miners).
		Error

	if err != nil {
		return nil, err
	}

	result := &PaginatedResult{
		Page:       p.Page,
		Limit:      p.Limit,
		TotalCount: count,
		Records:    miners,
	}

	return result.update(), nil
}

// FindTop100ByHeight retrieves top 100 miners for a given height
func (ms *minerStore) FindTop100ByHeight(height int64) ([]model.Miner, error) {
	var miners []model.Miner

	err := ms.db.
		Where("height = ?", height).
		Order("score DESC").
		Limit(100).
		Find(&miners).
		Error

	if err != nil {
		return nil, err
	}

	return miners, nil
}
