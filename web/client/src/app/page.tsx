import { getClient } from "@/lib/client";
import { gql } from "@apollo/client";
import { GetHomeQuery, Query } from "@/graphql/graphql";

export const GET_HOME = gql`
  query getHome($name: String!) {
    hello(name: $name) {
      message
    }
    users {
      id
      name
      email
      createdAt
      updatedAt
      deletedAt
    }
  }
`;

export default async function Home() {
  const client = getClient();

  const { data, error, loading } = await client.query<GetHomeQuery>({
    query: GET_HOME,
    variables: { name: "nextjs" },
  });
  if (error) return <b>Error: {error.message}</b>;
  if (loading) return <b>Loading...</b>;
  const hello = data?.hello;
  console.log(new Date(data?.users[0].createdAt.seconds.low * 1000));
  return (
    <div>
      <p>テスト</p>
      <p>{hello?.message}</p>
      {data?.users.length &&
        data?.users.map((user) => (
          <div key={user.id}>
            <p>{user.name}</p>
            <p>{user.email}</p>
            {/* <p>{user.createdAt}</p> */}
            {/* <p>{user.updatedAt}</p> */}
          </div>
        ))}
    </div>
  );
}
