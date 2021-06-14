package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/datalake"
	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/config"
	"github.com/figment-networks/filecoin-indexer/store"
)

// RunFetcherPipeline runs the fetching pipeline
func RunFetcherPipeline(height int64, client *client.Client, dl *datalake.DataLake) error {
	p := pipeline.NewDefault(NewPayloadFactory(dl))

	p.SetTasks(pipeline.StageFetcher,
		NewEpochFetcherTask(client),
		NewDealFetcherTask(client),
		NewMinerFetcherTask(client),
		NewTransactionFetcherTask(client),
	)

	_, err := p.Run(context.Background(), height, nil)
	if err != nil {
		return err
	}

	return nil
}

// StartIndexerPipeline starts the indexing pipeline
func StartIndexerPipeline(cfg *config.Config, client *client.Client, store *store.Store, dl *datalake.DataLake) error {
	source, err := NewSource(cfg, client, store)
	if err != nil {
		return err
	}

	p := pipeline.NewDefault(NewPayloadFactory(dl))

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

	err = p.Start(context.Background(), source, sink, nil)
	if err != nil {
		store.Rollback()

		return err
	}

	return nil
}
