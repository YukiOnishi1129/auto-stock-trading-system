import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import {
  GreetingServiceClient,
  HelloRequest,
  HelloResponse,
} from 'src/grpc/hello';

@Injectable()
export class AppService implements OnModuleInit {
  private greetingService: GreetingServiceClient;

  constructor(
    @Inject('HELLO_PACKAGE')
    private client: ClientGrpc,
  ) {}

  onModuleInit() {
    this.greetingService =
      this.client.getService<GreetingServiceClient>('GreetingService');
  }

  getHello(): Observable<HelloResponse> {
    return this.greetingService.hello({
      name: 'Nest',
    } as HelloRequest);
  }
}
