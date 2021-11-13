package server

import (
	"strings"
	"time"

	"github.com/figment-networks/indexing-engine/pipeline/metrics"
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/config"
)

// MetricsMiddleware logs the execution time of every request
func MetricsMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		elapsed := time.Since(t)

		if c.Request.URL.Path == _metricsPath {
			return
		}

		if path := strings.Split(strings.TrimPrefix(c.Request.URL.Path, "/"), "/"); len(path) > 0 {
			metrics.ServerRequestDuration.WithLabels(path[0]).Observe(elapsed.Seconds())
		}
	}
}

// RollbarMiddleware reports panics to Rollbar
func RollbarMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer config.LogPanic()
		c.Next()
	}
}
