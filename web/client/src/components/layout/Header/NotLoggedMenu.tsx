"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

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
