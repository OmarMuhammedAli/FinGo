package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/OmarMuhammedAli/FinGo/util"
	"github.com/stretchr/testify/require"
)

func createRandAccount(t *testing.T) Account {
	testAccountParams := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(0, 1000),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), testAccountParams)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, testAccountParams.Owner, account.Owner)
	require.Equal(t, testAccountParams.Balance, account.Balance)
	require.Equal(t, testAccountParams.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandAccount(t)

	retrievedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)

	require.Equal(t, account.ID, retrievedAccount.ID)
	require.Equal(t, account.Balance, retrievedAccount.Balance)
	require.Equal(t, account.Owner, retrievedAccount.Owner)
	require.Equal(t, account.CreatedAt, retrievedAccount.CreatedAt)
	require.Equal(t, account.Currency, retrievedAccount.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandAccount(t)

	args := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomInt(0, 10000),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), args)

	require.NoError(t, err)
	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.CreatedAt, updatedAccount.CreatedAt)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandAccount(t)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
