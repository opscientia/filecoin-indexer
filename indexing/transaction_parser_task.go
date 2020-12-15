package indexing

import (
	"context"

	"github.com/figment-networks/indexing-engine/pipeline"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/shopspring/decimal"

	"github.com/figment-networks/filecoin-indexer/model"
)

// TransactionParserTask transforms raw transaction data
type TransactionParserTask struct{}

// NewTransactionParserTask creates the task
func NewTransactionParserTask() pipeline.Task {
	return &TransactionParserTask{}
}

// GetName returns the task name
func (t *TransactionParserTask) GetName() string {
	return "TransactionParser"
}

// Run performs the task
func (t *TransactionParserTask) Run(ctx context.Context, p pipeline.Payload) error {
	payload := p.(*payload)

	for i, msg := range payload.TransactionsMessages {
		transaction := model.Transaction{
			CID:    payload.TransactionsCIDs[i].String(),
			Height: &payload.currentHeight,
			From:   msg.From.String(),
			To:     msg.To.String(),
			Value:  decimal.NewFromBigInt(msg.Value.Int, -18),
			Method: methodNames[msg.Method],
		}

		payload.Transactions = append(payload.Transactions, &transaction)
	}

	return nil
}

var methodNames = map[abi.MethodNum]string{
	builtin.MethodSend:                            "Send",
	builtin.MethodsMiner.Constructor:              "Constructor",
	builtin.MethodsMiner.ControlAddresses:         "ControlAddresses",
	builtin.MethodsMiner.ChangeWorkerAddress:      "ChangeWorkerAddress",
	builtin.MethodsMiner.ChangePeerID:             "ChangePeerID",
	builtin.MethodsMiner.SubmitWindowedPoSt:       "SubmitWindowedPoSt",
	builtin.MethodsMiner.PreCommitSector:          "PreCommitSector",
	builtin.MethodsMiner.ProveCommitSector:        "ProveCommitSector",
	builtin.MethodsMiner.ExtendSectorExpiration:   "ExtendSectorExpiration",
	builtin.MethodsMiner.TerminateSectors:         "TerminateSectors",
	builtin.MethodsMiner.DeclareFaults:            "DeclareFaults",
	builtin.MethodsMiner.DeclareFaultsRecovered:   "DeclareFaultsRecovered",
	builtin.MethodsMiner.OnDeferredCronEvent:      "OnDeferredCronEvent",
	builtin.MethodsMiner.CheckSectorProven:        "CheckSectorProven",
	builtin.MethodsMiner.ApplyRewards:             "ApplyRewards",
	builtin.MethodsMiner.ReportConsensusFault:     "ReportConsensusFault",
	builtin.MethodsMiner.WithdrawBalance:          "WithdrawBalance",
	builtin.MethodsMiner.ConfirmSectorProofsValid: "ConfirmSectorProofsValid",
	builtin.MethodsMiner.ChangeMultiaddrs:         "ChangeMultiaddrs",
	builtin.MethodsMiner.CompactPartitions:        "CompactPartitions",
	builtin.MethodsMiner.CompactSectorNumbers:     "CompactSectorNumbers",
	builtin.MethodsMiner.ConfirmUpdateWorkerKey:   "ConfirmUpdateWorkerKey",
	builtin.MethodsMiner.RepayDebt:                "RepayDebt",
	builtin.MethodsMiner.ChangeOwnerAddress:       "ChangeOwnerAddress",
}
