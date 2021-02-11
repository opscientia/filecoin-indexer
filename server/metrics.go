package server

import "github.com/figment-networks/indexing-engine/metrics"

var (
	serverRequestDuration = metrics.MustNewHistogramWithTags(metrics.HistogramOptions{
		Namespace: "indexer",
		Subsystem: "server",
		Name:      "request_duration",
		Desc:      "The total time spent handling an HTTP request",
		Tags:      []string{"path"},
	})
)
