package token

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
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

func (t *TokenManager) GenerateCode(length int) (string, error) {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(length)))),
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%0*d", length, bi), nil
}
