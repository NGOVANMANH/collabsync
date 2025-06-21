package models

import "github.com/google/uuid"

type Room struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	Name         string      `json:"name" db:"name"`
	Description  string      `json:"description" db:"description"`
	ThumbnailURL string      `json:"thumbnail_url" db:"thumbnail_url"`
	Participants []uuid.UUID `json:"participants" db:"participants"`
	CreatedAt    string      `json:"created_at" db:"created_at"`
	UpdatedAt    string      `json:"updated_at" db:"updated_at"`
}
