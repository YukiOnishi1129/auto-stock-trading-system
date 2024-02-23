/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from '@nestjs/microservices';
import { Observable } from 'rxjs';
import { Timestamp } from './google/protobuf/timestamp';

export const protobufPackage = 'auto.trading.hello.v1';

export interface HelloRequest {
  name: string;
}

export interface HelloResponse {
  message: string;
  createTime: Timestamp | undefined;
}

export const AUTO_TRADING_HELLO_V1_PACKAGE_NAME = 'auto.trading.hello.v1';

export interface GreetingServiceClient {
  /** unary */

  hello(request: HelloRequest): Observable<HelloResponse>;

  /** server streaming */

  helloServerStream(request: HelloRequest): Observable<HelloResponse>;

  /** client streaming */

  helloClientStream(
    request: Observable<HelloRequest>,
  ): Observable<HelloResponse>;

  /** bidirectional streaming */

  helloBiStreams(request: Observable<HelloRequest>): Observable<HelloResponse>;
}

export interface GreetingServiceController {
  /** unary */

  hello(
    request: HelloRequest,
  ): Promise<HelloResponse> | Observable<HelloResponse> | HelloResponse;

  /** server streaming */

  helloServerStream(request: HelloRequest): Observable<HelloResponse>;

  /** client streaming */

  helloClientStream(
    request: Observable<HelloRequest>,
  ): Promise<HelloResponse> | Observable<HelloResponse> | HelloResponse;

  /** bidirectional streaming */

  helloBiStreams(request: Observable<HelloRequest>): Observable<HelloResponse>;
}

export function GreetingServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ['hello', 'helloServerStream'];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcMethod('GreetingService', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
    const grpcStreamMethods: string[] = ['helloClientStream', 'helloBiStreams'];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(
        constructor.prototype,
        method,
      );
      GrpcStreamMethod('GreetingService', method)(
        constructor.prototype[method],
        method,
        descriptor,
      );
    }
  };
}

export const GREETING_SERVICE_NAME = 'GreetingService';
