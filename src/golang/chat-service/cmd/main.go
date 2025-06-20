package main

import (
	"chat-service/configs"
	"chat-service/internal/db"
	"chat-service/internal/kafka"
	"chat-service/internal/message"
	"chat-service/internal/redis"
	"chat-service/internal/ws"
	"log"
	"net/http"
)

func main() {
	redis.Init()

	// Database connection
	db, err := db.Init()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Create repository, service
	repo := message.NewMessageRepository(db)
	messageService := message.NewMessageService(repo)

	// Initialize Kafka
	config := &configs.KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "chat-messages",
		GroupID: "chat-service",
	}

	kafkaService, err := kafka.NewKafkaService(config, messageService)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaService.Close()

	go func() {
		err := kafkaService.StartConsume()
		if err != nil {
			log.Printf("Consumer error: %v", err)
		}
	}()

	ws.InitHub(kafkaService)
	http.HandleFunc("/ws", ws.HandleWS)

	log.Println("✅ WebSocket server started at :8081")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
