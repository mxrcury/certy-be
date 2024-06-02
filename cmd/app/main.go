package main

import (
	"fmt"
	"os"

	"github.com/mxrcury/certy/internal/config"
	"github.com/mxrcury/certy/internal/domain/http"
	v1 "github.com/mxrcury/certy/internal/domain/http/v1"
	"github.com/mxrcury/certy/internal/repository"
	"github.com/mxrcury/certy/internal/service"
	"github.com/mxrcury/certy/pkg/crypto/hash"
	"github.com/mxrcury/certy/pkg/crypto/token"
	"github.com/mxrcury/certy/pkg/database/postgres"
	"github.com/mxrcury/certy/pkg/mail"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	db, err := postgres.Init(cfg.DatabaseConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repos := repository.NewRepositories(db)

	tokenManager := token.NewTokenManager(cfg.AuthConfig.AccessTokenSecretKey)
	hasher := hash.NewHasher(cfg.AuthConfig.PasswordSalt)
	smtpSender := mail.NewSMTPSender(
		cfg.SMTPConfig.Username,
		cfg.SMTPConfig.Password,
		cfg.SMTPConfig.From,
		cfg.SMTPConfig.Host,
		cfg.SMTPConfig.Port,
	)

	services := service.NewServices(&service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
		Hasher:       hasher,
		SMTPSender:   smtpSender,
	})

	server := http.NewServer(cfg.ServerConfig)

	v1.InitHandlers(&v1.Deps{Router: server.Router, Services: services})

	if err := server.Run(); err != nil {
		// add graceful shutdown
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
