package types

// EventKind represents an event type
type EventKind string

const (
	// StorageCapacityChangeEvent represents a change of miner's storage capacity
	StorageCapacityChangeEvent EventKind = "storage_capacity_change"

	// NewDealEvent represents the act of signing a new deal by a miner
	NewDealEvent EventKind = "new_deal"
)
