import {
  Injectable,
  Logger,
  OnModuleDestroy,
  OnModuleInit,
} from '@nestjs/common';
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
      Logger.log('Kafka producer disconnected successfully.', 'KafkaService');
    } catch (error) {
      Logger.error(
        'Error disconnecting Kafka producer:',
        error,
        'KafkaService',
      );
    }
  }

  async onModuleInit() {
    try {
      await this.producer.connect();
      Logger.log('Kafka producer connected successfully.', 'KafkaService');
    } catch (error) {
      Logger.error('Error connecting Kafka producer:', error, 'KafkaService');
      throw new Error('Failed to connect Kafka producer');
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
      Logger.log(
        `Message sent to topic ${topic}: ${JSON.stringify(message)}`,
        'KafkaService',
      );
    } catch (error) {
      Logger.error(
        `Error sending message to topic ${topic}:`,
        error,
        'KafkaService',
      );
      throw new Error(`Failed to send message to topic ${topic}`);
    }
  }
}
