package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)



func HashPassword(password string) (string, error){
	hashpassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "",fmt.Errorf("not able to hash password %w",err)
	}
	return string(hashpassword),nil
}

func CheckPassoword(password string, hashPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password))
} 