package fetcher

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/pipeline"
	"github.com/figment-networks/filecoin-indexer/worker"
)

// Worker represents a fetcher worker
type Worker struct {
	cfg    *config.Config
	client *client.Client
}

// NewWorker creates a fetcher worker
func NewWorker(cfg *config.Config, client *client.Client) *Worker {
	return &Worker{
		cfg:    cfg,
		client: client,
	}
}

// Run starts the fetcher worker
func (w *Worker) Run() error {
	server := http.Server{
		Addr:    w.cfg.ServerListenAddr(),
		Handler: websocket.Handler(w.handleConnection),
	}

	return server.ListenAndServe()
}

func (w *Worker) handleConnection(conn *websocket.Conn) {
	server := worker.NewWebsocketServer(conn)
	loop := worker.NewLoop(server)

	loop.Run(w.handleRequest)
}

func (w *Worker) handleRequest(req worker.Request) error {
	return pipeline.RunFetcherPipeline(req.Height, w.client)
}
