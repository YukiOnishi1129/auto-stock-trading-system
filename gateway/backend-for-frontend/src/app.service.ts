import { Injectable } from '@nestjs/common';
// import { ClientGrpc } from '@nestjs/microservices';

// interface GreetingService {
//   Hello(request: { name: string }): { message: string };
// }

@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }
}
