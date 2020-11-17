package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"golang.org/x/sync/errgroup"

	"github.com/figment-networks/filecoin-indexer/client"
)

// MinerFetcherTask fetches raw miner data
type MinerFetcherTask struct {
	client *client.Client
}

// NewMinerFetcherTask creates the task
func NewMinerFetcherTask(client *client.Client) pipeline.Task {
	return &MinerFetcherTask{client: client}
}

// GetName returns the task name
func (t *MinerFetcherTask) GetName() string {
	return "MinerFetcher"
}

// Run performs the task
func (t *MinerFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)
	tsk := payload.EpochTipset.Key()

	deals, err := t.client.Miner.GetMarketDeals(tsk)
	if err != nil {
		return err
	}
	payload.MarketDeals = deals

	addresses, err := t.client.Miner.GetAddressesByTipset(tsk)
	if err != nil {
		return err
	}
	payload.MinersAddresses = addresses

	payload.MinersInfo = make([]*miner.MinerInfo, len(addresses))
	payload.MinersPower = make([]*api.MinerPower, len(addresses))
	payload.MinersFaults = make([]*bitfield.BitField, len(addresses))

	eg, _ := errgroup.WithContext(ctx)

	for i := range addresses {
		func(index int) {
			eg.Go(func() error {
				return fetchMinerData(index, t.client, payload)
			})
		}(i)
	}

	return eg.Wait()
}

func fetchMinerData(index int, c *client.Client, p *payload) error {
	address := p.MinersAddresses[index]
	tsk := p.EpochTipset.Key()

	info, err := c.Miner.GetInfoByTipset(address, tsk)
	if err != nil {
		return err
	}

	power, err := c.Miner.GetPowerByTipset(address, tsk)
	if err != nil {
		return err
	}

	faults, err := c.Miner.GetFaultsByTipset(address, tsk)
	if err != nil {
		return err
	}

	p.MinersInfo[index] = info
	p.MinersPower[index] = power
	p.MinersFaults[index] = faults

	return nil
}