package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


const secretKey_len = 32;

type JWTMaker struct{
	secretkey string 
}

func NewJwtMaker(secretKey string)(Maker,error){
	if len(secretKey)< secretKey_len{
		return nil, fmt.Errorf("secret key should be atleast %d character",secretKey_len)
	}
	return &(JWTMaker{secretkey: secretKey}), nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration)(string,error){
	payload,err:=NewPayload(username,duration)
	if err!=nil{
		return "",err
	}
	jwtToken:=jwt.NewWithClaims(jwt.SigningMethodHS256,payload)
	return jwtToken.SignedString([]byte(maker.secretkey))
}

func (maker *JWTMaker) VerifyToken(token string)(*Payload ,error){
	keyFunc := func (token *jwt.Token)(interface{},error)  {
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretkey),nil
	}
	jwtToken,err:=jwt.ParseWithClaims(token,&Payload{},keyFunc)
	if err!=nil{
		verr,ok:=err.(*jwt.ValidationError)
		if ok && verr.Is(ErrorExpiredToken){
			return nil,ErrorExpiredToken
		}
		return nil,ErrorInvalidToken
	}
	payload,ok:=jwtToken.Claims.(*Payload);
	if !ok{
		return nil,ErrorInvalidToken
	}
	return payload,nil
}