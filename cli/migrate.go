package cli

import (
	"github.com/pressly/goose"

	"github.com/figment-networks/filecoin-indexer/config"
)

func runMigrations(cfg *config.Config, cmd string) error {
	store, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	conn, err := store.Conn()
	if err != nil {
		return err
	}

	switch cmd {
	case "migrate":
		return goose.Up(conn, "migrations")
	case "rollback":
		return goose.Down(conn, "migrations")
	default:
		return nil
	}
}
