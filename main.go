package main

import (
	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/indexing"
	"github.com/figment-networks/filecoin-indexer/server"
	"github.com/figment-networks/filecoin-indexer/store"
)

func main() {
	cfg := config.New()

	err := config.FromEnv(cfg)
	if err != nil {
		panic(err)
	}

	client := client.New(cfg.NodeURL)

	store, err := store.New(cfg.DatabaseDSN)
	if err != nil {
		panic(err)
	}
	defer store.Close()

	err = indexing.StartPipeline(client, store)
	if err != nil {
		panic(err)
	}

	server.Run(cfg.ListenAddr(), store)
}
