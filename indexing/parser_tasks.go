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

	for _, address := range payload.MinersAddresses {
		sectorSize := uint64(payload.MinersInfo[address].SectorSize)
		rawBytePower := payload.MinersPower[address].MinerPower.RawBytePower.Uint64()
		qualityAdjPower := payload.MinersPower[address].MinerPower.QualityAdjPower.Uint64()
		totalPower := payload.MinersPower[address].TotalPower.QualityAdjPower.Uint64()
		relativePower := float64(qualityAdjPower) / float64(totalPower)

		miner := model.Miner{
			Address:         address.String(),
			SectorSize:      &sectorSize,
			RawBytePower:    &rawBytePower,
			QualityAdjPower: &qualityAdjPower,
			RelativePower:   &relativePower,
		}

		payload.Miners = append(payload.Miners, &miner)
	}

	return nil
}
