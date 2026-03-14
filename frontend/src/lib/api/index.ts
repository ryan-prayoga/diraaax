// ==========================================
// diraaax v2 — API endpoints
// ==========================================

import { get, post, patch, del } from "./client";
import type {
  AuthSession,
  TimelineEvent,
  Memory,
  BucketItem,
  LoveCapsule,
  CapsuleScene,
  Mood,
  LoveReason,
} from "$lib/types";

function asArray<T>(value: unknown): T[] {
  if (Array.isArray(value)) {
    return value as T[];
  }

  if (value && typeof value === "object") {
    const maybeItems = (value as { items?: unknown }).items;
    if (Array.isArray(maybeItems)) {
      return maybeItems as T[];
    }

    const maybeList = (value as { list?: unknown }).list;
    if (Array.isArray(maybeList)) {
      return maybeList as T[];
    }
  }

  return [];
}

function asObjectOrNull<T extends object>(value: unknown): T | null {
  if (value && typeof value === "object" && !Array.isArray(value)) {
    return value as T;
  }
  return null;
}

function normalizeTimelineEvent(event: any): TimelineEvent {
  return {
    ...event,
    event_date: event?.event_date ?? event?.date,
    event_type: event?.event_type ?? event?.type,
  } as TimelineEvent;
}

function normalizeMemory(memory: any): Memory {
  return {
    ...memory,
    memory_date: memory?.memory_date ?? memory?.date,
  } as Memory;
}

function normalizeCapsule(capsule: any): LoveCapsule {
  return {
    ...capsule,
    is_opened: Boolean(capsule?.is_opened),
    cover_image_url: capsule?.cover_image_url ?? capsule?.cover_url,
  } as LoveCapsule;
}

function normalizeBucketItem(item: any): BucketItem {
  const isDone = item?.is_done ?? item?.status === "done";
  return {
    ...item,
    is_done: Boolean(isDone),
  } as BucketItem;
}

function normalizeCapsuleScene(scene: any): CapsuleScene {
  const order = scene?.order_index ?? scene?.scene_order ?? 0;
  return {
    ...scene,
    order_index: Number(order),
    scene_order: Number(order),
  } as CapsuleScene;
}

function normalizeMood(mood: any): Mood {
  return {
    ...mood,
    date: mood?.date ?? mood?.mood_date,
  } as Mood;
}

function normalizeAuthSession(authSession: any): AuthSession {
  return {
    ...authSession,
    authenticated: Boolean(authSession?.session),
  } as AuthSession;
}

// ── Auth ──────────────────────────────────
export const auth = {
  me: async () => normalizeAuthSession(await get<AuthSession>("/api/auth/me")),
  verifyPin: (pin: string) =>
    post<AuthSession>("/api/auth/verify-pin", { pin }).then(
      normalizeAuthSession,
    ),
  logout: () => post<void>("/api/auth/logout"),
};

// ── Timeline ──────────────────────────────
export const timeline = {
  list: async () =>
    asArray<any>(await get<unknown>("/api/timeline")).map(
      normalizeTimelineEvent,
    ),
  create: (data: Partial<TimelineEvent>) =>
    post<any>("/api/timeline", data).then(normalizeTimelineEvent),
  remove: (id: number) => del<void>(`/api/timeline/${id}`),
};

// ── Memories ──────────────────────────────
export const memories = {
  list: async () =>
    asArray<any>(await get<unknown>("/api/memories")).map(normalizeMemory),
  random: async () => {
    const payload = await get<unknown>("/api/memories/random");
    const memory = asObjectOrNull<any>(payload);
    return memory ? normalizeMemory(memory) : null;
  },
  create: (data: FormData | Partial<Memory>) =>
    post<any>("/api/memories", data).then(normalizeMemory),
  remove: (id: number) => del<void>(`/api/memories/${id}`),
};

// ── Bucket List ───────────────────────────
export const bucketList = {
  list: async () =>
    asArray<any>(await get<unknown>("/api/bucket-list")).map(
      normalizeBucketItem,
    ),
  create: (data: Partial<BucketItem>) =>
    post<any>("/api/bucket-list", data).then(normalizeBucketItem),
  toggle: (id: number) =>
    patch<any>(`/api/bucket-list/${id}/toggle`).then(normalizeBucketItem),
  remove: (id: number) => del<void>(`/api/bucket-list/${id}`),
};

// ── Capsules ──────────────────────────────
export const capsules = {
  list: async () =>
    asArray<any>(await get<unknown>("/api/capsules")).map(normalizeCapsule),
  get: (id: number) => get<any>(`/api/capsules/${id}`).then(normalizeCapsule),
  create: (data: Partial<LoveCapsule>) =>
    post<any>("/api/capsules", data).then(normalizeCapsule),
  open: (id: number) =>
    post<any>(`/api/capsules/${id}/open`).then(normalizeCapsule),
  remove: (id: number) => del<void>(`/api/capsules/${id}`),
  scenes: (id: number) =>
    get<unknown>(`/api/capsules/${id}/scenes`).then((items) =>
      asArray<any>(items).map(normalizeCapsuleScene),
    ),
  addScene: (id: number, data: Partial<CapsuleScene>) =>
    post<any>(`/api/capsules/${id}/scenes`, data).then(normalizeCapsuleScene),
};

// ── Moods ─────────────────────────────────
export const moods = {
  list: async () =>
    asArray<any>(await get<unknown>("/api/moods")).map(normalizeMood),
  create: (data: { mood: string; note?: string }) =>
    post<any>("/api/moods", data).then(normalizeMood),
};

// ── Love Reasons ──────────────────────────
export const loveReasons = {
  list: async () =>
    asArray<LoveReason>(await get<unknown>("/api/love-reasons")),
  create: (data: { message: string }) =>
    post<LoveReason>("/api/love-reasons", data),
  remove: (id: number) => del<void>(`/api/love-reasons/${id}`),
};

// Re-export client utilities
export { imageUrl } from "./client";
