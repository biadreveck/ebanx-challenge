package contract

import (
	domainObj "ebanx/challenge/domain/object"
)

type AccountRepository interface {
	GetBalance(string) (*domainObj.Balance, error)
	Deposit(string, float64) *domainObj.Balance
	Withdraw(string, float64) (*domainObj.Balance, error)
	Reset()
}
