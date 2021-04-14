package pipeline

import (
	"time"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/figment-networks/filecoin-indexer/datalake"
	"github.com/figment-networks/filecoin-indexer/model"
)

var (
	_ pipeline.PayloadFactory = (*PayloadFactory)(nil)
	_ pipeline.Payload        = (*payload)(nil)
)

// PayloadFactory creates payloads
type PayloadFactory struct {
	dataLake *datalake.DataLake
}

// NewPayloadFactory creates a payload factory
func NewPayloadFactory(dl *datalake.DataLake) *PayloadFactory {
	return &PayloadFactory{dataLake: dl}
}

// GetPayload returns a payload for a given height
func (pf *PayloadFactory) GetPayload(height int64) pipeline.Payload {
	return &payload{
		startedAt:     time.Now(),
		currentHeight: height,
		dataLake:      pf.dataLake,
	}
}

type payload struct {
	startedAt     time.Time
	currentHeight int64
	processed     bool

	dataLake *datalake.DataLake

	// Fetcher stage
	EpochTipset          *types.TipSet
	DealsData            map[string]api.MarketDeal
	DealsCount           map[address.Address]uint32
	DealsSlashedCount    map[address.Address]uint32
	DealsSlashedIDs      []string
	MinersAddresses      []address.Address
	MinersInfo           []*miner.MinerInfo
	MinersPower          []*api.MinerPower
	MinersFaults         []*bitfield.BitField
	TransactionsCIDs     []cid.Cid
	TransactionsMessages []*types.Message

	// Parser stage
	Epoch        *model.Epoch
	Miners       []*model.Miner
	Transactions []*model.Transaction
	Events       []*model.Event

	// Sequencer stage
	StoredMiners         map[string]model.Miner
	StoredSlashedDealIDs []string
}

func (p *payload) SetCurrentHeight(height int64) {
	p.currentHeight = height
}

func (p *payload) GetCurrentHeight() int64 {
	return p.currentHeight
}

func (p *payload) MarkAsProcessed() {
	p.processed = true
}

func (p *payload) IsProcessed() bool {
	return p.processed
}

func (p *payload) Duration() float64 {
	return time.Since(p.startedAt).Seconds()
}

func (p *payload) Store(name string, obj interface{}) error {
	res, err := datalake.NewJSONResource(name, obj)
	if err != nil {
		return err
	}

	return p.dataLake.StoreResourceAtHeight(res, p.currentHeight)
}

func (p *payload) Retrieve(name string, obj interface{}) error {
	res, err := p.dataLake.RetrieveResourceAtHeight(name, p.currentHeight)
	if err != nil {
		return err
	}

	return res.ScanJSON(obj)
}
