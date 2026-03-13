<script lang="ts">
	import { onMount } from 'svelte';
	import { capsules as capsulesApi } from '$lib/api';
	import type { LoveCapsule } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import CapsuleCard from '$lib/components/capsules/CapsuleCard.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let allCapsules = $state<LoveCapsule[]>([]);
	let loading = $state(true);
	let error = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			allCapsules = await capsulesApi.list();
		} catch (err: any) {
			error = err.message || 'Gagal memuat capsule';
		} finally {
			loading = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Love Capsules — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<SectionHeader
		title="Love Capsules"
		subtitle="Pesan-pesan cinta yang menunggu untuk dibuka"
		emoji="💌"
	/>

	{#if loading}
		<LoadingState text="Memuat capsule..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else if allCapsules.length === 0}
		<EmptyState
			emoji="💌"
			title="Belum ada Love Capsule"
			subtitle="Buat kapsul cinta pertamamu!"
		/>
	{:else}
		<div class="space-y-4">
			{#each allCapsules as capsule, i}
				<div class="animate-fade-in-up" style="animation-delay: {i * 100}ms">
					<CapsuleCard {capsule} />
				</div>
			{/each}
		</div>
	{/if}
</div>
