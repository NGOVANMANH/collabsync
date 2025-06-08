import { Injectable } from '@nestjs/common';
import { RegisterDto } from 'src/auth/dtos/register.dto';
import { KafkaService } from 'src/kafka/kafka.service';

@Injectable()
export class UserEventProducer {
  constructor(private readonly kafkaService: KafkaService) {}

  async sendUserCreatedEvent(user: RegisterDto): Promise<void> {
    const message = {
      event: 'user.created',
      data: user,
    };

    await this.kafkaService.sendMessage('user-events', message);
  }
}
