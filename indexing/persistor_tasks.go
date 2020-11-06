package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

// EpochPersistorTask stores epochs in the database
type EpochPersistorTask struct {
	store *store.Store
}

// NewEpochPersistorTask creates the task
func NewEpochPersistorTask(store *store.Store) pipeline.Task {
	return &EpochPersistorTask{store: store}
}

// GetName returns the task name
func (t *EpochPersistorTask) GetName() string {
	return "MinerPersistor"
}

// Run performs the task
func (t *EpochPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	t.store.Db.Create(payload.Epoch)

	return nil
}

// MinerPersistorTask stores miners in the database
type MinerPersistorTask struct {
	store *store.Store
}

// NewMinerPersistorTask creates the task
func NewMinerPersistorTask(store *store.Store) pipeline.Task {
	return &MinerPersistorTask{store: store}
}

// GetName returns the task name
func (t *MinerPersistorTask) GetName() string {
	return "MinerPersistor"
}

// Run performs the task
func (t *MinerPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if payload.IsProcessed() {
		return nil
	}

	for _, miner := range payload.Miners {
		m := model.Miner{}

		t.store.Db.
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
			FirstOrCreate(&m)
	}

	return nil
}
