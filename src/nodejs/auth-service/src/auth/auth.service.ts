import {
  Injectable,
  NotFoundException,
  UnauthorizedException,
} from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { LoginDto } from 'src/auth/dtos/login.dto';
import { RegisterDto } from 'src/auth/dtos/register.dto';
import { User, UserStatus } from 'src/auth/schemas/user.schema';
import { comparePassword, hashPassword } from 'src/utils/hash';

@Injectable()
export class AuthService {
  constructor(
    @InjectModel(User.name) private readonly userModel: Model<User>,
    private readonly jwtService: JwtService,
  ) {}

  async login(dto: LoginDto): Promise<{
    user: {
      id: string;
      email: string;
      username?: string;
      firstName?: string;
      lastName?: string;
      displayName?: string;
      avatar?: string;
      locale?: string;
      timezone?: string;
      dateOfBirth?: Date;
    };
    accessToken: string;
  }> {
    const { email, password } = dto;

    if (!email || !password) {
      throw new UnauthorizedException('Email and password are required.');
    }

    const user = await this.userModel.findOne({ email }).select('+password');

    if (!user) {
      throw new NotFoundException('User not found.');
    }

    if (user.status !== UserStatus.ACTIVE) {
      throw new UnauthorizedException('User account is not active.');
    }

    if (!(await comparePassword(password, user.password || ''))) {
      throw new UnauthorizedException('Invalid email or password.');
    }

    const accessToken = await this.jwtService.signAsync({
      user_id: user._id.toString(),
      email: user.email,
      username: user.username,
      roles: user.roles,
    });

    return {
      user: {
        id: user._id.toString(),
        email: user.email,
        username: user.username,
        firstName: user.firstName,
        lastName: user.lastName,
        displayName: user.displayName,
        avatar: user.avatar,
        locale: user.locale,
        timezone: user.timezone,
        dateOfBirth: user.dateOfBirth,
      },
      accessToken: accessToken,
    };
  }

  async register(dto: RegisterDto): Promise<{
    user: {
      id: string;
      email: string;
      username?: string;
      firstName?: string;
      lastName?: string;
      displayName?: string;
      avatar?: string;
      locale?: string;
      timezone?: string;
      dateOfBirth?: Date;
    };
    verificationUrl?: string; // Optional, can be used for email verification
  }> {
    const {
      email,
      password,
      username,
      firstName,
      lastName,
      displayName,
      avatar,
      locale,
      timezone,
      dateOfBirth,
    } = dto;

    if (!email || !password) {
      throw new Error('Email and password are required.');
    }

    const existingUser = await this.userModel.find({ email });
    if (existingUser.length > 0) {
      throw new Error('Email already exists.');
    }

    const newUser = new this.userModel({
      email,
      password: await hashPassword(password), // Password should be hashed before saving
      username,
      firstName,
      lastName,
      displayName,
      avatar,
      locale,
      timezone,
      dateOfBirth,
      status: UserStatus.PENDING, // Default to pending until verified
    });

    await newUser.save();
    return {
      user: {
        id: newUser._id.toString(),
        email: newUser.email,
        username: newUser.username,
        firstName: newUser.firstName,
        lastName: newUser.lastName,
        displayName: newUser.displayName,
        avatar: newUser.avatar,
        locale: newUser.locale,
        timezone: newUser.timezone,
        dateOfBirth: newUser.dateOfBirth,
      },
      verificationUrl: 'mock-verification-url', // Replace with actual verification URL logic
    };
  }
}
