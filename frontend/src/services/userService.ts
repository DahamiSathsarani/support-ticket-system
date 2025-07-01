import api from "@/lib/api";

export const getUser = async (token: string) => {
  return await api.get("/api/users/get", {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
