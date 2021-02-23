package pipeline

import "github.com/figment-networks/indexing-engine/metrics"

var (
	pipelineDatabaseSizeAfterHeight = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "database_size",
		Desc:      "The size of the database after indexing a height",
	})
)
