import Image from "next/image";
import { gql, useQuery } from "@apollo/client";

export const GET_HELLO = gql`
  query getHello($name: String!) {
    hello(name: $name) {
      message
    }
  }
`;

export default function Home() {
  return <div>フロント</div>;
}
