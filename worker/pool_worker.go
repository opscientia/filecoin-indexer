package worker

import "sync"

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
func (pw *PoolWorker) Run(handler ResponseHandler, wg *sync.WaitGroup) {
	for height := range pw.channel {
		pw.Process(height, handler)

		if wg != nil {
			wg.Done()
		}
	}
}

// Process handles the processing of a given height
func (pw *PoolWorker) Process(height int64, handler ResponseHandler) {
	err := pw.client.Send(Request{Height: height})
	if err != nil {
		panic(err)
	}

	var res Response

	err = pw.client.Receive(&res)
	if err != nil {
		panic(err)
	}

	handler(res)
}

// Stop stops the pool worker
func (pw *PoolWorker) Stop() {
	close(pw.channel)
}
