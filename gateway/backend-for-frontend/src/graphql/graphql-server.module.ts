import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { Module } from '@nestjs/common';
import { GraphQLModule } from '@nestjs/graphql';
// import { ApolloServerPluginLandingPageLocalDefault } from '@apollo/server/plugin/landingPage/default';
import { join } from 'path/posix';

@Module({
  imports: [
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      typePaths: ['./src/**/*.graphql'],
      playground: true,
      definitions: {
        path: join(process.cwd(), 'src/graphql/types/graphql.ts'),
        outputAs: 'class',
      },
      //   useFactory: async () => ({
      //     typePaths: ['./src/**/*.graphql'],
      //     playground: false,
      //     definitions: {
      //       path: join(process.cwd(), 'src/graphql/types/graphql.ts'),
      //       outputAs: 'class',
      //     },
      //     plugins: [ApolloServerPluginLandingPageLocalDefault()],
      //   }),
    }),
  ],
})
export class GraphQLServerModule {}
