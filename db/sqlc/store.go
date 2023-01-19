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
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx error : %v, rb error : %v ", err, rberr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg AddTransferParams) (TransferTxResult, error) {

	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.AddTransfer(ctx, AddTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.AddEntry(ctx, AddEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.AddEntry(ctx, AddEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		if(arg.FromAccountID<arg.ToAccountID){
			result.FromAccount, result.ToAccount, err = updateAccountBalance(ctx,q,arg.FromAccountID,-arg.Amount,arg.ToAccountID,arg.Amount)

		} else {
			result.ToAccount, result.FromAccount, err = updateAccountBalance(ctx,q,arg.ToAccountID,arg.Amount,arg.FromAccountID,-arg.Amount)

		}
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

func updateAccountBalance(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
)(account1 Account, account2 Account, err error){
	account1,err=q.AddAccountBalance(ctx,AddAccountBalanceParams{
		Amount: amount1,
		ID: accountID1,
	})
	if err!=nil {
		return
	}
	account2,err=q.AddAccountBalance(ctx,AddAccountBalanceParams{
		Amount: amount2,
		ID: accountID2,
	})
	return

}
