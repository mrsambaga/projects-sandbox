package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	queryResult := fn(queries)
	if queryResult != nil {
		rollbackErr := tx.Rollback();
		if rollbackErr != nil {
			return fmt.Errorf("tx err: %s, rb err: %s", err, rollbackErr)
		}
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Ammount int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Ammount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount: arg.Ammount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount: arg.Ammount,
		})
		if err != nil {
			return err
		}

		if (arg.FromAccountID < arg.ToAccountID) {
			result.FromAccount, result.ToAccount, err = transferMoney(context.Background(), q, arg)
			if err != nil {
				return nil
			}
		} else {
			result.FromAccount, result.ToAccount, err = transferMoney(context.Background(), q, TransferTxParams{
				FromAccountID: arg.ToAccountID,
				ToAccountID: arg.FromAccountID,
				Ammount: -arg.Ammount,
			})
			if err != nil {
				return nil
			}
		}

		return nil
	})

	

	return result, err
}

func transferMoney(
	ctx context.Context,
	queries *Queries,
	arg TransferTxParams,
) (fromAccount Account, toAccount Account, err error) {
	fromAccount, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID: arg.FromAccountID,
		Amount: -arg.Ammount,
	})
	if err != nil {
		return
	}

	toAccount, err = queries.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID: arg.ToAccountID,
		Amount: arg.Ammount,
	})

	return
}