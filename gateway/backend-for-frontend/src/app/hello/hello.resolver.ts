import { Resolver, Query, Args } from '@nestjs/graphql';

import { Hello } from '../../graphql/types/graphql';

import { HelloService } from './hello.service';
@Resolver('Hello')
export class HelloResolver {
  constructor(private readonly helloService: HelloService) {}

  @Query('hello')
  async getHello(@Args('name') name: string): Promise<Hello> {
    return this.helloService.getHello(name);
  }
}
