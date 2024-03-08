"use client";

import { useRouter } from "next/navigation";

import { Button } from "@/components/ui/button";

export const NotLoggedMenu = () => {
  const router = useRouter();
  return (
    <>
      <Button
        variant={"outline"}
        onClick={() => {
          router.push("/login");
        }}
      >
        SignIn
      </Button>
      <Button>Menu</Button>
    </>
  );
};
