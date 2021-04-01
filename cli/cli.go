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
	var configPath string

	flag.StringVar(&command, "cmd", "", "Command to run")
	flag.StringVar(&configPath, "config", "", "Path to a config file")
	flag.Parse()

	cfg, err := initConfig(configPath)
	if err != nil {
		terminate(err)
	}

	config.InitRollbar(cfg)
	defer config.LogPanic()

	if command == "" {
		terminate("Command is required")
	}

	if err := runCommand(cfg, command); err != nil {
		terminate(err)
	}
}

func runCommand(cfg *config.Config, name string) error {
	switch name {
	case "migrate", "rollback":
		return runMigrations(cfg, name)
	case "sync":
		return runSync(cfg)
	case "worker":
		return runWorker(cfg)
	case "server":
		return runServer(cfg)
	default:
		return fmt.Errorf("%s is not a valid command", name)
	}
}

func terminate(message interface{}) {
	log.Fatal("ERROR: ", message)
}
