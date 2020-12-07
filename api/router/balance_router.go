package router

import (
	"net/http"

	domainI "ebanx/challenge/domain/contract"

	"github.com/gin-gonic/gin"
)

type BalanceRouter struct {
	balanceUsecase domainI.BalanceUsecase
}

func NewBalanceRouter(router *gin.Engine, uc domainI.BalanceUsecase) {
	balanceRouter := &BalanceRouter{
		balanceUsecase: uc,
	}

	balance := router.Group("/v1/balance")
	{
		balance.GET("", balanceRouter.GetBalance)
	}
}

func (r *BalanceRouter) GetBalance(c *gin.Context) {
	accountId := c.Query("account_id")

	if accountId == "" {
		c.JSON(http.StatusBadRequest, 0)
		return
	}

	balance, err := r.balanceUsecase.Get(accountId)
	if err != nil {
		c.JSON(getStatusCode(err), 0)
		return
	}

	c.JSON(http.StatusOK, balance.Balance)
}
