package contract

import (
	domainObj "ebanx/challenge/domain/object"
)

type BalanceUsecase interface {
	Get(string) (*domainObj.Balance, error)
}
