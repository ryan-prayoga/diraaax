<script lang="ts">
	import { onMount } from 'svelte';
	import { bucketList as bucketApi } from '$lib/api';
	import type { BucketItem } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import BucketProgressCard from '$lib/components/bucket/BucketProgressCard.svelte';
	import BucketListCard from '$lib/components/bucket/BucketListCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let items = $state<BucketItem[]>([]);
	let loading = $state(true);
	let error = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			items = await bucketApi.list();
		} catch (err: any) {
			error = err.message || 'Gagal memuat bucket list';
		} finally {
			loading = false;
		}
	}

	async function handleToggle(id: number) {
		try {
			const updated = await bucketApi.toggle(id);
			items = items.map(item => item.id === id ? updated : item);
		} catch (err: any) {
			// revert on error
			console.error('Toggle failed:', err);
		}
	}

	onMount(fetchData);

	const doneItems = $derived(items.filter(i => i.is_done));
	const todoItems = $derived(items.filter(i => !i.is_done));
</script>

<svelte:head>
	<title>Bucket List — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<SectionHeader
		title="Bucket List"
		subtitle="Impian dan rencana yang ingin kita wujudkan bersama"
		emoji="✨"
	/>

	{#if loading}
		<LoadingState text="Memuat bucket list..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else if items.length === 0}
		<EmptyState
			emoji="✨"
			title="Belum ada bucket list"
			subtitle="Tambahkan impian dan rencana kalian!"
		/>
	{:else}
		<!-- Progress Card -->
		<div class="animate-fade-in-up">
			<BucketProgressCard {items} />
		</div>

		<!-- Todo Items -->
		{#if todoItems.length > 0}
			<div class="space-y-3">
				<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold">Belum Tercapai 🎯</p>
				{#each todoItems as item, i}
					<div class="animate-fade-in-up" style="animation-delay: {i * 80}ms">
						<BucketListCard {item} ontoggle={handleToggle} />
					</div>
				{/each}
			</div>
		{/if}

		<!-- Done Items -->
		{#if doneItems.length > 0}
			<div class="space-y-3">
				<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold">Sudah Tercapai 🎉</p>
				{#each doneItems as item, i}
					<div class="animate-fade-in-up" style="animation-delay: {i * 80}ms">
						<BucketListCard {item} ontoggle={handleToggle} />
					</div>
				{/each}
			</div>
		{/if}
	{/if}
</div>
