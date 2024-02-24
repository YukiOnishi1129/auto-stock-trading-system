import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { HelloService } from './hello.service';
import { HelloResolver } from './hello.resolver';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'HELLO_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: 'auto_stock_trading_system_user_service:3001',
          package: 'auto.trading.hello.v1',
          protoPath: join(__dirname, '../proto/hello.proto'),
        },
      },
    ]),
  ],
  providers: [HelloResolver, HelloService],
})
export class HelloModule {}
