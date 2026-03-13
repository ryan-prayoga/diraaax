// ==========================================
// diraaax v2 — API Client
// ==========================================
// Uses same-origin relative URLs for production.
// In development, set PUBLIC_API_BASE_URL in .env to proxy to backend.

import { PUBLIC_API_BASE_URL } from '$env/static/public';

const API_BASE = (PUBLIC_API_BASE_URL || '').replace(/\/$/, '');

export class ApiError extends Error {
  status: number;
  code?: string;
  constructor(message: string, status: number, code?: string) {
    super(message);
    this.status = status;
    this.code = code;
    this.name = 'ApiError';
  }
}

async function request<T>(path: string, options: RequestInit = {}): Promise<T> {
  const url = `${API_BASE}${path}`;

  const res = await fetch(url, {
    ...options,
    credentials: 'include',
    headers: {
      ...(options.body instanceof FormData
        ? {}
        : { 'Content-Type': 'application/json' }),
      ...options.headers
    }
  });

  if (!res.ok) {
    const errorData = await res.json().catch(() => ({ error: 'Request failed' }));
    let errorMessage = `HTTP ${res.status}`;
    let errorCode = undefined;
    
    if (errorData && typeof errorData.error === 'object') {
      errorMessage = errorData.error.message || errorMessage;
      errorCode = errorData.error.code;
    } else if (errorData && typeof errorData.error === 'string') {
      errorMessage = errorData.error;
    } else if (errorData && errorData.message) {
      errorMessage = errorData.message;
    }

    throw new ApiError(errorMessage, res.status, errorCode);
  }

  // Handle 204 No Content
  if (res.status === 204) {
    return undefined as T;
  }

  return res.json();
}

export function get<T>(path: string): Promise<T> {
  return request<T>(path);
}

export function post<T>(path: string, body?: unknown): Promise<T> {
  return request<T>(path, {
    method: 'POST',
    body: body instanceof FormData ? body : JSON.stringify(body)
  });
}

export function patch<T>(path: string, body?: unknown): Promise<T> {
  return request<T>(path, {
    method: 'PATCH',
    body: body ? JSON.stringify(body) : undefined
  });
}

export function del<T>(path: string): Promise<T> {
  return request<T>(path, { method: 'DELETE' });
}

export function imageUrl(fileName: string): string {
  return `${API_BASE}/uploads/${fileName}`;
}
