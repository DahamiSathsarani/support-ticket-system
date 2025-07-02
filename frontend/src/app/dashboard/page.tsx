"use client";

import { useEffect, useState } from "react";
import { fetchUserStats } from "@/services/ticketService";
import { useRouter } from "next/navigation";

export default function DashboardPage() {
  const router = useRouter();
  const [stats, setStats] = useState({
    total: 0,
    pending: 0,
    resolved: 0,
    rejected: 0,
  });

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) return router.push("/login");

    fetchUserStats(token)
      .then((res) => setStats(res.data))
      .catch(() => router.push("/login"));
  }, []);

  return (
    <div className="p-6 space-y-6 min-h-screen bg-gray-100">
      <div className="bg-white rounded-2xl shadow-lg p-6">
        <h2 className="text-3xl font-bold text-indigo-700">Welcome Back 👋</h2>
        <p className="text-gray-600 text-lg mt-2">
          Here's an overview of your ticket activity. Track and manage your support requests easily.
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        <StatCard label="🎫 Total Tickets" value={stats.total} color="text-indigo-600" />
        <StatCard label="⏳ Pending" value={stats.pending} color="text-yellow-500" />
        <StatCard label="✅ Resolved" value={stats.resolved} color="text-green-500" />
        <StatCard label="❌ Rejected" value={stats.rejected} color="text-red-500" />
      </div>
    </div>
  );
}

function StatCard({ label, value, color }: { label: string; value: number; color: string }) {
  return (
    <div className="bg-white rounded-2xl shadow-md p-5 hover:shadow-xl transition">
      <p className="text-gray-500">{label}</p>
      <h3 className={`text-3xl font-semibold mt-2 ${color}`}>{value}</h3>
    </div>
  );
}
