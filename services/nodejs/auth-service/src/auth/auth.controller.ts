import { Body, Controller, Post, Res } from '@nestjs/common';
import { LoginDto } from 'src/auth/dtos/login.dto';
import { AuthService } from './auth.service';
import { RegisterDto } from './dtos/register.dto';
import { UserEventProducer } from './events/producers/user-event.producer';

@Controller('auth')
export class AuthController {
  constructor(
    private readonly authService: AuthService,
    private readonly userEventProducer: UserEventProducer,
  ) {}

  @Post('login')
  async login(@Body() dto: LoginDto, @Res() res) {
    const { email, password } = dto;

    if (!email || !password) {
      return res
        .status(400)
        .json({ message: 'Email and password are required.' });
    }

    const data = await this.authService.login(dto);
    return res.status(200).json(data);
  }

  @Post('register')
  async register(@Body() dto: RegisterDto, @Res() res) {
    const { email, password } = dto;

    if (!email || !password) {
      return res
        .status(400)
        .json({ message: 'Email and password are required.' });
    }

    const user = await this.authService.register(dto);

    this.userEventProducer.sendUserCreatedEvent(dto);

    return res.status(201).json(user);
  }
}
