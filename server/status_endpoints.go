package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHealth checks the health of the indexer
func (s *Server) GetHealth(c *gin.Context) {
	if err := s.store.Test(); err != nil {
		serverError(c, err)
		return
	}

	if _, err := s.client.Epoch.GetLatestHeight(); err != nil {
		serverError(c, err)
		return
	}

	c.Status(http.StatusOK)
}

// GetStatus returns the synchronization status
func (s *Server) GetStatus(c *gin.Context) {
	currentEpoch, err := s.client.Epoch.GetLatestHeight()
	if err != nil {
		serverError(c, err)
		return
	}

	lastSyncedEpoch, err := s.store.Epoch.LastHeight()
	if err != nil {
		serverError(c, err)
		return
	}

	jsonOK(c, gin.H{
		"current_epoch":     currentEpoch,
		"last_synced_epoch": lastSyncedEpoch,
	})
}
