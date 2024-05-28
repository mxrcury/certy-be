package service

import (
	"github.com/mxrcury/certy/internal/repository"
	"github.com/mxrcury/certy/pkg/crypto/hash"
	"github.com/mxrcury/certy/pkg/crypto/token"
)

type Services struct {
	AuthService   Auth
	TokensService Tokens
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager *token.TokenManager
	Hasher       *hash.Hasher
}

type (
	Auth interface {
		SignUp(*SignUpInput) error
		SignIn(*SignInInput) (*repository.User, error)
	}

	Tokens interface {
		GenerateJWT(userId string) (string, error)
	}
)

func NewServices(deps *Deps) *Services {
	tokensService := NewTokensService(deps.TokenManager)

	authService := NewAuthService(&AuthServiceDeps{
		repo:   deps.Repos.UsersRepository,
		hasher: deps.Hasher,
	})

	return &Services{AuthService: authService, TokensService: tokensService}
}
