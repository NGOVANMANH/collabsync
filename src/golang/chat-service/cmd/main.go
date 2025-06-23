package main

import (
	"chat-service/internal/kafka"
	"chat-service/internal/redis"
	"chat-service/internal/ws"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redis.Init()
	defer redis.Close()

	kafka.Init()
	defer kafka.Close()

	kafka.StartConsumer()

	hub := ws.InitHub()

	go hub.Run()
	http.HandleFunc("/ws", hub.HandleWs)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Serve static files for testing
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	log.Println("WebSocket server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
