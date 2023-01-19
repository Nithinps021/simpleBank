package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T){
	store:=NewStore(db)

	account1:=CreateRandomAccount(t)
	account2:=CreateRandomAccount(t)

	n:=5
	amount:=int64(10)

	errs:=make(chan error)
	results:=make(chan TransferTxResult)

	for i:=0;i<n;i++{
		go func () {
			result,err:=store.TransferTx(context.Background(),AddTransferParams{
				FromAccountID: account1.ID,
				ToAccountID: account2.ID,
				Amount: amount,
			})
	
			errs <-err
			results<-result
		}()
	}
	
	for i:=0;i<n;i++{
		fmt.Print(i)
		err:=<-errs
		require.NoError(t,err)
		result:=<-results
		require.NotEmpty(t,result)

		transfer:=result.Transfer
		require.NotEmpty(t,transfer.ID)
		require.Equal(t,transfer.FromAccountID,account1.ID)
		require.Equal(t,transfer.ToAccountID,account2.ID)
		require.Equal(t,transfer.Amount,amount)
		require.NotEmpty(t,transfer.CreatedAt)

		_,err=store.GetTransfer(context.Background(),transfer.ID)
		require.NoError(t,err) 
 
		fromEntry:=result.FromEntry
		require.NotEmpty(t,fromEntry)
		require.Equal(t,account1.ID,fromEntry.AccountID)
		require.Equal(t,-amount,fromEntry.Amount)
		require.NotEmpty(t,fromEntry.CreatedAt)

		_,err=store.GetEntry(context.Background(),fromEntry.ID)
		require.NoError(t,err)

		toEntry:=result.ToEntry
		require.NotEmpty(t,toEntry)
		require.Equal(t,account2.ID,toEntry.AccountID)
		require.Equal(t,amount,toEntry.Amount)
		require.NotEmpty(t,toEntry.CreatedAt)

		_,err=store.GetEntry(context.Background(),toEntry.ID )
		require.NoError(t,err)

		fromAccount:=result.FromAccount;
		require.NotEmpty(t,fromAccount)
		require.Equal(t,account1.ID,fromAccount.ID)

		toAccount:=result.ToAccount
		require.NotEmpty(t,toAccount)
		require.Equal(t,account2.ID,toAccount.ID)

		//account balance
		diff1:=account1.Balance-fromAccount.Balance
		diff2:=toAccount.Balance-account2.Balance

		require.Equal(t,diff1,diff2)
		require.True(t,diff1>0)
		require.True(t,diff1%amount==0)

		
	}

}

func TestTransferTxDeadLock(t *testing.T){
	store:=NewStore(db)

	account1:=CreateRandomAccount(t)
	account2:=CreateRandomAccount(t)

	n:=10
	amount:=int64(10)

	errs:=make(chan error)

	for i:=0;i<n;i++{
		fromAccount:=account1.ID
		toAccount:=account2.ID

		if i%2==0{
			fromAccount=account2.ID
			toAccount=account1.ID
		}
		go func () {
			_,err:=store.TransferTx(context.Background(),AddTransferParams{
				FromAccountID: fromAccount,
				ToAccountID: toAccount,
				Amount: amount,
			})
	
			errs <-err
		}()
	}
	
	for i:=0;i<n;i++{
		fmt.Print(i)
		err:=<-errs
		require.NoError(t,err)

	}

}