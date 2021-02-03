package server

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/store"
)

// GetTransactions lists all transactions
func (s *Server) GetTransactions(c *gin.Context) {
	height := c.Query("height")

	pagination := store.Pagination{}
	if err := c.Bind(&pagination); err != nil {
		badRequest(c, err)
		return
	}

	result, err := s.store.Transaction.FindAll(height, pagination)
	if err != nil {
		badRequest(c, err)
		return
	}

	jsonOK(c, result)
}

// GetAccountTransactions returns account transactions
func (s *Server) GetAccountTransactions(c *gin.Context) {
	addr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		badRequest(c, fmt.Sprintf("invalid address: %s", err))
		return
	}

	addresses := []string{
		s.client.Account.GetIDAddress(addr),
		s.client.Account.GetPublicKeyAddress(addr),
	}

	height := c.Query("height")

	pagination := store.Pagination{}
	if err := c.Bind(&pagination); err != nil {
		badRequest(c, err)
		return
	}

	result, err := s.store.Transaction.FindAllByAddresses(addresses, height, pagination)
	if err != nil {
		badRequest(c, err)
		return
	}

	jsonOK(c, result)
}
