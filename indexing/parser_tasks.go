package indexing

import (
	"context"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/indexing-engine/pipeline"
)

const (
	MinerParserTaskName = "MinerParser"
)

type MinerParserTask struct{}

func NewMinerParserTask() pipeline.Task {
	return &MinerParserTask{}
}

func (t *MinerParserTask) GetName() string {
	return MinerParserTaskName
}

func (t *MinerParserTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	payload.Miners = make([]*model.Miner, len(payload.MinersAddresses))

	for i, minerAddress := range payload.MinersAddresses {
		sectorSize := uint64(payload.MinersInfo[i].SectorSize)
		rawBytePower := payload.MinersPower[i].MinerPower.RawBytePower.Uint64()
		qualityAdjPower := payload.MinersPower[i].MinerPower.QualityAdjPower.Uint64()
		totalPower := payload.MinersPower[i].TotalPower.QualityAdjPower.Uint64()
		relativePower := float64(qualityAdjPower) / float64(totalPower)

		miner := model.Miner{
			Address:         minerAddress.String(),
			SectorSize:      &sectorSize,
			RawBytePower:    &rawBytePower,
			QualityAdjPower: &qualityAdjPower,
			RelativePower:   &relativePower,
		}

		payload.Miners[i] = &miner
	}

	return nil
}
