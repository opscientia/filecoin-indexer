package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// BeginTransactionTask starts a database transaction
type BeginTransactionTask struct {
	store    *store.Store
	observer metrics.Observer
}

// BeginTransactionTaskName represents the name of the task
const BeginTransactionTaskName = "BeginTransaction"

// NewBeginTransactionTask creates the task
func NewBeginTransactionTask(store *store.Store) pipeline.Task {
	return &BeginTransactionTask{
		store:    store,
		observer: pipelineTaskDuration.WithLabels(BeginTransactionTaskName),
	}
}

// GetName returns the task name
func (t *BeginTransactionTask) GetName() string {
	return BeginTransactionTaskName
}

// Run performs the task
func (t *BeginTransactionTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	err := t.store.Begin()
	if err != nil {
		return err
	}

	return nil
}
