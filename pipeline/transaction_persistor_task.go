package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// TransactionPersistorTask stores transactions in the database
type TransactionPersistorTask struct {
	store *store.Store
}

// TransactionPersistorTaskName represents the name of the task
const TransactionPersistorTaskName = "TransactionPersistor"

// NewTransactionPersistorTask creates the task
func NewTransactionPersistorTask(store *store.Store) pipeline.Task {
	return &TransactionPersistorTask{store: store}
}

// GetName returns the task name
func (t *TransactionPersistorTask) GetName() string {
	return TransactionPersistorTaskName
}

// Run performs the task
func (t *TransactionPersistorTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if len(payload.Transactions) == 0 {
		return nil
	}

	err := t.store.Transaction.Create(payload.Transactions)
	if err != nil {
		return err
	}

	return nil
}
