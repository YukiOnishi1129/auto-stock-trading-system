import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';

async function bootstrap() {
  const app = await NestFactory.createMicroservice(AppModule, {
    transport: Transport.GRPC,
    options: {
      package: 'app',
      protoPath: 'src/app.proto',
    },
  });

  await app.listen(3000);
}
bootstrap();
