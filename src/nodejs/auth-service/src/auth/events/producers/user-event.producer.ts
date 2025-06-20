import { Injectable } from '@nestjs/common';
import { KafkaService } from 'src/kafka/kafka.service';

export type UserRegisterEvent = {
  user_id: string;
  email: string;
  verification_url: string;
};

@Injectable()
export class UserEventProducer {
  constructor(private readonly kafkaService: KafkaService) {}

  async sendUserCreatedEvent(event: UserRegisterEvent): Promise<void> {
    const message = {
      event: 'user.registered',
      data: event,
    };

    await this.kafkaService.sendMessage('emails', message);
  }
}
