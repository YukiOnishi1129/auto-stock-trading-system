import { Module } from '@nestjs/common';

import { GraphQLServerModule } from '../graphql/graphql-server.module';

import { HelloModule } from './hello/hello.module';
import { ScalerModule } from './scaler/scaler.module';
import { UserModule } from './user/user.module';
@Module({
  controllers: [],
  imports: [
    GraphQLServerModule,
    // ClientsModule.register([
    //   {
    //     name: 'HELLO_PACKAGE',
    //     transport: Transport.GRPC,
    //     options: {
    //       url: 'auto_stock_trading_system_user_service:3001',
    //       package: 'auto.trading.hello.v1',
    //       protoPath: join(__dirname, 'proto/hello.proto'),
    //     },
    //   },
    // ]),
    HelloModule,
    UserModule,
    ScalerModule,
  ],
  providers: [],
})
export class AppModule {}
