package server

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"github.com/figment-networks/filecoin-indexer/model"
)

// GetAccount returns account details
func (s *Server) GetAccount(c *gin.Context) {
	addr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		badRequest(c, fmt.Sprintf("invalid address: %s", err))
		return
	}

	actor, err := s.client.Account.GetActor(addr)
	if err != nil {
		notFound(c, err)
		return
	}

	id := s.client.Account.GetIDAddress(addr)
	pubkey := s.client.Account.GetPublicKeyAddress(addr)
	addresses := []string{id, pubkey}

	sent, err := s.store.Transaction.CountSentByAddresses(addresses)
	if err != nil {
		serverError(c, err)
		return
	}

	received, err := s.store.Transaction.CountReceivedByAddresses(addresses)
	if err != nil {
		serverError(c, err)
		return
	}

	account := model.Account{
		ID:                   id,
		PublicKey:            pubkey,
		Balance:              decimal.NewFromBigInt(actor.Balance.Int, -18),
		Nonce:                actor.Nonce,
		TransactionsSent:     sent,
		TransactionsReceived: received,
	}

	jsonOK(c, account)
}
