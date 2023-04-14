package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	// Create a Store instance to run the transaction
	s := NewStore(testDb)

	// Create 2 dummy accounts to complete the transfers between them.
	fa := createRandAccount(t)
	ta := createRandAccount(t)

	// run n concurrent transactions in separate goroutines
	n := 5
	amount := int64(10)

	// Create errors and results channels to record err and result of each running
	// goroutine.

	es := make(chan error)
	rs := make(chan TransferTxResults)

	for i := 0; i < n; i++ {
		go func() {
			result, err := s.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fa.ID,
				ToAccountID:   ta.ID,
				Amount:        amount,
			})
			es <- err
			rs <- result
		}()
	}

	// Test results of running transactions
	for i := 0; i < n; i++ {
		err := <-es
		require.NoError(t, err)

		r := <-rs
		require.NotEmpty(t, r)

		// Check transfer
		transfer := r.Transfer
		require.NotEmpty(t, transfer)

		require.Equal(t, fa.ID, transfer.FromAccountID)
		require.Equal(t, ta.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = s.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check entries
		fe := r.FromEntry
		require.NotEmpty(t, fe)

		require.Equal(t, fe.AccountID, fa.ID)
		require.Equal(t, fe.Amount, -amount)

		require.NotZero(t, fe.ID)
		require.NotZero(t, fe.CreatedAt)

		_, err = s.GetEntry(context.Background(), fe.ID)
		require.NoError(t, err)

		te := r.ToEntry
		require.NotEmpty(t, te)

		require.Equal(t, te.AccountID, ta.ID)
		require.Equal(t, te.Amount, amount)

		require.NotZero(t, te.ID)
		require.NotZero(t, te.CreatedAt)

		_, err = s.GetEntry(context.Background(), te.ID)
		require.NoError(t, err)

		// TODO: Check accounts' balances
	}

}
