package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nithinps021/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJwtTMaker(t *testing.T) {
	maker,err:=NewJwtMaker(util.RandomString(32))
	require.NoError(t,err)

	username:=util.RandomOwner()
	duration:=time.Minute

	issueAt:=time.Now()
	expiredAt:=issueAt.Add(duration)

	token,err:=maker.CreateToken(username,duration)
	require.NoError(t,err)
	require.NotEmpty(t,token)

	payload,err:=maker.VerifyToken(token)
	require.NoError(t,err)
	require.NotEmpty(t,payload)

	require.NotZero(t,payload.ID)
	require.Equal(t,payload.Username,username)
	require.WithinDuration(t,payload.ExpiredAt,expiredAt,time.Second)
	require.WithinDuration(t,issueAt,payload.IssuedAt,time.Second)
}

func TestJwtExpiredToken(t *testing.T) {
	maker,err:=NewJwtMaker(util.RandomString(32))
	require.NoError(t,err)

	token,err:=maker.CreateToken(util.RandomOwner(),-time.Minute)
	require.NoError(t,err)
	require.NotEmpty(t,token)

	payload,err:=maker.VerifyToken(token)
	require.Nil(t,payload)
	require.Error(t,err)
	require.EqualError(t,err,ErrorExpiredToken.Error())
}

func TestInvalid(t *testing.T) {
	payload,err:=NewPayload(util.RandomOwner(),time.Minute)
	require.NoError(t,err)

	jwtoken:=jwt.NewWithClaims(jwt.SigningMethodNone,payload)
	token,err:=jwtoken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t,err)
	require.NotEmpty(t,token)

	maker,err:=NewJwtMaker(util.RandomString(32))
	require.NoError(t,err)

	payload,err=maker.VerifyToken(token)
	require.Nil(t,payload)
	require.Error(t,err)
	require.EqualError(t,err,ErrorInvalidToken.Error())
}
