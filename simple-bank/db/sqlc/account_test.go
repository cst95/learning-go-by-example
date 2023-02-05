package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/cst95/learning-go-by-example/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountOne := createRandomAccount(t)
	accountTwo, err := testQueries.GetAccount(context.Background(), accountOne.ID)

	require.NoError(t, err)
	require.NotEmpty(t, accountTwo)
	require.Equal(t, accountOne, accountTwo)
}

func TestUpdateAccount(t *testing.T) {
	accountOne := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      accountOne.ID,
		Balance: util.RandomMoney(),
	}

	accountTwo, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Balance, accountTwo.Balance)

	updatedAccount, err := testQueries.GetAccount(context.Background(), arg.ID)

	require.NoError(t, err)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
}

func TestDeleteAccount(t *testing.T) {
	accountOne := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), accountOne.ID)

	require.NoError(t, err)

	accountTwo, err := testQueries.GetAccount(context.Background(), accountOne.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountTwo)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
