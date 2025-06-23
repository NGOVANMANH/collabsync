package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin for development
		// In production, implement proper origin checking
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Hub) HandleWs(w http.ResponseWriter, r *http.Request) {
	// Get client ID from query params or generate one
	clientID := r.URL.Query().Get("id")
	if clientID == "" {
		http.Error(w, "Missing client ID", http.StatusBadRequest)
		return
	}

	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	// Create new client
	client := &Client{
		ID:   clientID,
		Hub:  h,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	// Register client with hub
	h.Register <- client

	// Start goroutines for reading and writing
	go client.WriteMessages()
	go client.ReadMessages()

	log.Printf("Client %s connected", clientID)
}
