package model

import "time"

// Job represents a worker job
type Job struct {
	ID         uint64     `json:"-"`
	Height     *int64     `json:"height"`
	RunCount   uint64     `json:"run_count"`
	LastError  string     `json:"last_error"`
	StartedAt  *time.Time `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"-"`
}
