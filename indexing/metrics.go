package indexing

import "github.com/figment-networks/indexing-engine/metrics"

var (
	pipelineTaskDuration = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "task_duration",
		Desc:      "The total time spent processing an indexing task",
		Tags:      []string{"task"},
	})

	pipelineHeightsTotal = metrics.MustNewCounterWithTags(metrics.Options{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "heights_total",
		Desc:      "The total number of successfully indexed heights",
	})

	pipelineErrorsTotal = metrics.MustNewCounterWithTags(metrics.Options{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "errors_total",
		Desc:      "The total number of indexing errors",
	})

	pipelineHeightDuration = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "height_duration",
		Desc:      "The total time spent indexing a height",
	})

	pipelineDatabaseSizeAfterHeight = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "pipeline",
		Name:      "database_size",
		Desc:      "The size of the database after indexing a height",
	})
)
