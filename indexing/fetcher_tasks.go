package indexing

import (
	"context"
	"fmt"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"

	"github.com/figment-networks/filecoin-indexer/client"
)

const (
	MinerFetcherTaskName = "MinerFetcher"
)

type MinerFetcherTask struct {
	client *client.Client
}

func NewMinerFetcherTask(client *client.Client) pipeline.Task {
	return &MinerFetcherTask{client: client}
}

func (t *MinerFetcherTask) GetName() string {
	return MinerFetcherTaskName
}

func (t *MinerFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	minersAddresses, err := t.client.Miner.GetAddressesByHeight(payload.currentHeight)
	if err != nil {
		return err
	}
	payload.MinersAddresses = minersAddresses

	payload.MinersInfo = make(map[address.Address]*api.MinerInfo)
	payload.MinersPower = make(map[address.Address]*api.MinerPower)

	for i, address := range minersAddresses {
		minerInfo, err := t.client.Miner.GetInfoByHeight(address, payload.currentHeight)
		if err != nil {
			return err
		}
		payload.MinersInfo[address] = minerInfo

		minerPower, err := t.client.Miner.GetPowerByHeight(address, payload.currentHeight)
		if err != nil {
			return err
		}
		payload.MinersPower[address] = minerPower

		fmt.Println("Fetched", i+1, "out of", len(minersAddresses), "miners")
	}

	return nil
}
