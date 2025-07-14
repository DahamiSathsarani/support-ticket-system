"use client";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getUser } from "@/services/userService";
import { FiHome, FiUser, FiTool, FiTag, FiMenu } from "react-icons/fi";

export default function Sidebar() {
  const pathname = usePathname();
  const router = useRouter();
  const [user, setUser] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login");
      return;
    }

    getUser(token)
      .then((res) => {
        setUser(res.data);
        setLoading(false);
      })
      .catch(() => {
        localStorage.removeItem("token");
        router.push("/login");
      });
  }, []);

  const navItem = (href: string, label: string, Icon: any) => (
    <Link
      href={href}
      className={`flex items-center gap-3 px-4 py-2 rounded-md font-medium transition ${
        pathname === href
          ? "bg-white text-indigo-600 shadow"
          : "text-white hover:bg-indigo-400/30"
      }`}
    >
      <Icon /> {label}
    </Link>
  );

  return (
    <aside className="min-h-screen w-64 bg-gradient-to-b from-indigo-700 to-blue-600 text-white p-6 space-y-6 rounded-r-3xl shadow-xl hidden md:block">
      {!loading && user && (
        <div>
          <div className="text-center">
            <div className="w-16 h-16 mx-auto rounded-full bg-white/20 flex items-center justify-center text-2xl font-bold">
              {user?.name?.charAt(0).toUpperCase()}
            </div>
            <p className="mt-2 font-semibold">{user?.name || "User"}</p>
            <p className="text-sm opacity-70 capitalize">{user.role}</p>
          </div>
        
          <nav className="space-y-3 mt-8">
            {user.role === "user" && navItem("/dashboard", "Dashboard", FiHome)}
            {user.role === "user" && navItem("/dashboard/my-tickets", "My Tickets", FiTag)}
            {user.role === "admin" && navItem("/admin", "Admin Panel", FiTool)}
            {user.role === "agent" && navItem("/assigned", "Assigned Tickets", FiTool)}
            {navItem("/profile", "Profile", FiUser)}
          </nav>
        </div>
      )}

    </aside>
  );
}
