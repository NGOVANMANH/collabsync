package email

import (
	"fmt"
	"net/smtp"
)

// EmailService provides methods to send emails.
type EmailService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

// NewEmailService creates a new EmailService.
func NewEmailService(host, port, username, password, from string) *EmailService {
	return &EmailService{
		SMTPHost: host,
		SMTPPort: port,
		Username: username,
		Password: password,
		From:     from,
	}
}

// Send sends an email to the specified recipient.
func (s *EmailService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost)
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", s.From, to, subject, body))
	addr := fmt.Sprintf("%s:%s", s.SMTPHost, s.SMTPPort)
	return smtp.SendMail(addr, auth, s.From, []string{to}, msg)
}