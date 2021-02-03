package server

import (
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/store"
)

// GetEvents lists all events
func (s *Server) GetEvents(c *gin.Context) {
	height := c.Query("height")
	kind := c.Query("kind")

	pagination := store.Pagination{}
	if err := c.Bind(&pagination); err != nil {
		badRequest(c, err)
		return
	}

	result, err := s.store.Event.FindAll(height, kind, pagination)
	if err != nil {
		badRequest(c, err)
		return
	}

	jsonOK(c, result)
}

// GetMinerEvents lists storage miner events
func (s *Server) GetMinerEvents(c *gin.Context) {
	address := c.Param("address")

	height := c.Query("height")
	kind := c.Query("kind")

	pagination := store.Pagination{}
	if err := c.Bind(&pagination); err != nil {
		badRequest(c, err)
		return
	}

	result, err := s.store.Event.FindAllByMinerAddress(address, height, kind, pagination)
	if err != nil {
		badRequest(c, err)
		return
	}

	jsonOK(c, result)
}
