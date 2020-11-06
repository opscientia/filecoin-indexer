package client

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)

// MinerClient fetches miner information
type MinerClient interface {
	GetAddressesByTipset(types.TipSetKey) ([]address.Address, error)
	GetInfoByTipset(address.Address, types.TipSetKey) (*miner.MinerInfo, error)
	GetPowerByTipset(address.Address, types.TipSetKey) (*api.MinerPower, error)
	GetFaultsByTipset(address.Address, types.TipSetKey) (*bitfield.BitField, error)
}

type minerClient struct {
	api *apistruct.FullNodeStruct
}

var (
	_ MinerClient = (*minerClient)(nil)
)

// NewMinerClient creates a miner client
func NewMinerClient(api *apistruct.FullNodeStruct) MinerClient {
	return &minerClient{api: api}
}

// GetAddresses fetches miners' addresses for a given tipset
func (mc *minerClient) GetAddressesByTipset(tsk types.TipSetKey) ([]address.Address, error) {
	return mc.api.StateListMiners(context.Background(), tsk)
}

// GetInfo fetches miner's information for a given tipset
func (mc *minerClient) GetInfoByTipset(address address.Address, tsk types.TipSetKey) (*miner.MinerInfo, error) {
	info, err := mc.api.StateMinerInfo(context.Background(), address, tsk)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetPower fetches miner's power for a given tipset
func (mc *minerClient) GetPowerByTipset(address address.Address, tsk types.TipSetKey) (*api.MinerPower, error) {
	return mc.api.StateMinerPower(context.Background(), address, tsk)
}

// GetFaults fetches miner's faults for a given tipset
func (mc *minerClient) GetFaultsByTipset(address address.Address, tsk types.TipSetKey) (*bitfield.BitField, error) {
	faults, err := mc.api.StateMinerFaults(context.Background(), address, tsk)
	if err != nil {
		return nil, err
	}

	return &faults, nil
}
