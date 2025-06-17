package message

import "github.com/google/uuid"

// Define the interface
type IMessageRepository interface {
	GetMessage() []byte
	GetMessageByID(id uuid.UUID) []byte
}

// Define a struct that implements the interface
type MessageRepository struct{}

// Implement GetMessage
func (r *MessageRepository) GetMessage() []byte {
	return []byte("Hello, World!")
}

// Implement GetMessageByID
func (r *MessageRepository) GetMessageByID(id uuid.UUID) []byte {
	return []byte("Hello, World! " + id.String())
}
