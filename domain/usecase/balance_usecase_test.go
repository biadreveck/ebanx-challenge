package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"ebanx/challenge/domain"
	"ebanx/challenge/domain/contract/mocks"
	domainObj "ebanx/challenge/domain/object"
	ucase "ebanx/challenge/domain/usecase"
)

func TestGet(t *testing.T) {
	mockAccountRepo := new(mocks.AccountRepository)
	mockBalance := &domainObj.Balance{
		AccountId: "100",
		Balance:   10,
	}

	t.Run("success", func(t *testing.T) {
		mockAccountRepo.On("GetBalance", mock.AnythingOfType("string")).Return(mockBalance, nil).Once()

		uc := ucase.NewBalanceUsecase(mockAccountRepo)
		balance, err := uc.Get("100")

		assert.NotNil(t, balance)
		assert.NoError(t, err)
		assert.Equal(t, mockBalance, balance)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("error-notfound", func(t *testing.T) {
		mockAccountRepo.On("GetBalance", mock.AnythingOfType("string")).Return(nil, domain.ErrNotFound).Once()

		uc := ucase.NewBalanceUsecase(mockAccountRepo)
		balance, err := uc.Get("100")

		assert.Error(t, err)
		assert.Nil(t, balance)

		mockAccountRepo.AssertExpectations(t)
	})
}
