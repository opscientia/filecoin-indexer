package server

import (
	"net/http"

	"github.com/filecoin-project/go-address"
	"github.com/gin-gonic/gin"
)

// GetTransactions lists all transactions
func (s *Server) GetTransactions(c *gin.Context) {
	transactions, _ := s.store.Transaction.FindAll(c.Query("height"))

	c.JSON(http.StatusOK, transactions)
}

// GetAccountTransactions returns account transactions
func (s *Server) GetAccountTransactions(c *gin.Context) {
	addr, err := address.NewFromString(c.Param("address"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	height := c.Query("height")

	id := s.client.Account.GetIDAddress(addr)
	pubkey := s.client.Account.GetPublicKeyAddress(addr)

	transactions, _ := s.store.Transaction.FindAllByAddress(height, id, pubkey)

	c.JSON(http.StatusOK, transactions)
}
