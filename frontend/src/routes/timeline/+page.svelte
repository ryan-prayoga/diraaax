<script lang="ts">
	import { onMount } from 'svelte';
	import { timeline as timelineApi } from '$lib/api';
	import type { TimelineEvent } from '$lib/types';
	import SectionHeader from '$lib/components/ui/SectionHeader.svelte';
	import TimelineList from '$lib/components/timeline/TimelineList.svelte';
	import LoadingState from '$lib/components/ui/LoadingState.svelte';
	import ErrorState from '$lib/components/ui/ErrorState.svelte';
	import ErrorAlert from '$lib/components/ui/ErrorAlert.svelte';

	let events = $state<TimelineEvent[]>([]);
	let loading = $state(true);
	let error = $state('');
	let creating = $state(false);
	let actionError = $state('');

	let newTitle = $state('');
	let newDescription = $state('');
	let newDate = $state('');
	let newType = $state('special');

	async function fetchData() {
		loading = true;
		error = '';
		try {
			const list = await timelineApi.list();
			events = Array.isArray(list) ? list : [];
		} catch (err: any) {
			error = err.message || 'Gagal memuat timeline';
			events = [];
		} finally {
			loading = false;
		}
	}

	const eventCount = $derived(events.length);
	const photoCount = $derived(events.filter(event => Boolean(event.image_url)).length);

	async function handleCreateEvent(e: Event) {
		e.preventDefault();
		actionError = '';
		if (!newTitle.trim() || !newDate || !newType.trim()) {
			actionError = 'Title, tanggal, dan tipe event wajib diisi ya 💕';
			return;
		}

		creating = true;
		try {
			const created = await timelineApi.create({
				title: newTitle.trim(),
				description: newDescription.trim() || undefined,
				event_date: newDate,
				event_type: newType.trim()
			});
			events = [...events, created].sort((a, b) => new Date(a.event_date).getTime() - new Date(b.event_date).getTime());
			newTitle = '';
			newDescription = '';
			newDate = '';
			newType = 'special';
		} catch (err: any) {
			actionError = err.message || 'Gagal menambah event timeline';
		} finally {
			creating = false;
		}
	}

	async function handleDeleteEvent(id: number) {
		actionError = '';
		try {
			await timelineApi.remove(id);
			events = events.filter(event => event.id !== id);
		} catch (err: any) {
			actionError = err.message || 'Gagal menghapus event timeline';
		}
	}

	onMount(fetchData);
</script>

<svelte:head>
	<title>Timeline — diraaax 💕</title>
</svelte:head>


<div class="space-y-6 md:space-y-8">
	<SectionHeader
		title="Our Timeline"
		subtitle="Scrapbook perjalanan cinta kita, dari momen kecil sampai milestone besar"
		emoji="📅"
	/>

	{#if loading}
		<LoadingState text="Memuat timeline..." />
	{:else if error}
		<ErrorState message={error} onretry={fetchData} />
	{:else}
		<div class="love-card-static p-4 md:p-5">
			<form class="space-y-3" onsubmit={handleCreateEvent}>
				<p class="text-sm font-semibold text-rose-deep">Tambah momen baru ✨</p>
				<div class="grid grid-cols-1 md:grid-cols-3 gap-3">
					<input
						type="text"
						bind:value={newTitle}
						placeholder="Judul momen"
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
					<input
						type="date"
						bind:value={newDate}
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					/>
					<select
						bind:value={newType}
						class="px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none"
					>
						<option value="special">Special</option>
						<option value="first_meet">First Meet</option>
						<option value="date">Date</option>
						<option value="anniversary">Anniversary</option>
						<option value="travel">Travel</option>
						<option value="milestone">Milestone</option>
					</select>
				</div>
				<textarea
					bind:value={newDescription}
					rows="2"
					placeholder="Deskripsi singkat momen ini (opsional)"
					class="w-full px-4 py-2.5 rounded-2xl border border-pink-100 focus:border-pink-400 focus:outline-none resize-none"
				></textarea>
				<button type="submit" class="btn-primary text-sm px-5 py-2.5 rounded-2xl" disabled={creating}>
					{creating ? 'Menyimpan...' : 'Simpan Momen'}
				</button>
			</form>
		</div>

		{#if actionError}
			<ErrorAlert message={actionError} />
		{/if}

		<div class="love-card-static p-4 md:p-5 bg-linear-to-br from-pink-50/70 to-white">
			<div class="grid grid-cols-2 gap-3 md:gap-4">
				<div class="rounded-2xl bg-white border border-pink-100 px-4 py-3">
					<p class="text-[11px] md:text-xs uppercase tracking-wider text-rose-muted font-semibold">Total Momen</p>
					<p class="text-xl md:text-2xl font-extrabold text-rose-deep mt-1">{eventCount}</p>
				</div>
				<div class="rounded-2xl bg-white border border-pink-100 px-4 py-3">
					<p class="text-[11px] md:text-xs uppercase tracking-wider text-rose-muted font-semibold">Dengan Foto</p>
					<p class="text-xl md:text-2xl font-extrabold text-pink-500 mt-1">{photoCount}</p>
				</div>
			</div>
		</div>

		<TimelineList {events} ondelete={handleDeleteEvent} />
	{/if}
</div>
