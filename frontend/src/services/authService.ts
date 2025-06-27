import api from "@/lib/api";

export const register = async (formData: {
  name: string;
  email: string;
  password: string;
}) => {
  return api.post("/register", formData);
};

export const login = async (formData: {
  email: string;
  password: string;
}) => {
  return api.post("/login", formData);
};
