package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"go.uber.org/multierr"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/model/types"
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

	if err := t.retrieveDeals(payload); err != nil {
		tsk := payload.EpochTipset.Key()

		deals, err := t.client.Deal.GetMarketDeals(tsk)
		if err != nil {
			return err
		}

		t.parseDeals(deals, payload)

		if err := t.storeDeals(payload); err != nil {
			return err
		}
	}

	return t.parseMinersAddresses(payload)
}

func (t *DealFetcherTask) retrieveDeals(payload *payload) error {
	return multierr.Combine(
		payload.Retrieve("deals_count", &payload.DealsCount),
		payload.Retrieve("deals_slashed_count", &payload.DealsSlashedCount),
		payload.Retrieve("deals_slashed", &payload.DealsSlashed),
	)
}

func (t *DealFetcherTask) storeDeals(payload *payload) error {
	return multierr.Combine(
		payload.Store("deals_count", payload.DealsCount),
		payload.Store("deals_slashed_count", payload.DealsSlashedCount),
		payload.Store("deals_slashed", payload.DealsSlashed),
	)
}

func (t *DealFetcherTask) parseDeals(deals map[string]api.MarketDeal, payload *payload) {
	payload.DealsCount = make(map[string]uint32)
	payload.DealsSlashedCount = make(map[string]uint32)
	payload.DealsSlashed = make(map[string]types.SlashedDeal)

	for dealID, deal := range deals {
		minerAddress := deal.Proposal.Provider.String()

		payload.DealsCount[minerAddress]++

		if deal.State.SlashEpoch != -1 {
			payload.DealsSlashedCount[minerAddress]++

			payload.DealsSlashed[dealID] = types.SlashedDeal{
				MinerAddress: minerAddress,
				SlashEpoch:   int64(deal.State.SlashEpoch),
			}
		}
	}
}

func (t *DealFetcherTask) parseMinersAddresses(payload *payload) error {
	for minerAddress := range payload.DealsCount {
		addr, err := address.NewFromString(minerAddress)
		if err != nil {
			return err
		}

		payload.MinersAddresses = append(payload.MinersAddresses, addr)
	}

	return nil
}
