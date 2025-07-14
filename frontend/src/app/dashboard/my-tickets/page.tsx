"use client";
import { useEffect, useState } from "react";
import { getMyTickets, createTicket, updateTicket, deleteTicket } from "@/services/ticketService";
import { useRouter } from "next/navigation";
import { Dialog } from "@headlessui/react";

export default function MyTicketsPage() {
  const router = useRouter();
  const [tickets, setTickets] = useState<any[]>([]);
  const [form, setForm] = useState({ title: "", description: "" });
  const [editingId, setEditingId] = useState<number | null>(null);
  const [isOpen, setIsOpen] = useState(false); 

  const token = typeof window !== "undefined" ? localStorage.getItem("token") : "";

  useEffect(() => {
    if (!token) return router.push("/login");
    loadTickets();
  }, []);

  const loadTickets = async () => {
    const res = await getMyTickets(token!);
    setTickets(res.data);
  };

  const handleSubmit = async () => {
    if (!form.title || !form.description) return;
    if (editingId) {
      await updateTicket(editingId, form, token!);
    } else {
      await createTicket(form, token!);
    }
    setForm({ title: "", description: "" });
    setEditingId(null);
    setIsOpen(false)
    loadTickets();
  };

  const handleEdit = (ticket: any) => {
    setIsOpen(true)
    setForm({ title: ticket.title, description: ticket.description });
    setEditingId(ticket.ID);
  };

  const handleDelete = async (id: number) => {
    await deleteTicket(id, token!);
    loadTickets();
  };

  return (
    <div className="p-6 space-y-6 min-h-screen bg-gray-100">
      <div className="bg-white shadow-md rounded-2xl p-6">
        <h2 className="text-2xl font-semibold text-indigo-700">My Tickets</h2>
        <p className="text-gray-500">Manage your support tickets</p>

        <button
          onClick={() => setIsOpen(true)}
          className="bg-indigo-600 text-white px-4 py-2 rounded hover:bg-indigo-700 mt-5"
        >
          + Create New Ticket
        </button>
      </div>

      <div className="bg-white shadow rounded-xl p-6">
        <h3 className="text-lg font-semibold text-gray-700 mb-4">Your Tickets</h3>
        {tickets.length === 0 ? (
          <p className="text-gray-500">No tickets found.</p>
        ) : (
          <div className="space-y-4">
            {tickets.map((ticket) => (
              <div
                key={ticket.ID}
                className="border rounded-lg p-4 flex justify-between items-start bg-gray-50 hover:shadow"
              >
                <div>
                  <h4 className="font-bold text-indigo-700">{ticket.title}</h4>
                  <p className="text-gray-600">{ticket.description}</p>
                  <p className="text-sm text-gray-400 mt-1">Status: {ticket.status}</p>
                </div>
                <div className="space-x-2">
                  <button
                    onClick={() => handleEdit(ticket)}
                    className="border-solid border-2 border-yellow-400 text-yellow-400 px-3 py-1 rounded-md hover:bg-yellow-400 hover:text-white"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(ticket.ID)}
                    className="bg-red-600 text-white px-3 py-1 rounded hover:bg-red-700"
                  >
                    Delete
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      <Dialog open={isOpen} onClose={() => setIsOpen(false)} className="relative z-50">
        <div className="fixed inset-0 bg-black/30" aria-hidden="true" />
        <div className="fixed inset-0 flex items-center justify-center p-4">
          <Dialog.Panel className="w-full max-w-md bg-white rounded-xl p-6 shadow-lg space-y-4">
            <Dialog.Title className="text-lg font-semibold text-indigo-700">Create New Ticket</Dialog.Title>

            <div>
              <label className="block text-sm font-medium text-gray-700">Title</label>
              <input
                name="title"
                value={form.title}
                onChange={(e) => setForm({ ...form, title: e.target.value })}
                className="w-full border border-gray-300 rounded-md p-2 placeholder:text-gray-200 text-black focus:outline-none focus:ring-1 focus:ring-indigo-600"
                placeholder="Enter ticket title"
                required
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700">Description</label>
              <textarea
                name="description"
                value={form.description}
                onChange={(e) => setForm({ ...form, description: e.target.value })}
                className="w-full border border-gray-300 rounded-md p-2 placeholder:text-gray-200 text-black focus:outline-none focus:ring-1 focus:ring-indigo-600"
                placeholder="Describe the issue"
                rows={4}
                required
              />
            </div>

            <div className="flex justify-end space-x-2 pt-2">
              <button
                onClick={() => setIsOpen(false)}
                className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
              >
                Cancel
              </button>
              <button
                onClick={handleSubmit}
                className="px-4 py-2 bg-indigo-600 text-white rounded hover:bg-indigo-700"
              >
                Create
              </button>
            </div>
          </Dialog.Panel>
        </div>
      </Dialog>
    </div>
  );
}
