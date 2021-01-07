package server

import (
	"net/http"

	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
)

// GetTransactions lists all transactions
func (s *Server) GetTransactions(c *gin.Context) {
	var transactions *[]model.Transaction

	height := c.Query("height")

	if height != "" {
		transactions, _ = s.store.Transaction.FindAllByHeight(height)
	} else {
		transactions, _ = s.store.Transaction.FindAll()
	}

	c.JSON(http.StatusOK, transactions)
}

// GetAccountTransactions returns account transactions
func (s *Server) GetAccountTransactions(c *gin.Context) {
	var transactions *[]model.Transaction

	addr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	id := s.client.Account.GetIDAddress(addr)
	pubkey := s.client.Account.GetPublicKeyAddress(addr)

	height := c.Query("height")

	if height != "" {
		transactions, _ = s.store.Transaction.FindAllByAddressAndHeight(height, id, pubkey)
	} else {
		transactions, _ = s.store.Transaction.FindAllByAddress(id, pubkey)
	}

	c.JSON(http.StatusOK, transactions)
}
