<script lang="ts">
	import { onMount } from 'svelte';
	import { loveReasons as reasonsApi } from '$lib/api';
	import type { LoveReason } from '$lib/types';
	import { shuffle } from '$lib/utils';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import LoveReasonsList from '$lib/components/reasons/LoveReasonsList.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';

	let reasons = $state<LoveReason[]>([]);
	let loading = $state(true);
	let error = $state('');

	// Add form
	let newMessage = $state('');
	let adding = $state(false);

	async function fetchData() {
		loading = true;
		error = '';
		try {
			reasons = await reasonsApi.list();
		} catch (err: any) {
			error = err.message || 'Gagal memuat love reasons';
		} finally {
			loading = false;
		}
	}

	function handleShuffle() {
		reasons = shuffle(reasons);
	}

	async function handleAdd(e: Event) {
		e.preventDefault();
		const msg = newMessage.trim();
		if (!msg) return;

		adding = true;
		try {
			const newReason = await reasonsApi.create({ message: msg });
			reasons = [newReason, ...reasons];
			newMessage = '';
		} catch (err: any) {
			error = err.message || 'Gagal menambahkan';
		} finally {
			adding = false;
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Love Reasons — diraaax 💕</title>
</svelte:head>

<div class="space-y-6">
	<div class="flex items-start justify-between gap-4">
		<SectionHeader
			title="Things I Love About You"
			subtitle="Semua alasan kenapa aku sayang kamu 💝"
			emoji="💝"
		/>
		{#if reasons.length > 1}
			<button
				onclick={handleShuffle}
				class="text-sm text-pink-400 hover:text-pink-500 transition-colors shrink-0 mt-1"
				title="Acak"
			>
				🔀
			</button>
		{/if}
	</div>

	{#if loading}
		<LoadingState text="Memuat alasan cinta..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<!-- Add new reason form -->
		<div class="love-card-static p-4 animate-fade-in-up">
			<form onsubmit={handleAdd} class="flex flex-col sm:flex-row gap-3">
				<input
					type="text"
					bind:value={newMessage}
					placeholder="Tulis alasan baru... 💕"
					class="w-full sm:flex-1 px-4 py-2.5 rounded-2xl border-2 border-pink-100 focus:border-pink-400 focus:outline-none text-sm bg-pink-50/30 placeholder:text-pink-200 transition-all"
				/>
				<button
					type="submit"
					disabled={adding || !newMessage.trim()}
					class="btn-primary text-sm px-5 py-2.5 rounded-2xl disabled:opacity-50 w-full sm:w-auto"
				>
					{adding ? '...' : '💝'}
				</button>
			</form>
		</div>

		<!-- Reasons List -->
		<div class="animate-fade-in-up" style="animation-delay: 150ms">
			<LoveReasonsList {reasons} />
		</div>
	{/if}
</div>
