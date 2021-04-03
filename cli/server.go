package cli

import "github.com/figment-networks/filecoin-indexer/config"

func runServer(cfg *config.Config) error {
	store, err := initStore(cfg)
	if err != nil {
		return err
	}
	defer store.Close()

	client, err := initClient(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	server, err := initServer(cfg, store, client)
	if err != nil {
		return err
	}

	return server.Start()
}
