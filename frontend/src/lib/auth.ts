import { writable } from "svelte/store";
import { api } from "./api";

export const user = writable<any>(null);
export const isAuthenticated = writable(false);
export const authLoading = writable(true);

export async function checkAuth(): Promise<boolean> {
  try {
    const data = await api.me();
    if (data.authenticated) {
      user.set(data.user);
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
