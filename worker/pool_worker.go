package worker

import (
	"errors"
	"sync"
)

// PoolWorker represents a pool worker
type PoolWorker struct {
	client  Client
	channel chan int64
}

// NewPoolWorker creates a pool worker
func NewPoolWorker(client Client) *PoolWorker {
	return &PoolWorker{
		client:  client,
		channel: make(chan int64),
	}
}

// Run starts the pool worker
func (pw *PoolWorker) Run(handler ClientHandler, wg *sync.WaitGroup) {
	for height := range pw.channel {
		pw.Process(height, handler)

		if wg != nil {
			wg.Done()
		}
	}
}

// Process handles the processing of a given height
func (pw *PoolWorker) Process(height int64, handler ClientHandler) {
	err := pw.client.Send(Request{Height: height})
	if err != nil {
		panic(err)
	}

	var res Response

	err = pw.client.Receive(&res)
	if err != nil {
		panic(err)
	}

	if res.Success {
		handler(res.Height, nil)
	} else {
		handler(res.Height, errors.New(res.Error))
	}
}

// Stop stops the pool worker
func (pw *PoolWorker) Stop() {
	close(pw.channel)
}
