package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	AccountID string `json:"accountId"`
	jwt.RegisteredClaims
}

func GenerateToken(accountId string, publicKey string, ttl int) (string, error) {
	secret := []byte(publicKey)

	claims := Claims{
		accountId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(ttl) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	return tokenString, err
}
