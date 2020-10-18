package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

// Run executes the command line interface
func Run() {
	var command string
	var configPath string

	flag.StringVar(&command, "cmd", "", "Command to run")
	flag.StringVar(&configPath, "config", "", "Path to a config file")
	flag.Parse()

	cfg, err := initConfig(configPath)
	if err != nil {
		terminate(err)
	}

	if command == "" {
		terminate("Command is required")
	}

	if err := runCommand(cfg, command); err != nil {
		terminate(err)
	}
}

func runCommand(cfg *config.Config, name string) error {
	switch name {
	case "migrate":
		return runMigrations(cfg)
	case "sync":
		return runSync(cfg)
	case "server":
		return runServer(cfg)
	default:
		return fmt.Errorf("%s is not a valid command", name)
	}
}

func terminate(message interface{}) {
	if message != nil {
		log.Fatal("ERROR: ", message)
	}
}

func initConfig(path string) (*config.Config, error) {
	cfg := config.New()

	if err := config.FromEnv(cfg); err != nil {
		return nil, err
	}

	if path != "" {
		if err := config.FromFile(path, cfg); err != nil {
			return nil, err
		}
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func initStore(cfg *config.Config) (*store.Store, error) {
	store, err := store.New(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func initClient(cfg *config.Config) (*client.Client, error) {
	return client.New(cfg.NodeURL), nil
}
