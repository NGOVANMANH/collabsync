package handler

import (
	"email-service/internal/email"
	"email-service/pkg/models"
	"encoding/json"
	"errors"
)

type UserRegisteredEvent struct {
	UserID          string `json:"user_id"`
	Email           string `json:"email"`
	VerificationURL string `json:"verification_url"`
}

type UserVerifiedEvent struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func HandleIncomingEvent(eventType string, data json.RawMessage, emailService *email.EmailService) error {
	switch eventType {
	case "user.registered":
		handleRegisteredEvent(data, emailService)
	case "user.verified":
		handleVerifiedEvent(data, emailService)
	default:
		return errors.New("unknown event type: " + eventType)
	}
	return nil
}

func handleRegisteredEvent(data json.RawMessage, emailService *email.EmailService) error {
	var userEvent UserRegisteredEvent
	if err := json.Unmarshal(data, &userEvent); err != nil {
		return err
	}
	println("User Registered:", userEvent.UserID, userEvent.Email, userEvent.VerificationURL)

	emailTemplataData := email.EmailTemplateData{
		UserName:        userEvent.Email,
		Title:           "Welcome to Our Service!",
		VerificationURL: userEvent.VerificationURL,
	}

	emailTemplate, err := email.GetEmailTemplate(&emailTemplataData)

	if err != nil {
		println("Error getting email template:", err)
		return err
	}

	emailMessage := models.EmailMessage{
		To:          userEvent.Email,
		Subject:     "Please verify your email",
		Body:        emailTemplate,
		Attachments: nil,
	}

	emailService.Send(&emailMessage)

	return nil
}

func handleVerifiedEvent(data json.RawMessage, emailService *email.EmailService) error {
	var userEvent UserVerifiedEvent
	if err := json.Unmarshal(data, &userEvent); err != nil {
		return err
	}
	println("User Verified:", userEvent.UserID, userEvent.Email)
	return nil
}
