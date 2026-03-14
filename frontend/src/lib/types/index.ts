// ==========================================
// diraaax v2 — Type Definitions
// ==========================================

export interface User {
  id: number;
  display_name: string;
  nickname?: string;
  slug?: string;
  avatar_url?: string;
  is_active?: boolean;
  created_at: string;
  updated_at?: string;
}

export interface Session {
  id: number;
  user_id?: number;
  access_code_id?: number;
  expires_at: string;
  last_seen_at?: string;
  created_at: string;
}

export interface AuthSession {
  authenticated: boolean;
  session: Session;
  user?: User;
}

export interface TimelineEvent {
  id: number;
  title: string;
  description: string;
  event_date: string;
  event_type?: string;
  image_url?: string;
  created_by: number;
  created_at: string;
}

export interface Memory {
  id: number;
  title?: string;
  description?: string;
  image_url?: string;
  memory_date?: string;
  created_by?: number;
  is_featured?: boolean;
  is_memory_of_day_enabled?: boolean;
  created_at: string;
  updated_at?: string;
}

export interface BucketItem {
  id: number;
  title: string;
  description: string;
  category?: string;
  is_done: boolean;
  status?: "pending" | "done" | string;
  target_date?: string;
  created_by?: number;
  completed_by?: number;
  completed_at?: string;
  sort_order?: number;
  created_at: string;
  updated_at?: string;
}

export interface LoveCapsule {
  id: number;
  title: string;
  message?: string;
  open_date: string;
  is_opened: boolean;
  cover_url?: string;
  cover_image_url?: string;
  music_url?: string;
  theme_variant?: string;
  visible_to?: string;
  created_by?: number;
  created_at: string;
  updated_at?: string;
}

export interface CapsuleScene {
  id: number;
  capsule_id: number;
  scene_type: "intro" | "photo" | "message" | "quote" | "ending";
  title?: string;
  content?: string;
  image_url?: string;
  order_index: number;
  scene_order?: number;
  animation_key?: string;
  created_at?: string;
  updated_at?: string;
}

export interface Mood {
  id: number;
  user_id: number;
  mood: MoodValue;
  note?: string;
  date: string;
  mood_date?: string;
  created_at?: string;
  user?: User;
}

export type MoodValue =
  | "happy"
  | "excited"
  | "calm"
  | "sleepy"
  | "sad"
  | "miss_you"
  | "love"
  | "angry"
  | "neutral";

export interface LoveReason {
  id: number;
  message: string;
  created_by?: number;
  visible_to?: string;
  is_pinned?: boolean;
  created_at: string;
  updated_at?: string;
}

export const MOOD_EMOJI: Record<MoodValue, string> = {
  happy: "😊",
  excited: "🤩",
  calm: "😌",
  sleepy: "😴",
  sad: "😢",
  miss_you: "🥺",
  love: "🥰",
  angry: "😤",
  neutral: "😐",
};

export const MOOD_LABELS: Record<MoodValue, string> = {
  happy: "Happy",
  excited: "Excited",
  calm: "Calm",
  sleepy: "Sleepy",
  sad: "Sad",
  miss_you: "Miss You",
  love: "In Love",
  angry: "Angry",
  neutral: "Neutral",
};

export const MOOD_COLORS: Record<MoodValue, string> = {
  happy: "bg-yellow-100 text-yellow-700 border-yellow-200",
  excited: "bg-orange-100 text-orange-700 border-orange-200",
  calm: "bg-green-100 text-green-700 border-green-200",
  sleepy: "bg-indigo-100 text-indigo-700 border-indigo-200",
  sad: "bg-blue-100 text-blue-700 border-blue-200",
  miss_you: "bg-purple-100 text-purple-700 border-purple-200",
  love: "bg-pink-100 text-pink-700 border-pink-200",
  angry: "bg-red-100 text-red-700 border-red-200",
  neutral: "bg-gray-100 text-gray-700 border-gray-200",
};
