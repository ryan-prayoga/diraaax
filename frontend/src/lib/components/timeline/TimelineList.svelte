<script lang="ts">
	import type { TimelineEvent } from '$lib/types';
	import TimelineItemCard from './TimelineItemCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';

	let {
		events = [],
		ondelete
	}: {
		events: TimelineEvent[];
		ondelete?: (id: number) => void;
	} = $props();
</script>

{#if events.length === 0}
	<EmptyState
		emoji="📅"
		title="Belum ada timeline"
		subtitle="Perjalanan cinta kalian akan ditampilkan di sini"
	/>
{:else}
	<div class="relative rounded-3xl border border-pink-100/80 bg-white/70 backdrop-blur-sm p-4 md:p-6">
		<!-- Vertical line -->
		<div class="absolute left-[1.45rem] md:left-[1.85rem] top-5 bottom-5 w-0.5 bg-linear-to-b from-pink-300 via-pink-200 to-transparent"></div>

		<div class="space-y-5 md:space-y-6">
			{#each events as event, i}
				<div class="animate-fade-in-up" style="animation-delay: {i * 120}ms">
					<TimelineItemCard {event} {ondelete} />
				</div>
			{/each}
		</div>
	</div>
{/if}
