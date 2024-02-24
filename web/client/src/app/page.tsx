import { getClient } from "@/lib/client";
import { gql } from "@apollo/client";
import { GetHelloQuery } from "@/graphql/graphql";

export const GET_HELLO = gql`
  query getHello($name: String!) {
    hello(name: $name) {
      message
    }
  }
`;

export default async function Home() {
  const client = getClient();

  const { data, error, loading } = await client.query<GetHelloQuery>({
    query: GET_HELLO,
    variables: { name: "nextjs" },
  });
  if (error) return <b>Error: {error.message}</b>;
  if (loading) return <b>Loading...</b>;
  const hello = data?.hello;
  return (
    <div>
      <p>テスト</p>
      <p>{hello?.message}</p>
    </div>
  );
}
