package router

import (
	"net/http"

	domainI "ebanx/challenge/domain/contract"

	"github.com/gin-gonic/gin"
)

type ResetRouter struct {
	resetUsecase domainI.ResetUsecase
}

func NewResetRouter(router *gin.Engine, uc domainI.ResetUsecase) {
	resetRouter := &ResetRouter{
		resetUsecase: uc,
	}

	reset := router.Group("/v1/reset")
	{
		reset.GET("", resetRouter.PostReset)
	}
}

func (r *ResetRouter) PostReset(c *gin.Context) {
	r.resetUsecase.Reset()
	c.JSON(http.StatusOK, nil)
}
