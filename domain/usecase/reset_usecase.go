package usecase

import (
	domainI "ebanx/challenge/domain/contract"
)

type resetUsecase struct {
	accountRepo domainI.AccountRepository
}

func NewResetUsecase(r domainI.AccountRepository) domainI.ResetUsecase {
	return &resetUsecase{
		accountRepo: r,
	}
}

func (uc *resetUsecase) Reset() {
	uc.accountRepo.Reset()
}
