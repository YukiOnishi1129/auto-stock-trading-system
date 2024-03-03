import { Injectable, OnModuleInit, Inject } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { User } from '../../graphql/types/graphql';
import { UserServiceClient } from '../../grpc/user';
import { lastValueFrom } from 'rxjs';
@Injectable()
export class UserService implements OnModuleInit {
  private userService: UserServiceClient;

  constructor(@Inject('USER_PACKAGE') private client: ClientGrpc) {}

  onModuleInit() {
    this.userService = this.client.getService<UserServiceClient>('UserService');
  }

  async getUsers(): Promise<Array<User>> {
    const rpcResponse = this.userService.getUsers({});
    const response = await lastValueFrom(rpcResponse);
    const resUsers = response.users.map((user) => {
      const date = user.createdAt;

      console.log(new Date(date.seconds * 1000));

      const resUser: User = {
        id: user.id,
        name: user.name,
        email: user.email,
        createdAt: new Date(user.createdAt.seconds * 1000).toString(),
        updatedAt: new Date(user.updatedAt.seconds * 1000).toString(),
        deletedAt: new Date(user.deletedAt.seconds * 1000).toString(),
      };
      return resUser;
    });
    return resUsers;
  }
}
