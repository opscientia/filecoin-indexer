package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"

	"github.com/figment-networks/filecoin-indexer/client"
)

// DealFetcherTask fetches raw deal data
type DealFetcherTask struct {
	client   *client.Client
	observer metrics.Observer
}

// DealFetcherTaskName represents the name of the task
const DealFetcherTaskName = "DealFetcher"

// NewDealFetcherTask creates the task
func NewDealFetcherTask(client *client.Client) pipeline.Task {
	return &DealFetcherTask{
		client:   client,
		observer: pipelineTaskDuration.WithLabels(DealFetcherTaskName),
	}
}

// GetName returns the task name
func (t *DealFetcherTask) GetName() string {
	return DealFetcherTaskName
}

// Run performs the task
func (t *DealFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	timer := metrics.NewTimer(t.observer)
	defer timer.ObserveDuration()

	payload := p.(*payload)
	tsk := payload.EpochTipset.Key()

	deals, err := t.client.Deal.GetMarketDeals(tsk)
	if err != nil {
		return err
	}
	payload.DealsData = deals

	payload.DealsCount = make(map[address.Address]uint32)
	payload.DealsSlashedCount = make(map[address.Address]uint32)

	for dealID, deal := range deals {
		minerAddress := deal.Proposal.Provider

		payload.DealsCount[minerAddress]++

		if deal.State.SlashEpoch != -1 {
			payload.DealsSlashedCount[minerAddress]++
			payload.DealsSlashedIDs = append(payload.DealsSlashedIDs, dealID)
		}
	}

	for address := range payload.DealsCount {
		payload.MinersAddresses = append(payload.MinersAddresses, address)
	}

	return nil
}
