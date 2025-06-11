package kafka

import (
	"context"
	"log"

	kafka_lib "github.com/segmentio/kafka-go"
)

func ConsumeUserRegisterEvents(brokers []string, topic string, groupID string, handleFunc func(message kafka_lib.Message)) {
	r := kafka_lib.NewReader(kafka_lib.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, 
		MaxBytes: 10e6, 
	})
	defer r.Close()

	ctx := context.Background()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("error while reading message: %v", err)
			continue
		}
		handleFunc(m)
	}
}