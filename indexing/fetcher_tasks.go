package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

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

	minerAddresses, err := t.client.Miner.GetAddressesByHeight(payload.currentHeight)
	if err != nil {
		return err
	}

	payload.MinerAddresses = minerAddresses

	return nil
}
