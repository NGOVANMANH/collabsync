package configs

type KafkaConfig struct {
	Brokers   []string
	Topic     string
	GroupID   string
	Partition int
}

func NewKafkaConfig(brokers []string, topic, groupID string, partition int) *KafkaConfig {
	return &KafkaConfig{
		Brokers:   brokers,
		Topic:     topic,
		GroupID:   groupID,
		Partition: partition,
	}
}
