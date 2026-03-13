// ==========================================
// diraaax v2 — API endpoints
// ==========================================

import { get, post, patch, del } from './client';
import type {
  AuthResponse,
  User,
  TimelineEvent,
  Memory,
  BucketItem,
  LoveCapsule,
  CapsuleScene,
  Mood,
  LoveReason
} from '$lib/types';

// ── Auth ──────────────────────────────────
export const auth = {
  me: () => get<AuthResponse>('/api/auth/me'),
  verifyPin: (pin: string) =>
    post<{ message: string; user: User }>('/api/auth/verify-pin', { pin }),
  logout: () => post<void>('/api/auth/logout')
};

// ── Timeline ──────────────────────────────
export const timeline = {
  list: () => get<TimelineEvent[]>('/api/timeline'),
  create: (data: Partial<TimelineEvent>) =>
    post<TimelineEvent>('/api/timeline', data),
  remove: (id: number) => del<void>(`/api/timeline/${id}`)
};

// ── Memories ──────────────────────────────
export const memories = {
  list: () => get<Memory[]>('/api/memories'),
  random: () => get<Memory>('/api/memories/random'),
  create: (data: FormData) => post<Memory>('/api/memories', data),
  remove: (id: number) => del<void>(`/api/memories/${id}`)
};

// ── Bucket List ───────────────────────────
export const bucketList = {
  list: () => get<BucketItem[]>('/api/bucket-list'),
  create: (data: Partial<BucketItem>) =>
    post<BucketItem>('/api/bucket-list', data),
  toggle: (id: number) =>
    patch<BucketItem>(`/api/bucket-list/${id}/toggle`),
  remove: (id: number) => del<void>(`/api/bucket-list/${id}`)
};

// ── Capsules ──────────────────────────────
export const capsules = {
  list: () => get<LoveCapsule[]>('/api/capsules'),
  get: (id: number) => get<LoveCapsule>(`/api/capsules/${id}`),
  create: (data: Partial<LoveCapsule>) =>
    post<LoveCapsule>('/api/capsules', data),
  open: (id: number) =>
    post<LoveCapsule>(`/api/capsules/${id}/open`),
  scenes: (id: number) =>
    get<CapsuleScene[]>(`/api/capsules/${id}/scenes`),
  addScene: (id: number, data: Partial<CapsuleScene>) =>
    post<CapsuleScene>(`/api/capsules/${id}/scenes`, data)
};

// ── Moods ─────────────────────────────────
export const moods = {
  list: () => get<Mood[]>('/api/moods'),
  create: (data: { mood: string; note?: string }) =>
    post<Mood>('/api/moods', data)
};

// ── Love Reasons ──────────────────────────
export const loveReasons = {
  list: () => get<LoveReason[]>('/api/love-reasons'),
  create: (data: { message: string }) =>
    post<LoveReason>('/api/love-reasons', data),
  remove: (id: number) => del<void>(`/api/love-reasons/${id}`)
};

// Re-export client utilities
export { imageUrl } from './client';
