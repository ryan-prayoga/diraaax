<script lang="ts">
	import { onMount } from 'svelte';
	import { moods as moodsApi } from '$lib/api';
	import type { Mood, MoodValue } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import MoodPicker from '$lib/components/moods/MoodPicker.svelte';
	import MoodHistoryStrip from '$lib/components/moods/MoodHistoryStrip.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let allMoods = $state<Mood[]>([]);
	let loading = $state(true);
	let error = $state('');
	let selectedMood = $state<MoodValue | null>(null);
	let submitting = $state(false);
	let submitted = $state(false);

	async function fetchData() {
		loading = true;
		error = '';
		try {
			allMoods = await moodsApi.list();
		} catch (err: any) {
			error = err.message || 'Gagal memuat moods';
		} finally {
			loading = false;
		}
	}

	async function handleSelectMood(mood: MoodValue) {
		selectedMood = mood;
		submitting = true;
		submitted = false;
		try {
			const newMood = await moodsApi.create({ mood });
			allMoods = [newMood, ...allMoods];
			submitted = true;
			setTimeout(() => {
				submitted = false;
				selectedMood = null;
			}, 2000);
		} catch (err: any) {
			error = err.message || 'Gagal menyimpan mood';
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
			<MoodHistoryStrip moods={allMoods} />
		</div>
	{/if}
</div>
