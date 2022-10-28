package config

import (
	"log"

	"github.com/ezzycreative1/svc-url-shortener/pkg/envar"
	"github.com/joho/godotenv"
)

type Group struct {
	Server Server `json:"server,omitempty"`
	//Database Database `json:"database,omitempty"`
	Redis        Redis              `json:"redis,omitempty"`
	URLShortener URLShortenerConfig `json:"urlshortener,omitempty"`
}

type Server struct {
	ENV string `json:"env"`
}

func LoadConfig() *Group {
	if err := godotenv.Load(); err != nil {
		// in prod we will not use this,use os env instead
		log.Print(".env notfound")
	}

	env = envar.GetEnv("ENV", "dev")

	return &Group{
		Server: Server{
			ENV: env,
		},
		// Database: Database{
		// 	Engine:          envar.GetEnv("DATABASE_ENGINE", "mysqli"),
		// 	Host:            envar.GetEnv("DATABASE_HOST", "localhost"),
		// 	Port:            envar.GetEnv("DATABASE_PORT", 3306),
		// 	Username:        envar.GetEnv("DATABASE_USERNAME", "root"),
		// 	Password:        envar.GetEnv("DATABASE_PASSWORD", ""),
		// 	Schema:          envar.GetEnv("DATABASE_SCHEMA", "inventory"),
		// 	MaxIdle:         envar.GetEnv("DATABASE_MAX_IDLE", 20),
		// 	MaxConn:         envar.GetEnv("DATABASE_MAX_CONN", 100),
		// 	ConnMaxLifetime: envar.GetEnv("DATABASE_CONN_LIFETIME", 180),
		// 	Environment:     env,
		// },
		Redis: Redis{
			Host:     envar.GetEnv("REDIS_HOST", "localhost"),
			Port:     envar.GetEnv("REDIS_PORT", 6379),
			Username: envar.GetEnv("REDIS_USERNAME", ""),
			Password: envar.GetEnv("REDIS_PASSWORD", ""),
			DB:       envar.GetEnv("REDIS_DB", 4),
			UseTLS:   envar.GetEnv("REDIS_USE_TLS", false),
		},
	}
}
