import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { JwtService } from '@nestjs/jwt';
import { Request } from 'express';
import { Observable } from 'rxjs';
import { JwtPayloadDto } from './dtos/jwt-payload.dto';

@Injectable()
export class AuthGuard implements CanActivate {
  constructor(private readonly _jwtService: JwtService) {}
  canActivate(
    context: ExecutionContext,
  ): boolean | Promise<boolean> | Observable<boolean> {
    const request = context
      .switchToHttp()
      .getRequest<Request & { user?: JwtPayloadDto }>();

    const authHeader = request.headers.authorization;

    const [type, token] = authHeader?.split(' ') || [];

    if (type !== 'Bearer' || !token) {
      return false;
    }

    // Here you would typically validate the token
    try {
      const decoded: JwtPayloadDto = this._jwtService.verify(token);
      request.user = decoded; // Attach user info to the request
    } catch {
      return false; // Token is invalid
    }

    return true;
  }
}
