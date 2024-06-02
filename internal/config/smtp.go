package config

import (
	"errors"
	"strconv"
)

func NewSMTPConfig() (*SMTPConfig, error) {
	port, err := getEnv("SMTP_PORT")
	if err != nil {
		return nil, err
	}

	parsedPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil, errors.New("error parsing port")
	}

	host, err := getEnv("SMTP_HOST")
	if err != nil {
		return nil, err
	}

	from, err := getEnv("SMTP_FROM")
	if err != nil {
		return nil, err
	}

	pass, err := getEnv("SMTP_PASS")
	if err != nil {
		return nil, err
	}

	username, err := getEnv("SMTP_USERNAME")
	if err != nil {
		return nil, err
	}

	return &SMTPConfig{
		Port:     int(parsedPort),
		Password: pass,
		From:     from,
		Host:     host,
		Username: username,
	}, nil
}
