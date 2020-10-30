package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

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

	for _, miner := range payload.Miners {
		m := model.Miner{}

		t.store.Db.
			Where(model.Miner{
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
