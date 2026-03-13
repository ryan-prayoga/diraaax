// ==========================================
// diraaax v2 — Utility Functions
// ==========================================

/**
 * Format date to a readable string
 */
export function formatDate(dateStr: string, style: 'short' | 'long' | 'relative' = 'long'): string {
  const date = new Date(dateStr);

  if (style === 'relative') {
    return getRelativeTime(date);
  }

  if (style === 'short') {
    return date.toLocaleDateString('id-ID', {
      day: 'numeric',
      month: 'short',
      year: 'numeric'
    });
  }

  return date.toLocaleDateString('id-ID', {
    weekday: 'long',
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  });
}

/**
 * Get relative time string
 */
function getRelativeTime(date: Date): string {
  const now = new Date();
  const diffMs = now.getTime() - date.getTime();
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));

  if (diffDays === 0) return 'Hari ini';
  if (diffDays === 1) return 'Kemarin';
  if (diffDays < 7) return `${diffDays} hari lalu`;
  if (diffDays < 30) return `${Math.floor(diffDays / 7)} minggu lalu`;
  if (diffDays < 365) return `${Math.floor(diffDays / 30)} bulan lalu`;
  return `${Math.floor(diffDays / 365)} tahun lalu`;
}

/**
 * Calculate days between two dates
 */
export function daysBetween(date1: string, date2: string): number {
  const d1 = new Date(date1);
  const d2 = new Date(date2);
  const diffMs = Math.abs(d2.getTime() - d1.getTime());
  return Math.floor(diffMs / (1000 * 60 * 60 * 24));
}

/**
 * Calculate anniversary details
 */
export function getAnniversaryInfo(anniversaryDate: string) {
  const anniversary = new Date(anniversaryDate);
  const now = new Date();

  const totalDays = Math.floor((now.getTime() - anniversary.getTime()) / (1000 * 60 * 60 * 24));
  const totalMonths = (now.getFullYear() - anniversary.getFullYear()) * 12 + (now.getMonth() - anniversary.getMonth());
  const years = Math.floor(totalMonths / 12);
  const months = totalMonths % 12;
  const days = now.getDate() - anniversary.getDate();

  return {
    totalDays: Math.max(0, totalDays),
    years,
    months,
    days: days < 0 ? 0 : days,
    formattedDate: formatDate(anniversaryDate, 'long')
  };
}

/**
 * Check if a capsule can be opened
 */
export function canOpenCapsule(openDate: string): boolean {
  return new Date() >= new Date(openDate);
}

/**
 * Truncate text with ellipsis
 */
export function truncate(text: string, maxLength: number): string {
  if (text.length <= maxLength) return text;
  return text.slice(0, maxLength).trimEnd() + '...';
}

/**
 * Stagger delay for animations
 */
export function staggerDelay(index: number, baseMs: number = 100): string {
  return `${index * baseMs}ms`;
}

/**
 * Shuffle an array (Fisher-Yates)
 */
export function shuffle<T>(array: T[]): T[] {
  const arr = [...array];
  for (let i = arr.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [arr[i], arr[j]] = [arr[j], arr[i]];
  }
  return arr;
}
