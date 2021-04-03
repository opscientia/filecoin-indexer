package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	errEndpointRequired    = errors.New("RPC endpoint is required")
	errDatabaseRequired    = errors.New("database credentials are required")
	errRPCTimeoutInvalid   = errors.New("RPC timeout is invalid")
	errSyncIntervalInvalid = errors.New("sync interval is invalid")
)

// Config holds the configuration data
type Config struct {
	RPCEndpoint   string `json:"rpc_endpoint" envconfig:"RPC_ENDPOINT"`
	RPCTimeout    string `json:"rpc_timeout" envconfig:"RPC_TIMEOUT" default:"30s"`
	DatabaseDSN   string `json:"database_dsn" envconfig:"DATABASE_DSN"`
	InitialHeight int64  `json:"initial_height" envconfig:"INITIAL_HEIGHT"`
	BatchSize     int64  `json:"batch_size" envconfig:"BATCH_SIZE"`
	SyncInterval  string `json:"sync_interval" envconfig:"SYNC_INTERVAL" default:"1s"`
	ServerAddr    string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort    uint16 `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
	RollbarToken  string `json:"rollbar_token" envconfig:"ROLLBAR_TOKEN"`
	RollbarEnv    string `json:"rollbar_env" envconfig:"ROLLBAR_ENV" default:"development"`
	MetricsAddr   string `json:"metrics_addr" envconfig:"METRICS_ADDR" default:"127.0.0.1"`
	MetricsPort   uint16 `json:"metrics_port" envconfig:"METRICS_PORT" default:"8090"`
	Debug         bool   `json:"debug" envconfig:"DEBUG"`

	rpcTimeout   time.Duration
	syncInterval time.Duration
}

// NewConfig creates a configuration
func NewConfig() *Config {
	return &Config{}
}

// FromEnv reads the configuration from environment variables
func FromEnv(config *Config) error {
	return envconfig.Process("", config)
}

// FromFile reads the configuration from a file
func FromFile(path string, config *Config) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, config)
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.RPCEndpoint == "" {
		return errEndpointRequired
	}

	if c.DatabaseDSN == "" {
		return errDatabaseRequired
	}

	d, err := time.ParseDuration(c.RPCTimeout)
	if err != nil {
		return errRPCTimeoutInvalid
	}
	c.rpcTimeout = d

	d, err = time.ParseDuration(c.SyncInterval)
	if err != nil {
		return errSyncIntervalInvalid
	}
	c.syncInterval = d

	return nil
}

// ServerListenAddr returns the listen address for the API server
func (c *Config) ServerListenAddr() string {
	return listenAddr(c.ServerAddr, c.ServerPort)
}

// MetricsListenAddr returns the listen address for the metrics server
func (c *Config) MetricsListenAddr() string {
	return listenAddr(c.MetricsAddr, c.MetricsPort)
}

func listenAddr(addr string, port uint16) string {
	return fmt.Sprintf("%s:%d", addr, port)
}

// ClientRPCTimeout returns the timeout for the RPC client
func (c *Config) ClientRPCTimeout() time.Duration {
	return c.rpcTimeout
}

// PipelineSyncInterval returns the interval between synchronization jobs
func (c *Config) PipelineSyncInterval() time.Duration {
	return c.syncInterval
}
