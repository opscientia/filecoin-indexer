package cli

import (
	"path"
	"runtime"

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

	dir := "migrations"

	_, filename, _, ok := runtime.Caller(1)
	if ok {
		dir = path.Join(path.Dir(filename), "../migrations")
	}

	switch cmd {
	case "migrate":
		return goose.Up(conn, dir)
	case "rollback":
		return goose.Down(conn, dir)
	default:
		return nil
	}
}
