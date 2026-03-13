<script lang="ts">
	import type { TimelineEvent } from '$lib/types';
	import { formatDate } from '$lib/utils';
	import { imageUrl } from '$lib/api';
	import TimelineItemCard from './TimelineItemCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';

	let {
		events = []
	}: {
		events: TimelineEvent[];
	} = $props();
</script>

{#if events.length === 0}
	<EmptyState
		emoji="📅"
		title="Belum ada timeline"
		subtitle="Perjalanan cinta kalian akan ditampilkan di sini"
	/>
{:else}
	<div class="relative">
		<!-- Vertical line -->
		<div class="absolute left-5 top-0 bottom-0 w-0.5 bg-gradient-to-b from-pink-300 via-pink-200 to-transparent"></div>

		<div class="space-y-6">
			{#each events as event, i}
				<div class="animate-fade-in-up" style="animation-delay: {i * 120}ms">
					<TimelineItemCard {event} />
				</div>
			{/each}
		</div>
	</div>
{/if}
