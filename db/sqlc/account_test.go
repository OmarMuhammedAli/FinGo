package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	testAccountParams := CreateAccountParams{
		Owner:    "Omar",
		Balance:  200,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), testAccountParams)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, testAccountParams.Owner, account.Owner)
	require.Equal(t, testAccountParams.Balance, account.Balance)
	require.Equal(t, testAccountParams.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}
