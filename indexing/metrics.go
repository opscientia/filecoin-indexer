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
)
