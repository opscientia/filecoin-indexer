package client

import (
	"context"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/types"
)

type dealClient struct {
	api *apistruct.FullNodeStruct
}

// GetMarketDeals fetches market deals for a given tipset
func (dc *dealClient) GetMarketDeals(tsk types.TipSetKey) (map[string]api.MarketDeal, error) {
	return dc.api.StateMarketDeals(context.Background(), tsk)
}
