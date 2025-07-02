import api from "@/lib/api";

export const fetchUserStats = async (token: string) => {
  return api.get("/api/tickets/get-tickets-stats", {
    headers: { Authorization: `Bearer ${token}` },
  });
};
