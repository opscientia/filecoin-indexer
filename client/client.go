package client

import (
	"context"
	"net/http"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/apistruct"
)

// Client fetches data from a Lotus node
type Client struct {
	endpoint string
	timeout  time.Duration

	api    *apistruct.FullNodeStruct
	closer jsonrpc.ClientCloser

	Epoch       epochClient
	Deal        dealClient
	Miner       minerClient
	Transaction transactionClient
	Account     accountClient
}

// NewClient creates a Lotus client
func NewClient(endpoint string, timeout time.Duration) (*Client, error) {
	client := Client{
		endpoint: endpoint,
		timeout:  timeout,
	}

	err := client.Connect()
	if err != nil {
		return nil, err
	}

	return &client, nil
}

// Connect establishes a websocket connection
func (c *Client) Connect() error {
	var api apistruct.FullNodeStruct

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		"ws://"+c.endpoint+"/rpc/v0",
		"Filecoin",
		[]interface{}{&api.Internal, &api.CommonStruct.Internal},
		http.Header{},
		jsonrpc.WithTimeout(c.timeout))

	if err != nil {
		return err
	}

	c.setConnection(&api, closer)

	return nil
}

func (c *Client) setConnection(api *apistruct.FullNodeStruct, closer jsonrpc.ClientCloser) {
	c.api = api
	c.closer = closer

	c.Epoch = epochClient{api: api}
	c.Deal = dealClient{api: api}
	c.Miner = minerClient{api: api}
	c.Transaction = transactionClient{api: api}
	c.Account = accountClient{api: api}
}

// Close attempts to close the websocket connection
func (c *Client) Close() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = v.(error)
		}
	}()
	c.closer()

	return nil
}

// Reconnect restablishes a websocket connection
func (c *Client) Reconnect() error {
	c.Close()
	return c.Connect()
}
