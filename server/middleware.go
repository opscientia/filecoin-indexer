package server

import (
	"github.com/figment-networks/indexing-engine/metrics"
	m "github.com/figment-networks/indexing-engine/pipeline/metrics"
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/config"
)

// MetricsMiddleware logs the execution time of every request
func MetricsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		if path != _metricsPath {
			observer := m.ServerRequestDuration.WithLabels(path)

			t := metrics.NewTimer(observer)
			defer t.ObserveDuration()
		}

		c.Next()
	}
}

// RollbarMiddleware reports panics to Rollbar
func RollbarMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer config.LogPanic()
		c.Next()
	}
}
