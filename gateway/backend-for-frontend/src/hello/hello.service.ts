import { Injectable } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';

@Injectable()
export class HelloService {
  @GrpcMethod('HelloService', 'Hello')
  hello(data: { name: string }) {
    return { message: `Hello, ${data.name}!` };
  }
}
