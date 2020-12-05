package router

import (
	"net/http"

	"ebanx/challenge/api/payload"
	domainI "ebanx/challenge/domain/contract"
	domainObj "ebanx/challenge/domain/object"

	"github.com/gin-gonic/gin"
)

type EventRouter struct {
	eventUsecase domainI.EventUsecase
}

func NewEventRouter(router *gin.Engine, uc domainI.EventUsecase) {
	eventRouter := &EventRouter{
		eventUsecase: uc,
	}

	event := router.Group("/v1/event")
	{
		event.POST("", eventRouter.PostEvent)
	}
}

func (r *EventRouter) PostEvent(c *gin.Context) {
	var event payload.EventPayload
	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, 0)
		return
	}

	switch event.Type {
	case payload.EVENT_TYPE_DEPOSIT:
		r.processDeposit(c, event)
		return
	case payload.EVENT_TYPE_WITHDRAW:
		r.processWithdraw(c, event)
		return
	case payload.EVENT_TYPE_TRANSFER:
		r.processTransfer(c, event)
		return
	default:
		c.JSON(http.StatusBadRequest, 0)
		return
	}
}

func (r *EventRouter) processDeposit(c *gin.Context, event payload.EventPayload) {
	balance, err := r.eventUsecase.Deposit(domainObj.Transaction{
		AccountId: event.DestinationAccountId,
		Amount:    event.Amount,
	})

	if err != nil {
		c.JSON(getStatusCode(err), 0)
		return
	}

	c.JSON(http.StatusCreated, balance)
}

func (r *EventRouter) processWithdraw(c *gin.Context, event payload.EventPayload) {
	balance, err := r.eventUsecase.Withdraw(domainObj.Transaction{
		AccountId: event.OriginAccountId,
		Amount:    event.Amount,
	})

	if err != nil {
		c.JSON(getStatusCode(err), 0)
		return
	}

	c.JSON(http.StatusCreated, balance)
}

func (r *EventRouter) processTransfer(c *gin.Context, event payload.EventPayload) {
	transaction := domainObj.Transaction{
		AccountId: event.DestinationAccountId,
		Amount:    event.Amount,
	}
	balance, err := r.eventUsecase.Transfer(domainObj.TransferTransaction{
		OriginAccountId: event.OriginAccountId,
		Transaction:     transaction,
	})

	if err != nil {
		c.JSON(getStatusCode(err), 0)
		return
	}

	c.JSON(http.StatusCreated, balance)
}
