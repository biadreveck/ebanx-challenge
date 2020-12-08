package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ebanx/challenge/data/repository"
	"ebanx/challenge/domain"
)

func TestDeposit(t *testing.T) {
	t.Run("deposit-success", func(t *testing.T) {
		repo := repository.NewAccountInMemoryRepository()

		balance := repo.Deposit("100", 10)

		assert.NotNil(t, balance)
		assert.Equal(t, "100", balance.AccountId)
		assert.Equal(t, 10.0, balance.Balance)

		balance = repo.Deposit("200", 20)

		assert.NotNil(t, balance)
		assert.Equal(t, "200", balance.AccountId)
		assert.Equal(t, 20.0, balance.Balance)
	})
}

func TestWithdraw(t *testing.T) {
	t.Run("withdraw-success", func(t *testing.T) {
		repo := repository.NewAccountInMemoryRepository()

		balance := repo.Deposit("100", 10)

		assert.NotNil(t, balance)
		assert.Equal(t, "100", balance.AccountId)
		assert.Equal(t, 10.0, balance.Balance)

		balance, err := repo.Withdraw("100", 5)

		assert.NoError(t, err)
		assert.NotNil(t, balance)
		assert.Equal(t, "100", balance.AccountId)
		assert.Equal(t, 5.0, balance.Balance)
	})

	t.Run("withdraw-error-notenough", func(t *testing.T) {
		repo := repository.NewAccountInMemoryRepository()

		balance := repo.Deposit("100", 10)

		assert.NotNil(t, balance)
		assert.Equal(t, "100", balance.AccountId)
		assert.Equal(t, 10.0, balance.Balance)

		balance, err := repo.Withdraw("100", 15)

		assert.Error(t, err)
		assert.Nil(t, balance)
		assert.Equal(t, domain.ErrNotEnoughBalance, err)
	})

	t.Run("withdraw-error-notfound", func(t *testing.T) {
		repo := repository.NewAccountInMemoryRepository()

		balance, err := repo.Withdraw("100", 5)

		assert.Error(t, err)
		assert.Nil(t, balance)
		assert.Equal(t, domain.ErrNotFound, err)
	})
}
