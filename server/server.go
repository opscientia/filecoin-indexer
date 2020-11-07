package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/store"
)

// Run starts an HTTP server
func Run(listenAddr string, store *store.Store) error {
	router := gin.Default()

	router.GET("/miners", func(c *gin.Context) {
		height := getHeight(c, store)
		miners, _ := store.Miner.FindAllByHeight(height)

		c.JSON(http.StatusOK, miners)
	})

	router.GET("/top_miners", func(c *gin.Context) {
		height := getHeight(c, store)
		miners, _ := store.Miner.FindTop100ByHeight(height)

		c.JSON(http.StatusOK, miners)
	})

	return router.Run(listenAddr)
}

func getHeight(c *gin.Context, store *store.Store) int64 {
	height, err := strconv.ParseInt(c.Query("height"), 10, 64)
	if err != nil {
		height, _ = store.Epoch.LastHeight()
	}

	return height
}
