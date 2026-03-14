<script lang="ts">
	import type { Mood, MoodValue } from '$lib/types';
	import { MOOD_EMOJI, MOOD_LABELS, MOOD_COLORS } from '$lib/types';
	import { formatDate } from '$lib/utils';

	let {
		moods = []
	}: {
		moods: Mood[];
	} = $props();

	const safeMoods = $derived(Array.isArray(moods) ? moods : []);
</script>

{#if safeMoods.length > 0}
	<div class="space-y-2">
		{#each safeMoods as mood, i}
			<div
				class="flex items-center gap-3 love-card-static p-3 animate-fade-in-up"
				style="animation-delay: {i * 80}ms"
			>
				<span class="text-2xl">{MOOD_EMOJI[mood.mood]}</span>
				<div class="flex-1 min-w-0">
					<div class="flex items-center gap-2">
						<span class="text-sm font-semibold text-rose-deep">{MOOD_LABELS[mood.mood]}</span>
						{#if mood.user}
							<span class="text-xs text-rose-muted">— {mood.user.nickname || mood.user.display_name}</span>
						{/if}
					</div>
						<p class="text-xs text-rose-muted">{formatDate(mood.date || mood.mood_date || new Date().toISOString(), 'relative')}</p>
					{#if mood.note}
						<p class="text-xs text-rose-deep/70 mt-1">{mood.note}</p>
					{/if}
				</div>
			</div>
		{/each}
	</div>
{:else}
	<p class="text-center text-rose-muted text-sm py-8">Belum ada mood yang dicatat 💭</p>
{/if}
