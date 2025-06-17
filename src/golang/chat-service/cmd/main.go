package main

import (
	"chat-service/internal/redis"
	"chat-service/internal/ws"
	"log"
	"net/http"
)

func main() {
	redis.Init()

	ws.InitHub()
	http.HandleFunc("/ws", ws.HandleWS)

	log.Println("✅ WebSocket server started at :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
