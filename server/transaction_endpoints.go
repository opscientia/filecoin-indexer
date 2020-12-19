package server

import (
	"net/http"

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
