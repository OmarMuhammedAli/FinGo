package db

import (
	"context"
	"testing"

	"github.com/OmarMuhammedAli/FinGo/util"
	"github.com/stretchr/testify/require"
)

func createRandEntry(t *testing.T, account Account) Entry {

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomInt(0, 10000),
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, account.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandAccount(t)
	createRandEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandAccount(t)
	entry := createRandEntry(t, account)

	retrievedEntry, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.Equal(t, retrievedEntry.AccountID, entry.AccountID)
	require.Equal(t, retrievedEntry.ID, entry.ID)
	require.Equal(t, retrievedEntry.Amount, entry.Amount)
	require.Equal(t, retrievedEntry.CreatedAt, entry.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account := createRandAccount(t)
	for i := 0; i < 10; i++ {
		createRandEntry(t, account)
	}

	args := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
