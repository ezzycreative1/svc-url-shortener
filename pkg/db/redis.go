package db

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/config"
	"github.com/go-redis/redis/v8"
)

func NewRedis(config *config.Redis) *redis.Client {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		Username: config.Username,
		DB:       config.DB,
	}

	if config.UseTLS {
		opt.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := redis.NewClient(opt)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Panic("REDIS " + err.Error())
	}

	return client
}

// RedisNewClient open redis session with connection pooling, adjustment timeout and custome options
func RedisNewClient(config *config.Redis) (*redis.Client, error) {
	// Redis connection options
	options := &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Username:     config.Username,
		Password:     config.Password,
		DB:           config.DB,
		MaxRetries:   config.MaxRetries,
		MinIdleConns: config.MinIdleConns,
		PoolSize:     config.PoolSize,
		PoolTimeout:  time.Second * time.Duration(config.PoolTimeout),  // Seconds
		MaxConnAge:   time.Second * time.Duration(config.MaxConnAge),   // Seconds
		ReadTimeout:  time.Second * time.Duration(config.ReadTimeout),  // Seconds
		WriteTimeout: time.Second * time.Duration(config.WriteTimeout), // Seconds
	}
	if config.UseTLS {
		options.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Open New Session
	rdb := redis.NewClient(options)

	// Test Connection And Auth with PING
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping command: %w", err)
	}

	return rdb, nil
}
