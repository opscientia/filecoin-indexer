package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// EventPersistorTask stores events in the database
type EventPersistorTask struct {
	store    *store.Store
	observer metrics.Observer
}

// EventPersistorTaskName represents the name of the task
const EventPersistorTaskName = "EventPersistor"

// NewEventPersistorTask creates the task
func NewEventPersistorTask(store *store.Store) pipeline.Task {
	return &EventPersistorTask{
		store:    store,
		observer: pipelineTaskDuration.WithLabels(EventPersistorTaskName),
	}
}

// GetName returns the task name
func (t *EventPersistorTask) GetName() string {
	return EventPersistorTaskName
}

// Run performs the task
func (t *EventPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	payload := p.(*payload)

	if len(payload.Events) == 0 {
		return nil
	}

	err := t.store.Event.Create(payload.Events)
	if err != nil {
		return err
	}

	return nil
}
