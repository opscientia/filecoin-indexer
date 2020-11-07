package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

// Run starts an HTTP server
func Run(listenAddr string, store *store.Store) error {
	router := gin.Default()

	router.GET("/miners", func(ctx *gin.Context) {
		height := ctx.DefaultQuery("height", lastHeight(store))

		var miners []model.Miner

		store.Db.Where("height = ?", height).Find(&miners)

		ctx.JSON(http.StatusOK, miners)
	})

	router.GET("/top_miners", func(ctx *gin.Context) {
		height := ctx.DefaultQuery("height", lastHeight(store))

		var miners []model.Miner

		store.Db.
			Where("height = ?", height).
			Order("score DESC").
			Limit(100).
			Find(&miners)

		ctx.JSON(http.StatusOK, miners)
	})

	return router.Run(listenAddr)
}

func lastHeight(store *store.Store) string {
	result, _ := store.LastHeight()
	return fmt.Sprintf("%d", result)
}
