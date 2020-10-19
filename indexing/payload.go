package indexing

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/api"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/indexing-engine/pipeline"
)

var (
	_ pipeline.PayloadFactory = (*payloadFactory)(nil)
	_ pipeline.Payload        = (*payload)(nil)
)

func NewPayloadFactory() *payloadFactory {
	return &payloadFactory{}
}

type payloadFactory struct{}

func (pf *payloadFactory) GetPayload(height int64) pipeline.Payload {
	return &payload{currentHeight: height}
}

type payload struct {
	currentHeight int64

	MinersAddresses []address.Address
	MinersInfo      []*api.MinerInfo
	MinersPower     []*api.MinerPower

	Miners []*model.Miner
}

func (p *payload) SetCurrentHeight(height int64) {
	p.currentHeight = height
}

func (p *payload) GetCurrentHeight() int64 {
	return p.currentHeight
}

func (p *payload) MarkAsProcessed() {}
