package client

import (
	"context"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
)

// Client fetches data from a Filecoin node
type Client struct {
	api    *apistruct.FullNodeStruct
	closer jsonrpc.ClientCloser

	Epoch epochClient
	Miner minerClient
}

// New creates a Filecoin client
func New(endpoint string) (*Client, error) {
	var api apistruct.FullNodeStruct

	ctx := context.Background()
	addr := "ws://" + endpoint + "/rpc/v0"
	outs := []interface{}{&api.Internal, &api.CommonStruct.Internal}

	closer, err := jsonrpc.NewMergeClient(ctx, addr, "Filecoin", outs, http.Header{})
	if err != nil {
		return nil, err
	}

	return &Client{
		api:    &api,
		closer: closer,

		Epoch: epochClient{api: &api},
		Miner: minerClient{api: &api},
	}, nil
}

// Close closes the websocket connection
func (c *Client) Close() {
	c.closer()
}
