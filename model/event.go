package model

import (
	"time"

	"github.com/figment-networks/filecoin-indexer/model/types"
)

// Event represents a network event
type Event struct {
	ID           uint64      `json:"-"`
	Height       *int64      `json:"height"`
	MinerAddress string      `json:"miner_address"`
	Kind         string      `json:"kind"`
	Data         types.JSONB `json:"data"`
	CreatedAt    time.Time   `json:"-"`
	UpdatedAt    time.Time   `json:"-"`
}
