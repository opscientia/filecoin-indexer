package types

// EventKind represents an event type
type EventKind string

const (
	// StorageCapacityChangeEvent represents a change of miner's storage capacity
	StorageCapacityChangeEvent EventKind = "storage_capacity_change"

	// SlashedDealEvent represents a slash of a miner's deal
	SlashedDealEvent EventKind = "slashed_deal"

	// SectorFaultEvent represents an increase in miner's faults
	SectorFaultEvent EventKind = "sector_fault"

	// SectorRecoveryEvent represents a decrease in miner's faults
	SectorRecoveryEvent EventKind = "sector_recovery"
)
