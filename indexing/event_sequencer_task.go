package indexing

import (
	"context"
	"strconv"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/shopspring/decimal"
	"github.com/stretchr/stew/slice"

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

	err := t.fetchMinersAtPreviousHeight(payload)
	if err != nil {
		return err
	}

	err = t.fetchDealIDs(payload)
	if err != nil {
		return err
	}

	err = t.fetchSlashedDealIDs(payload)
	if err != nil {
		return err
	}

	t.trackStorageCapacityChanges(payload)
	t.trackSectorFaults(payload)
	t.trackNewDeals(payload)
	t.trackSlashedDeals(payload)

	return nil
}

func (t *EventSequencerTask) fetchMinersAtPreviousHeight(p *payload) error {
	miners, err := t.store.Miner.FindAllAtPreviousHeight(p.currentHeight)
	if err != nil {
		return err
	}

	p.StoredMiners = make(map[string]model.Miner)

	for _, miner := range miners {
		p.StoredMiners[miner.Address] = miner
	}

	return nil
}

func (t *EventSequencerTask) fetchDealIDs(p *payload) error {
	dealIDs, err := t.store.Event.DealIDsByKind(types.NewDealEvent)
	if err != nil {
		return err
	}

	p.StoredDealIDs = dealIDs

	return nil
}

func (t *EventSequencerTask) fetchSlashedDealIDs(p *payload) error {
	slashedDealIDs, err := t.store.Event.DealIDsByKind(types.SlashedDealEvent)
	if err != nil {
		return err
	}

	p.StoredSlashedDealIDs = slashedDealIDs

	return nil
}

func (t *EventSequencerTask) trackStorageCapacityChanges(p *payload) {
	for _, miner := range p.Miners {
		oldMiner, ok := p.StoredMiners[miner.Address]
		if !ok {
			continue
		}

		if *miner.RawBytePower == *oldMiner.RawBytePower {
			continue
		}

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

func (t *EventSequencerTask) trackSectorFaults(p *payload) {
	for _, miner := range p.Miners {
		oldMiner, ok := p.StoredMiners[miner.Address]
		if !ok {
			continue
		}

		faultsDiff := int32(*miner.FaultsCount - *oldMiner.FaultsCount)
		if faultsDiff == 0 {
			continue
		}

		var eventKind types.EventKind
		var sectorsCount uint64

		if faultsDiff > 0 {
			eventKind = types.SectorFaultEvent
			sectorsCount = uint64(faultsDiff)
		} else {
			eventKind = types.SectorRecoveryEvent
			sectorsCount = uint64(-faultsDiff)
		}

		event := model.Event{
			Height:       &p.currentHeight,
			MinerAddress: miner.Address,
			Kind:         eventKind,

			Data: map[string]interface{}{
				"sectors_count": strconv.FormatUint(sectorsCount, 10),
			},
		}

		p.Events = append(p.Events, &event)
	}
}

func (t *EventSequencerTask) trackNewDeals(p *payload) {
	for dealID, deal := range p.DealsData {
		if deal.State.SectorStartEpoch == -1 {
			continue
		}

		if slice.Contains(p.StoredDealIDs, dealID) {
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

		if slice.Contains(p.StoredSlashedDealIDs, dealID) {
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
