package configs

import (
	"os"
	"strconv"
	"strings"
)

type KafkaConfig struct {
	Brokers   []string
	Topic     string
	GroupID   string
	Partition int
}

func GetKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		Topic:   os.Getenv("KAFKA_TOPIC"),
		GroupID: os.Getenv("KAFKA_GROUP_ID"),
		Partition: func() int {
			partition, err := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
			if err != nil {
				return 0
			}
			return partition
		}(),
	}
}
