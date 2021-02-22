package types

// EventKind represents an event type
type EventKind string

const (
	// StorageCapacityChangeEvent represents a change of miner's storage capacity
	StorageCapacityChangeEvent EventKind = "storage_capacity_change"

	// SlashedDealEvent represents a slash of a miner's deal
	SlashedDealEvent EventKind = "slashed_deal"

	// SectorFaultEvent represents a sector fault
	SectorFaultEvent EventKind = "sector_fault"

	// SectorRecoveryEvent represents a sector recovery
	SectorRecoveryEvent EventKind = "sector_recovery"
)
