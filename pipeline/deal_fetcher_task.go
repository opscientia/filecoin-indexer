package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"

	"github.com/figment-networks/filecoin-indexer/client"
)

// DealFetcherTask fetches raw deal data
type DealFetcherTask struct {
	client *client.Client
}

// DealFetcherTaskName represents the name of the task
const DealFetcherTaskName = "DealFetcher"

// NewDealFetcherTask creates the task
func NewDealFetcherTask(client *client.Client) pipeline.Task {
	return &DealFetcherTask{client: client}
}

// GetName returns the task name
func (t *DealFetcherTask) GetName() string {
	return DealFetcherTaskName
}

// Run performs the task
func (t *DealFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	err := payload.Retrieve("deals", &payload.DealsData)
	if err != nil {
		if err := t.fetchDeals(payload); err != nil {
			return err
		}

		err := payload.Store("deals", payload.DealsData)
		if err != nil {
			return err
		}
	}

	t.parseDeals(payload)
	t.parseMinersAddresses(payload)

	return nil
}

func (t *DealFetcherTask) fetchDeals(payload *payload) error {
	tsk := payload.EpochTipset.Key()

	deals, err := t.client.Deal.GetMarketDeals(tsk)
	if err != nil {
		return err
	}

	payload.DealsData = deals

	return nil
}

func (t *DealFetcherTask) parseDeals(payload *payload) {
	payload.DealsCount = make(map[address.Address]uint32)
	payload.DealsSlashedCount = make(map[address.Address]uint32)

	for dealID, deal := range payload.DealsData {
		minerAddress := deal.Proposal.Provider

		payload.DealsCount[minerAddress]++

		if deal.State.SlashEpoch != -1 {
			payload.DealsSlashedCount[minerAddress]++
			payload.DealsSlashedIDs = append(payload.DealsSlashedIDs, dealID)
		}
	}
}

func (t *DealFetcherTask) parseMinersAddresses(payload *payload) {
	for address := range payload.DealsCount {
		payload.MinersAddresses = append(payload.MinersAddresses, address)
	}
}
