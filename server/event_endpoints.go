package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
)

// GetEvents lists all events
func (s *Server) GetEvents(c *gin.Context) {
	var events *[]model.Event

	height := c.Query("height")

	if height != "" {
		events, _ = s.store.Event.FindAllByHeight(height)
	} else {
		events, _ = s.store.Event.FindAll()
	}

	c.JSON(http.StatusOK, events)
}

// GetMinerEvents lists storage miner events
func (s *Server) GetMinerEvents(c *gin.Context) {
	var events *[]model.Event

	address := c.Param("address")
	height := c.Query("height")

	if height != "" {
		events, _ = s.store.Event.FindAllByMinerAddressAndHeight(address, height)
	} else {
		events, _ = s.store.Event.FindAllByMinerAddress(address)
	}

	c.JSON(http.StatusOK, events)
}
