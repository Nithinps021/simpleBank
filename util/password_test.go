package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestCreatePassword(t *testing.T) {
	password:=RandomString(6);
	hashedPassword1,err:=HashPassword(password)
	require.NoError(t,err)

	err=CheckPassoword(password,hashedPassword1)
	require.NoError(t,err)

	wrongPassword:=RandomString(7)
	err=CheckPassoword(wrongPassword,hashedPassword1)
	require.EqualError(t,err,bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2,err:=HashPassword(password);
	require.NoError(t,err)
	require.NotEqual(t,hashedPassword1,hashedPassword2)
}
