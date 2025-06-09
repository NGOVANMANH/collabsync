package main

import (
	"email-service/config"
	"email-service/internal/email"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize the email service with the loaded configuration
	fmt.Printf("Email Service Configuration:\n")
	fmt.Printf("Kafka Broker: %s\n", cfg.KafkaBroker)
	fmt.Printf("Kafka Topic: %s\n", cfg.KafkaTopic)

	emailTemplate, err := email.GetEmailTemplate(email.NewEmailTemplateData("Name", "Subject", "Body content"))
	if err != nil {
		fmt.Printf("Error generating email template: %v\n", err)
		return
	}

	fmt.Println("Generated Email Template:")
	fmt.Println(emailTemplate)
}