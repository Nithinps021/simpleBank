package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrorInvalidToken=fmt.Errorf("token is invalid")
	ErrorExpiredToken = fmt.Errorf("token has expired")
)
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration)(*Payload,error){
	tokenId,err:=uuid.NewRandom()
	if err!=nil{
		return  nil,err
	}
	paylod:= &Payload{
		ID: tokenId,
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return paylod, nil;
}

func (payload *Payload) Valid () error{
	if time.Now().After(payload.ExpiredAt){
		return ErrorExpiredToken
	}
	return nil
}