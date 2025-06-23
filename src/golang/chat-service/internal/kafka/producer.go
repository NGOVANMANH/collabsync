package kafka

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce(key, msg []byte) {
	go func() error {
		log.Println("Producing message to Kafka...")

		if writer == nil {
			return kafka.ErrGenerationEnded
		}

		err := writer.WriteMessages(Ctx, kafka.Message{
			Key:   key,
			Value: msg,
			Time:  time.Now(),
		})

		if err != nil {
			return err
		}

		return nil
	}()
}
