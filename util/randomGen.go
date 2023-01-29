package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)


const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init(){
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min , max int64) int64{
	return min + rand.Int63n(max-min+1);
}

func RandomString(length int) string{
	var sb strings.Builder;
	n := len(alphabets)
	for i:=0; i<length ; i++ {
		c := alphabets[rand.Intn(n)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string{
	return RandomString(6);
}

func RandomMoney() int64{
	return RandInt(0,1000);
}

func RandomCurrency() string {
	currency:=[]string{"EUR","USD","CAD"}
	return currency[rand.Intn(3)]
}

func RandomEmail() string{
	return fmt.Sprintf("%s@gmial.com",RandomString(6))
}

