package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/stretchr/stew/slice"

	"github.com/figment-networks/filecoin-indexer/model/types"
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

	dealIDs, err := t.store.Event.DealIDsByKind(types.NewDealEvent)
	if err != nil {
		return err
	}

	for _, event := range payload.Events {
		if event.Kind == types.NewDealEvent && slice.Contains(dealIDs, event.Data["deal_id"]) {
			continue
		}

		err := t.store.Event.Create(event)
		if err != nil {
			return err
		}
	}

	return nil
}
