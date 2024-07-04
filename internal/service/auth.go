package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mxrcury/certy/internal/repository"
	"github.com/mxrcury/certy/pkg/cache"
	"github.com/mxrcury/certy/pkg/crypto/hash"
	"github.com/mxrcury/certy/pkg/crypto/token"
	"github.com/mxrcury/certy/pkg/mail"
)

type AuthService struct {
	repo         repository.Users
	hasher       *hash.Hasher
	SMTPSender   *mail.SMTPSender
	TokenManager *token.TokenManager
	CacheClient  *cache.Client
}

type AuthServiceDeps struct {
	repo         repository.Users
	hasher       *hash.Hasher
	SMTPSender   *mail.SMTPSender
	TokenManager *token.TokenManager
	CacheClient  *cache.Client
}

type SignUpInput struct {
	FirstName string `json:"firstName" binding:"required,min=4,max=14,alphanum"`
	LastName  string `json:"lastName" binding:"omitempty,min=4,max=14,alphanum"`
	Username  string `json:"username" binding:"required,min=4,max=14,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=5"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4"`
}

func NewAuthService(deps *AuthServiceDeps) Auth {
	return &AuthService{
		repo:         deps.repo,
		hasher:       deps.hasher,
		SMTPSender:   deps.SMTPSender,
		TokenManager: deps.TokenManager,
		CacheClient:  deps.CacheClient,
	}
}

func (s *AuthService) SignUp(input *SignUpInput) error {
	isExistingUser := s.repo.GetByEmailOrUsername(input.Email, input.Username)

	if isExistingUser != nil {
		return errors.New("user with this email or username already exists")
	}

	hashedPassword := s.hasher.Hash([]byte(input.Password))
	user := &repository.User{
		ID:        uuid.New(),
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName: sql.NullString{
			String: input.LastName,
			Valid:  input.LastName != "",
		},
		Email:     input.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err := s.repo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(input *SignInInput) (*repository.User, error) {
	isExistingUser := s.repo.GetByEmail(input.Email)
	if isExistingUser == nil {
		return nil, errors.New("user with this email does not exist")
	}

	isValidPassword := s.hasher.Verify([]byte(input.Password), isExistingUser.Password)
	if !isValidPassword {
		return nil, errors.New("you entered wrong password")
	}

	return isExistingUser, nil
}

func (s *AuthService) SendVerificationCode(email string) error {
	verificationCode, err := s.TokenManager.GenerateCode(6)

	if err != nil {
		return err
	}

	input := &mail.SendEmailInput{
		To:      email,
		Subject: "Please verify your email",
		Content: fmt.Sprintf("<p>Hi! Your verification code is <b>%s</b></p>\n", verificationCode),
	}

	ttl := time.Minute * 5
	if err := s.CacheClient.Set(verificationCode, verificationCode, ttl); err != nil {
		return err
	}

	if err := s.SMTPSender.SendEmail(input); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) VerifyCode(code string) error {
	_, err := s.CacheClient.Get(code)

	if err != nil {
		return err
	}

	return nil
}
