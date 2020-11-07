package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/store"
)

type source struct {
	startHeight   int64
	currentHeight int64
	endHeight     int64

	err error
}

// NewSource creates a pipeline source
func NewSource(client *client.Client, store *store.Store) (pipeline.Source, error) {
	startHeight := startHeight(store)

	endHeight, err := client.Epoch.GetCurrentHeight()
	if err != nil {
		return nil, err
	}

	return &source{
		startHeight:   startHeight,
		currentHeight: startHeight,
		endHeight:     endHeight,
	}, nil
}

func startHeight(store *store.Store) int64 {
	lastHeight, err := store.LastHeight()
	if err != nil {
		return 0
	}

	return lastHeight + 1
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
