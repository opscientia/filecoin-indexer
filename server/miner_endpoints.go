package server

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/store"
)

// GetMiners lists all storage miners
func (s *Server) GetMiners(c *gin.Context) {
	height := getHeight(c, s.store)
	miners, _ := s.store.Miner.FindAllByHeight(height)

	c.JSON(http.StatusOK, miners)
}

// GetMiner returns storage miner details
func (s *Server) GetMiner(c *gin.Context) {
	address := c.Param("address")
	height := getHeight(c, s.store)

	miner, err := s.store.Miner.FindByHeight(address, height)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, miner)
}

// GetTopMiners lists top 100 storage miners
func (s *Server) GetTopMiners(c *gin.Context) {
	height := getHeight(c, s.store)
	miners, _ := s.store.Miner.FindTop100ByHeight(height)

	c.JSON(http.StatusOK, miners)
}

func getHeight(c *gin.Context, store *store.Store) int64 {
	height, err := strconv.ParseInt(c.Query("height"), 10, 64)
	if err != nil {
		height, _ = store.Epoch.LastHeight()
	}

	return height
}
