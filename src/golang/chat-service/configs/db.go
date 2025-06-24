package configs

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func GetDbConfig() *DbConfig {
	port, err := strconv.Atoi(os.Getenv("CHAT_DB_HOST"))
	if err != nil {
		port = 5432
	}
	return &DbConfig{
		Host:     os.Getenv("CHAT_DB_HOST"),
		Port:     port,
		Username: os.Getenv("CHAT_DB_USER"),
		Password: os.Getenv("CHAT_DB_PASSWORD"),
		Database: os.Getenv("CHAT_DB_NAME"),
	}
}
