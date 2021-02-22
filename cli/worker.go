package cli

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/pipeline"
)

func runWorker(cfg *config.Config) error {
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

	err = initMetrics(cfg)
	if err != nil {
		return err
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())
	ticker := time.NewTicker(cfg.PipelineSyncInterval())

	go func() {
		defer wg.Done()
		defer config.LogPanic()

		for {
			select {
			case <-ticker.C:
				pipeline.StartPipeline(cfg, client, store)
			case <-ctx.Done():
				return
			}
		}
	}()

	<-interrupt()

	ticker.Stop()
	cancel()
	wg.Wait()

	return nil
}

func interrupt() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	return c
}
