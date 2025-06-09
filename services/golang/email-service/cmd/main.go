package main

import (
	"email-service/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize the email service with the loaded configuration
	fmt.Printf("Email Service Configuration:\n")
	fmt.Printf("Kafka Broker: %s\n", cfg.KafkaBroker)
	fmt.Printf("Kafka Topic: %s\n", cfg.KafkaTopic)
}
