package configs

import "os"

type ServiceConfig struct {
	ServiceToken string
}

func GetServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		ServiceToken: os.Getenv("SERVICE_TOKEN"),
	}
}
