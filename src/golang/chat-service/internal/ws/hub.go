package ws

import (
	"chat-service/internal/kafka"
	"log"
	"sync"
)

type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte  // For broadcasting to all clients
	mu         sync.RWMutex // For thread-safe operations
}

func InitHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client.ID] = client
			h.mu.Unlock()
			if err := kafka.Produce([]byte("ws:client:connected"), []byte(client.ID)); err != nil {
				log.Printf("Failed to produce message for client %s: %v", client.ID, err)
			}
			log.Printf("Client %s registered. Total clients: %d", client.ID, len(h.Clients))

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client.ID]; ok {
				close(client.Send)
				delete(h.Clients, client.ID)
				log.Printf("Client %s unregistered. Total clients: %d", client.ID, len(h.Clients))
			}
			h.mu.Unlock()

		case message := <-h.Broadcast:
			h.mu.RLock()
			for _, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client.ID)
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) SendToClients(message []byte, clientIds []string) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, id := range clientIds {
		if client, ok := h.Clients[id]; ok {
			select {
			case client.Send <- message:
			default:
				// Client's send channel is full, close it
				close(client.Send)
				delete(h.Clients, id)
				log.Printf("Client %s removed due to full send buffer", id)
			}
		}
	}
}

func (h *Hub) BroadcastToAll(message []byte) {
	h.Broadcast <- message
}

func (h *Hub) GetConnectedClients() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	clients := make([]string, 0, len(h.Clients))
	for id := range h.Clients {
		clients = append(clients, id)
	}
	return clients
}

func (h *Hub) IsClientConnected(clientID string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	_, exists := h.Clients[clientID]
	return exists
}
