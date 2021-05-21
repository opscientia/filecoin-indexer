package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/figment-networks/filecoin-indexer/config"
)

// Run executes the command line interface
func Run() {
	var command string
	var mode string
	var configPath string

	flag.StringVar(&command, "cmd", "", "Command to run")
	flag.StringVar(&mode, "mode", "worker", "Fetcher mode")
	flag.StringVar(&configPath, "config", "", "Path to a config file")
	flag.Parse()

	cfg, err := initConfig(configPath, command, mode)
	if err != nil {
		terminate(err)
	}

	config.InitRollbar(cfg)
	defer config.LogPanic()

	if command == "" {
		terminate("Command is required")
	}

	if err := runCommand(cfg, command, mode); err != nil {
		terminate(err)
	}
}

func runCommand(cfg *config.Config, name string, mode string) error {
	switch name {
	case "migrate", "rollback":
		return runMigrations(cfg, name)
	case "fetcher":
		return runFetcher(cfg, mode)
	case "indexer":
		return runIndexer(cfg)
	case "server":
		return runServer(cfg)
	default:
		return fmt.Errorf("%s is not a valid command", name)
	}
}

func terminate(message interface{}) {
	log.Fatal("ERROR: ", message)
}
