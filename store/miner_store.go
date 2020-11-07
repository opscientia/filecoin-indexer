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
			SectorSize:      miner.SectorSize,
			RawBytePower:    miner.RawBytePower,
			QualityAdjPower: miner.QualityAdjPower,
			RelativePower:   miner.RelativePower,
			FaultsCount:     miner.FaultsCount,
			Score:           miner.Score,
		}).
		FirstOrCreate(&result).
		Error

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FindAllByHeight retrieves all miners for a given height
func (ms *minerStore) FindAllByHeight(height int64) (*[]model.Miner, error) {
	var miners []model.Miner

	err := ms.db.Where("height = ?", height).Find(&miners).Error
	if err != nil {
		return nil, err
	}

	return &miners, nil
}

// FindTop100ByHeight retrieves top 100 miners for a given height
func (ms *minerStore) FindTop100ByHeight(height int64) (*[]model.Miner, error) {
	var miners []model.Miner

	err := ms.db.Where("height = ?", height).Order("score DESC").Limit(100).Find(&miners).Error
	if err != nil {
		return nil, err
	}

	return &miners, nil
}
