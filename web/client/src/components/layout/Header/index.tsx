import Link from "next/link";
import { getServerSession } from "next-auth/next";

import { LoggedMenu } from "@/components/layout/Header/LoggedMenu";
import { NotLoggedMenu } from "@/components/layout/Header/NotLoggedMenu";
import { authOptions } from "@/lib/auth";

export const Header = async () => {
  const session = await getServerSession(authOptions);
  return (
    <div className="fixed flex justify-between px-8 w-screen h-16 bg-white shadow-md items-center z-10">
      <h1 className="font-bold text-2xl">
        <Link href="/">auto trading system</Link>
      </h1>
      <div className="flex gap-3">
        {session ? <LoggedMenu /> : <NotLoggedMenu />}
      </div>
    </div>
  );
};
