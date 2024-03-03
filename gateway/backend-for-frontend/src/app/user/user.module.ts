import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { UserService } from './user.service';
import { UserResolver } from './user.resolver';
import { join } from 'path';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'USER_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: 'auto_stock_trading_system_user_service:3001',
          package: 'auto.trading.user.v1',
          protoPath: join(__dirname, '../../proto/user.proto'),
        },
      },
    ]),
  ],
  providers: [UserResolver, UserService],
})
export class UserModule {}
