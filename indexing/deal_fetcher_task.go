package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
)

// DealFetcherTask fetches raw deal data
type DealFetcherTask struct {
	client *client.Client
}

// NewDealFetcherTask creates the task
func NewDealFetcherTask(client *client.Client) pipeline.Task {
	return &DealFetcherTask{client: client}
}

// GetName returns the task name
func (t *DealFetcherTask) GetName() string {
	return "DealFetcher"
}

// Run performs the task
func (t *DealFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)
	tsk := payload.EpochTipset.Key()

	deals, err := t.client.Deal.GetMarketDeals(tsk)
	if err != nil {
		return err
	}

	payload.DealsData = deals

	return nil
}
