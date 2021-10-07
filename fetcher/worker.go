package fetcher

import (
	"log"
	"net/http"

	"github.com/figment-networks/indexing-engine/datalake"
	"github.com/figment-networks/indexing-engine/pipeline/worker"
	"github.com/rollbar/rollbar-go"
	"golang.org/x/net/websocket"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/pipeline"
)

// Worker represents a fetcher worker
type Worker struct {
	cfg    *config.Config
	client *client.Client
	dl     *datalake.DataLake
}

// NewWorker creates a fetcher worker
func NewWorker(cfg *config.Config, client *client.Client, dl *datalake.DataLake) *Worker {
	return &Worker{
		cfg:    cfg,
		client: client,
		dl:     dl,
	}
}

// Run starts the fetcher worker
func (w *Worker) Run() error {
	server := http.Server{
		Addr:    w.cfg.WorkerListenAddr(),
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
	log.Printf("job started [height=%d]", req.Height)

	err := pipeline.RunFetcherPipeline(req.Height, w.client, w.dl)
	if err != nil {
		log.Printf("job error [height=%d]: %v", req.Height, err)
		rollbar.Error(err)

		log.Println("reconnecting to the node...")
		w.client.Reconnect()

		return err
	}

	log.Printf("job finished [height=%d]", req.Height)

	return nil
}
