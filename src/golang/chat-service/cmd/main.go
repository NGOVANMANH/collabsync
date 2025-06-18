package main

import (
	"chat-service/configs"
	"chat-service/internal/kafka"
	"chat-service/internal/redis"
	"chat-service/internal/ws"
	"context"
	"log"
	"net/http"
)

func main() {
	redis.Init()

	// Initialize
	config := &configs.KafkaConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "chat-messages",
		GroupID: "chat-service",
	}

	kafkaService, err := kafka.NewKafkaService(config)
	if err != nil {
		log.Fatal(err)
	}
	defer kafkaService.Close()

	// Start consuming
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := kafkaService.StartConsume(ctx, kafka.HandleMessage)
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
