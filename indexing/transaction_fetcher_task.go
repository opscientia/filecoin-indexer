package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/figment-networks/filecoin-indexer/client"
)

// TransactionFetcherTask fetches raw transaction
type TransactionFetcherTask struct {
	client *client.Client
}

// NewTransactionFetcherTask creates the task
func NewTransactionFetcherTask(client *client.Client) pipeline.Task {
	return &TransactionFetcherTask{client: client}
}

// GetName returns the task name
func (t *TransactionFetcherTask) GetName() string {
	return "TransactionFetcher"
}

// Run performs the task
func (t *TransactionFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)
	tsk := payload.EpochTipset.Key()

	for _, addr := range payload.MinersAddresses {
		cids, err := t.client.Transaction.GetCIDsByAddress(addr, tsk, payload.currentHeight)
		if err != nil {
			return err
		}
		payload.TransactionsCIDs = append(payload.TransactionsCIDs, cids...)
	}

	payload.TransactionsMessages = make([]*types.Message, len(payload.TransactionsCIDs))

	for i, cid := range payload.TransactionsCIDs {
		msg, err := t.client.Transaction.GetMessage(cid)
		if err != nil {
			return err
		}
		payload.TransactionsMessages[i] = msg
	}

	return nil
}
