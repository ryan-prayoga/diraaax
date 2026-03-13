<script lang="ts">
	import type { MoodValue } from '$lib/types';
	import { MOOD_EMOJI, MOOD_LABELS, MOOD_COLORS } from '$lib/types';

	let {
		selected = null as MoodValue | null,
		onselect
	}: {
		selected?: MoodValue | null;
		onselect?: (mood: MoodValue) => void;
	} = $props();

	const allMoods: MoodValue[] = [
		'happy', 'excited', 'calm', 'sleepy', 'sad',
		'miss_you', 'love', 'angry', 'neutral'
	];
</script>

<div class="love-card-static p-5">
	<p class="text-sm font-bold text-rose-deep mb-3">Gimana perasaanmu hari ini? 💕</p>
	<div class="grid grid-cols-3 gap-2">
		{#each allMoods as mood}
			<button
				onclick={() => onselect?.(mood)}
				class="flex flex-col items-center gap-1 p-3 rounded-2xl border-2 transition-all
					{selected === mood
						? 'border-pink-400 bg-pink-50 scale-105 shadow-md'
						: 'border-transparent hover:bg-pink-50/50 hover:border-pink-100'}"
			>
				<span class="text-2xl">{MOOD_EMOJI[mood]}</span>
				<span class="text-[10px] font-medium text-rose-muted">{MOOD_LABELS[mood]}</span>
			</button>
		{/each}
	</div>
</div>
