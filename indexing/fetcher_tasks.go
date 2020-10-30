package indexing

import (
	"context"
	"sync"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"

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

	addresses, err := t.client.Miner.GetAddresses()
	if err != nil {
		return err
	}
	payload.MinersAddresses = addresses

	payload.MinersInfo = make([]*miner.MinerInfo, len(addresses))
	payload.MinersPower = make([]*api.MinerPower, len(addresses))
	payload.MinersFaults = make([]*bitfield.BitField, len(addresses))

	wg := &sync.WaitGroup{}

	for i := range addresses {
		wg.Add(1)
		go fetchMiner(i, t.client.Miner, payload, wg)
	}

	wg.Wait()

	return nil
}

func fetchMiner(index int, mc client.MinerClient, p *payload, wg *sync.WaitGroup) {
	defer wg.Done()

	address := p.MinersAddresses[index]

	info, err := mc.GetInfo(address)
	if err != nil {
		panic(err)
	}

	power, err := mc.GetPower(address)
	if err != nil {
		panic(err)
	}

	faults, err := mc.GetFaults(address)
	if err != nil {
		panic(err)
	}

	p.MinersInfo[index] = info
	p.MinersPower[index] = power
	p.MinersFaults[index] = faults
}
