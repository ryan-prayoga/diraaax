<script lang="ts">
	import { onMount } from 'svelte';
	import { timeline as timelineApi } from '$lib/api';
	import type { TimelineEvent } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import TimelineList from '$lib/components/timeline/TimelineList.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let events = $state<TimelineEvent[]>([]);
	let loading = $state(true);
	let error = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			events = await timelineApi.list();
		} catch (err: any) {
			error = err.message || 'Gagal memuat timeline';
		} finally {
			loading = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Timeline — diraaax 💕</title>
</svelte:head>

<div>
	<SectionHeader
		title="Our Timeline"
		subtitle="Perjalanan cinta kita dari hari ke hari"
		emoji="📅"
	/>

	{#if loading}
		<LoadingState text="Memuat timeline..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<TimelineList {events} />
	{/if}
</div>
