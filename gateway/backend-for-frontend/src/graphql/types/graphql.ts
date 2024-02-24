
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
}

export class User {
    name: string;
    email: string;
}

type Nullable<T> = T | null;
