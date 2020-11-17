package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// EpochPersistorTask stores epochs in the database
type EpochPersistorTask struct {
	store *store.Store
}

// NewEpochPersistorTask creates the task
func NewEpochPersistorTask(store *store.Store) pipeline.Task {
	return &EpochPersistorTask{store: store}
}

// GetName returns the task name
func (t *EpochPersistorTask) GetName() string {
	return "MinerPersistor"
}

// Run performs the task
func (t *EpochPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if err := t.store.Epoch.Create(payload.Epoch); err != nil {
		return err
	}

	return nil
}
