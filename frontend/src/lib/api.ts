import { PUBLIC_API_BASE_URL } from "$env/static/public";

const API_BASE = PUBLIC_API_BASE_URL.replace(/\/$/, "");

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    credentials: "include",
    headers: {
      ...(options.body instanceof FormData
        ? {}
        : { "Content-Type": "application/json" }),
      ...options.headers,
    },
  });

  if (!res.ok) {
    const error = await res.json().catch(() => ({ error: "Request failed" }));
    throw new Error(error.error || `HTTP ${res.status}`);
  }

  return res.json();
}

export const api = {
  // Auth
  verifyPin: (pin: string) =>
    request<{ message: string; user: any }>("/api/auth/verify-pin", {
      method: "POST",
      body: JSON.stringify({ pin }),
    }),
  logout: () => request("/api/auth/logout", { method: "POST" }),
  me: () => request<{ authenticated: boolean; user: any }>("/api/auth/me"),

  // Plans
  getPlans: () => request<any[]>("/api/plans"),
  createPlan: (data: {
    title: string;
    description?: string;
    category?: string;
  }) => request("/api/plans", { method: "POST", body: JSON.stringify(data) }),
  togglePlan: (id: number) =>
    request(`/api/plans/${id}/toggle`, { method: "PATCH" }),
  deletePlan: (id: number) => request(`/api/plans/${id}`, { method: "DELETE" }),

  // Gallery
  getGallery: () => request<any[]>("/api/gallery"),
  uploadImage: (formData: FormData) =>
    request("/api/gallery", { method: "POST", body: formData }),
  deleteGalleryItem: (id: number) =>
    request(`/api/gallery/${id}`, { method: "DELETE" }),

  // Secret Notes
  getNotes: () => request<any[]>("/api/secret-notes"),
  createNote: (data: {
    title: string;
    content: string;
    note_type?: string;
    visible_to?: string;
  }) =>
    request("/api/secret-notes", {
      method: "POST",
      body: JSON.stringify(data),
    }),

  // Image URL helper
  imageUrl: (fileName: string) => `${API_BASE}/uploads/${fileName}`,
};
