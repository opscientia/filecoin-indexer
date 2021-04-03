package cli

import (
	"net/http"

	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/metrics/prometheusmetrics"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/server"
	"github.com/figment-networks/filecoin-indexer/store"
)

func initConfig(path string) (*config.Config, error) {
	cfg := config.NewConfig()

	if path != "" {
		if err := config.FromFile(path, cfg); err != nil {
			return nil, err
		}
	}

	if err := config.FromEnv(cfg); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func initClient(cfg *config.Config) (*client.Client, error) {
	return client.NewClient(cfg.RPCEndpoint, cfg.RPCTimeoutDuration)
}

func initStore(cfg *config.Config) (*store.Store, error) {
	logMode := logger.Warn
	if cfg.Debug {
		logMode = logger.Info
	}

	store, err := store.NewStore(cfg.DatabaseDSN, logMode)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func initServer(cfg *config.Config, store *store.Store, client *client.Client) (*server.Server, error) {
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	return server.NewServer(cfg, store, client)
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
