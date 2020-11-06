package client

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/types"
)

// EpochClient fetches epoch information
type EpochClient interface {
	GetCurrentHeight() (int64, error)
	GetTipsetByHeight(int64) (*types.TipSet, error)
}

type epochClient struct {
	api *apistruct.FullNodeStruct
}

var (
	_ EpochClient = (*epochClient)(nil)
)

// NewEpochClient creates an epoch client
func NewEpochClient(api *apistruct.FullNodeStruct) EpochClient {
	return &epochClient{api: api}
}

// GetCurrentHeight fetches the height of the current epoch
func (ec *epochClient) GetCurrentHeight() (int64, error) {
	tipset, err := ec.api.ChainHead(context.Background())
	if err != nil {
		return 0, err
	}

	return int64(tipset.Height()), nil
}

// GetTipsetByHeight fetches a tipset for a given height
func (ec *epochClient) GetTipsetByHeight(height int64) (*types.TipSet, error) {
	return ec.api.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(height), types.EmptyTSK)
}
