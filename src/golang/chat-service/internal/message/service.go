package message

import (
	"chat-service/pkg/models"

	"github.com/google/uuid"
)

// IMessageService defines the interface for the message service
type IMessageService interface {
	GetMessages() ([]*models.Message, error)
	GetMessageByID(id uuid.UUID) (*models.Message, error)
	CreateMessage(msg *models.Message) (*models.Message, error)
	UpdateMessage(id uuid.UUID, msg *models.Message) error
	DeleteMessage(id uuid.UUID) error
	GetMessagesByRoomID(roomID string) ([]*models.Message, error)
}

// MessageService is the concrete implementation
type MessageService struct {
	repo IMessageRepository
}

// NewMessageService is a constructor function for MessageService
func NewMessageService(repo IMessageRepository) IMessageService {
	return &MessageService{repo: repo}
}

// GetMessages retrieves all messages
func (s *MessageService) GetMessages() ([]*models.Message, error) {
	return s.repo.GetMessages()
}

// GetMessageByID retrieves a message by its ID
func (s *MessageService) GetMessageByID(id uuid.UUID) (*models.Message, error) {
	return s.repo.GetMessageByID(id)
}

// CreateMessage creates a new message
func (s *MessageService) CreateMessage(msg *models.Message) (*models.Message, error) {
	return s.repo.CreateMessage(msg)
}

// UpdateMessage updates an existing message
func (s *MessageService) UpdateMessage(id uuid.UUID, msg *models.Message) error {
	return s.repo.UpdateMessage(id, msg)
}

// DeleteMessage deletes a message by its ID
func (s *MessageService) DeleteMessage(id uuid.UUID) error {
	return s.repo.DeleteMessage(id)
}

// GetMessagesByRoomID retrieves messages by room ID
func (s *MessageService) GetMessagesByRoomID(roomID string) ([]*models.Message, error) {
	return s.repo.GetMessagesByRoomID(roomID)
}
