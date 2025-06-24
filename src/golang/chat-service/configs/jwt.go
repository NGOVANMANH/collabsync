package configs

import "os"

type JWTConfig struct {
	SecretKey string
}

func GetJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET"),
	}
}
