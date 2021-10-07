package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/pipeline"
	m "github.com/figment-networks/indexing-engine/pipeline/metrics"

	"github.com/figment-networks/filecoin-indexer/store"
)

type sink struct {
	store        *store.Store
	databaseSize metrics.Gauge
}

// NewSink creates a pipeline sink
func NewSink(store *store.Store) pipeline.Sink {
	return &sink{
		store:        store,
		databaseSize: m.DatabaseSizeBytes.WithLabels(),
	}
}

func (s sink) Consume(ctx context.Context, p pipeline.Payload) error {
	size, err := s.store.DatabaseSize()
	if err != nil {
		return err
	}

	s.databaseSize.Set(float64(size))

	return nil
}
