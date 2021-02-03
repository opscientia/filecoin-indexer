package server

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/store"
)

// GetMiners lists all storage miners
func (s *Server) GetMiners(c *gin.Context) {
	height := getHeight(c, s.store)

	pagination := store.Pagination{}
	if err := c.Bind(&pagination); err != nil {
		badRequest(c, err)
		return
	}

	result, err := s.store.Miner.FindAllByHeight(height, pagination)
	if err != nil {
		serverError(c, err)
		return
	}

	jsonOK(c, result)
}

// GetMiner returns storage miner details
func (s *Server) GetMiner(c *gin.Context) {
	address := c.Param("address")
	height := getHeight(c, s.store)

	miner, err := s.store.Miner.FindByHeight(address, height)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			notFound(c, nil)
		} else {
			serverError(c, err)
		}
		return
	}

	jsonOK(c, miner)
}

// GetTopMiners lists top 100 storage miners
func (s *Server) GetTopMiners(c *gin.Context) {
	height := getHeight(c, s.store)

	miners, err := s.store.Miner.FindTop100ByHeight(height)
	if err != nil {
		serverError(c, err)
		return
	}

	jsonOK(c, miners)
}

func getHeight(c *gin.Context, store *store.Store) int64 {
	height, err := strconv.ParseInt(c.Query("height"), 10, 64)
	if err != nil {
		height, _ = store.Epoch.LastHeight()
	}

	return height
}
