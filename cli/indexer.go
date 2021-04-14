package cli

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rollbar/rollbar-go"

	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/pipeline"
)

func runIndexer(cfg *config.Config) error {
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

	dl, err := initDataLake(cfg)
	if err != nil {
		return err
	}

	err = initMetrics(cfg)
	if err != nil {
		return err
	}

	interval := cfg.SyncIntervalDuration
	ticker := time.NewTicker(interval)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			ticker.Stop()

			err := pipeline.StartIndexerPipeline(cfg, client, store, dl)
			if err != nil {
				rollbar.Error(err)
			}

			ticker.Reset(interval)
		case <-interrupt:
			return nil
		}
	}
}
