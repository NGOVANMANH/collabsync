package redis

import (
	"chat-service/configs"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func Init(conf *configs.RedisConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%d",
			os.Getenv("CHAT_REDIS_HOST"),
			func() int {
				port := os.Getenv("CHAT_REDIS_PORT")
				if port == "" {
					return 6379
				}
				p, err := strconv.Atoi(port)
				if err != nil {
					log.Fatalf("❌ Invalid Redis port: %v", err)
				}
				return p
			}()),
		Password: os.Getenv("CHAT_REDIS_PASSWORD"),
		DB: func() int {
			db := os.Getenv("CHAT_REDIS_DB")
			if db == "" {
				return 0 // default Redis DB
				// return 0
			}
			dbNum, err := strconv.Atoi(db)
			if err != nil {
				log.Fatalf("❌ Invalid Redis DB number: %v", err)
			}
			return dbNum
		}(),
	})

	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}
	log.Println("✅ Redis connection established successfully")
}

func Close() {
	if err := Client.Close(); err != nil {
		log.Fatalf("❌ Failed to close Redis client: %v", err)
	}
	log.Println("✅ Redis client closed successfully")
}
