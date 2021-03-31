package worker

// Request respresents a worker request
type Request struct {
	Height int64
}

// Response represents a worker response
type Response struct {
	Height  int64
	Success bool
	Error   string
}

// Client interacts with a worker
type Client interface {
	Send(Request) error
	Receive(*Response) error
	Close() error
}

// Server interacts with a manager
type Server interface {
	Receive(*Request) error
	Send(Response) error
}

// ServerHandler handles a request from a manager
type ServerHandler func(height int64) error

// ClientHandler handles a response from a worker
type ClientHandler func(height int64, err error)
