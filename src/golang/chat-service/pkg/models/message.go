package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	Content     string      `json:"content" db:"content"`
	SenderID    string      `json:"sender_id" db:"sender_id"`
	RoomID      string      `json:"room_id" db:"room_id"`
	Attachments []uuid.UUID `json:"attachments,omitempty" db:"attachments"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}

func NewMessage(content, senderID, roomID string, attachments []uuid.UUID) *Message {
	return &Message{
		ID:          uuid.New(),
		Content:     content,
		SenderID:    senderID,
		RoomID:      roomID,
		Attachments: attachments,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type MessageRequest struct {
	Content     string      `json:"content"`
	SenderID    string      `json:"sender_id"`
	RoomID      string      `json:"room_id"`
	Attachments []uuid.UUID `json:"attachments,omitempty"`
}

func (mr *MessageRequest) ToMessage() *Message {
	return NewMessage(mr.Content, mr.SenderID, mr.RoomID, mr.Attachments)
}
