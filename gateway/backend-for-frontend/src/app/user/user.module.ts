import { join } from 'path';

import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';

import { UserResolver } from './user.resolver';
import { UserService } from './user.service';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'USER_PACKAGE',
        options: {
          package: 'auto.trading.user.v1',
          protoPath: join(__dirname, '../../proto/user.proto'),
          url: 'auto_stock_trading_system_user_service:3001',
        },
        transport: Transport.GRPC,
      },
    ]),
  ],
  providers: [UserResolver, UserService],
})
export class UserModule {}
