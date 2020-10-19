package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

const (
	MinerPersistorTaskName = "MinerPersistor"
)

type MinerPersistorTask struct {
	store *store.Store
}

func NewMinerPersistorTask(store *store.Store) pipeline.Task {
	return &MinerPersistorTask{store: store}
}

func (t *MinerPersistorTask) GetName() string {
	return MinerPersistorTaskName
}

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
			}).
			FirstOrCreate(&m)
	}

	return nil
}
