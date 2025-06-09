import { Injectable } from '@nestjs/common';
import { KafkaService } from 'src/kafka/kafka.service';

@Injectable()
export class UserEventProducer {
  constructor(private readonly kafkaService: KafkaService) {}

  async sendUserCreatedEvent(event: {
    user: { id: string; email: string };
    verificationUrl: string;
  }): Promise<void> {
    const message = {
      event: 'user.registered',
      data: event,
    };

    await this.kafkaService.sendMessage('user-events', message);
  }
}
