package server

import (
	"github.com/figment-networks/indexing-engine/metrics"
	"github.com/figment-networks/indexing-engine/metrics/prometheusmetrics"
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

// Server handles HTTP requests
type Server struct {
	engine *gin.Engine
	cfg    *config.Config
	store  *store.Store
	client *client.Client
}

// NewServer creates an HTTP server
func NewServer(cfg *config.Config, store *store.Store, client *client.Client) (*Server, error) {
	server := Server{
		engine: gin.Default(),
		cfg:    cfg,
		store:  store,
		client: client,
	}

	server.setMiddleware()
	server.setRoutes()

	err := server.initMetrics()
	if err != nil {
		return nil, err
	}

	return &server, nil
}

func (s *Server) setMiddleware() {
	s.engine.Use(MetricsMiddleware(s.cfg))
	s.engine.Use(RollbarMiddleware())
}

func (s *Server) setRoutes() {
	s.engine.GET("/health", s.GetHealth)
	s.engine.GET("/status", s.GetStatus)
	s.engine.GET("/miners", s.GetMiners)
	s.engine.GET("/miners/:address", s.GetMiner)
	s.engine.GET("/miners/:address/events", s.GetMinerEvents)
	s.engine.GET("/top_miners", s.GetTopMiners)
	s.engine.GET("/transactions", s.GetTransactions)
	s.engine.GET("/accounts/:address", s.GetAccount)
	s.engine.GET("/accounts/:address/transactions", s.GetAccountTransactions)
	s.engine.GET("/events", s.GetEvents)
}

func (s *Server) initMetrics() error {
	prom := prometheusmetrics.New()

	err := metrics.AddEngine(prom)
	if err != nil {
		return err
	}

	err = metrics.Hotload(prom.Name())
	if err != nil {
		return err
	}

	s.engine.GET(s.cfg.MetricsPath, gin.WrapH(metrics.Handler()))

	return nil
}

// Start runs the server
func (s *Server) Start() error {
	return s.engine.Run(s.cfg.ServerListenAddr())
}
