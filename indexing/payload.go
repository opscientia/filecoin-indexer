package indexing

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/indexing-engine/pipeline"
)

var (
	_ pipeline.PayloadFactory = (*PayloadFactory)(nil)
	_ pipeline.Payload        = (*payload)(nil)
)

// NewPayloadFactory creates a payload factory
func NewPayloadFactory() *PayloadFactory {
	return &PayloadFactory{}
}

// PayloadFactory creates payloads
type PayloadFactory struct{}

// GetPayload returns a payload for a given height
func (pf *PayloadFactory) GetPayload(height int64) pipeline.Payload {
	return &payload{currentHeight: height}
}

type payload struct {
	currentHeight int64

	MinersAddresses []address.Address
	MinersInfo      []*miner.MinerInfo
	MinersPower     []*api.MinerPower
	MinersFaults    []*bitfield.BitField

	Miners []*model.Miner
}

func (p *payload) SetCurrentHeight(height int64) {
	p.currentHeight = height
}

func (p *payload) GetCurrentHeight() int64 {
	return p.currentHeight
}

func (p *payload) MarkAsProcessed() {}
