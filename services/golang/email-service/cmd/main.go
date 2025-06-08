package main

import (
	"email-service/config"
	"fmt"
)

func main() {
	cfg := config.LoadConfig();

	fmt.Println("Kafka Broker:", cfg.KafkaBroker)
	fmt.Println("Kafka Topic:", cfg.KafkaTopic)
	fmt.Println("SMTP Host:", cfg.SMTPHost)
	fmt.Println("SMTP Port:", cfg.SMTPPort)
	fmt.Println("SMTP User:", cfg.SMTPUser)
	fmt.Println("SMTP Password:", cfg.SMTPPassword)
	fmt.Println("Email Service is running with the above configuration.")
}