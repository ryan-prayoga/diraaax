<script lang="ts">
	import { onMount } from 'svelte';
	import { memories as memoriesApi } from '$lib/api';
	import type { Memory } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import MemoryOfDayHero from '$lib/components/memories/MemoryOfDayHero.svelte';
	import PolaroidGrid from '$lib/components/memories/PolaroidGrid.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let allMemories = $state<Memory[]>([]);
	let randomMemory = $state<Memory | null>(null);
	let loading = $state(true);
	let error = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const [list, random] = await Promise.allSettled([
				memoriesApi.list(),
				memoriesApi.random()
			]);

			if (list.status === 'fulfilled') {
				allMemories = list.value;
			}
			if (random.status === 'fulfilled') {
				randomMemory = random.value;
			}
		} catch (err: any) {
			error = err.message || 'Gagal memuat kenangan';
		} finally {
			loading = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Memories — diraaax 💕</title>
</svelte:head>

<div class="space-y-6 md:space-y-8">
	<SectionHeader
		title="Our Memories"
		subtitle="Setiap momen bersama adalah harta 💎"
		emoji="📸"
	/>

	{#if loading}
		<LoadingState text="Memuat kenangan..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<!-- Memory of the Day -->
		{#if randomMemory}
			<div class="animate-fade-in-up">
				<MemoryOfDayHero memory={randomMemory} />
			</div>
		{/if}

		<!-- All Memories Grid -->
		<div class="animate-fade-in-up" style="animation-delay: 200ms">
			<PolaroidGrid memories={allMemories} />
		</div>
	{/if}
</div>
