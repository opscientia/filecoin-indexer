package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/store"
)

type sink struct {
	store *store.Store

	heightCounter          metrics.Counter
	heightDurationObserver metrics.Observer
	databaseSizeObserver   metrics.Observer
}

// NewSink creates a pipeline sink
func NewSink(store *store.Store) pipeline.Sink {
	return &sink{
		store: store,

		heightCounter:          pipelineHeightsTotal.WithLabels(),
		heightDurationObserver: pipelineHeightDuration.WithLabels(),
		databaseSizeObserver:   pipelineDatabaseSizeAfterHeight.WithLabels(),
	}
}

func (s sink) Consume(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	s.heightCounter.Inc()
	s.heightDurationObserver.Observe(payload.Duration())

	size, err := s.store.DatabaseSize()
	if err != nil {
		return err
	}
	s.databaseSizeObserver.Observe(float64(size))

	return nil
}
