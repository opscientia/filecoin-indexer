package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	errRPCEndpointRequired = errors.New("RPC endpoint is required")
	errDatabaseDSNRequired = errors.New("database DSN is required")
	errRPCTimeoutInvalid   = errors.New("RPC timeout is invalid")
	errSyncIntervalInvalid = errors.New("sync interval is invalid")
)

// Config holds the configuration data
type Config struct {
	RPCEndpoint   string `json:"rpc_endpoint" envconfig:"RPC_ENDPOINT"`
	RPCTimeout    string `json:"rpc_timeout" envconfig:"RPC_TIMEOUT" default:"30s"`
	DatabaseDSN   string `json:"database_dsn" envconfig:"DATABASE_DSN"`
	RedisURL      string `json:"redis_url" envconfig:"REDIS_URL" default:"127.0.0.1:6379"`
	InitialHeight int64  `json:"initial_height" envconfig:"INITIAL_HEIGHT"`
	BatchSize     int64  `json:"batch_size" envconfig:"BATCH_SIZE"`
	SyncInterval  string `json:"sync_interval" envconfig:"SYNC_INTERVAL" default:"1s"`
	WorkerAddr    string `json:"worker_addr" envconfig:"WORKER_ADDR" default:"127.0.0.1"`
	WorkerPort    uint16 `json:"worker_port" envconfig:"WORKER_PORT" default:"7000"`
	Workers       string `json:"workers" envconfig:"WORKERS"`
	ServerAddr    string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort    uint16 `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
	RollbarToken  string `json:"rollbar_token" envconfig:"ROLLBAR_TOKEN"`
	RollbarEnv    string `json:"rollbar_env" envconfig:"ROLLBAR_ENV" default:"development"`
	MetricsAddr   string `json:"metrics_addr" envconfig:"METRICS_ADDR" default:"127.0.0.1"`
	MetricsPort   uint16 `json:"metrics_port" envconfig:"METRICS_PORT" default:"8090"`
	Debug         bool   `json:"debug" envconfig:"DEBUG"`

	RPCTimeoutDuration   time.Duration
	SyncIntervalDuration time.Duration
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
func (c *Config) Validate(cmd, mode string) error {
	if c.RPCEndpoint == "" {
		if cmd != "migrate" && cmd != "rollback" {
			return errRPCEndpointRequired
		}
	}

	if c.DatabaseDSN == "" {
		if cmd != "fetcher" || mode != "worker" {
			return errDatabaseDSNRequired
		}
	}

	d, err := time.ParseDuration(c.RPCTimeout)
	if err != nil {
		return errRPCTimeoutInvalid
	}
	c.RPCTimeoutDuration = d

	d, err = time.ParseDuration(c.SyncInterval)
	if err != nil {
		return errSyncIntervalInvalid
	}
	c.SyncIntervalDuration = d

	return nil
}

// WorkerEndpoints returns worker endpoints
func (c *Config) WorkerEndpoints() []string {
	return strings.Fields(c.Workers)
}

// WorkerListenAddr returns the listen address for the worker server
func (c *Config) WorkerListenAddr() string {
	return listenAddr(c.WorkerAddr, c.WorkerPort)
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
