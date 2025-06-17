package configs

type KafkaConfig struct {
	BootstrapServers []string
	Topic            string
	ConsumerGroup    string
}
