package indexing

import (
	"context"
	"errors"
	"strconv"

	"github.com/figment-networks/indexing-engine/pipeline"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/model/types"
	"github.com/figment-networks/filecoin-indexer/store"
)

// EventSequencerTask creates network events
type EventSequencerTask struct {
	store *store.Store
}

// NewEventSequencerTask creates the task
func NewEventSequencerTask(store *store.Store) pipeline.Task {
	return &EventSequencerTask{store: store}
}

// GetName returns the task name
func (t *EventSequencerTask) GetName() string {
	return "EventSequencer"
}

// Run performs the task
func (t *EventSequencerTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	err := t.trackStorageCapacityChanges(payload)
	if err != nil {
		return err
	}

	return nil
}

func (t *EventSequencerTask) trackStorageCapacityChanges(p *payload) error {
	for _, miner := range p.Miners {
		oldMiner, err := t.store.Miner.FindAtPreviousHeight(miner.Address, p.currentHeight)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			return err
		}

		if *miner.RawBytePower != *oldMiner.RawBytePower {
			event := model.Event{
				Height:       &p.currentHeight,
				MinerAddress: miner.Address,
				Kind:         types.StorageCapacityChangeEvent,

				Data: map[string]interface{}{
					"from": strconv.FormatUint(*oldMiner.RawBytePower, 10),
					"to":   strconv.FormatUint(*miner.RawBytePower, 10),
				},
			}

			p.Events = append(p.Events, &event)
		}
	}

	return nil
}
