package server

import (
	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware logs the execution time of every request
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := metrics.NewTimer(serverRequestDuration.WithLabels(c.Request.URL.Path))
		defer t.ObserveDuration()
		c.Next()
	}
}
