package configs

import (
	"os"
	"strconv"
)

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Database int
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Host: os.Getenv("REDIS_HOST"),
		Port: func() int {
			port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
			if err != nil {
				return 6379
			}
			return port
		}(),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: func() int {
			db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
			if err != nil {
				return 0
			}
			return db
		}(),
	}
}
