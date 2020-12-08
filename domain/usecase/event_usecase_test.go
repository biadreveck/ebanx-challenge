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

func TestDeposit(t *testing.T) {
	mockAccountRepo := new(mocks.AccountRepository)
	mockBalance := &domainObj.Balance{
		AccountId: "100",
		Balance:   10,
	}

	t.Run("deposit-success", func(t *testing.T) {
		mockAccountRepo.On("Deposit", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(mockBalance).Once()

		transaction := domainObj.Transaction{
			AccountId: "100",
			Amount:    10,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Deposit(transaction)

		assert.NotNil(t, balance)
		assert.NoError(t, err)
		assert.Equal(t, mockBalance, balance.Destination)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("deposit-error-amount", func(t *testing.T) {
		transaction := domainObj.Transaction{
			AccountId: "100",
			Amount:    -10,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Deposit(transaction)

		assert.Error(t, err)
		assert.Nil(t, balance)
	})
}

func TestWithdraw(t *testing.T) {
	mockAccountRepo := new(mocks.AccountRepository)
	mockBalance := &domainObj.Balance{
		AccountId: "100",
		Balance:   10,
	}

	t.Run("withdraw-success", func(t *testing.T) {
		mockAccountRepo.On("Withdraw", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(mockBalance, nil).Once()

		transaction := domainObj.Transaction{
			AccountId: "100",
			Amount:    10,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Withdraw(transaction)

		assert.NotNil(t, balance)
		assert.NoError(t, err)
		assert.Equal(t, mockBalance, balance.Origin)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("withdraw-error-amount", func(t *testing.T) {
		transaction := domainObj.Transaction{
			AccountId: "100",
			Amount:    -10,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Withdraw(transaction)

		assert.Error(t, err)
		assert.Nil(t, balance)
	})

	t.Run("withdraw-error-notfound", func(t *testing.T) {
		mockAccountRepo.On("Withdraw", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(nil, domain.ErrNotFound).Once()

		transaction := domainObj.Transaction{
			AccountId: "200",
			Amount:    10,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Withdraw(transaction)

		assert.Error(t, err)
		assert.Nil(t, balance)

		mockAccountRepo.AssertExpectations(t)
	})
}

func TestTransfer(t *testing.T) {
	mockAccountRepo := new(mocks.AccountRepository)
	mockBalance := &domainObj.Balance{
		AccountId: "100",
		Balance:   10,
	}
	mockBalanceDestination := &domainObj.Balance{
		AccountId: "200",
		Balance:   20,
	}

	t.Run("transfer-success", func(t *testing.T) {
		mockAccountRepo.On("Deposit", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(mockBalanceDestination).Once()
		mockAccountRepo.On("Withdraw", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(mockBalance, nil).Once()

		transaction := domainObj.Transaction{
			AccountId: "200",
			Amount:    10,
		}
		transferTransaction := domainObj.TransferTransaction{
			OriginAccountId: "100",
			Transaction:     transaction,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Transfer(transferTransaction)

		assert.NotNil(t, balance)
		assert.NoError(t, err)
		assert.Equal(t, mockBalance, balance.Origin)
		assert.Equal(t, mockBalanceDestination, balance.Destination)

		mockAccountRepo.AssertExpectations(t)
	})

	t.Run("transfer-error-amount", func(t *testing.T) {
		transaction := domainObj.Transaction{
			AccountId: "200",
			Amount:    -10,
		}
		transferTransaction := domainObj.TransferTransaction{
			OriginAccountId: "100",
			Transaction:     transaction,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Transfer(transferTransaction)

		assert.Error(t, err)
		assert.Nil(t, balance)
	})

	t.Run("transfer-error-notfound", func(t *testing.T) {
		mockAccountRepo.On("Withdraw", mock.AnythingOfType("string"), mock.AnythingOfType("float64")).Return(nil, domain.ErrNotFound).Once()

		transaction := domainObj.Transaction{
			AccountId: "200",
			Amount:    10,
		}
		transferTransaction := domainObj.TransferTransaction{
			OriginAccountId: "100",
			Transaction:     transaction,
		}

		uc := ucase.NewEventUsecase(mockAccountRepo)
		balance, err := uc.Transfer(transferTransaction)

		assert.Error(t, err)
		assert.Nil(t, balance)

		mockAccountRepo.AssertExpectations(t)
	})
}
