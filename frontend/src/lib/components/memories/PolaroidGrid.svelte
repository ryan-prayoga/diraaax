<script lang="ts">
	import type { Memory } from '$lib/types';
	import MemoryCard from './MemoryCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';

	let {
		memories = []
	}: {
		memories: Memory[];
	} = $props();
</script>

{#if memories.length === 0}
	<EmptyState
		emoji="📸"
		title="Belum ada kenangan"
		subtitle="Tambahkan foto kenangan kalian di sini"
	/>
{:else}
	<div class="grid grid-cols-2 md:grid-cols-3 gap-4">
		{#each memories as memory, i}
			<div class="animate-fade-in-up" style="animation-delay: {i * 100}ms">
				<MemoryCard {memory} rotation={((i % 5) - 2) * 1.5} />
			</div>
		{/each}
	</div>
{/if}
