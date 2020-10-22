package indexing

import (
	"context"
	"fmt"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
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

	minersAddresses, err := t.client.Miner.GetAddresses()
	if err != nil {
		return err
	}
	payload.MinersAddresses = minersAddresses

	payload.MinersInfo = make(map[address.Address]*miner.MinerInfo)
	payload.MinersPower = make(map[address.Address]*api.MinerPower)

	for i, address := range minersAddresses {
		minerInfo, err := t.client.Miner.GetInfo(address)
		if err != nil {
			return err
		}
		payload.MinersInfo[address] = minerInfo

		minerPower, err := t.client.Miner.GetPower(address)
		if err != nil {
			return err
		}
		payload.MinersPower[address] = minerPower

		fmt.Println("Fetched", i+1, "out of", len(minersAddresses), "miners")
	}

	return nil
}
