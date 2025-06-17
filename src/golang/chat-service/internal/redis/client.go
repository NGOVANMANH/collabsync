package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123123123", // if none
		DB:       0,
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}
	log.Println("✅ Redis connection established successfully")
}
