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

	for i, minerAddress := range payload.MinersAddresses {
		miner := model.Miner{}

		sectorSize := uint64(payload.MinersInfo[i].SectorSize)
		rawBytePower := payload.MinersPower[i].MinerPower.RawBytePower.Uint64()
		qualityAdjPower := payload.MinersPower[i].MinerPower.QualityAdjPower.Uint64()

		t.store.Db.
			Where(model.Miner{
				Address: minerAddress.String(),
			}).
			Assign(model.Miner{
				SectorSize:      &sectorSize,
				RawBytePower:    &rawBytePower,
				QualityAdjPower: &qualityAdjPower,
			}).
			FirstOrCreate(&miner)
	}

	return nil
}
