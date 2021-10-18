package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"go.uber.org/multierr"
	"golang.org/x/sync/errgroup"

	"github.com/figment-networks/filecoin-indexer/client"
)

// MinerFetcherTask fetches raw miner data
type MinerFetcherTask struct {
	client *client.Client
}

// MinerFetcherTaskName represents the name of the task
const MinerFetcherTaskName = "MinerFetcher"

// NewMinerFetcherTask creates the task
func NewMinerFetcherTask(client *client.Client) pipeline.Task {
	return &MinerFetcherTask{client: client}
}

// GetName returns the task name
func (t *MinerFetcherTask) GetName() string {
	return MinerFetcherTaskName
}

// Run performs the task
func (t *MinerFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	if err := t.retrieveMiners(payload); err != nil {
		if err := t.fetchMiners(ctx, payload); err != nil {
			return err
		}

		if err := t.storeMiners(payload); err != nil {
			return err
		}
	}

	return nil
}

func (t *MinerFetcherTask) retrieveMiners(payload *payload) error {
	return multierr.Combine(
		payload.Retrieve("miners_info", &payload.MinersInfo),
		payload.Retrieve("miners_power", &payload.MinersPower),
		payload.Retrieve("miners_faults", &payload.MinersFaults),
	)
}

func (t *MinerFetcherTask) storeMiners(payload *payload) error {
	return multierr.Combine(
		payload.Store("miners_info", payload.MinersInfo),
		payload.Store("miners_power", payload.MinersPower),
		payload.Store("miners_faults", payload.MinersFaults),
	)
}

func (t *MinerFetcherTask) fetchMiners(ctx context.Context, payload *payload) error {
	minersCount := len(payload.MinersAddresses)

	payload.MinersInfo = make([]*miner.MinerInfo, minersCount)
	payload.MinersPower = make([]*api.MinerPower, minersCount)
	payload.MinersFaults = make([]*bitfield.BitField, minersCount)

	eg, _ := errgroup.WithContext(ctx)

	for i := 0; i < minersCount; i++ {
		func(index int) {
			eg.Go(func() error {
				return t.fetchMinerData(index, payload)
			})
		}(i)
	}

	return eg.Wait()
}

func (t *MinerFetcherTask) fetchMinerData(index int, payload *payload) error {
	address := payload.MinersAddresses[index]
	tsk := payload.EpochTipset.Key()

	info, err := t.client.Miner.GetInfoByTipset(address, tsk)
	if err != nil {
		return err
	}

	power, err := t.client.Miner.GetPowerByTipset(address, tsk)
	if err != nil {
		return err
	}

	faults, err := t.client.Miner.GetFaultsByTipset(address, tsk)
	if err != nil {
		return err
	}

	payload.MinersInfo[index] = info
	payload.MinersPower[index] = power
	payload.MinersFaults[index] = faults

	return nil
}
