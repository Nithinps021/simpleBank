package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/nithinps021/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomEntry(t *testing.T) Entry{
	account1:=CreateRandomAccount(t);
	args:=AddEntryParams{
		AccountID: account1.ID,
		Amount: util.RandomMoney(),
	}
	entry1,err:=testQueries.AddEntry(context.Background(),args)
	require.NoError(t,err)
	require.Equal(t,entry1.AccountID,account1.ID)
	require.Equal(t,args.Amount,entry1.Amount)
	require.NotZero(t,entry1.ID)
	require.NotZero(t,entry1.CreatedAt)
	return entry1;
}

func TestAddEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T){
	entry1:=createRandomEntry(t)
	entry2,err:=testQueries.GetEntry(context.Background(),entry1.ID)
	require.NoError(t,err)
	require.NotEmpty(t,entry2)
	require.Equal(t,entry1.ID,entry2.ID)
	require.Equal(t,entry1.AccountID,entry2.AccountID)
	require.Equal(t,entry1.Amount,entry2.Amount)
	require.WithinDuration(t,entry1.CreatedAt,entry2.CreatedAt,time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1:=createRandomEntry(t)
	err:=testQueries.DeleteEntry(context.Background(),entry1.ID)
	require.NoError(t,err)
	entry2,err:=testQueries.GetEntry(context.Background(),entry1.ID)
	require.Empty(t,entry2)
	require.EqualError(t,err,sql.ErrNoRows.Error())
}

func TestListEntries(t *testing.T){
	for i:=0;i<7;i++{
		createRandomEntry(t)
	}
	args:=ListEntriesParams{
		Limit:4 ,
		Offset: 2,
	}
	entries,err:=testQueries.ListEntries(context.Background(),args)
	require.NoError(t,err)
	require.Len(t,entries,4)
	for _, entry := range entries{
		require.NotEmpty(t,entry)
	}
}