package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
	"github.com/figment-networks/filecoin-indexer/store"
)

func StartPipeline(client *client.Client, store *store.Store) error {
	p := pipeline.NewDefault(NewPayloadFactory())

	p.SetTasks(pipeline.StageFetcher, NewMinerFetcherTask(client))
	p.SetTasks(pipeline.StageParser, NewMinerParserTask())
	p.SetTasks(pipeline.StagePersistor, NewMinerPersistorTask(store))

	err := p.Start(context.Background(), NewSource(), NewSink(), &pipeline.Options{})
	if err != nil {
		return err
	}

	return nil
}
