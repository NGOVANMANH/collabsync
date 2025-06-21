package ws

import (
	"chat-service/internal/kafka"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var hub *Hub
var kafkaService *kafka.KafkaService

func InitHub(kafka *kafka.KafkaService) {
	hub = NewHub()
	kafkaService = kafka
	go hub.Run()
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		ID:   r.RemoteAddr,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  hub,
	}

	// log full client information
	log.Printf("New client connected: %s", client.ID)

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
