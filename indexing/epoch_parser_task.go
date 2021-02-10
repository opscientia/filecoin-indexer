package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/model"
)

// EpochParserTask transforms raw epoch data
type EpochParserTask struct {
	observer metrics.Observer
}

// EpochParserTaskName represents the name of the task
const EpochParserTaskName = "EpochParser"

// NewEpochParserTask creates the task
func NewEpochParserTask() pipeline.Task {
	return &EpochParserTask{
		observer: pipelineTaskDuration.WithLabels(EpochParserTaskName),
	}
}

// GetName returns the task name
func (t *EpochParserTask) GetName() string {
	return EpochParserTaskName
}

// Run performs the task
func (t *EpochParserTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	payload := p.(*payload)
	epochHeight := int64(payload.EpochTipset.Height())

	var blocksCount uint16
	if epochHeight == payload.currentHeight {
		blocksCount = uint16(len(payload.EpochTipset.Blocks()))
	} else {
		blocksCount = 0 // Null-block height
	}

	payload.Epoch = &model.Epoch{
		Height:      &payload.currentHeight,
		BlocksCount: &blocksCount,
	}

	return nil
}
