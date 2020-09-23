package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/store"
)

func Run(listenAddr string, store *store.Store) {
	router := gin.Default()

	router.GET("/miners", func(ctx *gin.Context) {
		var miners []model.Miner

		store.Db.Find(&miners)

		ctx.JSON(http.StatusOK, miners)
	})

	router.Run(listenAddr)
}
