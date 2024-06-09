package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, val interface{}, ttl time.Duration) error
}

type Client struct {
	client *redis.Client

	ctx context.Context
}

type ClientConfig struct {
	Host string
	Pass string
	Port string
	DB   int
}

func NewClient(config *ClientConfig) *Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Pass,
		DB:       config.DB,
	})
	ctx := context.Background()

	return &Client{client: client, ctx: ctx}
}

func (c *Client) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *Client) Set(key string, val interface{}, ttl time.Duration) error {
	return c.client.Set(c.ctx, key, val, ttl).Err()
}
