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
      imports: [ConfigModule],
      useFactory: (configService: ConfigService) => {
        const host = configService.get<string>('MONGO_HOST', 'localhost');
        const port = configService.get<string>('MONGO_PORT', '27017');
        const db = configService.get<string>('MONGO_DBNAME', 'mydatabase');
        const user = configService.get<string>('MONGO_USER');
        const pass = configService.get<string>('MONGO_PASSWORD');

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
  providers: [AppService, RoleService],
})
export class AppModule {}
