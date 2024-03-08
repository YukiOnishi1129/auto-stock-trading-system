"use client";
import { Button } from "@/components/ui/button";
import { signOut } from "next-auth/react";

export const LoggedMenu = () => {
  return <Button onClick={() => signOut()}>Logout</Button>;
};
