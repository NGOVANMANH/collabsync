import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument, Types } from 'mongoose';
import { IsEmail, IsEnum } from 'class-validator';

export type UserDocument = HydratedDocument<User>;

export enum UserStatus {
  ACTIVE = 'active',
  PENDING = 'pending',
  SUSPENDED = 'suspended',
  DELETED = 'deleted',
}

@Schema({
  timestamps: true,
  collection: 'users',
})
export class User {
  _id: Types.ObjectId;

  // Primary identifiers
  @Prop({
    required: true,
    unique: true,
    lowercase: true,
    trim: true,
    index: true,
  })
  @IsEmail()
  email: string;

  @Prop({ required: false, unique: true, sparse: true })
  username?: string;

  // Authentication
  @Prop({ required: false, select: false }) // Not required for OAuth users
  password?: string;

  @Prop({ type: String, enum: UserStatus, default: UserStatus.PENDING })
  @IsEnum(UserStatus)
  status: UserStatus;

  // Profile information
  @Prop({ required: false, trim: true })
  firstName?: string;

  @Prop({ required: false, trim: true })
  lastName?: string;

  @Prop({ required: false, trim: true })
  displayName?: string;

  @Prop({ required: false })
  avatar?: string;

  @Prop({ required: false })
  locale?: string; // 'en-US', 'vi-VN', etc.

  @Prop({ required: false })
  timezone?: string; // 'Asia/Ho_Chi_Minh'

  @Prop({ required: false })
  dateOfBirth?: Date;

  // Verification status
  @Prop({ default: false })
  emailVerified: boolean;

  @Prop({ required: false })
  emailVerifiedAt?: Date;

  // Security & Access
  @Prop({ type: [String], default: ['user'] })
  roles: string[];

  @Prop({ type: Object, default: {} }) // Custom user attributes
  metadata: Record<string, any>;

  // Login tracking
  @Prop({ required: false })
  lastLogin?: Date;

  @Prop({ required: false })
  lastLoginIp?: string;

  @Prop({ default: 0 })
  loginCount: number;

  // Account security
  @Prop({ default: 0 })
  failedLoginAttempts: number;

  @Prop({ required: false })
  lockedUntil?: Date;

  @Prop({ required: false })
  passwordChangedAt?: Date;

  // Multi-factor authentication
  @Prop({ default: false })
  mfaEnabled: boolean;

  @Prop({ required: false, select: false })
  mfaSecret?: string;

  @Prop({ type: [String], default: [] })
  mfaBackupCodes: string[];

  // Terms and privacy
  @Prop({ required: false })
  termsAcceptedAt?: Date;

  @Prop({ required: false })
  privacyAcceptedAt?: Date;

  createdAt: Date;
  updatedAt: Date;
}

export const UserSchema = SchemaFactory.createForClass(User);

UserSchema.index({ email: 1 }, { unique: true });
UserSchema.index({ phoneNumber: 1 }, { sparse: true, unique: true });
UserSchema.index({ username: 1 }, { sparse: true, unique: true });
UserSchema.index({ status: 1 });
UserSchema.index({ roles: 1 });
UserSchema.index({ emailVerified: 1 });
UserSchema.index({ createdAt: -1 });
UserSchema.index({ lastLogin: -1 });

UserSchema.virtual('fullName').get(function () {
  return `${this.firstName || ''} ${this.lastName || ''}`.trim();
});

UserSchema.virtual('isLocked').get(function () {
  return !!(this.lockedUntil && this.lockedUntil > new Date());
});
