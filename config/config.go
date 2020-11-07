package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
)

var (
	errEndpointRequired = errors.New("RPC endpoint is required")
	errDatabaseRequired = errors.New("database credentials are required")
)

// Config holds the configuration data
type Config struct {
	RPCEndpoint   string `json:"rpc_endpoint" envconfig:"RPC_ENDPOINT"`
	DatabaseDSN   string `json:"database_dsn" envconfig:"DATABASE_DSN"`
	ServerAddr    string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort    uint16 `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
	InitialHeight int64  `json:"initial_height" envconfig:"INITIAL_HEIGHT"`
	BatchSize     int64  `json:"batch_size" envconfig:"BATCH_SIZE"`
	Debug         bool   `json:"debug" envconfig:"DEBUG"`
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

	return nil
}

// ListenAddr returns a listen address
func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}
