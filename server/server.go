package server

import (
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/store"
)

// Server handles HTTP requests
type Server struct {
	engine *gin.Engine
	store  *store.Store
}

// New creates an HTTP server
func New(store *store.Store) *Server {
	server := Server{
		engine: gin.Default(),
		store:  store,
	}

	server.setRoutes()

	return &server
}

func (s *Server) setRoutes() {
	s.engine.GET("/miners", s.GetMiners)
	s.engine.GET("/top_miners", s.GetTopMiners)
}

// Start runs the server
func (s *Server) Start(listenAddr string) error {
	return s.engine.Run(listenAddr)
}
