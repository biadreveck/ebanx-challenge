package repository

import (
	"ebanx/challenge/domain"
	domainI "ebanx/challenge/domain/contract"
	domainObj "ebanx/challenge/domain/object"
)

type accountInMemoryRepo struct {
	Data map[string]float64
}

func NewAccountInMemoryRepository() domainI.AccountRepository {
	return &accountInMemoryRepo{
		Data: make(map[string]float64),
	}
}

func (r *accountInMemoryRepo) GetBalance(accountId string) (balance *domainObj.Balance, err error) {
	accountValue, hasAccount := r.Data[accountId]
	if !hasAccount {
		err = domain.ErrNotFound
		return
	}

	balance = &domainObj.Balance{
		AccountId: accountId,
		Balance:   accountValue,
	}

	return
}

func (r *accountInMemoryRepo) Deposit(accountId string, value float64) (balance *domainObj.Balance) {
	accountValue, hasAccount := r.Data[accountId]
	if !hasAccount {
		accountValue = value
		r.Data[accountId] = value
	} else {
		accountValue += value
		r.Data[accountId] = accountValue
	}

	balance = &domainObj.Balance{
		AccountId: accountId,
		Balance:   accountValue,
	}

	return
}

func (r *accountInMemoryRepo) Withdraw(accountId string, value float64) (balance *domainObj.Balance, err error) {
	accountValue, hasAccount := r.Data[accountId]
	if !hasAccount {
		err = domain.ErrNotFound
		return
	}
	if accountValue < value {
		err = domain.ErrNotEnoughBalance
		return
	}

	accountValue -= value
	r.Data[accountId] = accountValue

	balance = &domainObj.Balance{
		AccountId: accountId,
		Balance:   accountValue,
	}

	return
}

func (r *accountInMemoryRepo) Reset() {
	r.Data = make(map[string]float64)
}
