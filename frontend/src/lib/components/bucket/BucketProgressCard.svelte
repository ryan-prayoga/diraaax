<script lang="ts">
	import type { BucketItem } from '$lib/types';

	let {
		items = []
	}: {
		items: BucketItem[];
	} = $props();

		const total = $derived(items.length);
		const done = $derived(items.filter(i => i.is_done).length);
		const percent = $derived(total > 0 ? Math.round((done / total) * 100) : 0);
</script>

<div class="love-card-static p-5 bg-linear-to-br from-pink-50 to-white">
	<div class="flex items-center justify-between mb-3">
		<div>
			<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold">Progress Kita</p>
			<p class="text-2xl font-extrabold text-pink-500">{done}/{total}</p>
		</div>
		<div class="text-3xl">
			{#if percent >= 100}
				🏆
			{:else if percent >= 50}
				🌟
			{:else}
				✨
			{/if}
		</div>
	</div>

	<!-- Progress bar -->
	<div class="h-3 bg-pink-100 rounded-full overflow-hidden">
		<div
			class="h-full bg-linear-to-r from-pink-400 to-pink-500 rounded-full transition-all duration-1000 ease-out"
			style="width: {percent}%"
		></div>
	</div>
	<p class="text-xs text-rose-muted mt-2 text-right">{percent}% tercapai 💪</p>
</div>
