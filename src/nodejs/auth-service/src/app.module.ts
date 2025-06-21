import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { AuthModule } from './auth/auth.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { KafkaModule } from './kafka/kafka.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    MongooseModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: async (configService: ConfigService) => {
        const host = configService.get('MONGO_HOST', 'localhost');
        const port = configService.get('MONGO_PORT', '27017');
        const db = configService.get('MONGO_DBNAME', 'mydatabase');
        const user = configService.get('MONGO_USER');
        const pass = configService.get('MONGO_PASSWORD');

        const authPart = user && pass ? `${user}:${pass}@` : '';
        const uri = `mongodb://${authPart}${host}:${port}/${db}?authSource=admin`;

        return {
          uri,
        };
      },
      inject: [ConfigService],
    }),
    AuthModule,
    KafkaModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
