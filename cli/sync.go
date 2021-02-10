package cli

import (
	"net/http"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/metrics/prometheusmetrics"

	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/indexing"
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

	prom := prometheusmetrics.New()
	err = metrics.AddEngine(prom)
	if err != nil {
		return err
	}

	err = metrics.Hotload(prom.Name())
	if err != nil {
		return err
	}

	s := &http.Server{
		Addr:    cfg.ListenAddr(),
		Handler: metrics.Handler(),
	}

	go func() {
		s.ListenAndServe()
	}()

	return indexing.StartPipeline(cfg, client, store)
}
