package client

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/types"
)

type accountClient struct {
	api *apistruct.FullNodeStruct
}

// GetActor fetches account details for a given address
func (ac *accountClient) GetActor(addr address.Address) (*types.Actor, error) {
	return ac.api.StateGetActor(context.Background(), addr, types.EmptyTSK)
}
