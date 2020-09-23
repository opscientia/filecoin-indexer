package model

import "time"

type Miner struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
