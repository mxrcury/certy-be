package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		ServerConfig   *ServerConfig
		DatabaseConfig *DatabaseConfig
		AuthConfig     *AuthConfig
		SMTPConfig     *SMTPConfig
	}

	ServerConfig struct {
		Port string
	}

	DatabaseConfig struct {
		DataSourceName string
		MigrationsDir  string
		DatabaseDriver string
	}

	AuthConfig struct {
		AccessTokenSecretKey string
		PasswordSalt         string
	}

	SMTPConfig struct {
		Password string
		Port     int
		From     string
		Host     string
		Username string
	}
)

func Init() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	databaseConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	serverConfig, err := NewServerConfig()
	if err != nil {
		return nil, err
	}

	authConfig, err := NewAuthConfig()
	if err != nil {
		return nil, err
	}

	smtpConfig, err := NewSMTPConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		DatabaseConfig: databaseConfig,
		ServerConfig:   serverConfig,
		AuthConfig:     authConfig,
		SMTPConfig:     smtpConfig,
	}, nil
}

func getEnv(key string) (string, error) {
	env := os.Getenv(key)
	if env == "" {
		return "", fmt.Errorf(`variable "%s" was not found `, key)
	}

	return env, nil
}
