package message

import (
	"github.com/google/uuid"
)

// IMessageService defines the interface for the message service
type IMessageService interface {
	GetWelcomeMessage() []byte
	GetPersonalizedMessage(id uuid.UUID) []byte
}

// MessageService is the concrete implementation
type MessageService struct {
	repo IMessageRepository
}

// NewMessageService is a constructor function for MessageService
func NewMessageService(repo IMessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

// GetWelcomeMessage returns a basic message
func (s *MessageService) GetWelcomeMessage() []byte {
	return s.repo.GetMessage()
}

// GetPersonalizedMessage returns a message with a UUID
func (s *MessageService) GetPersonalizedMessage(id uuid.UUID) []byte {
	return s.repo.GetMessageByID(id)
}
