package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
) 


func TestGetTransfer(t *testing.T){
	account1:=CreateRandomAccount(t)
	account2:=CreateRandomAccount(t)

	amount:=int64(10)

	transfer,err:=testQueries.AddTransfer(context.Background(),AddTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: amount,
	})
	require.NoError(t,err)
	getTransfer,err:=testQueries.GetTransfer(context.Background(),transfer.ID)
	require.NoError(t,err)
	require.NotEmpty(t,getTransfer)
	require.NotZero(t,getTransfer.ID)
	require.NotZero(t,getTransfer.CreatedAt)
	require.Equal(t,getTransfer.Amount,amount)
	require.Equal(t,getTransfer.FromAccountID,account1.ID)
	require.Equal(t,getTransfer.ToAccountID,account2.ID)
}


func TestAddTransfer(t *testing.T) {
	account1:=CreateRandomAccount(t)
	account2:=CreateRandomAccount(t)

	amount:=int64(10)

	transfer,err:=testQueries.AddTransfer(context.Background(),AddTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: amount,
	})
	require.NoError(t,err)
	require.NotEmpty(t,transfer)
	require.NotZero(t,transfer.ID)
	require.NotZero(t,transfer.CreatedAt)
	require.Equal(t,transfer.Amount,amount)
	require.Equal(t,transfer.FromAccountID,account1.ID)
	require.Equal(t,transfer.ToAccountID,account2.ID)
}

func TestListTransfer(t *testing.T){
	account1:=CreateRandomAccount(t)
	account2:=CreateRandomAccount(t)
	amount:=int64(4)
	n:=5
	for i:=0;i<n;i++{
		testQueries.AddTransfer(context.Background(),AddTransferParams{
			FromAccountID: account1.ID,
			ToAccountID: account2.ID,
			Amount: amount,
		})
	}
	transfers,err:=testQueries.ListTransfers(context.Background(),ListTransfersParams{
		Limit: 3,
		Offset: 2,
	})
	require.NoError(t,err)
	require.Len(t,transfers,3)
	for _,transfer:= range transfers{
		require.NotEmpty(t,transfer)
	}
}