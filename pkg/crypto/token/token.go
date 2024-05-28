package token

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	SecretKey string
}

func NewTokenManager(secretKey string) *TokenManager {
	return &TokenManager{secretKey}
}

func (t *TokenManager) GenerateJWT(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 6)),
		Subject:   userId,
	})

	return token.SignedString([]byte(t.SecretKey))
}
