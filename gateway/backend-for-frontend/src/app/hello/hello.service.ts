import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';

import { Hello } from '../../graphql/types/graphql';
import { GreetingServiceClient, HelloRequest } from '../../grpc/hello';

@Injectable()
export class HelloService implements OnModuleInit {
  private greetingService: GreetingServiceClient;

  constructor(@Inject('HELLO_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.greetingService =
      this.client.getService<GreetingServiceClient>('GreetingService');
  }

  async getHello(name: string): Promise<Hello> {
    const rpcResponse = this.greetingService.hello({ name } as HelloRequest);
    const response = await lastValueFrom(rpcResponse);
    return {
      message: response.message,
    };
  }
}
