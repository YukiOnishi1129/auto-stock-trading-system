import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { GraphQLServerModule } from '../graphql/graphql-server.module';
import { join } from 'path';
import { HelloModule } from './hello/hello.module';
import { UserModule } from './user/user.module';

@Module({
  imports: [
    GraphQLServerModule,
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
    HelloModule,
    UserModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
