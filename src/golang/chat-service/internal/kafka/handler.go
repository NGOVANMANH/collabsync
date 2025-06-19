package kafka

import (
	"chat-service/pkg/models"
	"log"
)

func handleMessage(msg *models.Message) error {
	log.Printf("Received message: %s from sender: %s in room: %s", msg.Content, msg.SenderID, msg.RoomID)

	_, err := messageService.CreateMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
