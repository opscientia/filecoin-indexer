package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

// StartPipeline runs the indexing pipeline
func StartPipeline(cfg *config.Config, client *client.Client, store *store.Store) error {
	p := pipeline.NewDefault(NewPayloadFactory())

	p.SetTasks(pipeline.StageFetcher,
		NewEpochFetcherTask(client),
		NewMinerFetcherTask(client),
		NewTransactionFetcherTask(client),
	)

	p.SetTasks(pipeline.StageParser,
		NewEpochParserTask(),
		NewMinerParserTask(),
		NewTransactionParserTask(),
	)

	p.SetTasks(pipeline.StagePersistor,
		NewMinerPersistorTask(store),
		NewTransactionPersistorTask(store),
		NewEpochPersistorTask(store),
	)

	source, err := NewSource(cfg, client, store)
	if err != nil {
		return err
	}

	err = p.Start(context.Background(), source, NewSink(), &pipeline.Options{})
	if err != nil {
		return err
	}

	return nil
}
