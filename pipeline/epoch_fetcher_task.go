package pipeline

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"

	"github.com/figment-networks/filecoin-indexer/client"
)

// EpochFetcherTask fetches raw epoch data
type EpochFetcherTask struct {
	client *client.Client
}

// EpochFetcherTaskName represents the name of the task
const EpochFetcherTaskName = "EpochFetcher"

// NewEpochFetcherTask creates the task
func NewEpochFetcherTask(client *client.Client) pipeline.Task {
	return &EpochFetcherTask{client: client}
}

// GetName returns the task name
func (t *EpochFetcherTask) GetName() string {
	return EpochFetcherTaskName
}

// Run performs the task
func (t *EpochFetcherTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	tipset, err := t.client.Epoch.GetTipsetByHeight(payload.currentHeight)
	if err != nil {
		return err
	}

	payload.EpochTipset = tipset

	return nil
}
