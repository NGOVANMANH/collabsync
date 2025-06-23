package ws

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer
	maxMessageSize = 512
)

// Message types
const (
	MessageTypeChat     = "chat"
	MessageTypePing     = "ping"
	MessageTypePong     = "pong"
	MessageTypeUserList = "user_list"
	MessageTypeError    = "error"
)

// Enhanced envelope structure
type envelope struct {
	Type      string   `json:"type"`           // message type
	To        []string `json:"to,omitempty"`   // target client IDs (empty for broadcast)
	From      string   `json:"from,omitempty"` // sender ID
	Body      string   `json:"body"`           // the message text
	Timestamp int64    `json:"timestamp"`      // unix timestamp
}

type Client struct {
	ID   string
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadMessages() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	// Set read deadline and message size limit
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, frame, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WebSocket error for client %s: %v", c.ID, err)
			}
			break
		}

		// Parse incoming JSON
		var env envelope
		if err := json.Unmarshal(frame, &env); err != nil {
			log.Printf("Invalid envelope from client %s: %v", c.ID, err)
			c.sendError("Invalid message format")
			continue
		}

		// Set sender and timestamp
		env.From = c.ID
		env.Timestamp = time.Now().Unix()

		// Handle different message types
		switch env.Type {
		case MessageTypeChat:
			c.handleChatMessage(env)
		case MessageTypePing:
			c.sendPong()
		case MessageTypeUserList:
			c.sendUserList()
		default:
			c.sendError("Unknown message type")
		}
	}
}

func (c *Client) WriteMessages() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Write error for client %s: %v", c.ID, err)
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) handleChatMessage(env envelope) {
	// Re-marshal the message to send to other clients
	messageBytes, err := json.Marshal(env)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}

	// Send to specific clients or broadcast
	if len(env.To) > 0 {
		c.Hub.SendToClients(messageBytes, env.To)
	} else {
		c.Hub.BroadcastToAll(messageBytes)
	}
}

func (c *Client) sendError(errorMsg string) {
	env := envelope{
		Type:      MessageTypeError,
		Body:      errorMsg,
		Timestamp: time.Now().Unix(),
	}

	if msg, err := json.Marshal(env); err == nil {
		select {
		case c.Send <- msg:
		default:
			// Channel full, ignore
		}
	}
}

func (c *Client) sendPong() {
	env := envelope{
		Type:      MessageTypePong,
		Timestamp: time.Now().Unix(),
	}

	if msg, err := json.Marshal(env); err == nil {
		select {
		case c.Send <- msg:
		default:
			// Channel full, ignore
		}
	}
}

func (c *Client) sendUserList() {
	clients := c.Hub.GetConnectedClients()
	clientsJSON, _ := json.Marshal(clients)

	env := envelope{
		Type:      MessageTypeUserList,
		Body:      string(clientsJSON),
		Timestamp: time.Now().Unix(),
	}

	if msg, err := json.Marshal(env); err == nil {
		select {
		case c.Send <- msg:
		default:
			// Channel full, ignore
		}
	}
}
