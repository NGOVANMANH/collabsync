package message

import (
	"chat-service/pkg/models"
	"database/sql"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// Define the interface
type IMessageRepository interface {
	GetMessages() ([]*models.Message, error)
	GetMessageByID(id uuid.UUID) (*models.Message, error)
	CreateMessage(msg *models.Message) (*models.Message, error)
	UpdateMessage(id uuid.UUID, msg *models.Message) error
	DeleteMessage(id uuid.UUID) error
	GetMessagesByRoomID(roomID string) ([]*models.Message, error)
}

// Define a struct that implements the interface
type MessageRepository struct {
	db *sql.DB
}

// Constructor
func NewMessageRepository(db *sql.DB) IMessageRepository {
	return &MessageRepository{db: db}
}

// Implement GetMessages
func (r *MessageRepository) GetMessages() ([]*models.Message, error) {
	query := `
		SELECT id, content, sender_id, room_id, created_at, updated_at 
		FROM messages 
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		msg := &models.Message{}
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&msg.ID,
			&msg.Content,
			&msg.SenderID,
			&msg.RoomID,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		msg.CreatedAt = createdAt
		msg.UpdatedAt = updatedAt
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

// Implement GetMessageByID
func (r *MessageRepository) GetMessageByID(id uuid.UUID) (*models.Message, error) {
	query := `
		SELECT id, content, sender_id, room_id, created_at, updated_at 
		FROM messages 
		WHERE id = $1
	`

	msg := &models.Message{}
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(query, id).Scan(
		&msg.ID,
		&msg.Content,
		&msg.SenderID,
		&msg.RoomID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Message not found
		}
		return nil, err
	}

	msg.CreatedAt = createdAt
	msg.UpdatedAt = updatedAt

	return msg, nil
}

// Implement CreateMessage
func (r *MessageRepository) CreateMessage(msg *models.Message) (*models.Message, error) {
	// Generate ID if not provided
	if msg.ID == uuid.Nil {
		msg.ID = uuid.New()
	}

	now := time.Now()
	msg.CreatedAt = now
	msg.UpdatedAt = now

	query := `
		INSERT INTO messages (id, content, sender_id, room_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, content, sender_id, room_id, created_at, updated_at
	`

	var createdAt, updatedAt time.Time
	err := r.db.QueryRow(
		query,
		msg.ID,
		msg.Content,
		msg.SenderID,
		msg.RoomID,
		msg.CreatedAt,
		msg.UpdatedAt,
	).Scan(
		&msg.ID,
		&msg.Content,
		&msg.SenderID,
		&msg.RoomID,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	msg.CreatedAt = createdAt
	msg.UpdatedAt = updatedAt

	return msg, nil
}

// Implement UpdateMessage
func (r *MessageRepository) UpdateMessage(id uuid.UUID, msg *models.Message) error {
	msg.UpdatedAt = time.Now()

	query := `
		UPDATE messages 
		SET content = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := r.db.Exec(query, id, msg.Content, msg.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Implement DeleteMessage
func (r *MessageRepository) DeleteMessage(id uuid.UUID) error {
	query := `DELETE FROM messages WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Additional helper method - Get messages by room
func (r *MessageRepository) GetMessagesByRoomID(roomID string) ([]*models.Message, error) {
	query := `
		SELECT id, content, sender_id, room_id, created_at, updated_at 
		FROM messages 
		WHERE room_id = $1 
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		msg := &models.Message{}
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&msg.ID,
			&msg.Content,
			&msg.SenderID,
			&msg.RoomID,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		msg.CreatedAt = createdAt
		msg.UpdatedAt = updatedAt
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}
