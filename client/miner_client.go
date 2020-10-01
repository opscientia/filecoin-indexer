package client

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/ybbus/jsonrpc"
)

type MinerClient interface {
	GetAddressesByHeight(int64) ([]address.Address, error)
	GetInfoByHeight(address.Address, int64) (*api.MinerInfo, error)
	GetPowerByHeight(address.Address, int64) (*api.MinerPower, error)
}

type minerClient struct {
	client *jsonrpc.RPCClient
}

var (
	_ MinerClient = (*minerClient)(nil)
)

func NewMinerClient(client *jsonrpc.RPCClient) MinerClient {
	return &minerClient{client: client}
}

func (mc *minerClient) GetAddressesByHeight(height int64) ([]address.Address, error) {
	var addresses []address.Address

	err := (*mc.client).CallFor(&addresses, "Filecoin.StateListMiners", nil)

	return addresses, err
}

func (mc *minerClient) GetInfoByHeight(address address.Address, height int64) (*api.MinerInfo, error) {
	var minerInfo api.MinerInfo

	err := (*mc.client).CallFor(&minerInfo, "Filecoin.StateMinerInfo", address, nil)

	return &minerInfo, err
}

func (mc *minerClient) GetPowerByHeight(address address.Address, height int64) (*api.MinerPower, error) {
	var minerPower api.MinerPower

	err := (*mc.client).CallFor(&minerPower, "Filecoin.StateMinerPower", address, nil)

	return &minerPower, err
}
