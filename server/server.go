package server

import (
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/store"
)

// Server handles HTTP requests
type Server struct {
	engine *gin.Engine
	store  *store.Store
	client *client.Client
}

// New creates an HTTP server
func New(store *store.Store, client *client.Client) *Server {
	server := Server{
		engine: gin.Default(),
		store:  store,
		client: client,
	}

	server.setRoutes()

	return &server
}

func (s *Server) setRoutes() {
	s.engine.GET("/miners", s.GetMiners)
	s.engine.GET("/miners/:address", s.GetMiner)
	s.engine.GET("/miners/:address/events", s.GetMinerEvents)
	s.engine.GET("/top_miners", s.GetTopMiners)
	s.engine.GET("/transactions", s.GetTransactions)
	s.engine.GET("/accounts/:address", s.GetAccount)
	s.engine.GET("/accounts/:address/transactions", s.GetAccountTransactions)
	s.engine.GET("/events", s.GetEvents)
}

// Start runs the server
func (s *Server) Start(listenAddr string) error {
	return s.engine.Run(listenAddr)
}
