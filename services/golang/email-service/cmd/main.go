package main

import (
	"email-service/config"
	"email-service/internal/email"
	"email-service/internal/handler"
	"email-service/internal/kafka"
)

func main() {
	cfg := config.LoadConfig()

	emailService := &email.EmailService{
		SMTPHost: cfg.SMTPHost,
		SMTPPort: cfg.SMTPPort,
		Username: cfg.SMTPUser,
		Password: cfg.SMTPPassword,
		From:     cfg.SMTPFrom,
	}

	kafkaService := kafka.KafkaService{
		Brokers:      []string{cfg.KafkaBroker},
		Topic:        cfg.KafkaTopic,
		GroupID:      "email-service-group",
		EmailService: emailService,
	}

	kafkaService.StartConsumer(handler.HandleIncomingEvent)
}
