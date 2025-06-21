export class UserDto {
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
}
