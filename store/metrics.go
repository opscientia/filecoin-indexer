package store

import "github.com/figment-networks/indexing-engine/metrics"

var (
	databaseQueryDuration = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "database",
		Name:      "query_duration",
		Desc:      "The total time spent executing a database query",
		Tags:      []string{"query"},
	})
)
