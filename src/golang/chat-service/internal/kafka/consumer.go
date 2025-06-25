package kafka

import (
	"log"
)

func StartConsumer() {
	if reader == nil {
		return
	}

	log.Println("Starting Kafka consumer...")

	go func() {
		for {
			msg, err := reader.ReadMessage(Ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			log.Printf("Received message: %s:%s", string(msg.Key), string(msg.Value))
		}
	}()
}
