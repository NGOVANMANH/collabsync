package email

import (
	"bytes"
	"email-service/pkg/models"
	"fmt"
	"net/smtp"
)

type EmailService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

func NewEmailService(host, port, username, password, from string) *EmailService {
	return &EmailService{
		SMTPHost: host,
		SMTPPort: port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (s *EmailService) Send(emailMessage *models.EmailMessage) error {
	if emailMessage == nil {
		return fmt.Errorf("email message is nil")
	}

	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("From: %s\r\n", s.From))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", emailMessage.To))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", emailMessage.Subject))
	msg.WriteString("MIME-Version: 1.0\r\n")
	msg.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	msg.WriteString("\r\n")
	msg.WriteString(emailMessage.Body)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost)

	err := smtp.SendMail(
		s.SMTPHost+":"+s.SMTPPort,
		auth,
		s.From,
		[]string{emailMessage.To},
		msg.Bytes(),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
