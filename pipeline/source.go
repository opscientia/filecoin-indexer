package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

type source struct {
	startHeight   int64
	currentHeight int64
	endHeight     int64

	err error
}

// NewSource creates a pipeline source
func NewSource(cfg *config.Config, client *client.Client, store *store.Store) (pipeline.Source, error) {
	latestHeight, err := client.Epoch.GetLatestHeight()
	if err != nil {
		return nil, err
	}

	lastHeight, err := store.Epoch.LastHeight()
	if err != nil {
		lastHeight = -1
	}

	hr := HeightRange{
		LatestHeight:  latestHeight,
		LastHeight:    lastHeight,
		InitialHeight: cfg.InitialHeight,
		BatchSize:     cfg.BatchSize,
	}

	if err := hr.Validate(true); err != nil {
		return nil, err
	}

	return &source{
		startHeight:   hr.StartHeight(),
		currentHeight: hr.StartHeight(),
		endHeight:     hr.EndHeight(),
	}, nil
}

func (s *source) Current() int64 {
	return s.currentHeight
}

func (s *source) Next(context.Context, pipeline.Payload) bool {
	if s.err == nil && s.currentHeight < s.endHeight {
		s.currentHeight++
		return true
	}

	return false
}

func (s *source) Skip(stageName pipeline.StageName) bool {
	return false
}

func (s *source) Err() error {
	return s.err
}
