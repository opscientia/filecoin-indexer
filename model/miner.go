package model

import "time"

type Miner struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Address         string    `json:"address"`
	RawBytePower    *uint64   `json:"raw_byte_power"`
	QualityAdjPower *uint64   `json:"quality_adj_power"`
	SectorSize      *uint64   `json:"sector_size"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}
