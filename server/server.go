package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

// Run starts an HTTP server
func Run(listenAddr string, store *store.Store) error {
	router := gin.Default()

	router.GET("/miners", func(ctx *gin.Context) {
		var miners []model.Miner

		store.Db.Find(&miners)

		ctx.JSON(http.StatusOK, miners)
	})

	router.GET("/top_miners", func(ctx *gin.Context) {
		var miners []model.Miner

		store.Db.Order("score DESC").Limit(100).Find(&miners)

		ctx.JSON(http.StatusOK, miners)
	})

	return router.Run(listenAddr)
}
