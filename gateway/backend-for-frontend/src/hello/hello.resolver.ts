import { Resolver, Query, Args } from '@nestjs/graphql';
import { HelloService } from './hello.service';
import { Hello } from '../graphql/types/graphql';
import { Observable } from 'rxjs';
import { HelloResponse } from '../grpc/hello';
@Resolver('Hello')
export class HelloResolver {
  constructor(private readonly helloService: HelloService) {}

  @Query(() => Hello)
  hello(@Args('name') name: string): Observable<HelloResponse> {
    return this.helloService.getHello(name);
  }
}
