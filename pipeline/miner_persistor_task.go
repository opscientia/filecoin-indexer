package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// MinerPersistorTask stores miners in the database
type MinerPersistorTask struct {
	store *store.Store
}

// MinerPersistorTaskName represents the name of the task
const MinerPersistorTaskName = "MinerPersistor"

// NewMinerPersistorTask creates the task
func NewMinerPersistorTask(store *store.Store) pipeline.Task {
	return &MinerPersistorTask{store: store}
}

// GetName returns the task name
func (t *MinerPersistorTask) GetName() string {
	return MinerPersistorTaskName
}

// Run performs the task
func (t *MinerPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if len(payload.Miners) == 0 {
		return nil
	}

	err := t.store.Miner.Create(payload.Miners)
	if err != nil {
		return err
	}

	return nil
}
