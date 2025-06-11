export class RegisterDto {
  email: string;
  username: string;
  password: string;
  firstName?: string;
  lastName?: string;
  displayName?: string;
  avatar?: string;
  locale?: string; // 'en-US', 'vi-VN', etc.
  timezone?: string; // 'Asia/Ho_Chi_Minh'
  dateOfBirth?: Date; // ISO date string or Date object
}
