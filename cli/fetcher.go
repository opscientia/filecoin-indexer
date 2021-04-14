package cli

import (
	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/fetcher"
)

func runFetcher(cfg *config.Config, mode string) error {
	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	switch mode {
	case "worker":
		return runFetcherWorker(cfg, client)
	case "manager":
		return runFetcherManager(cfg, client)
	default:
		return nil
	}
}

func runFetcherWorker(cfg *config.Config, client *client.Client) error {
	dl, err := initDataLake(cfg)
	if err != nil {
		return err
	}

	err = initMetrics(cfg)
	if err != nil {
		return err
	}

	return fetcher.NewWorker(cfg, client, dl).Run()
}

func runFetcherManager(cfg *config.Config, client *client.Client) error {
	pool, close, err := initWorkerPool(cfg)
	if err != nil {
		return err
	}
	defer close()

	store, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	manager, err := fetcher.NewManager(cfg, pool, store, client)
	if err != nil {
		return err
	}

	return manager.Run()
}
