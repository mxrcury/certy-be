package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mxrcury/certy/internal/repository"
	"github.com/mxrcury/certy/pkg/crypto/hash"
)

type AuthService struct {
	repo   repository.Users
	hasher *hash.Hasher
}

type AuthServiceDeps struct {
	repo   repository.Users
	hasher *hash.Hasher
}

type SignUpInput struct {
	Username string `json:"username" binding:"required,min=4,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required,min=4,alphanum"`
	Password string `json:"password" binding:"required,min=5"`
}

func NewAuthService(deps *AuthServiceDeps) Auth {
	return &AuthService{repo: deps.repo, hasher: deps.hasher}
}

func (s *AuthService) SignUp(input *SignUpInput) error {
	isExistingUser := s.repo.GetByEmailOrUsername(input.Email, input.Username)

	if isExistingUser != nil {
		return errors.New("user with this email already exists")
	}

	hashedPassword := s.hasher.Hash([]byte(input.Password))
	user := &repository.User{
		ID:        uuid.New(),
		Username:  input.Username,
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format(time.RFC3339Nano),
	}

	err := s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(input *SignInInput) (*repository.User, error) {
	isExistingUser := s.repo.GetByUsername(input.Username)
	if isExistingUser == nil {
		return nil, errors.New("user with this username does not exist")
	}

	isValidPassword := s.hasher.Verify([]byte(input.Password), isExistingUser.Password)
	if !isValidPassword {
		return nil, errors.New("you entered wrong password")
	}

	return isExistingUser, nil
}
