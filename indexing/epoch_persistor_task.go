package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// EpochPersistorTask stores epochs in the database
type EpochPersistorTask struct {
	store    *store.Store
	observer metrics.Observer
}

// EpochPersistorTaskName represents the name of the task
const EpochPersistorTaskName = "EpochPersistor"

// NewEpochPersistorTask creates the task
func NewEpochPersistorTask(store *store.Store) pipeline.Task {
	return &EpochPersistorTask{
		store:    store,
		observer: pipelineTaskDuration.WithLabels(EpochPersistorTaskName),
	}
}

// GetName returns the task name
func (t *EpochPersistorTask) GetName() string {
	return EpochPersistorTaskName
}

// Run performs the task
func (t *EpochPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	payload := p.(*payload)

	if err := t.store.Epoch.Create(payload.Epoch); err != nil {
		return err
	}

	return nil
}
