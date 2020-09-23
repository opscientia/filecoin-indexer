package client

import (
	"github.com/ybbus/jsonrpc"
)

type Client struct {
	client *jsonrpc.RPCClient

	Miner MinerClient
}

func New(endpoint string) *Client {
	client := jsonrpc.NewClient(endpoint)

	return &Client{
		client: &client,

		Miner: NewMinerClient(&client),
	}
}
