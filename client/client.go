package client

import (
	"context"
	"net/http"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
)

// Client fetches data from a Filecoin node
type Client struct {
	api    *apistruct.FullNodeStruct
	closer jsonrpc.ClientCloser

	Epoch       epochClient
	Deal        dealClient
	Miner       minerClient
	Transaction transactionClient
	Account     accountClient
}

// NewClient creates a JSON RPC client
func NewClient(endpoint string, timeout time.Duration) (*Client, error) {
	var api apistruct.FullNodeStruct

	ctx := context.Background()
	addr := "ws://" + endpoint + "/rpc/v0"
	outs := []interface{}{&api.Internal, &api.CommonStruct.Internal}

	closer, err := jsonrpc.NewMergeClient(
		ctx,
		addr,
		"Filecoin",
		outs,
		http.Header{},
		jsonrpc.WithTimeout(timeout))

	if err != nil {
		return nil, err
	}

	return &Client{
		api:    &api,
		closer: closer,

		Epoch:       epochClient{api: &api},
		Deal:        dealClient{api: &api},
		Miner:       minerClient{api: &api},
		Transaction: transactionClient{api: &api},
		Account:     accountClient{api: &api},
	}, nil
}

// Close closes the websocket connection
func (c *Client) Close() {
	c.closer()
}
