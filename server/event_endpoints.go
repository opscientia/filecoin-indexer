package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvents lists all events
func (s *Server) GetEvents(c *gin.Context) {
	events, _ := s.store.Event.FindAll(c.Query("height"), c.Query("kind"))

	c.JSON(http.StatusOK, events)
}

// GetMinerEvents lists storage miner events
func (s *Server) GetMinerEvents(c *gin.Context) {
	address := c.Param("address")

	height := c.Query("height")
	kind := c.Query("kind")

	events, _ := s.store.Event.FindAllByMinerAddress(address, height, kind)

	c.JSON(http.StatusOK, events)
}
