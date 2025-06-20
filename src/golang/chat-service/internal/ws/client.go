package ws

import (
	"chat-service/pkg/models"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID         string
	Conn       *websocket.Conn
	Send       chan []byte
	Hub        *Hub
	ClientInfo *models.ClientInfo
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		go func() {
			msg := &models.MessageRequest{}
			if err := json.Unmarshal(message, msg); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				return
			}

			// ✅ Produce to Kafka
			if err := kafkaService.Produce(msg.ToMessage()); err != nil {
				// Optional: log or retry
				log.Printf("Failed to produce message to Kafka: %v", err)
			}
		}()

		// ✅ Broadcast to other WebSocket clients
		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for {
		msg, ok := <-c.Send
		if !ok {
			return
		}
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (c *Client) SendMessage(msg *models.Message, toClientIds []string) {
	message, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}
}
