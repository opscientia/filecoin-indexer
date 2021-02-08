package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

// CommitTransactionTask commits the database transaction
type CommitTransactionTask struct {
	store *store.Store
}

// NewCommitTransactionTask creates the task
func NewCommitTransactionTask(store *store.Store) pipeline.Task {
	return &CommitTransactionTask{store: store}
}

// GetName returns the task name
func (t *CommitTransactionTask) GetName() string {
	return "CommitTransaction"
}

// Run performs the task
func (t *CommitTransactionTask) Run(ctx context.Context, p pipeline.Payload) error {
	err := t.store.Commit()
	if err != nil {
		return err
	}

	return nil
}
