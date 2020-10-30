package indexing

import (
	"context"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/indexing-engine/pipeline"
)

// MinerParserTask transforms raw miner data
type MinerParserTask struct{}

// NewMinerParserTask creates the task
func NewMinerParserTask() pipeline.Task {
	return &MinerParserTask{}
}

// GetName returns the task name
func (t *MinerParserTask) GetName() string {
	return "MinerParser"
}

// Run performs the task
func (t *MinerParserTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	for i, address := range payload.MinersAddresses {
		sectorSize := uint64(payload.MinersInfo[i].SectorSize)
		rawBytePower := payload.MinersPower[i].MinerPower.RawBytePower.Uint64()
		qualityAdjPower := payload.MinersPower[i].MinerPower.QualityAdjPower.Uint64()
		totalPower := payload.MinersPower[i].TotalPower.QualityAdjPower.Uint64()
		relativePower := float32(float64(qualityAdjPower) / float64(totalPower))

		score := calculateScore(relativePower, sectorSize)

		miner := model.Miner{
			Address:         address.String(),
			SectorSize:      &sectorSize,
			RawBytePower:    &rawBytePower,
			QualityAdjPower: &qualityAdjPower,
			RelativePower:   &relativePower,
			Score:           &score,
		}

		payload.Miners = append(payload.Miners, &miner)
	}

	return nil
}

func calculateScore(relativePower float32, sectorSize uint64) uint32 {
	const sectorSizeBaseline = 32 * 1024 * 1024 * 1024

	powerScore := relativePower
	sectorSizeScore := sectorSize / sectorSizeBaseline

	return uint32(powerScore*1000) + uint32(sectorSizeScore*10)
}
