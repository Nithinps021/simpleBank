package db

import (
	"context"
	"testing"
	"time"

	"github.com/nithinps021/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		HashPassword: "secret",
		FullName: util.RandomOwner(),
		EmialID: util.RandomEmail(),

	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashPassword, user.HashPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotZero(t, user.PasswordLastChanged)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {  
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashPassword, user2.HashPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t,user1.EmialID,user2.EmialID)
	require.WithinDuration(t, user1.PasswordLastChanged, user2.PasswordLastChanged, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
