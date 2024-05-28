package config

func NewServerConfig() (*ServerConfig, error) {
	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}

	return &ServerConfig{Port: port}, nil
}
