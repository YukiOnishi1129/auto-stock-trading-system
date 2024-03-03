import { Resolver, Query } from '@nestjs/graphql';
import { UserService } from './user.service';
import { User } from '../../graphql/types/graphql';

@Resolver('User')
export class UserResolver {
  constructor(private readonly userService: UserService) {}

  @Query('users')
  async getUsers(): Promise<Array<User>> {
    return this.userService.getUsers();
  }
}
