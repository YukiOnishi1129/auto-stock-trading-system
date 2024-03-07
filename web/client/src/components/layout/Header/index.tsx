import { Button } from "@/components/ui/button";
import Link from "next/link";

export const Header = () => {
  return (
    <div className="fixed flex justify-between px-8 w-screen h-16 bg-white shadow-md items-center z-10">
      <h1 className="font-bold text-2xl">
        <Link href="/">auto trading system</Link>
      </h1>
      <div className="flex gap-3">
        <Button variant={"outline"}>SignIn</Button>
        <Button>Menu</Button>
      </div>
    </div>
  );
};
