import { Injectable } from '@nestjs/common';
import { Hello } from '../graphql/types/graphql';

@Injectable()
export class HelloService {
  async getHello(name: string): Promise<Hello> {
    return {
      message: `This action returns a #${name} hello`,
    };
  }
}
