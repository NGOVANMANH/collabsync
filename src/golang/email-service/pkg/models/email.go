package models

type EmailMessage struct {
	To          string   `json:"to"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	Attachments []string `json:"attachments,omitempty"`
}

func NewEmailMessage(to, subject, body string, attachments ...string) *EmailMessage {
	return &EmailMessage{
		To:          to,
		Subject:     subject,
		Body:        body,
		Attachments: attachments,
	}
}
