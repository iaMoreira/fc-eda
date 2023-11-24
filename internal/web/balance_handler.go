package web

import (
	"fmt"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com/gin-gonic/gin"
)

type WebBalanceHandler struct {
	db *database.BalanceDB
}

func NewWebBalanceHandler(balanceDB *database.BalanceDB) *WebBalanceHandler {
	return &WebBalanceHandler{
		db: balanceDB,
	}
}

func (h *WebBalanceHandler) GetBalance(c *gin.Context) {
	AccountId := c.Params.ByName("AccountId")
	fmt.Println(AccountId)
	balance, err := h.db.Get(AccountId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		c.Writer.WriteString(err.Error())
		return
	}

	if balance == nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.Writer.WriteString("not fount")
		return
	}

	c.IndentedJSON(http.StatusOK, balance)
}
