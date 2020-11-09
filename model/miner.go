package model

import "time"

// Miner represents a storage miner
type Miner struct {
	ID                uint64    `json:"-"`
	Height            *int64    `json:"-"`
	Address           string    `json:"address"`
	SectorSize        *uint64   `json:"sector_size"`
	RawBytePower      *uint64   `json:"raw_byte_power"`
	QualityAdjPower   *uint64   `json:"quality_adj_power"`
	RelativePower     *float32  `json:"relative_power"`
	DealsCount        *uint32   `json:"deals_count"`
	SlashedDealsCount *uint32   `json:"slashed_deals_count"`
	FaultsCount       *uint32   `json:"faults_count"`
	Score             *uint32   `json:"score"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}
