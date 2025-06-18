package kafka

import (
	"chat-service/configs"
	"chat-service/pkg/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_kafka "github.com/segmentio/kafka-go"
)

type KafkaService struct {
	kafkaConfig *configs.KafkaConfig
	writer      *_kafka.Writer
	reader      *_kafka.Reader
	mu          sync.RWMutex
	closed      bool
}

// NewKafkaService creates a new KafkaService instance
func NewKafkaService(config *configs.KafkaConfig) (*KafkaService, error) {
	if config == nil {
		return nil, errors.New("kafka config cannot be nil")
	}

	if len(config.Brokers) == 0 {
		return nil, errors.New("no brokers provided")
	}

	ks := &KafkaService{
		kafkaConfig: config,
	}

	// Initialize writer
	ks.writer = _kafka.NewWriter(_kafka.WriterConfig{
		Brokers:      config.Brokers,
		Topic:        config.Topic,
		BatchSize:    100, // Better batch size for performance
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: int(_kafka.RequireOne), // Ensure at least one ack
		Async:        false,                  // Synchronous for reliability
	})

	return ks, nil
}

// StartConsume starts consuming messages with proper error handling and graceful shutdown
func (ks *KafkaService) StartConsume(ctx context.Context, handleFunc func(*models.Message) error) error {
	ks.mu.Lock()
	if ks.closed {
		ks.mu.Unlock()
		return errors.New("kafka service is closed")
	}

	// Initialize reader if not already done
	if ks.reader == nil {
		ks.reader = _kafka.NewReader(_kafka.ReaderConfig{
			Brokers:        ks.kafkaConfig.Brokers,
			Topic:          ks.kafkaConfig.Topic,
			GroupID:        ks.kafkaConfig.GroupID,
			Partition:      ks.kafkaConfig.Partition,
			MaxBytes:       10e6, // 10MB
			CommitInterval: time.Second,
			StartOffset:    _kafka.LastOffset,
		})
	}
	ks.mu.Unlock()

	log.Printf("Starting Kafka consumer for topic: %s", ks.kafkaConfig.Topic)

	for {
		select {
		case <-ctx.Done():
			log.Println("Context cancelled, stopping consumer")
			return ctx.Err()
		default:
			// Set a timeout for reading messages
			msgCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			message, err := ks.reader.FetchMessage(msgCtx)
			cancel()

			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					continue // Timeout is expected, continue polling
				}
				log.Printf("Error reading message: %v", err)
				continue // Continue consuming even on errors
			}

			if err := ks.processMessage(&message, handleFunc); err != nil {
				log.Printf("Error processing message: %v", err)
				// Don't commit the message on processing error
				continue
			}

			// Commit the message after successful processing
			if err := ks.reader.CommitMessages(ctx, message); err != nil {
				log.Printf("Error committing message: %v", err)
			}
		}
	}
}

// processMessage handles the message processing logic
func (ks *KafkaService) processMessage(kafkaMsg *_kafka.Message, handleFunc func(*models.Message) error) error {
	if len(kafkaMsg.Value) == 0 {
		return errors.New("received empty message value")
	}

	var msg models.Message

	// Try JSON unmarshaling first (assuming this is the intended format)
	if err := json.Unmarshal(kafkaMsg.Value, &msg); err != nil {
		// If JSON fails, try the original binary unmarshaling
		if binaryErr := msg.ID.UnmarshalBinary(kafkaMsg.Value); binaryErr != nil {
			return fmt.Errorf("failed to unmarshal message as JSON: %v, and as binary: %v", err, binaryErr)
		}
	}

	// Call the handler function
	return handleFunc(&msg)
}

// Produce sends a message to Kafka with better error handling
func (ks *KafkaService) Produce(ctx context.Context, message *models.Message) error {
	if message == nil {
		return errors.New("message cannot be nil")
	}

	ks.mu.RLock()
	if ks.closed {
		ks.mu.RUnlock()
		return errors.New("kafka service is closed")
	}
	writer := ks.writer
	ks.mu.RUnlock()

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	kafkaMessage := _kafka.Message{
		Key:   []byte(message.ID.String()), // Use message ID as key for partitioning
		Value: messageBytes,
		Time:  time.Now(),
	}

	// Use context with timeout for production
	prodCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := writer.WriteMessages(prodCtx, kafkaMessage); err != nil {
		return fmt.Errorf("failed to write message to kafka: %w", err)
	}

	return nil
}

// ProduceBatch sends multiple messages in a single batch for better performance
func (ks *KafkaService) ProduceBatch(ctx context.Context, messages []*models.Message) error {
	if len(messages) == 0 {
		return errors.New("no messages to produce")
	}

	ks.mu.RLock()
	if ks.closed {
		ks.mu.RUnlock()
		return errors.New("kafka service is closed")
	}
	writer := ks.writer
	ks.mu.RUnlock()

	kafkaMessages := make([]_kafka.Message, 0, len(messages))

	for _, msg := range messages {
		if msg == nil {
			continue // Skip nil messages
		}

		messageBytes, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to marshal message %s: %v", msg.ID.String(), err)
			continue // Skip messages that can't be marshaled
		}

		kafkaMessages = append(kafkaMessages, _kafka.Message{
			Key:   []byte(msg.ID.String()),
			Value: messageBytes,
			Time:  time.Now(),
		})
	}

	if len(kafkaMessages) == 0 {
		return errors.New("no valid messages to produce")
	}

	prodCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	if err := writer.WriteMessages(prodCtx, kafkaMessages...); err != nil {
		return fmt.Errorf("failed to write batch messages to kafka: %w", err)
	}

	return nil
}

// Close gracefully shuts down the Kafka service
func (ks *KafkaService) Close() error {
	ks.mu.Lock()
	defer ks.mu.Unlock()

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

// Health check methods
func (ks *KafkaService) IsHealthy(ctx context.Context) error {
	ks.mu.RLock()
	defer ks.mu.RUnlock()

	if ks.closed {
		return errors.New("kafka service is closed")
	}

	// Try to get metadata to check connectivity
	conn, err := _kafka.DialContext(ctx, "tcp", ks.kafkaConfig.Brokers[0])
	if err != nil {
		return fmt.Errorf("failed to connect to kafka: %w", err)
	}
	defer conn.Close()

	return nil
}
