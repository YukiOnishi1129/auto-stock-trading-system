import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { HelloModule } from './hello/hello.module';
import { HelloService } from './hello/hello.service';

@Module({
  imports: [
    ClientsModule.register([
      {
        name: 'HELLO_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: 'localhost:4001',
          package: 'hello',
          protoPath: 'src/hello/hello.proto',
        },
      },
    ]),
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
