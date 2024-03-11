import { join } from 'path';

import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';

import { HelloResolver } from './hello.resolver';
import { HelloService } from './hello.service';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'HELLO_PACKAGE',
        options: {
          package: 'auto.trading.hello.v1',
          protoPath: join(__dirname, '../../proto/hello.proto'),
          url: 'auto_stock_trading_system_user_service:3001',
        },
        transport: Transport.GRPC,
      },
    ]),
  ],
  providers: [HelloResolver, HelloService],
})
export class HelloModule {}
