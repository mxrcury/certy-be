package config

func NewAuthConfig() (*AuthConfig, error) {
	accessTokenSecretKey, err := getEnv("ACCESS_TOKEN_SECRET_KEY")
	if err != nil {
		return nil, err
	}

	passwordSalt, err := getEnv("PASSWORD_SALT")
	if err != nil {
		return nil, err
	}

	return &AuthConfig{
		AccessTokenSecretKey: accessTokenSecretKey,
		PasswordSalt:         passwordSalt,
	}, nil
}
