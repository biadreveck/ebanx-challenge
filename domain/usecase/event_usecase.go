package usecase

import (
	"ebanx/challenge/domain"
	domainI "ebanx/challenge/domain/contract"
	domainObj "ebanx/challenge/domain/object"
)

type eventUsecase struct {
	accountRepo domainI.AccountRepository
}

func NewEventUsecase(r domainI.AccountRepository) domainI.EventUsecase {
	return &eventUsecase{
		accountRepo: r,
	}
}

func (uc *eventUsecase) Deposit(transaction domainObj.Transaction) (b *domainObj.DepositBalance, err error) {
	if transaction.Amount <= 0 {
		err = domain.ErrBadParamInput
		return
	}

	balance := uc.accountRepo.Deposit(transaction.AccountId, transaction.Amount)
	b = &domainObj.DepositBalance{
		Destination: balance,
	}
	return
}

func (uc *eventUsecase) Withdraw(transaction domainObj.Transaction) (b *domainObj.WithdrawBalance, err error) {
	if transaction.Amount <= 0 {
		err = domain.ErrBadParamInput
		return
	}

	balance, err := uc.accountRepo.Withdraw(transaction.AccountId, transaction.Amount)
	if err != nil {
		return
	}

	b = &domainObj.WithdrawBalance{
		Origin: balance,
	}
	return
}

func (uc *eventUsecase) Transfer(transaction domainObj.TransferTransaction) (b *domainObj.TransferBalance, err error) {
	if transaction.Amount <= 0 {
		err = domain.ErrBadParamInput
		return
	}

	originBalance, err := uc.accountRepo.Withdraw(transaction.OriginAccountId, transaction.Amount)
	if err != nil {
		return
	}

	destinationBalance := uc.accountRepo.Deposit(transaction.AccountId, transaction.Amount)

	withdrawBalance := domainObj.WithdrawBalance{
		Origin: originBalance,
	}
	depositBalance := domainObj.DepositBalance{
		Destination: destinationBalance,
	}

	b = &domainObj.TransferBalance{
		withdrawBalance,
		depositBalance,
	}
	return
}
