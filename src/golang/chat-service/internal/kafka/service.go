package kafka

import (
	"chat-service/configs"
	"chat-service/internal/message"
	"chat-service/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	_kafka "github.com/segmentio/kafka-go"
)

var messageService message.IMessageService

type KafkaService struct {
	kafkaConfig *configs.KafkaConfig
	writer      *_kafka.Writer
	reader      *_kafka.Reader
	closed      bool
}

func NewKafkaService(config *configs.KafkaConfig, msgService message.IMessageService) (*KafkaService, error) {
	messageService = msgService

	if config == nil {
		return nil, errors.New("kafka config cannot be nil")
	}

	if len(config.Brokers) == 0 {
		return nil, errors.New("no brokers provided")
	}

	ks := &KafkaService{
		kafkaConfig: config,
	}

	ks.writer = _kafka.NewWriter(_kafka.WriterConfig{
		Brokers:  config.Brokers,
		Topic:    config.Topic,
		Balancer: &_kafka.LeastBytes{},
	})

	return ks, nil
}

func (ks *KafkaService) StartConsume() error {
	if ks.reader == nil {
		ks.reader = _kafka.NewReader(_kafka.ReaderConfig{
			Brokers:   ks.kafkaConfig.Brokers,
			Topic:     ks.kafkaConfig.Topic,
			GroupID:   ks.kafkaConfig.GroupID,
			Partition: ks.kafkaConfig.Partition,
			MaxBytes:  10e6, // 10MB
		})
	}

	log.Printf("Starting Kafka consumer for topic: %s", ks.kafkaConfig.Topic)

	for {
		m, err := ks.reader.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("failed to read message from kafka: %w", err)
		}

		message := &models.Message{}
		if err := json.Unmarshal(m.Value, message); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			continue // Skip malformed messages
		}

		if err := handleMessage(message); err != nil {
			log.Printf("Error handling message: %v", err)
			continue
		}

		log.Printf("Consumed message from Kafka topic %s: %s", ks.kafkaConfig.Topic, message.ID)

	}
}

func (ks *KafkaService) Produce(message *models.Message) error {
	if message == nil {
		return errors.New("message cannot be nil")
	}

	writer := ks.writer

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	kafkaMessage := _kafka.Message{
		Key:   []byte(message.ID.String()),
		Value: messageBytes,
		Time:  time.Now(),
	}

	if err := writer.WriteMessages(context.Background(), kafkaMessage); err != nil {
		return fmt.Errorf("failed to write message to kafka: %w", err)
	}
	log.Printf("Produced message to Kafka topic %s: %s", ks.kafkaConfig.Topic, message.ID)
	return nil
}

func (ks *KafkaService) Close() error {
	if ks.closed {
		return nil
	}

	ks.closed = true

	var errs []error

	if ks.writer != nil {
		if err := ks.writer.Close(); err != nil {
			errs = append(errs, fmt.Errorf("error closing writer: %w", err))
		}
	}

	if ks.reader != nil {
		if err := ks.reader.Close(); err != nil {
			errs = append(errs, fmt.Errorf("error closing reader: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors during close: %v", errs)
	}

	return nil
}
