import { Injectable, OnModuleDestroy, OnModuleInit } from '@nestjs/common';
import { Kafka, Producer } from 'kafkajs';

@Injectable()
export class KafkaService implements OnModuleInit, OnModuleDestroy {
  private readonly producer: Producer;

  constructor() {
    const kafka = new Kafka({
      clientId: 'auth-service',
      brokers: [process.env.KAFKA_BROKER || 'localhost:9092'],
    });

    this.producer = kafka.producer();
  }

  async onModuleDestroy() {
    try {
      await this.producer.disconnect();
      console.log('Kafka producer disconnected successfully.');
    } catch (error) {
      console.error('Error disconnecting Kafka producer:', error);
    }
  }

  async onModuleInit() {
    try {
      await this.producer.connect();
      console.log('Kafka producer connected successfully.');
    } catch (error) {
      console.error('Error connecting Kafka producer:', error);
    }
  }

  async sendMessage(topic: string, message: any) {
    try {
      await this.producer.send({
        topic,
        messages: [
          {
            value: JSON.stringify(message),
          },
        ],
      });
      console.log(`Message sent to topic ${topic}:`, message);
    } catch (error) {
      console.error(`Error sending message to topic ${topic}:`, error);
    }
  }
}
