// ==========================================
// diraaax v2 — Auth Store (Svelte 5 runes)
// ==========================================

import { writable } from "svelte/store";
import { auth } from "$lib/api";
import type { User } from "$lib/types";

export const user = writable<User | null>(null);
export const isAuthenticated = writable(false);
export const authLoading = writable(true);

export async function checkAuth(): Promise<boolean> {
  try {
    const response = (await auth.me()) as any;
    if (response?.authenticated) {
      user.set(response.user ?? null);
      isAuthenticated.set(true);
      authLoading.set(false);
      return true;
    }
  } catch {
    // not authenticated
  }
  user.set(null);
  isAuthenticated.set(false);
  authLoading.set(false);
  return false;
}

export async function logout(): Promise<void> {
  try {
    await auth.logout();
  } catch {
    // ignore
  }
  user.set(null);
  isAuthenticated.set(false);
}
