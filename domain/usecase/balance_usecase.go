package usecase

import (
	domain "ebanx/challenge/domain/contract"
	domainObj "ebanx/challenge/domain/object"
)

type balanceUsecase struct {
	accountRepo domain.AccountRepository
}

func NewBalanceUsecase(r domain.AccountRepository) domain.BalanceUsecase {
	return &balanceUsecase{
		accountRepo: r,
	}
}

func (uc *balanceUsecase) Get(accountId string) (balance *domainObj.Balance, err error) {
	balance, err = uc.accountRepo.GetBalance(accountId)
	return
}
