import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument, Types } from 'mongoose';

export type RoleDocument = HydratedDocument<Role>;

@Schema({
  timestamps: true,
  collection: 'roles',
})
export class Role {
  _id: Types.ObjectId;

  @Prop({
    required: true,
    unique: true,
    lowercase: true,
    trim: true,
  })
  name: string;

  @Prop({
    required: false,
    trim: true,
    default: null,
  })
  description?: string;

  @Prop({
    type: [String],
    default: [],
  })
  permissions: string[];

  @Prop({
    default: true,
  })
  isActive: boolean;

  @Prop({
    default: false,
  })
  isSystem: boolean;

  createdAt: Date;
  updatedAt: Date;
}

export const RoleSchema = SchemaFactory.createForClass(Role);

RoleSchema.index({ name: 1 }, { unique: true });
RoleSchema.index({ isActive: 1 });
