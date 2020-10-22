package client

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
)

// MinerClient fetches miner information
type MinerClient interface {
	GetAddresses() ([]address.Address, error)
	GetInfo(address.Address) (*miner.MinerInfo, error)
	GetPower(address.Address) (*api.MinerPower, error)
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

// GetAddresses fetches miners' addresses
func (mc *minerClient) GetAddresses() ([]address.Address, error) {
	return mc.api.StateListMiners(context.Background(), types.EmptyTSK)
}

// GetInfo fetches miner's information
func (mc *minerClient) GetInfo(address address.Address) (*miner.MinerInfo, error) {
	info, err := mc.api.StateMinerInfo(context.Background(), address, types.EmptyTSK)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetPower fetches miner's power
func (mc *minerClient) GetPower(address address.Address) (*api.MinerPower, error) {
	return mc.api.StateMinerPower(context.Background(), address, types.EmptyTSK)
}
