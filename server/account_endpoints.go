package server

import (
	"net/http"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// GetAccount returns account details
func (s *Server) GetAccount(c *gin.Context) {
	addr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	actor, err := s.client.Account.GetActor(addr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id := s.client.Account.GetIDAddress(addr)
	pubkey := s.client.Account.GetPublicKeyAddress(addr)

	sent, _ := s.store.Transaction.CountSentByAddress(id, pubkey)
	received, _ := s.store.Transaction.CountReceivedByAddress(id, pubkey)

	account := model.Account{
		ID:                   id,
		PublicKey:            pubkey,
		Balance:              decimal.NewFromBigInt(actor.Balance.Int, -18),
		Nonce:                actor.Nonce,
		TransactionsSent:     sent,
		TransactionsReceived: received,
	}

	c.JSON(http.StatusOK, account)
}
