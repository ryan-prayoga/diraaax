<script lang="ts">
	import { onMount } from 'svelte';
	import { moods as moodsApi } from '$lib/api';
	import type { Mood, MoodValue } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import MoodPicker from '$lib/components/moods/MoodPicker.svelte';
	import MoodHistoryStrip from '$lib/components/moods/MoodHistoryStrip.svelte';
	import EmptyState from '$lib/components/ui/EmptyState.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import ErrorAlert from '$lib/components/ui/ErrorAlert.svelte';

	let allMoods = $state<Mood[]>([]);
	let loading = $state(true);
	let error = $state('');
	let selectedMood = $state<MoodValue | null>(null);
	let submitting = $state(false);
	let submitted = $state(false);
	let submitError = $state('');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const list = await moodsApi.list();
			allMoods = Array.isArray(list) ? list : [];
		} catch (err: any) {
			error = err.message || 'Gagal memuat moods';
			allMoods = [];
		} finally {
			loading = false;
		}
	}

	async function handleSelectMood(mood: MoodValue) {
		selectedMood = mood;
		submitting = true;
		submitted = false;
		submitError = '';
		try {
			const newMood = await moodsApi.create({ mood });
			allMoods = [newMood, ...(allMoods ?? [])];
			submitted = true;
			setTimeout(() => {
				submitted = false;
				selectedMood = null;
			}, 2000);
		} catch (err: any) {
			submitError = err.message || 'Gagal menyimpan mood';
		} finally {
			submitting = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Moods — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<SectionHeader
		title="Moods"
		subtitle="Gimana perasaan kita hari ini?"
		emoji="🥰"
	/>

	{#if loading}
		<LoadingState text="Memuat moods..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		{#if submitError}
			<ErrorAlert message={submitError} />
		{/if}

		<!-- Mood Picker -->
		<div class="animate-fade-in-up">
			<MoodPicker selected={selectedMood} onselect={handleSelectMood} />
			{#if submitting}
				<p class="text-center text-rose-muted text-sm mt-3 animate-pulse">Menyimpan mood... 💭</p>
			{/if}
			{#if submitted}
				<p class="text-center text-green-500 text-sm mt-3 animate-fade-in">Mood tersimpan! 💕</p>
			{/if}
		</div>

		<!-- Mood History -->
		<div class="animate-fade-in-up" style="animation-delay: 200ms">
			<p class="text-xs text-rose-muted uppercase tracking-wider font-semibold mb-3">Riwayat Mood 💭</p>
			{#if allMoods.length === 0}
				<EmptyState
					emoji="💭"
					title="Belum ada mood tercatat"
					subtitle="Pilih mood pertamamu hari ini, lalu biarkan kisah hati ini bertumbuh"
				/>
			{:else}
				<MoodHistoryStrip moods={allMoods} />
			{/if}
		</div>
	{/if}
</div>
