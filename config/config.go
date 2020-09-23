package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	NodeURL     string `json:"node_url" envconfig:"NODE_URL" required:"true"`
	DatabaseDSN string `json:"database_dsn" envconfig:"DATABASE_DSN" required:"true"`
	ServerAddr  string `json:"server_addr" envconfig:"SERVER_ADDR" default:"0.0.0.0"`
	ServerPort  int64  `json:"server_port" envconfig:"SERVER_PORT" default:"8080"`
}

func New() *Config {
	return &Config{}
}

func FromEnv(config *Config) error {
	return envconfig.Process("", config)
}

func (c *Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.ServerAddr, c.ServerPort)
}
