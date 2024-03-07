"use client";

import { Button } from "@/components/ui/button";
import { signIn } from "next-auth/react";

export default function Login() {
  return (
    <>
      <div>ログインページ</div>
      <Button onClick={() => signIn("google", { callbackUrl: "/" })}>
        ログイン
      </Button>
    </>
  );
}
