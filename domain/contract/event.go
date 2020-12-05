package contract

import (
	domainObj "ebanx/challenge/domain/object"
)

type EventUsecase interface {
	Deposit(domainObj.Transaction) (*domainObj.DepositBalance, error)
	Withdraw(domainObj.Transaction) (*domainObj.WithdrawBalance, error)
	Transfer(domainObj.TransferTransaction) (*domainObj.TransferBalance, error)
}
