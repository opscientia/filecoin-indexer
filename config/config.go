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
	errEndpointRequired  = errors.New("RPC endpoint is required")
	errDatabaseRequired  = errors.New("database credentials are required")
	errRPCTimeoutInvalid = errors.New("RPC timeout is invalid")
)

// Config holds the configuration data
type Config struct {
	RPCEndpoint   string `json:"rpc_endpoint" envconfig:"RPC_ENDPOINT"`
	RPCTimeout    string `json:"rpc_timeout" envconfig:"RPC_TIMEOUT" default:"30s"`
	DatabaseDSN   string `json:"database_dsn" envconfig:"DATABASE_DSN"`
	ServerAddr    string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort    uint16 `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
	InitialHeight int64  `json:"initial_height" envconfig:"INITIAL_HEIGHT"`
	BatchSize     int64  `json:"batch_size" envconfig:"BATCH_SIZE"`
	Debug         bool   `json:"debug" envconfig:"DEBUG"`

	rpcTimeout time.Duration
}

// New creates a new configuration
func New() *Config {
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

	return nil
}

// ListenAddr returns the listen address for the API server
func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}

// ClientTimeout returns the timeout for the RPC client
func (c *Config) ClientTimeout() time.Duration {
	return c.rpcTimeout
}
