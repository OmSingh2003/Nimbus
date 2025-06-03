package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// Input parameters for transfer transaction
type TransferTXParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// Result of the transfer transaction
type TransferTXResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTX performs a money transfer from one account to another.
// It creates a transfer record, account entries, and updates account balances within a transaction.
func (store *Store) TransferTX(ctx context.Context, arg TransferTXParams) (TransferTXResult, error) {
	var result TransferTXResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// 1. Create Transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		// 2. Create entries
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// 3. Update balances
		// To avoid deadlock, we need to acquire locks in consistent order
		// Always lock the account with smaller ID first
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:      arg.FromAccountID,
				Balance: -arg.Amount,
			})
			if err != nil {
				return err
			}

			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:      arg.ToAccountID,
				Balance: arg.Amount,
			})
		} else {
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:      arg.ToAccountID,
				Balance: arg.Amount,
			})
			if err != nil {
				return err
			}

			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:      arg.FromAccountID,
				Balance: -arg.Amount,
			})
		}
		return err
	})

	return result, err
}
