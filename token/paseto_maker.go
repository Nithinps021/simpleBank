package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size it must have %d number of characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error){
	payload, err:= NewPayload(username,duration)
	if err!=nil{
		return "",err
	}
	return maker.paseto.Encrypt(maker.symmetricKey,payload,nil) 
}

func (maker *PasetoMaker) VerifyToken(token string)(*Payload ,error){
	payload := &Payload{};
	
	err:=maker.paseto.Decrypt(token,maker.symmetricKey,payload,nil)
	if err!=nil{
		return nil,ErrorInvalidToken
	}
	err= payload.Valid();
	if err!=nil{
		return nil, ErrorExpiredToken
	}
	return payload,nil
}