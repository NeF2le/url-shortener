package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"redis"`
	Port     uint16 `yaml:"port" env:"PORT" env-default:"6379"`
	Password string `yaml:"password" env:"PASSWORD"`
	Username string `yaml:"user" env:"USER"`

	PoolSize int `yaml:"pool_size" env:"POOL_SIZE" env-default:"10"`
}

func NewRedisClient(ctx context.Context, cfg *Config, dbNum int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		Username: cfg.Username,
		DB:       dbNum,
		PoolSize: cfg.PoolSize,
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
