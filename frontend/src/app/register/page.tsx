"use client";
import { useState } from "react";
import { register } from "@/services/authService";
import { useRouter } from "next/navigation";

export default function RegisterPage() {
    const router = useRouter();
    const [form, setForm] = useState({ name: "", email: "", password: "" });
    const [error, setError] = useState("");

    const handleChange = (e: any) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleSubmit = async () => {
        try {
            console.log("Formdata", form);
            await register(form);
            router.push("/login");
        } catch (err: any) {
            setError(err.response?.data?.error || "Registration failed");
        }
    };

    return (
        <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-100 to-indigo-200 px-4">
        <div className="bg-white shadow-xl rounded-2xl p-8 sm:p-10 w-full max-w-md">
            <h2 className="text-3xl font-bold text-center text-indigo-700 mb-6">Create Account</h2>

            {error && <p className="text-red-600 text-sm text-center mb-4">{error}</p>}

            <div className="space-y-4">
            <div>
                <label className="block text-sm font-medium text-gray-600 mb-1">Name</label>
                <input
                type="text"
                name="name"
                onChange={handleChange}
                className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400 text-gray-700"
                placeholder="Jane Doe"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-600 mb-1">Email</label>
                <input
                type="email"
                name="email"
                onChange={handleChange}
                className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400 text-gray-700"
                placeholder="jane@example.com"
                />
            </div>

            <div>
                <label className="block text-sm font-medium text-gray-600 mb-1">Password</label>
                <input
                type="password"
                name="password"
                onChange={handleChange}
                className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-400 text-gray-700"
                placeholder="••••••••"
                />
            </div>

            <button
                onClick={handleSubmit}
                className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-2 rounded-md transition duration-200"
            >
                Register
            </button>
            </div>

            <p className="text-sm text-center text-gray-600 mt-6">
            Already have an account?{" "}
            <a href="/login" className="text-indigo-600 hover:underline font-medium">
                Log in
            </a>
            </p>
        </div>
        </div>
    );
}
