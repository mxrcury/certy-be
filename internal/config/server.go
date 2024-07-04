package config

func NewServerConfig() (*ServerConfig, error) {
	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}

	domain, err := getEnv("DOMAIN")
	if err != nil {
		return nil, err
	}

	clientURL, err := getEnv("CLIENT_URL")
	if err != nil {
		return nil, err
	}

	return &ServerConfig{Port: port, Domain: domain, ClientURL: clientURL}, nil
}
