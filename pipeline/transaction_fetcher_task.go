package pipeline

import (
	"context"
	"sync"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"

	"github.com/figment-networks/filecoin-indexer/client"
)

// TransactionFetcherTask fetches raw transaction data
type TransactionFetcherTask struct {
	client *client.Client
	mutex  sync.Mutex
}

// TransactionFetcherTaskName represents the name of the task
const TransactionFetcherTaskName = "TransactionFetcher"

// NewTransactionFetcherTask creates the task
func NewTransactionFetcherTask(client *client.Client) pipeline.Task {
	return &TransactionFetcherTask{client: client}
}

// GetName returns the task name
func (t *TransactionFetcherTask) GetName() string {
	return TransactionFetcherTaskName
}

// Run performs the task
func (t *TransactionFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if err := t.retrieveTransactions(payload); err != nil {
		if err := t.fetchTransactions(ctx, payload); err != nil {
			return err
		}

		if err := t.storeTransactions(payload); err != nil {
			return err
		}
	}

	return nil
}

func (t *TransactionFetcherTask) retrieveTransactions(payload *payload) error {
	return multierr.Combine(
		payload.Retrieve("transactions_cids", &payload.TransactionsCIDs),
		payload.Retrieve("transactions_messages", &payload.TransactionsMessages),
	)
}

func (t *TransactionFetcherTask) storeTransactions(payload *payload) error {
	return multierr.Combine(
		payload.Store("transactions_cids", payload.TransactionsCIDs),
		payload.Store("transactions_messages", payload.TransactionsMessages),
	)
}

func (t *TransactionFetcherTask) fetchTransactions(ctx context.Context, payload *payload) error {
	eg, _ := errgroup.WithContext(ctx)

	for _, minerAddress := range payload.MinersAddresses {
		func(addr address.Address) {
			eg.Go(func() error {
				return t.fetchTransactionCIDs(addr, payload)
			})
		}(minerAddress)
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	payload.TransactionsMessages = make([]*types.Message, len(payload.TransactionsCIDs))

	for i := range payload.TransactionsCIDs {
		func(index int) {
			eg.Go(func() error {
				return t.fetchTransactionMessage(index, payload)
			})
		}(i)
	}

	return eg.Wait()
}

func (t *TransactionFetcherTask) fetchTransactionCIDs(addr address.Address, payload *payload) error {
	tsk := payload.EpochTipset.Key()

	cids, err := t.client.Transaction.GetCIDsByAddress(addr, tsk, payload.currentHeight)
	if err != nil {
		return err
	}

	t.mutex.Lock()
	payload.TransactionsCIDs = append(payload.TransactionsCIDs, cids...)
	t.mutex.Unlock()

	return nil
}

func (t *TransactionFetcherTask) fetchTransactionMessage(index int, payload *payload) error {
	cid := payload.TransactionsCIDs[index]

	msg, err := t.client.Transaction.GetMessage(cid)
	if err != nil {
		return err
	}

	payload.TransactionsMessages[index] = msg

	return nil
}
