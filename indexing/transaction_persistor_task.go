package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// TransactionPersistorTask stores miners in the database
type TransactionPersistorTask struct {
	store *store.Store
}

// NewTransactionPersistorTask creates the task
func NewTransactionPersistorTask(store *store.Store) pipeline.Task {
	return &TransactionPersistorTask{store: store}
}

// GetName returns the task name
func (t *TransactionPersistorTask) GetName() string {
	return "TransactionPersistor"
}

// Run performs the task
func (t *TransactionPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	for _, transaction := range payload.Transactions {
		err := t.store.Transaction.FindOrCreate(transaction)
		if err != nil {
			return err
		}
	}

	return nil
}
