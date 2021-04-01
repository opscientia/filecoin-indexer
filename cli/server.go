package cli

import (
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/server"
)

func runServer(cfg *config.Config) error {
	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	store, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	server, err := server.NewServer(cfg, store, client)
	if err != nil {
		return err
	}

	return server.Start()
}
