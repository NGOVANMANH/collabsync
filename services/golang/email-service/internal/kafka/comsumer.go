package kafka

import (
	"context"
	"email-service/pkg/models"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic string, handler func(models.EmailMessage)) {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers: []string{broker},
        Topic:   topic,
        GroupID: "email-service-group",
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            log.Println("Failed to read message:", err)
            continue
        }

        var email models.EmailMessage
        if err := json.Unmarshal(m.Value, &email); err != nil {
            log.Println("Failed to unmarshal message:", err)
            continue
        }

        go handler(email)
    }
}
