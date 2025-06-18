package kafka

import (
	"chat-service/pkg/models"
	"log"
)

func HandleMessage(msg *models.Message) error {
	// Process the message
	// For example, you can log it or send it to a WebSocket hub
	log.Printf("Received message: %s", msg.Content)

	// Here you can add your logic to handle the message
	// For example, sending it to a WebSocket hub or saving it to a database

	return nil // Return nil if processed successfully
}
