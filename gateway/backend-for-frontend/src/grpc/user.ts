/* eslint-disable */
import { GrpcMethod, GrpcStreamMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { Timestamp } from "./google/protobuf/timestamp";

export const protobufPackage = "auto.trading.user.v1";

export interface CreateUserRequest {
  email: string;
  password: string;
  name: string;
}

export interface GetUserRequest {
  id: string;
}

export interface UpdateUserRequest {
  id: string;
  email: string;
  password: string;
  name: string;
}

export interface DeleteUserRequest {
  id: string;
}

export interface Empty {
}

export interface User {
  id: string;
  email: string;
  name: string;
  createdAt: Timestamp | undefined;
  updatedAt: Timestamp | undefined;
}

export interface UserResponse {
  user: User | undefined;
}

export interface UsersResponse {
  users: User[];
}

export const AUTO_TRADING_USER_V1_PACKAGE_NAME = "auto.trading.user.v1";

export interface UserServiceClient {
  getUser(request: GetUserRequest): Observable<UserResponse>;

  getUsers(request: Empty): Observable<UsersResponse>;

  createUser(request: CreateUserRequest): Observable<UserResponse>;

  updateUser(request: UpdateUserRequest): Observable<UserResponse>;

  deleteUser(request: DeleteUserRequest): Observable<UserResponse>;
}

export interface UserServiceController {
  getUser(request: GetUserRequest): Promise<UserResponse> | Observable<UserResponse> | UserResponse;

  getUsers(request: Empty): Promise<UsersResponse> | Observable<UsersResponse> | UsersResponse;

  createUser(request: CreateUserRequest): Promise<UserResponse> | Observable<UserResponse> | UserResponse;

  updateUser(request: UpdateUserRequest): Promise<UserResponse> | Observable<UserResponse> | UserResponse;

  deleteUser(request: DeleteUserRequest): Promise<UserResponse> | Observable<UserResponse> | UserResponse;
}

export function UserServiceControllerMethods() {
  return function (constructor: Function) {
    const grpcMethods: string[] = ["getUser", "getUsers", "createUser", "updateUser", "deleteUser"];
    for (const method of grpcMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcMethod("UserService", method)(constructor.prototype[method], method, descriptor);
    }
    const grpcStreamMethods: string[] = [];
    for (const method of grpcStreamMethods) {
      const descriptor: any = Reflect.getOwnPropertyDescriptor(constructor.prototype, method);
      GrpcStreamMethod("UserService", method)(constructor.prototype[method], method, descriptor);
    }
  };
}

export const USER_SERVICE_NAME = "UserService";
