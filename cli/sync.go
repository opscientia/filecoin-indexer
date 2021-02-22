package cli

import (
	"net/http"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/metrics/prometheusmetrics"

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

func initMetrics(cfg *config.Config) error {
	prom := prometheusmetrics.New()

	err := metrics.AddEngine(prom)
	if err != nil {
		return err
	}

	err = metrics.Hotload(prom.Name())
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    cfg.MetricsListenAddr(),
		Handler: metrics.Handler(),
	}

	go func() {
		defer config.LogPanic()

		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}
