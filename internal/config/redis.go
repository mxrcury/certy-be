package config

import "strconv"

func NewRedisConfig() (*RedisConfig, error) {
	port, err := getEnv("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	host, err := getEnv("REDIS_HOST")
	if err != nil {
		return nil, err
	}

	db, err := getEnv("REDIS_DB")
	if err != nil {
		return nil, err
	}
	parsedDB, err := strconv.ParseInt(db, 10, 64)
	if err != nil {
		return nil, err
	}

	pass, err := getEnv("REDIS_PASSWORD")
	if err != nil {
		return nil, err
	}

	return &RedisConfig{
		Pass: pass,
		DB:   int(parsedDB),
		Host: host,
		Port: port,
	}, nil
}
