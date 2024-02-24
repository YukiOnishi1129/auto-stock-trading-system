import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
// import { Hello } from '../graphql/types/graphql';
import {
  GreetingServiceClient,
  HelloRequest,
  HelloResponse,
} from '../grpc/hello';
import { Observable } from 'rxjs';

@Injectable()
export class HelloService implements OnModuleInit {
  private greetingService: GreetingServiceClient;

  constructor(@Inject('HELLO_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.greetingService =
      this.client.getService<GreetingServiceClient>('GreetingService');
  }

  getHello(name: string): Observable<HelloResponse> {
    return this.greetingService.hello({ name } as HelloRequest);
    // const massageList: Array<string> = [];
    // response.forEach((res) => {
    //   massageList.push(res.message);
    // });
    // return {
    //   message: massageList[0],
    // };
  }
}
