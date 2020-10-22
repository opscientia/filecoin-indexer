package cli

import (
	"path"
	"runtime"
	"strings"

	"github.com/pressly/goose"

	"github.com/figment-networks/filecoin-indexer/config"
)

func runMigrations(cfg *config.Config, cmd string) error {
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

	conn, err := store.Conn()
	if err != nil {
		return err
	}

	subcmd := "up"
	if chunks := strings.Split(cmd, ":"); len(chunks) > 1 {
		subcmd = chunks[1]
	}

	switch subcmd {
	case "up":
		return goose.Up(conn, dir)
	case "down":
		return goose.Down(conn, dir)
	default:
		return nil
	}
}
