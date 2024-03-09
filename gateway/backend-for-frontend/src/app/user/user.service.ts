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
    try {
      const rpcResponse = this.userService.getUsers({});
      if (!rpcResponse) return [];
      const response = await lastValueFrom(rpcResponse);
      if (!response.users) return [];

      const resUsers = response.users.map((user) => {
        const resUser: User = {
          id: user.id,
          name: user.name,
          email: user.email,
          image: user.image,
          createdAt: new Date(user.createdAt.seconds * 1000).toString(),
          updatedAt: new Date(user.updatedAt.seconds * 1000).toString(),
          deletedAt: user?.deletedAt?.nanos
            ? new Date(user.deletedAt.seconds * 1000).toString()
            : null,
        };
        return resUser;
      });
      return resUsers;
    } catch (error) {
      throw new Error(error.message);
    }
  }
}
