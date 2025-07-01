"use client";
import { useRouter } from "next/navigation";
import { FiLogOut, FiBell } from "react-icons/fi";

export default function Header() {
  const router = useRouter();

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    router.push("/login");
  };

  return (
    <header className="bg-white shadow-md rounded-b-2xl px-6 py-4 flex justify-between items-center sticky top-0 z-50">
      <h1 className="text-2xl font-bold text-indigo-700">Dashboard</h1>
      <div className="flex items-center gap-6">
        <button className="text-indigo-600 text-xl hover:text-indigo-800 transition">
          <FiBell />
        </button>
        <button
          onClick={handleLogout}
          className="flex items-center gap-2 text-red-600 hover:text-red-700 hover:underline font-medium transition"
        >
          <FiLogOut />
          Logout
        </button>
      </div>
    </header>
  );
}
