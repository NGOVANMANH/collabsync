import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MongooseModule } from '@nestjs/mongoose';
import { AuthModule } from './auth/auth.module';
import { KafkaModule } from './kafka/kafka.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { RoleService } from './role/role.service';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    MongooseModule.forRootAsync({
      inject: [ConfigService],
      useFactory: (config: ConfigService) => {
        const user = config.get<string>('MONGO_USER');
        const pass = config.get<string>('MONGO_PASSWORD');
        const db = config.get<string>('MONGO_DBNAME');
        const host = config.get<string>('MONGO_HOST') || 'localhost';
        const port = config.get<string>('MONGO_PORT') || '27017';

        const uri = `mongodb://${user}:${pass}@${host}:${port}/${db}?authSource=admin`;

        return { uri };
      },
    }),
    AuthModule,
    KafkaModule,
  ],
  controllers: [AppController],
  providers: [AppService, RoleService],
})
export class AppModule {}
