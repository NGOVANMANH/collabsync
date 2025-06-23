package kafka

import (
	"context"
	"os"
	"strings"

	"github.com/segmentio/kafka-go"
)

var (
	writer *kafka.Writer
	reader *kafka.Reader
	Ctx    = context.Background()
)

func Init() {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   strings.Split(os.Getenv("CHAT_KAFKA_BROKERS"), ","),
		GroupID:   os.Getenv("CHAT_KAFKA_GROUP_ID"),
		Topic:     os.Getenv("CHAT_KAFKA_TOPIC"),
		Partition: 0,
	})

	writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(os.Getenv("CHAT_KAFKA_BROKERS"), ","),
		Topic:    os.Getenv("CHAT_KAFKA_TOPIC"),
		Balancer: &kafka.LeastBytes{},
	})
}

func Close() {
	if writer != nil {
		if err := writer.Close(); err != nil {
			panic(err)
		}
	}

	if reader != nil {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}
}
