package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
)

var (
	errNodeRequired     = errors.New("Filecoin node URL is required")
	errDatabaseRequired = errors.New("Database credentials are required")
)

// Config holds the configuration data
type Config struct {
	NodeURL     string `json:"node_url" envconfig:"NODE_URL"`
	DatabaseDSN string `json:"database_dsn" envconfig:"DATABASE_DSN"`
	ServerAddr  string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort  int64  `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
}

// New returns a new configuration
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

// Validate returns an error if the configuration is invalid
func (c *Config) Validate() error {
	if c.NodeURL == "" {
		return errNodeRequired
	}

	if c.DatabaseDSN == "" {
		return errDatabaseRequired
	}

	return nil
}

// ListenAddr returns a full listen address
func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}
