package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/score"
)

// MinerParserTask transforms raw miner data
type MinerParserTask struct{}

// MinerParserTaskName represents the name of the task
const MinerParserTaskName = "MinerParser"

// NewMinerParserTask creates the task
func NewMinerParserTask() pipeline.Task {
	return &MinerParserTask{}
}

// GetName returns the task name
func (t *MinerParserTask) GetName() string {
	return MinerParserTaskName
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

		fc, err := payload.MinersFaults[i].Count()
		if err != nil {
			return err
		}
		faultsCount := uint32(fc)

		dealsCount := payload.DealsCount[address.String()]
		slashedDealsCount := payload.DealsSlashedCount[address.String()]

		score := score.CalculateScore(score.Variables{
			SectorSize:        sectorSize,
			RelativePower:     relativePower,
			FaultsCount:       faultsCount,
			SlashedDealsCount: slashedDealsCount,
		})

		miner := model.Miner{
			Height:            &payload.currentHeight,
			Address:           address.String(),
			SectorSize:        &sectorSize,
			RawBytePower:      &rawBytePower,
			QualityAdjPower:   &qualityAdjPower,
			RelativePower:     &relativePower,
			FaultsCount:       &faultsCount,
			DealsCount:        &dealsCount,
			SlashedDealsCount: &slashedDealsCount,
			Score:             &score,
		}

		payload.Miners = append(payload.Miners, &miner)
	}

	return nil
}
