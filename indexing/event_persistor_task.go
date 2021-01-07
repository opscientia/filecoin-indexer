package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// EventPersistorTask stores events in the database
type EventPersistorTask struct {
	store *store.Store
}

// NewEventPersistorTask creates the task
func NewEventPersistorTask(store *store.Store) pipeline.Task {
	return &EventPersistorTask{store: store}
}

// GetName returns the task name
func (t *EventPersistorTask) GetName() string {
	return "EventPersistor"
}

// Run performs the task
func (t *EventPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	for _, event := range payload.Events {
		err := t.store.Event.Create(event)
		if err != nil {
			return err
		}
	}

	return nil
}
