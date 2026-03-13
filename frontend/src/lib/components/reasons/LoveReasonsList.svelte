<script lang="ts">
	import type { LoveReason } from '$lib/types';
	import LoveReasonCard from './LoveReasonCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';

	let {
		reasons = []
	}: {
		reasons: LoveReason[];
	} = $props();
</script>

{#if reasons.length === 0}
	<EmptyState
		emoji="💝"
		title="Belum ada alasan cinta"
		subtitle="Tulis hal-hal yang kamu cintai dari dia"
	/>
{:else}
	<div class="space-y-3">
		{#each reasons as reason, i}
			<LoveReasonCard {reason} index={i} />
		{/each}
	</div>
{/if}
