package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var TokenKey = []byte("canwinter")

type TokenClaims struct {
	UUID string `json:"uuid"`
	Name string
	jwt.RegisteredClaims
}

func GenerateUserToken(uuid string, name string) (string, error) {
	claims := TokenClaims{
		UUID: uuid,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 90)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(TokenKey)
}
