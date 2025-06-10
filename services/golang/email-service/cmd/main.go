package main

import (
	"email-service/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	// emailService := email.NewEmailService(
	// 	cfg.SMTPHost,
	// 	cfg.SMTPPort,
	// 	cfg.SMTPUser,
	// 	cfg.SMTPPassword,
	// 	cfg.SMTPUser, // from address
	// )

	// emailTemplateData, err := email.GetEmailTemplate(email.NewEmailTemplateData(
	// 	"Test Email",
	// 	"Hello, this is a test email sent from the email service.",
	// 	"https://example.com/verify?token",
	// ))

	// if err != nil {
	// 	log.Fatalf("Failed to get email template: %v", err)
	// }

	// emailMessage := models.NewEmailMessage(
	// 	"21522328@gm.uit.edu.vn",              // recipient
	// 	"Test Email",                          // subject
	// 	emailTemplateData,                     // body
	// 	"https://example.com/attachment1.pdf", // attachment URL or file path
	// )

	// err = emailService.Send(emailMessage)
	// if err != nil {
	// 	log.Fatalf("Failed to send email: %v", err)
	// }

	// log.Println("Email sent successfully")

	// event.StartConsumer(cfg)

	log.Println(cfg.KafkaBroker)
}
