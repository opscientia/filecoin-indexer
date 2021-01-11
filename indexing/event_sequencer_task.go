package indexing

import (
	"context"
	"errors"
	"strconv"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/filecoin-indexer/model/types"
	"github.com/figment-networks/filecoin-indexer/store"
)

// EventSequencerTask creates network events
type EventSequencerTask struct {
	store *store.Store
}

// NewEventSequencerTask creates the task
func NewEventSequencerTask(store *store.Store) pipeline.Task {
	return &EventSequencerTask{store: store}
}

// GetName returns the task name
func (t *EventSequencerTask) GetName() string {
	return "EventSequencer"
}

// Run performs the task
func (t *EventSequencerTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	err := t.trackStorageCapacityChanges(payload)
	if err != nil {
		return err
	}

	t.trackNewDeals(payload)
	t.trackSlashedDeals(payload)

	return nil
}

func (t *EventSequencerTask) trackStorageCapacityChanges(p *payload) error {
	for _, miner := range p.Miners {
		oldMiner, err := t.store.Miner.FindAtPreviousHeight(miner.Address, p.currentHeight)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			return err
		}

		if *miner.RawBytePower != *oldMiner.RawBytePower {
			event := model.Event{
				Height:       &p.currentHeight,
				MinerAddress: miner.Address,
				Kind:         types.StorageCapacityChangeEvent,

				Data: map[string]interface{}{
					"from": strconv.FormatUint(*oldMiner.RawBytePower, 10),
					"to":   strconv.FormatUint(*miner.RawBytePower, 10),
				},
			}

			p.Events = append(p.Events, &event)
		}
	}

	return nil
}

func (t *EventSequencerTask) trackNewDeals(p *payload) {
	for dealID, deal := range p.DealsData {
		if deal.State.SectorStartEpoch == -1 {
			continue
		}

		event := model.Event{
			Height:       &p.currentHeight,
			MinerAddress: deal.Proposal.Provider.String(),
			Kind:         types.NewDealEvent,

			Data: map[string]interface{}{
				"deal_id":        dealID,
				"client_address": deal.Proposal.Client.String(),
				"piece_size":     strconv.FormatUint(uint64(deal.Proposal.PieceSize), 10),
				"storage_price":  decimal.NewFromBigInt(deal.Proposal.StoragePricePerEpoch.Int, -18),
				"is_verified":    deal.Proposal.VerifiedDeal,
			},
		}

		p.Events = append(p.Events, &event)
	}
}

func (t *EventSequencerTask) trackSlashedDeals(p *payload) {
	for dealID, deal := range p.DealsData {
		if deal.State.SlashEpoch == -1 {
			continue
		}

		event := model.Event{
			Height:       &p.currentHeight,
			MinerAddress: deal.Proposal.Provider.String(),
			Kind:         types.SlashedDealEvent,

			Data: map[string]interface{}{"deal_id": dealID},
		}

		p.Events = append(p.Events, &event)
	}
}
