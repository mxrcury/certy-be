package main

import (
	"fmt"
	"os"

	"github.com/mxrcury/dragonsage/internal/config"
	"github.com/mxrcury/dragonsage/internal/domain/http"
	v1 "github.com/mxrcury/dragonsage/internal/domain/http/v1"
	"github.com/mxrcury/dragonsage/internal/repository"
	"github.com/mxrcury/dragonsage/internal/service"
	"github.com/mxrcury/dragonsage/pkg/crypto/hash"
	"github.com/mxrcury/dragonsage/pkg/crypto/token"
	"github.com/mxrcury/dragonsage/pkg/database/postgres"
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

	services := service.NewServices(&service.Deps{
		Repos:        repos,
		TokenManager: tokenManager,
		Hasher:       hasher,
	})

	server := http.NewServer(cfg.ServerConfig)

	v1.InitHandlers(&v1.Deps{Router: server.Router, Services: services})

	if err := server.Run(); err != nil {
		// add graceful shutdown
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
