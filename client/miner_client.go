package client

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"
	"github.com/ybbus/jsonrpc"
)

type MinerClient interface {
	GetAddressesByHeight(int64) (*[]address.Address, error)
	GetByHeight(address.Address, int64) (*api.MinerInfo, error)
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

func (mc *minerClient) GetAddressesByHeight(height int64) (*[]address.Address, error) {
	var addresses []address.Address

	err := (*mc.client).CallFor(&addresses, "Filecoin.StateListMiners", nil)

	return &addresses, err
}

func (mc *minerClient) GetByHeight(address address.Address, height int64) (*api.MinerInfo, error) {
	var miner api.MinerInfo

	err := (*mc.client).CallFor(&miner, "Filecoin.StateMinerInfo", address, nil)

	return &miner, err
}
