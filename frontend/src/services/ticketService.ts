import api from "@/lib/api";

export const fetchUserStats = async (token: string) => {
  return api.get("/api/tickets/get-tickets-stats", {
    headers: { Authorization: `Bearer ${token}` },
  });
};

export const getMyTickets = async (token: string) => {
  return api.get("/api/tickets/get-my", {
    headers: { Authorization: `Bearer ${token}` },
  });
}

export const createTicket = async (data: any, token: string) =>
  api.post("/api/tickets/create", data, {
    headers: { Authorization: `Bearer ${token}` },
  });

export const updateTicket = async (id: number, data: any, token: string) =>
  api.put(`/api/tickets/update/${id}`, data, {
    headers: { Authorization: `Bearer ${token}` },
  });

export const deleteTicket = async (id: number, token: string) =>
  api.delete(`/api/tickets/delete/${id}`, {
    headers: { Authorization: `Bearer ${token}` },
  });
