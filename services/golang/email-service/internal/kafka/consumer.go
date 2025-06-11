package kafka

import (
	"context"
	"email-service/internal/email"
	"encoding/json"
	"fmt"
	"log"

	_kafka "github.com/segmentio/kafka-go"
)

type GenericEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type KafkaService struct {
	Brokers      []string
	Topic        string
	GroupID      string
	EmailService *email.EmailService
}

func (k *KafkaService) StartConsumer(handelFunc func(eventType string, data json.RawMessage, emailService *email.EmailService) error) {
	r := _kafka.NewReader(_kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "emails",
		GroupID:   "email-service-group",
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer r.Close()

	ctx := context.Background()

	fmt.Println("Consumer started...")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}

		var event GenericEvent
		err = json.Unmarshal(m.Value, &event)
		if err != nil {
			log.Printf("failed to unmarshal message: %v", err)
			continue
		}

		// TODO: Pass to email service
		handelFunc(event.Event, event.Data, k.EmailService)
		log.Printf("received message: %s, event: %s", m.Value, event.Event)
	}
}
