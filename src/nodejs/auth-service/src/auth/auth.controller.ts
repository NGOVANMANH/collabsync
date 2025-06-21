import {
  Body,
  Controller,
  Get,
  HttpStatus,
  Logger,
  Post,
  Req,
  Res,
  UseGuards,
} from '@nestjs/common';
import { Request, Response } from 'express';
import { LoginDto } from 'src/auth/dtos/login.dto';
import { AuthService } from './auth.service';
import { RegisterDto } from './dtos/register.dto';
import { UserEventProducer } from './events/producers/user-event.producer';
import { JwtPayloadDto } from './dtos/jwt-payload.dto';
import { AuthGuard } from './auth.guard';

@Controller('auth')
export class AuthController {
  constructor(
    private readonly authService: AuthService,
    private readonly userEventProducer: UserEventProducer,
  ) {}

  @Post('login')
  async login(@Body() dto: LoginDto, @Res() res: Response): Promise<any> {
    const { email, password } = dto;

    if (!email || !password) {
      return res
        .status(HttpStatus.BAD_REQUEST)
        .json({ message: 'Email and password are required.' });
    }

    const data = await this.authService.login(dto);
    return res.status(HttpStatus.OK).json(data);
  }

  @Post('register')
  async register(@Body() dto: RegisterDto, @Res() res: Response): Promise<any> {
    const { email, password } = dto;

    if (!email || !password) {
      return res
        .status(HttpStatus.BAD_REQUEST)
        .json({ message: 'Email and password are required.' });
    }

    const { user, verificationUrl } = await this.authService.register(dto);

    this.userEventProducer
      .sendUserCreatedEvent({
        user_id: user.id,
        email: user.email,
        verification_url: verificationUrl || '',
      })
      .catch((error) => {
        Logger.error(
          'Error sending user created event:',
          error,
          'AuthController',
        );
      });

    return res.status(HttpStatus.CREATED).json(user);
  }

  @Get('test-auth')
  @UseGuards(AuthGuard)
  testAuth(@Req() req: Request & { user: JwtPayloadDto }): any {
    return req.user;
  }
}
