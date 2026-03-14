<script lang="ts">
	import type { LoveReason } from '$lib/types';

	let {
		reason,
		index = 0
	}: {
		reason: LoveReason;
		index?: number;
	} = $props();

	const hearts = ['💕', '💗', '💖', '💝', '🩷', '💘', '💓', '♥️'];
	const heart = $derived(hearts[index % hearts.length]);
</script>

<div class="love-card-static p-4 md:p-5 animate-fade-in-up" style="animation-delay: {index * 80}ms">
	<div class="flex items-start gap-3">
		<span class="text-xl shrink-0 mt-0.5">{heart}</span>
		<p class="text-rose-deep text-sm leading-relaxed">{reason.message}</p>
	</div>
</div>
