import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { IsEnum } from 'class-validator';
import { HydratedDocument, Types } from 'mongoose';

export enum IdentityProvider {
  LOCAL = 'local',
  GOOGLE = 'google',
  FACEBOOK = 'facebook',
  GITHUB = 'github',
  APPLE = 'apple',
  MICROSOFT = 'microsoft',
  TWITTER = 'twitter',
}

export type AccountDocument = HydratedDocument<Account>;

@Schema({
  timestamps: true,
  collection: 'accounts',
})
export class Account {
  _id: Types.ObjectId;

  @Prop({ type: Types.ObjectId, ref: 'User', required: true, index: true })
  userId: Types.ObjectId;

  @Prop({ required: true, enum: Object.values(IdentityProvider) })
  @IsEnum(IdentityProvider)
  provider: IdentityProvider;

  @Prop({ required: true })
  providerAccountId: string; // Provider's user ID

  @Prop({ required: false })
  providerEmail?: string;

  @Prop({ required: false })
  providerName?: string; // Display name from provider

  @Prop({ required: false })
  providerAvatar?: string;

  // OAuth tokens
  @Prop({ required: false, select: false })
  accessToken?: string;

  @Prop({ required: false, select: false })
  refreshToken?: string;

  @Prop({ required: false })
  accessTokenExpires?: Date;

  @Prop({ required: false })
  tokenType?: string; // 'Bearer'

  @Prop({ required: false })
  scope?: string;

  @Prop({ required: false })
  idToken?: string;

  @Prop({ type: Object, default: {} })
  providerData: Record<string, any>; // Raw provider response

  @Prop({ default: true })
  isActive: boolean;

  createdAt: Date;
  updatedAt: Date;
}

export const AccountSchema = SchemaFactory.createForClass(Account);

AccountSchema.index({ userId: 1 });
AccountSchema.index({ provider: 1, providerAccountId: 1 }, { unique: true });
AccountSchema.index({ providerEmail: 1 });
