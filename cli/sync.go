package cli

import (
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/pipeline"
)

func runSync(cfg *config.Config) error {
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

	return pipeline.StartPipeline(cfg, client, store)
}
