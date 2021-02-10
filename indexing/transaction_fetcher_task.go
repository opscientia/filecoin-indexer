package indexing

import (
	"context"
	"sync"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"golang.org/x/sync/errgroup"

	"github.com/figment-networks/filecoin-indexer/client"
)

// TransactionFetcherTask fetches raw transaction data
type TransactionFetcherTask struct {
	client   *client.Client
	observer metrics.Observer
}

// TransactionFetcherTaskName represents the name of the task
const TransactionFetcherTaskName = "TransactionFetcher"

// NewTransactionFetcherTask creates the task
func NewTransactionFetcherTask(client *client.Client) pipeline.Task {
	return &TransactionFetcherTask{
		client:   client,
		observer: pipelineTaskDuration.WithLabels(TransactionFetcherTaskName),
	}
}

// GetName returns the task name
func (t *TransactionFetcherTask) GetName() string {
	return TransactionFetcherTaskName
}

// Run performs the task
func (t *TransactionFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	payload := p.(*payload)

	eg, _ := errgroup.WithContext(ctx)
	mutex := sync.Mutex{}

	for _, minerAddress := range payload.MinersAddresses {
		func(addr address.Address) {
			eg.Go(func() error {
				return fetchTransactionCIDs(addr, t.client, payload, &mutex)
			})
		}(minerAddress)
	}

	err := eg.Wait()
	if err != nil {
		return err
	}

	payload.TransactionsMessages = make([]*types.Message, len(payload.TransactionsCIDs))

	for i := range payload.TransactionsCIDs {
		func(index int) {
			eg.Go(func() error {
				return fetchTransactionMessage(index, t.client, payload)
			})
		}(i)
	}

	return eg.Wait()
}

func fetchTransactionCIDs(addr address.Address, c *client.Client, p *payload, m *sync.Mutex) error {
	tsk := p.EpochTipset.Key()

	cids, err := c.Transaction.GetCIDsByAddress(addr, tsk, p.currentHeight)
	if err != nil {
		return err
	}

	m.Lock()
	p.TransactionsCIDs = append(p.TransactionsCIDs, cids...)
	m.Unlock()

	return nil
}

func fetchTransactionMessage(index int, c *client.Client, p *payload) error {
	cid := p.TransactionsCIDs[index]

	msg, err := c.Transaction.GetMessage(cid)
	if err != nil {
		return err
	}

	p.TransactionsMessages[index] = msg

	return nil
}
