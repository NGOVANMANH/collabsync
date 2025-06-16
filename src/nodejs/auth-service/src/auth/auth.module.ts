import { Module } from '@nestjs/common';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';
import { MongooseModule } from '@nestjs/mongoose';
import { User, UserSchema } from './schemas/user.schema';
import { JwtModule } from '@nestjs/jwt';
import { UserEventProducer } from './events/producers/user-event.producer';
import { KafkaModule } from 'src/kafka/kafka.module';
import { ConfigService } from '@nestjs/config';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: User.name, schema: UserSchema }]),
    JwtModule.registerAsync({
      inject: [ConfigService],
      useFactory: (config: ConfigService) => ({
        secret: config.get<string>('JWT_SECRET') || 'defaultSecret',
        signOptions: {
          expiresIn: config.get<string>('JWT_EXPIRATION') || '1h',
        },
      }),
    }),
    KafkaModule,
  ],
  controllers: [AuthController],
  providers: [AuthService, UserEventProducer],
})
export class AuthModule {}
