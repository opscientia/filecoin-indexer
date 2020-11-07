package indexing

import (
	"context"
	"errors"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/pkg/math"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

var errNothingToProcess = errors.New("nothing to process")

type source struct {
	batchSize int64

	startHeight   int64
	currentHeight int64
	endHeight     int64

	err error
}

// NewSource creates a pipeline source
func NewSource(cfg *config.Config, client *client.Client, store *store.Store) (pipeline.Source, error) {
	source := source{
		batchSize: cfg.BatchSize,
	}

	source.setStartHeight(store)

	if err := source.setEndHeight(client); err != nil {
		return nil, err
	}

	if err := source.validate(); err != nil {
		return nil, err
	}

	return &source, nil
}

func (s *source) Current() int64 {
	return s.currentHeight
}

func (s *source) Len() int64 {
	return s.endHeight - s.startHeight + 1
}

func (s *source) Next(context.Context, pipeline.Payload) bool {
	if s.err == nil && s.currentHeight < s.endHeight {
		s.currentHeight++
		return true
	}
	return false
}

func (s *source) Err() error {
	return s.err
}

func (s *source) setStartHeight(store *store.Store) {
	lastHeight, err := store.Epoch.LastHeight()
	if err != nil {
		return // Keep zero values
	}

	s.startHeight = lastHeight + 1
	s.currentHeight = s.startHeight
}

func (s *source) setEndHeight(client *client.Client) error {
	currentHeight, err := client.Epoch.GetCurrentHeight()
	if err != nil {
		return err
	}

	if s.batchSize > 0 {
		batchEnd := s.startHeight + s.batchSize - 1
		s.endHeight = math.MinInt64(batchEnd, currentHeight)
	} else {
		s.endHeight = currentHeight
	}

	return nil
}

func (s *source) validate() error {
	if s.Len() <= 0 {
		return errNothingToProcess
	}
	return nil
}
