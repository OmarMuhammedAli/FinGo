package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	// Composition of the Queries struct to perform singular queries within a tx.
	*Queries
	// Essential for starting new DB transactions
	db *sql.DB
}

// Create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Executes a function within a transaction.
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// 1. Start the transaction using the Store's db field.
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}

	// Create a new query instance with the reference to the tx. This way, all queries
	// run using this instance will be run within the same tx.
	q := New(tx)
	err = fn(q)

	// Rollback or Commit tx.
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rbErr: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// Parameters of a transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResults struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// transaction key to add to a context to hold a value of the transaction name.
// var txKey = struct{}{}

// TransferTx performs a money transfer from one account to another.
// 1. Create account entries records
// 2. Create a transfer record
// 3. Update accounts' balances
func (s *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResults, error) {
	// result variable to be returned
	var result TransferTxResults

	// Execute all steps within the same transaction
	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey).(string)

		// 1. Create account entries
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})
		if err != nil {
			return err
		}

		// 2. Create a transfer record
		transferParams := TransferTxParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		}
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(transferParams))
		if err != nil {
			return err
		}

		// 3. TODO: Update both accounts' balances
		// a. Retrieve From Account from which the amount will be deducted

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.FromAccountID,
			Amount: -args.Amount,
		})
		if err != nil {
			return err
		}

		// b. Retrieve ToAccount to which the amount will be added.
		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     args.ToAccountID,
			Amount: args.Amount,
		})
		if err != nil {
			return err
		}
		return err
	})

	return result, err
}
