package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/rollbar/rollbar-go"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

// StartPipeline runs the indexing pipeline
func StartPipeline(cfg *config.Config, client *client.Client, store *store.Store) error {
	source, err := NewSource(cfg, client, store)
	if err != nil {
		return err
	}

	p := pipeline.NewDefault(NewPayloadFactory())

	p.SetTasks(pipeline.StageFetcher,
		NewEpochFetcherTask(client),
		NewDealFetcherTask(client),
		NewMinerFetcherTask(client),
		NewTransactionFetcherTask(client),
	)

	p.SetTasks(pipeline.StageParser,
		NewEpochParserTask(),
		NewMinerParserTask(),
		NewTransactionParserTask(),
	)

	p.SetTasks(pipeline.StageSequencer,
		NewEventSequencerTask(store),
	)

	p.SetTasks(pipeline.StagePersistor,
		NewBeginTransactionTask(store),
		NewMinerPersistorTask(store),
		NewTransactionPersistorTask(store),
		NewEventPersistorTask(store),
		NewEpochPersistorTask(store),
		NewCommitTransactionTask(store),
	)

	sink := NewSink(store)

	err = p.Start(context.Background(), source, sink, &pipeline.Options{})
	if err != nil {
		store.Rollback()

		rollbar.Error(err)
		rollbar.Wait()

		return err
	}

	return nil
}