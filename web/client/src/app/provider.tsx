"use client";

import { SessionProvider } from "next-auth/react";
import { ApolloWrapper } from "@/lib/apollo-wrapper";

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <ApolloWrapper>
      <SessionProvider>{children}</SessionProvider>
    </ApolloWrapper>
  );
}
