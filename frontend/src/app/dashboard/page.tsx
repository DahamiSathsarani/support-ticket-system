"use client";

export default function DashboardPage() {
  return (
    <div className="p-6 grid grid-cols-1 gap-6">
      <div className="bg-white rounded-2xl shadow-xl p-6">
        <h2 className="text-2xl font-semibold text-indigo-700">Welcome</h2>
        <p className="text-lg text-gray-600 mt-2">This is your support dashboard. Use the menu to manage tickets.</p>
      </div>

      <div className="bg-gradient-to-br from-indigo-500 to-blue-500 text-white rounded-2xl shadow-xl p-6">
        <h3 className="text-xl font-semibold">Quick Stats</h3>
        <div className="mt-4 space-y-3 text-base">
          <p>📝 Open Tickets: <span className="font-bold">3</span></p>
          <p>✅ Closed Tickets: <span className="font-bold">5</span></p>
          <p>📥 Assigned to You: <span className="font-bold">2</span></p>
        </div>
      </div>
    </div>
  );
}
