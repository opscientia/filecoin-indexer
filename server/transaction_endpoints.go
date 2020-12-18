package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/figment-networks/filecoin-indexer/model"
)

// GetTransactions lists all transactions
func (s *Server) GetTransactions(c *gin.Context) {
	var transactions *[]model.Transaction

	address := c.Query("address")

	if address != "" {
		transactions, _ = s.store.Transaction.FindAllByAddress(address)
	} else {
		transactions, _ = s.store.Transaction.FindAll()
	}

	c.JSON(http.StatusOK, transactions)
}
