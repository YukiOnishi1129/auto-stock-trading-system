import { Resolver, Query, Args } from '@nestjs/graphql';
import { HelloService } from './hello.service';
import { Hello, IQuery } from '../graphql/types/graphql';
@Resolver('Hello')
export class HelloResolver implements IQuery {
  constructor(private readonly helloService: HelloService) {}

  @Query(() => Hello)
  async hello(@Args('name') name: string) {
    return this.helloService.getHello(name);
  }
}
