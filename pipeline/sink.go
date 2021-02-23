package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

type sink struct {
	store    *store.Store
	observer metrics.Observer
}

// NewSink creates a pipeline sink
func NewSink(store *store.Store) pipeline.Sink {
	return &sink{
		store:    store,
		observer: pipelineDatabaseSizeAfterHeight.WithLabels(),
	}
}

func (s sink) Consume(ctx context.Context, p pipeline.Payload) error {
	size, err := s.store.DatabaseSize()
	if err != nil {
		return err
	}

	s.observer.Observe(float64(size))

	return nil
}
