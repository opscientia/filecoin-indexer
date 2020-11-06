package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/score"
)

// EpochParserTask transforms raw epoch data
type EpochParserTask struct{}

// NewEpochParserTask creates the task
func NewEpochParserTask() pipeline.Task {
	return &EpochParserTask{}
}

// GetName returns the task name
func (t *EpochParserTask) GetName() string {
	return "EpochParser"
}

// Run performs the task
func (t *EpochParserTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	var blocksCount uint16
	if payload.EpochTipset != nil {
		blocksCount = uint16(len(payload.EpochTipset.Blocks()))
	}

	payload.Epoch = &model.Epoch{
		Height:      &payload.currentHeight,
		BlocksCount: &blocksCount,
	}

	return nil
}

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

	if payload.IsProcessed() {
		return nil
	}

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

		score := score.CalculateScore(score.Variables{
			SectorSize:    sectorSize,
			RelativePower: relativePower,
			FaultsCount:   faultsCount,
		})

		miner := model.Miner{
			Height:          &payload.currentHeight,
			Address:         address.String(),
			SectorSize:      &sectorSize,
			RawBytePower:    &rawBytePower,
			QualityAdjPower: &qualityAdjPower,
			RelativePower:   &relativePower,
			FaultsCount:     &faultsCount,
			Score:           &score,
		}

		payload.Miners = append(payload.Miners, &miner)
	}

	return nil
}
