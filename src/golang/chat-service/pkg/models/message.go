package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	SenderID  string    `json:"sender_id"`
	RoomID    string    `json:"room_id"`
	Timestamp int64     `json:"timestamp"`
}

func NewMessage(content, senderID, roomID string) *Message {
	return &Message{
		ID:        uuid.New(),
		Content:   content,
		SenderID:  senderID,
		RoomID:    roomID,
		Timestamp: getCurrentTimestamp(),
	}
}

func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
