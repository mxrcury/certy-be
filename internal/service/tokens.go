package service

import "github.com/mxrcury/dragonsage/pkg/crypto/token"

type JWTTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"-"`
}

type TokensService struct {
	TokenManager *token.TokenManager
}

func NewTokensService(tokenManager *token.TokenManager) Tokens {
	return &TokensService{TokenManager: tokenManager}
}

func (t *TokensService) GenerateJWT(userId string) (string, error) {
	return t.TokenManager.GenerateJWT(userId)
}
