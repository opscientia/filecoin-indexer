package cli

import (
	"path"
	"runtime"

	"github.com/pressly/goose"

	"github.com/figment-networks/filecoin-indexer/config"
)

func runMigrations(cfg *config.Config) error {
	store, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	dir := "migrations"

	_, filename, _, ok := runtime.Caller(1)
	if ok {
		dir = path.Join(path.Dir(filename), "../migrations")
	}

	return goose.Up(store.Conn(), dir)
}
