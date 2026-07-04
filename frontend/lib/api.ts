import { HealthResponse } from "@/types/health";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export async function getHealth(): Promise<HealthResponse> {
    const response = await fetch(`${API_URL}/health`);
    if(!response.ok) {
        throw  new Error("Failed to fetch health status");
    }

    return response.json();
}