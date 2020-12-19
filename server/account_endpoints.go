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

	account := model.Account{
		Address: addr.String(),
		Balance: decimal.NewFromBigInt(actor.Balance.Int, -18),
	}

	c.JSON(http.StatusOK, account)
}
