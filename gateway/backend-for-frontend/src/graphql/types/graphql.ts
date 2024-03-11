/*
 * -------------------------------------------------------
 * THIS FILE WAS AUTOMATICALLY GENERATED (DO NOT MODIFY)
 * -------------------------------------------------------
 */

/* tslint:disable */
/* eslint-disable */

export class Hello {
  message: string;
}

export abstract class IQuery {
  abstract hello(name: string): Hello | Promise<Hello>;

  abstract user(id: string): User | Promise<User>;

  abstract users(): User[] | Promise<User[]>;

  abstract _empty(): Nullable<string> | Promise<Nullable<string>>;
}

export class User {
  id: string;
  name: string;
  email: string;
  image: string;
  createdAt: string;
  updatedAt: string;
  deletedAt?: Nullable<string>;
}

export abstract class IMutation {
  abstract _empty(): Nullable<string> | Promise<Nullable<string>>;
}

export type Upload = any;
type Nullable<T> = T | null;
