package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth checks the health of the indexer
func (s *Server) GetHealth(c *gin.Context) {
	storeErr := s.store.Test()
	_, clientErr := s.client.Epoch.GetCurrentHeight()

	if storeErr != nil || clientErr != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

// GetStatus returns the synchronization status
func (s *Server) GetStatus(c *gin.Context) {
	lastSyncedEpoch, _ := s.store.Epoch.LastHeight()

	currentEpoch, err := s.client.Epoch.GetCurrentHeight()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"current_epoch":     currentEpoch,
		"last_synced_epoch": lastSyncedEpoch,
	})
}
