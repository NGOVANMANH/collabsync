package ws

import (
	"chat-service/pkg/models"
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
	Hub  *Hub
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

		msg := &models.Message{}
		if err := json.Unmarshal(message, msg); err != nil {
			continue // Ignore malformed messages
		}

		// ✅ Produce to Kafka
		if err := kafkaService.Produce(context.Background(), msg); err != nil {
			// Optional: log or retry
			continue
		}

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
